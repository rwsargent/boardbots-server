package testingutils

import (
	"boardbots/server/player"
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"io"
	"net/http/httptest"
)

var TestUUID = uuid.MustParse("c67a791f-1d1b-41ae-b21b-14f79d4fdf66")
var TestMissingUUID = uuid.MustParse("df834d72-8d96-4010-bb37-36c60d9309cd")

func GetPayloadFromResult(httpError *echo.HTTPError, response *interface{}) interface{} {
	return httpError.Message.(interface{})
}

func ToJson(body interface{}) io.Reader {
	b, err := json.Marshal(body)
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(b)
}

func ReadBodyFromRecorder(payload *httptest.ResponseRecorder, response interface{}) {
	json.NewDecoder(payload.Body).Decode(response)
}

/**
 *
 */

var (
	GameIds = []uuid.UUID{
		uuid.MustParse("98ae983e-3f04-42ab-928a-c399d6d18375"),
		uuid.MustParse("5341acab-6e28-4d28-8530-8716e0c3dd2e"),
		uuid.MustParse("790bcc3f-6e72-4a0e-a6ea-bc806aa8aa03"),
		uuid.MustParse("6c8420b5-e7f5-4328-ae29-4dbdf7537612"),
		uuid.MustParse("f3282245-f546-4c71-92ca-5bada1f9c037"),
		uuid.MustParse("9cae2aa0-d21a-48ab-a877-4b78942259e4"),
		uuid.MustParse("0ad943b2-6ea9-45ad-9098-f67714652fcd"),
		uuid.MustParse("93ded37f-57d3-4b43-8933-1164e086a881"),
		uuid.MustParse("5b399bd3-aa3e-4754-bb51-175b30b77400"),
		uuid.MustParse("f7ea9019-033b-41e7-a671-26231952cd8c"),
	}

	PlayerIds = []uuid.UUID{
		uuid.MustParse("cb021c4f-b85c-4923-972c-d0a130282c41"),
		uuid.MustParse("37e45b58-034d-4805-8cef-f32e87d20403"),
		uuid.MustParse("d72c8ded-a088-4e75-aad5-62ccf3162df4"),
		uuid.MustParse("3bacbc0b-3a72-4a13-9038-03c37eeea55e"),
		uuid.MustParse("2ed86929-b645-424c-b448-2af610f00318"),
		uuid.MustParse("917b610f-8fb9-47bd-85bd-ecd3844c300d"),
		uuid.MustParse("944470f1-2bb4-4992-95d0-036fc5133a8e"),
		uuid.MustParse("4e2b5467-4cd3-4dd1-a24e-b0624e317e3e"),
		uuid.MustParse("26e72be6-6030-45c9-9382-4f5d8df531e3"),
		uuid.MustParse("f2c5ec6b-3839-4dfa-bf40-233340af904c"),
		uuid.MustParse("7e66e953-e09f-40e0-90b8-e7650834d73b"),
		uuid.MustParse("a54a5487-a15e-48fc-90e5-c1fd4deb4fdf"),
		uuid.MustParse("ca13b702-0e94-4aa3-8484-ea73827ef6cf"),
		uuid.MustParse("ad93696f-e226-49d4-a3b4-36dc8b4cfa64"),
		uuid.MustParse("055ea5e1-cd8d-4265-87f4-d258d43f796b"),
		uuid.MustParse("c8ad7c2f-f9d8-45b7-b992-99e9a808e404"),
		uuid.MustParse("b9bcabff-0cb9-4009-978d-a083071ef400"),
		uuid.MustParse("11479a55-ef3e-41b5-bd74-979d90739de2"),
		uuid.MustParse("8ab34e6b-e1bd-45c8-9fc3-9c5dbb73774f"),
		uuid.MustParse("ed4a1d00-9838-4627-89f1-ceaba8c8144c"),
	}

	PlayerPrinciples = []player.PlayerPrinciple{
		{
			UserName: "user0",
			Password: "password0",
			UserId:   PlayerIds[0],
		},
		{
			UserName: "user1",
			Password: "password1",
			UserId:   PlayerIds[1],
		},
		{
			UserName: "user2",
			Password: "password2",
			UserId:   PlayerIds[2],
		},
		{
			UserName: "user3",
			Password: "password3",
			UserId:   PlayerIds[3],
		},
		{
			UserName: "user4",
			Password: "password4",
			UserId:   PlayerIds[4],
		},
		{
			UserName: "user5",
			Password: "password5",
			UserId:   PlayerIds[5],
		},
		{
			UserName: "user6",
			Password: "password6",
			UserId:   PlayerIds[6],
		},
		{
			UserName: "user7",
			Password: "password7",
			UserId:   PlayerIds[7],
		}, {
			UserName: "user8",
			Password: "password8",
			UserId:   PlayerIds[8],
		}, {
			UserName: "user9",
			Password: "password9",
			UserId:   PlayerIds[9],
		},
		{
			UserName: "user10",
			Password: "password10",
			UserId:   PlayerIds[10],
		}, {
			UserName: "user11",
			Password: "password11",
			UserId:   PlayerIds[11],
		}, {
			UserName: "user12",
			Password: "password12",
			UserId:   PlayerIds[12],
		},
		{
			UserName: "user13",
			Password: "password13",
			UserId:   PlayerIds[13],
		},
		{
			UserName: "user14",
			Password: "password14",
			UserId:   PlayerIds[14],
		}, {
			UserName: "user15",
			Password: "password15",
			UserId:   PlayerIds[15],
		},
		{
			UserName: "user16",
			Password: "password16",
			UserId:   PlayerIds[16],
		},
		{
			UserName: "user17",
			Password: "password17",
			UserId:   PlayerIds[17],
		},
		{
			UserName: "user18",
			Password: "password18",
			UserId:   PlayerIds[18],
		},
		{
			UserName: "user19",
			Password: "password19",
			UserId:   PlayerIds[19],
		},
	}
)
