package test_book_test

import (
	"NewsBack/internal/api/Router"
	"NewsBack/internal/db"
	"NewsBack/internal/domain"
	"NewsBack/internal/repository"
	"NewsBack/internal/usecase"
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("NewsPass", func() {
	var Repository repository.NewsRepository
	var configDB string = "host=localhost user=postgres password=passwordtest dbname=Todos port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	TestDataOneObject := domain.News{Title: "Nik", Description: "password", ID: 1, UserID: 1, NameImage: "image/src:image.png"}
	TestDataDelete := domain.News{Title: "", Description: "", ID: 0, UserID: 0, NameImage: ""}
	TestDataFindAll := []domain.News{domain.News{Title: "Nik", Description: "password", ID: 1, UserID: 1, NameImage: "image/src:image.png"}}

	BeforeEach(func() {
		bd, _ := db.ConnectPGX(configDB)
		Repository = repository.NewNewsRepository(bd)
	})

	Describe("work to object in the database", func() {
		It("creating should return the same object", func() {
			Expect(Repository.Save(TestDataOneObject)).To(Equal(TestDataOneObject))
		})
		It("find all should return the array objects", func() {
			Expect(Repository.FindAll()).To(Equal(TestDataFindAll))
		})
		It("find one should return the one object", func() {
			Expect(Repository.FindOne(1)).To(Equal(TestDataOneObject))
		})

		It("delete should return the nil object", func() {
			Expect(Repository.DeleteById(1)).To(Equal(TestDataDelete))
		})
	})

	var UseCase Router.NewsUseCase

	BeforeEach(func() {
		UseCase = usecase.NewNewsUseCase(Repository)
	})

	Describe("work to object in the database", func() {
		It("creating should return the same object", func() {
			Expect(UseCase.Save(TestDataOneObject)).To(Equal(TestDataOneObject))
		})
		It("find all should return the array objects", func() {
			Expect(UseCase.FindAll()).To(Equal(TestDataFindAll))
		})
		It("find one should return the one object", func() {
			Expect(UseCase.FindOne(1)).To(Equal(TestDataOneObject))
		})

		It("delete should return the nil object", func() {
			Expect(UseCase.DeleteById(1)).To(Equal(TestDataDelete))
		})

	})
	type test struct {
		expectedCode int // expected HTTP status code
	}

	var testReqPass test
	var NewsRouter *Router.NewsHandler
	var app *fiber.App
	var reqSave, reqGetOne, reqAll, reqDelete *http.Request

	BeforeEach(func() {
		var TestDataOneObjectBuf bytes.Buffer
		_ = json.NewEncoder(&TestDataOneObjectBuf).Encode(TestDataOneObject)

		app = fiber.New()

		NewsRouter = Router.NewNewsRouter(UseCase)

		app.Post("/News", NewsRouter.Save)
		app.Get("/News/:id", NewsRouter.FindOne)
		app.Get("/Newss", NewsRouter.FindAll)
		app.Delete("/News/:id", NewsRouter.DeleteById)

		reqSave, _ = http.NewRequest("POST", "/News", &TestDataOneObjectBuf)
		reqGetOne, _ = http.NewRequest("GET", "/News/1", nil)
		reqAll, _ = http.NewRequest("GET", "/Newss", nil)
		reqDelete, _ = http.NewRequest("DELETE", "/News/1", nil)

		testReqPass = test{expectedCode: 200}
	})

	Describe("pass work to handlers ", func() {
		It("creating should return 200 statusCode", func() {
			resp, _ := app.Test(reqSave)
			Expect(resp.StatusCode).To(Equal(testReqPass.expectedCode))
		})
		It("find all should return 200 statusCode", func() {

			resp, _ := app.Test(reqAll)

			Expect(resp.StatusCode).To(Equal(testReqPass.expectedCode))
		})
		It("find one should return 200 statusCode", func() {

			resp, _ := app.Test(reqGetOne)

			Expect(resp.StatusCode).To(Equal(testReqPass.expectedCode))
		})

		It("delete should return 200 statusCode", func() {
			resp, _ := app.Test(reqDelete)
			Expect(resp.StatusCode).To(Equal(testReqPass.expectedCode))
		})
	})
})

var _ = Describe("NewsFail", func() {
	var Repository repository.NewsRepository
	var configDB string = "host=localhost user=postgres password=passwordtest dbname=Todos port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	TestDataOneObject := domain.News{Title: "Nik", Description: "password", ID: 1, UserID: 1, NameImage: "image/src:image.png"}
	TestNilDataOneObject := domain.News{Title: "", Description: "", ID: 0, UserID: 0, NameImage: ""}
	TestFailDataOneObject := struct{ id int }{id: 0}

	BeforeEach(func() {
		bd, _ := db.ConnectPGX(configDB)
		Repository = repository.NewNewsRepository(bd)
	})

	var UseCase Router.NewsUseCase

	BeforeEach(func() {
		UseCase = usecase.NewNewsUseCase(Repository)
	})

	type test struct {
		expectedCode int // expected HTTP status code
	}

	var testReqFail, testReqFailId, testReqNilData, testReqFailData test
	var NewsRouter *Router.NewsHandler
	var app *fiber.App
	var reqSave, reqGetOne, reqAll, reqDelete, reqGetOneFailId, reqDeleteFailId, reqSaveNilData, reqSaveFailData *http.Request

	BeforeEach(func() {
		var TestDataOneObjectBuf, TestNilDataOneObjectBuf, TestFailDataOneObjectBuf bytes.Buffer
		_ = json.NewEncoder(&TestDataOneObjectBuf).Encode(TestDataOneObject)
		_ = json.NewEncoder(&TestNilDataOneObjectBuf).Encode(TestNilDataOneObject)
		_ = json.NewEncoder(&TestFailDataOneObjectBuf).Encode(TestFailDataOneObject)

		app = fiber.New()

		NewsRouter = Router.NewNewsRouter(UseCase)

		app.Post("/News", NewsRouter.Save)
		app.Get("/News/:id", NewsRouter.FindOne)
		app.Get("/Newss", NewsRouter.FindAll)
		app.Delete("/News/:id", NewsRouter.DeleteById)

		reqSave, _ = http.NewRequest("POST", "/NewsOne", &TestDataOneObjectBuf)
		reqSaveNilData, _ = http.NewRequest("POST", "/News", &TestNilDataOneObjectBuf)
		reqSaveFailData, _ = http.NewRequest("POST", "/News", &TestFailDataOneObjectBuf)
		reqGetOne, _ = http.NewRequest("GET", "/News/12", nil)
		reqAll, _ = http.NewRequest("GET", "/Newss/1", nil)
		reqDelete, _ = http.NewRequest("DELETE", "/News/12", nil)
		reqGetOneFailId, _ = http.NewRequest("GET", "/News/fas12", nil)
		reqDeleteFailId, _ = http.NewRequest("DELETE", "/News/jd11", nil)

		testReqFail = test{expectedCode: 404}
		testReqFailId = test{expectedCode: 500}
		testReqNilData = test{expectedCode: 400}
		testReqFailData = test{expectedCode: 400}
	})

	Describe("fail work to handlers ", func() {
		It("creating should return 404 statusCode", func() {
			resp, _ := app.Test(reqSave)
			Expect(resp.StatusCode).To(Equal(testReqFail.expectedCode))
		})
		It("find all should return 404 statusCode", func() {

			resp, _ := app.Test(reqAll)

			Expect(resp.StatusCode).To(Equal(testReqFail.expectedCode))
		})
		It("find one should return 404 statusCode", func() {

			resp, _ := app.Test(reqGetOne)

			Expect(resp.StatusCode).To(Equal(testReqFail.expectedCode))
		})

		It("delete should return 404 statusCode", func() {
			resp, _ := app.Test(reqDelete)
			Expect(resp.StatusCode).To(Equal(testReqFail.expectedCode))
		})

		It("find one with failId should return 500 statusCode", func() {
			resp, _ := app.Test(reqGetOneFailId)
			Expect(resp.StatusCode).To(Equal(testReqFailId.expectedCode))
		})

		It("delete with failId should return 500 statusCode", func() {
			resp, _ := app.Test(reqDeleteFailId)
			Expect(resp.StatusCode).To(Equal(testReqFailId.expectedCode))
		})

		It("save with fail data should return 400 statusCode", func() {
			resp, _ := app.Test(reqSaveNilData)
			Expect(resp.StatusCode).To(Equal(testReqNilData.expectedCode))
		})

		It("save with fail data should return 400 statusCode", func() {
			resp, _ := app.Test(reqSaveFailData)
			Expect(resp.StatusCode).To(Equal(testReqFailData.expectedCode))
		})

	})
})
