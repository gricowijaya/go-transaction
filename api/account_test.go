package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/gricowijaya/go-transaction/db/mock"
	db "github.com/gricowijaya/go-transaction/db/sqlc"
	"github.com/gricowijaya/go-transaction/util"
	"github.com/stretchr/testify/require"
)

// to generate a random account for testing
func randomAccount() db.Accounts {
	return db.Accounts{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, accounts db.Accounts) {
	data, err := ioutil.ReadAll(body) // call the ioutil to read all the response from the body
	require.NoError(t, err)           // check the error from the response

	var gotAccount db.Accounts              // store the object we got from the response
	err = json.Unmarshal(data, &gotAccount) // unmarshal the data from the gotAccount
	require.NoError(t, err)
	require.Equal(t, accounts, gotAccount)
}

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	testCases := []struct {
		name          string
		accountID     int64
		buildStubs    func(store *mockdb.MockStore)                           // take the mock store as an input
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder) // for checking the output of the API
	}{
		{ // test data
			name:      "OK",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{ // test data with internal server error case
			name:      "InternalError",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Accounts{}, sql.ErrConnDone) // return ErrConnectionDone because it is possible the connection is returned from the pool
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code) // the status is internal error because the connection from the database server is broken 
			},
		},
		{ // test data with Bad Request server error case
			name:      "BadRequest",
			accountID: 0, // update the data to 0 because the minimum id is 1
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()). 					
          Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code) // the status is internal error because the connection from the database server is broken 
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)
			server := NewServer(store)
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/accounts/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
	// create new controller mocker
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // check all method should called when it's called

	// store the data into mock db
	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).            // expected to be called one time
		Return(account, nil) // should be the same as the function of Querier

	// creates an http test
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	// store the response into recorder
	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	// add the requireBodyMatchAccount
	requireBodyMatchAccount(t, recorder.Body, account)
}
