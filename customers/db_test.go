package customers_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/xerardoo/go-microservice-example/customers"
	"gorm.io/gorm"
	"regexp"
	"syreclabs.com/go/faker"
	"testing"
)

func TestCustomer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Customer Suite")
}

var _ = Describe("Customer", func() {
	var db *sql.DB
	var mock sqlmock.Sqlmock
	var repo *Repos
	var customer Customer
	var ctx context.Context

	BeforeEach(func() {
		var err error
		db, mock, err = sqlmock.New() // mock sql.DB
		Expect(err).ShouldNot(HaveOccurred())
		repo = &Repos{Db: db}
		ctx = context.Background()

		customer = Customer{
			Model:       gorm.Model{ID: uint(faker.Number().NumberInt32(4))},
			FirstName:   faker.Name().FirstName(),
			LastName:    faker.Name().LastName(),
			PhoneNumber: faker.PhoneNumber().String(),
			Address:     faker.Address().String(),
			City:        faker.Address().City(),
			State:       faker.Address().State(),
			ZipCode:     faker.Address().ZipCode(),
		}
	})

	AfterEach(func() {
		defer db.Close()
		err := mock.ExpectationsWereMet()
		Expect(err).NotTo(HaveOccurred())
	})

	Context("Get All Customers", func() {

		It("Empty", func() {
			query := "SELECT * FROM customers"
			rows := sqlmock.NewRows(nil)
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

			c, err := repo.GetAllCustomers(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(c).To(BeEmpty())
		})

		It("Exist", func() {
			query := "SELECT * FROM `customers`"
			newId := uint(faker.Number().NumberInt32(4))
			rows := sqlmock.
				NewRows([]string{"id", "first_name", "last_name", "phone_number", "address", "city", "state", "zip_code"}).
				AddRow(customer.ID, customer.FirstName, customer.LastName, customer.PhoneNumber, customer.Address, customer.City, customer.State, customer.ZipCode).
				AddRow(newId, customer.FirstName, customer.LastName, customer.PhoneNumber, customer.Address, customer.City, customer.State, customer.ZipCode)
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

			c, err := repo.GetAllCustomers(ctx)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(c[0].ID).Should(Equal(customer.ID))
			Expect(c[0].FirstName).Should(Equal(customer.FirstName))
			Expect(c[1].ID).Should(Equal(newId))
		})
	})

	Context("Find One", func() {

		It("Empty", func() {
			query := "SELECT * FROM `customers` WHERE `customers`.`id` = ?"
			rows := sqlmock.NewRows(nil)
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

			_, err := repo.GetCustomerById(ctx, customer.ID)
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})

		It("Exist", func() {
			query := "SELECT * FROM `customers` WHERE `customers`.`id` = ?"
			rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "phone_number", "address", "city", "state", "zip_code"}).
				AddRow(customer.ID, customer.FirstName, customer.LastName, customer.PhoneNumber, customer.Address, customer.City, customer.State, customer.ZipCode)
			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

			c, err := repo.GetCustomerById(ctx, customer.ID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(c.ID).Should(Equal(customer.ID))
			Expect(c.FirstName).Should(Equal(customer.FirstName))
		})
	})

	Context("Save", func() {

		It("Add", func() {
			query := "INSERT INTO `customers` (`created_at`,`updated_at`,`deleted_at`,`first_name`,`last_name`,`phone_number`,`address`,`city`,`state`,`zip_code`,`id`) VALUES (?,?,?,?,?,?,?,?,?,?,?)"

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(query)).
				WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					customer.FirstName,
					customer.LastName,
					customer.PhoneNumber,
					customer.Address,
					customer.City,
					customer.State,
					customer.ZipCode,
					customer.ID,
				).
				WillReturnResult(sqlmock.NewResult(int64(customer.ID), 1))
			mock.ExpectCommit()

			_, err := repo.CreateCustomer(ctx, customer)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Update", func() {
			query := "UPDATE `customers` SET `created_at`=?,`updated_at`=?,`deleted_at`=?,`first_name`=?,`last_name`=?,`phone_number`=?,`address`=?,`city`=?,`state`=?,`zip_code`=? WHERE `id` = ?"
			customer.ID = 1

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(query)).
				WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					customer.FirstName,
					customer.LastName,
					customer.PhoneNumber,
					customer.Address,
					customer.City,
					customer.State,
					customer.ZipCode,
					customer.ID,
				).WillReturnResult(sqlmock.NewResult(int64(customer.ID), 1))
			mock.ExpectCommit()

			_, err := repo.UpdateCustomer(ctx, customer)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Context("Delete", func() {
		It("Soft-Delete", func() {
			query := " UPDATE `customers` SET `deleted_at`=? WHERE `customers`.`id` = ? AND `customers`.`deleted_at` IS NULL"

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(query)).
				WithArgs(
					sqlmock.AnyArg(),
					customer.ID,
				).
				WillReturnResult(sqlmock.NewResult(0, 1))
			mock.ExpectCommit()

			err := repo.DeleteCustomer(ctx, customer.ID)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
