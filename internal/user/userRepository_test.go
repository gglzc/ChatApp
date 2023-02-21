package user

import (
	"context"
	"testing"


	sqlmock"github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/assert"
)

// Define a mock implementation of the DBTX interface
func TestCreateUser(t *testing.T) {
	db , mock ,err:=sqlmock.New()
	if err!=nil{
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	//testcase
	var user=[] *User{
		{ID:100,Username: "examm",Email:"examm@gmail.com",Password:"example",},
		{ID:102,Username: "examm2",Email:"examm2@gmail.com",Password:"example",},
		{ID:101,Username: "examm1",Email:"examm1@gmail.com",Password:"example",},
		{ID:103,Username: "examm3",Email:"examm3@gmail.com",Password:"example",},
		{ID:104,Username: "examm4",Email:"examm4@gmail.com",Password:"example",},
		{ID:105,Username: "examm5",Email:"examm5@gmail.com",Password:"example",},
	}
	nr := NewRepository(db, nil)
	
	for r,_ := range user{
		s:=user[r]
		row:=sqlmock.NewRows([]string{"id"}).AddRow(s.ID)
		// Expect the query to be executed with the correct arguments
		mock.ExpectQuery("INSERT INTO public.user").WithArgs(s.Username, s.Password, s.Email).WillReturnRows(row)
	
		// Create a new instance of the repository and call the CreateUser method
		nrInstance, err := nr.CreateUser(context.Background(), s)
		assert.NoError(t, err)
	
		// Check that the returned user instance matches the original one
		assert.Equal(t, user[r], nrInstance)
	
	}
	
    // Check that there are no more pending mock expectations
    assert.NoError(t, mock.ExpectationsWereMet())
}

	func TestGetUserByEmail(t *testing.T) {
		db , mock ,err:=sqlmock.New()
		if err!=nil{
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		user:=&User{
			ID:       100,
			Username: "examm",
			Email:    "examm@gmail.com",
			Password: "example",
		}
		query:="SELECT id, email, username, password FROM public.user WHERE email = ?"
		
		row:=sqlmock.NewRows([]string{"id","email","username","password"}).
		AddRow(user.ID,user.Email,user.Username,user.Password)
		
		mock.ExpectQuery(query).WithArgs(user.Email).WillReturnRows(row)

		nr := NewRepository(db, nil)
		nrInstance ,err :=nr.GetUserByEmail(context.Background(),user.Email)
		assert.NoError(t, err)

		assert.Equal(t,user,nrInstance)
		assert.NoError(t, mock.ExpectationsWereMet())
	}

	func TestCheckUsernameExist(t *testing.T) {
		// create a mock database for testing
		mockDB, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create mock database: %v", err)
		}
		defer mockDB.Close()
	
		// create a repository using the mock database
		repo := NewRepository(mockDB,nil)
		//test data
		var user=[] *User{
			{ID:100,Username: "examm",Email:"examm@gmail.com",Password:"example",},
			{ID:102,Username: "examm2",Email:"examm2@gmail.com",Password:"example",},
			{ID:101,Username: "examm1",Email:"examm1@gmail.com",Password:"example",},
			{ID:103,Username: "examm3",Email:"examm3@gmail.com",Password:"example",},
			{ID:104,Username: "examm4",Email:"examm4@gmail.com",Password:"example",},
			{ID:105,Username: "examm5",Email:"examm5@gmail.com",Password:"example",},
		}
		query:="SELECT username FROM public.user WHERE username = ?"
		// set up test cases
		row:=sqlmock.NewRows([]string{"username"})
		for r,_:=range user{
			s:=user[r]
			row.AddRow(s.Username)
					
			mock.ExpectQuery(query).WithArgs(s.Username).WillReturnRows(row)
			
			c,err:=repo.CheckUsernameExist(context.Background(),s.Username)
			assert.NoError(t,err)

			assert.True(t,c)	
		}
		assert.NoError(t, mock.ExpectationsWereMet())
	}


	func TestCheckEmailExist(t *testing.T) {
		// create a mock database for testing
		mockDB, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("failed to create mock database: %v", err)
		}
		defer mockDB.Close()
	
		// create a repository using the mock database
		repo := NewRepository(mockDB,nil)
		//test data
		var user=[] *User{
			{ID:100,Username: "examm",Email:"examm@gmail.com",Password:"example",},
			{ID:102,Username: "examm2",Email:"examm2@gmail.com",Password:"example",},
			{ID:101,Username: "examm1",Email:"examm1@gmail.com",Password:"example",},
			{ID:103,Username: "examm3",Email:"examm3@gmail.com",Password:"example",},
			{ID:104,Username: "examm4",Email:"examm4@gmail.com",Password:"example",},
			{ID:105,Username: "examm5",Email:"examm5@gmail.com",Password:"example",},
		}
		query:="SELECT email FROM public.user WHERE email = ?"
		// set up test cases
		row:=sqlmock.NewRows([]string{"email"})
		for r,_:=range user{
			s:=user[r]
			row.AddRow(s.Email)
					
			mock.ExpectQuery(query).WithArgs(s.Email).WillReturnRows(row)
			
			c,err:=repo.CheckEmailExist(context.Background(),s.Email)
			assert.NoError(t,err)

			assert.True(t,c)	
		}
		assert.NoError(t, mock.ExpectationsWereMet())
	}
	
	func TestCreateUserByCache(t *testing.T) {
		
	}