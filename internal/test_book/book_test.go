package test_book_test

//import (
//	. "github.com/onsi/ginkgo/v2"
//	. "github.com/onsi/gomega"
//)
//
//type Students struct {
//	Name  string
//	Year  int
//	Major string
//	CGPA  float64
//}
//
//var _ = Describe("Student", func() {
//	var Student1, Student2 *Students
//
//	BeforeEach(func() {
//		Student1 = &Students{
//			Name:  "Alex",
//			Year:  3,
//			Major: "Electronics & Communication",
//			CGPA:  9.4,
//		}
//
//		Student2 = &Students{
//			Name:  "Arjun",
//			Year:  4,
//			Major: "Computer Science",
//			CGPA:  9.7,
//		}
//
//	})
//
//	Describe("Categorizing students", func() {
//		Context("with CGPA above 9", func() {
//			It("must be a topper", func() {
//				Expect(Student1.CGPA).To(BeNumerically(">=", 10))
//			})
//		})
//
//		Context("with computer science as Major", func() {
//			It("must be a Computer Science Student", func() {
//				Expect(Student2.Major).To(Equal("Computer Science"))
//			})
//		})
//
//	})
//})
