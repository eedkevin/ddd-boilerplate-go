package usecase_test

import (
	"ddd-boilerplate/internal/app/application/service"
	"ddd-boilerplate/internal/app/application/usecase"
	"ddd-boilerplate/internal/app/domain/repository"
	"ddd-boilerplate/internal/app/domain/vo"
	"ddd-boilerplate/testdata"
	"fmt"
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v3"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	testcases       []texture
	twilioMock      service.ITelecom
	servicePoolMock repository.IServicePool
)

type texture struct {
	Input struct {
		Data struct {
			Paticipants []vo.Paticipant `yaml:"paticipants"`
		} `yaml:"data"`
	} `yaml:"input"`
	Expect struct {
		Error bool `yaml:"error"`
		Data  struct {
			Paticipants []vo.Paticipant `yaml:"paticipants"`
		} `yaml:"data"`
	} `yaml:"expect"`
}

func setup(t *testing.T) {
	initTestCase(t)
	mock(t)
}

func mock(t *testing.T) {
	twilioMock = testdata.MTwilio{}
	servicePoolMock = testdata.MTwilioServicePool{}
}

func initTestCase(t *testing.T) {
	file, err := ioutil.ReadFile("createsession.fixture.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(file), &testcases)
	if err != nil {
		panic(err)
	}
}

func TestCreateSession(t *testing.T) {
	setup(t)

	Convey("USECASE CreateSession", t, func() {
		for _, test := range testcases {
			Convey(fmt.Sprintf("Given a valid request: %v", test.Input.Data), func() {
				session, err := usecase.CreateSession(twilioMock, servicePoolMock, test.Input.Data.Paticipants...)

				Convey("Should create a valid session", func() {
					So(err, ShouldBeNil)
					So(session.ID, ShouldNotBeNil)
					So(session.Paticipants, ShouldNotBeNil)
					So(session.Expiry, ShouldNotBeEmpty)
				})
			})
		}
	})
}
