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

var _ = Describe("TagPass", func() {
	var Repository repository.TagRepository
	var configDB string = "host=localhost user=postgres password=passwordtest dbname=Todos port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	TestDataOneObject := domain.Tag{Name: "Nik", ID: 1}
	TestDataDelete := domain.Tag{Name: "", ID: 0}
	TestDataFindAll := []domain.Tag{domain.Tag{Name: "Nik", ID: 1}}

	BeforeEach(func() {
		bd, _ := db.Connect(configDB)
		Repository = repository.NewTagRepository(bd)
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

	var UseCase Router.TagUseCase

	BeforeEach(func() {
		UseCase = usecase.NewTagUseCase(Repository)
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
	var TagRouter *Router.TagHandler
	var app *fiber.App
	var reqSave, reqGetOne, reqAll, reqDelete *http.Request

	BeforeEach(func() {
		var TestDataOneObjectBuf bytes.Buffer
		_ = json.NewEncoder(&TestDataOneObjectBuf).Encode(TestDataOneObject)

		app = fiber.New()

		TagRouter = Router.NewTagRouter(UseCase)

		app.Post("/Tag", TagRouter.Save)
		app.Get("/Tag/:id", TagRouter.FindOne)
		app.Get("/Tags", TagRouter.FindAll)
		app.Delete("/Tag/:id", TagRouter.DeleteById)

		reqSave, _ = http.NewRequest("POST", "/Tag", &TestDataOneObjectBuf)
		reqGetOne, _ = http.NewRequest("GET", "/Tag/1", nil)
		reqAll, _ = http.NewRequest("GET", "/Tags", nil)
		reqDelete, _ = http.NewRequest("DELETE", "/Tag/1", nil)

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

var _ = Describe("TagFail", func() {
	var Repository repository.TagRepository
	var configDB string = "host=localhost user=postgres password=passwordtest dbname=Todos port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	TestDataOneObject := domain.Tag{Name: "Nik", ID: 1}
	TestNilDataOneObject := domain.Tag{Name: "", ID: 0}
	TestFailDataOneObject := struct{ id int }{id: 0}

	BeforeEach(func() {
		bd, _ := db.Connect(configDB)
		Repository = repository.NewTagRepository(bd)
	})

	var UseCase Router.TagUseCase

	BeforeEach(func() {
		UseCase = usecase.NewTagUseCase(Repository)
	})

	type test struct {
		expectedCode int // expected HTTP status code
	}

	var testReqFail, testReqFailId, testReqNilData, testReqFailData test
	var TagRouter *Router.TagHandler
	var app *fiber.App
	var reqSave, reqGetOne, reqAll, reqDelete, reqGetOneFailId, reqDeleteFailId, reqSaveNilData, reqSaveFailData *http.Request

	BeforeEach(func() {
		var TestDataOneObjectBuf, TestNilDataOneObjectBuf, TestFailDataOneObjectBuf bytes.Buffer
		_ = json.NewEncoder(&TestDataOneObjectBuf).Encode(TestDataOneObject)
		_ = json.NewEncoder(&TestNilDataOneObjectBuf).Encode(TestNilDataOneObject)
		_ = json.NewEncoder(&TestFailDataOneObjectBuf).Encode(TestFailDataOneObject)

		app = fiber.New()

		TagRouter = Router.NewTagRouter(UseCase)

		app.Post("/Tag", TagRouter.Save)
		app.Get("/Tag/:id", TagRouter.FindOne)
		app.Get("/Tags", TagRouter.FindAll)
		app.Delete("/Tag/:id", TagRouter.DeleteById)

		reqSave, _ = http.NewRequest("POST", "/TagOne", &TestDataOneObjectBuf)
		reqSaveNilData, _ = http.NewRequest("POST", "/Tag", &TestNilDataOneObjectBuf)
		reqSaveFailData, _ = http.NewRequest("POST", "/Tag", &TestFailDataOneObjectBuf)
		reqGetOne, _ = http.NewRequest("GET", "/Tag/12", nil)
		reqAll, _ = http.NewRequest("GET", "/Tags/1", nil)
		reqDelete, _ = http.NewRequest("DELETE", "/Tag/12", nil)
		reqGetOneFailId, _ = http.NewRequest("GET", "/Tag/fas12", nil)
		reqDeleteFailId, _ = http.NewRequest("DELETE", "/Tag/jd11", nil)

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
