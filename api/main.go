package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type queryResult struct {
	QueryText  string
	Action     string
	Parameters struct {
		Number       int
		UnitCurrency string
	}
}

type dialogRequest struct {
	ResponseID  string
	QueryResult queryResult
}

// Handler is my serverless function
func Handler(w http.ResponseWriter, r *http.Request) {
	var info dialogRequest
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Test print")
	Spew.dump(info.QueryResult.QueryText)
	fmt.Fprintf(w, info.QueryResult.QueryText)
}

//     "allRequiredParamsPresent": true,
//     "fulfillmentText": "Withhold $100. Any other income to calculate?",
//     "fulfillmentMessages": [
//       {
//         "text": {
//           "text": [
//             "Withhold $100. Any other income to calculate?"
//           ]
//         }
//       }
//     ],

//   },
//   "originalDetectIntentRequest": {
//     "source": "google",
//     "version": "2",
//     "payload": {
//       "user": {
//         "locale": "en-US",
//         "lastSeen": "2020-01-25T18:08:51Z",
//         "userVerificationStatus": "VERIFIED"
//       },
//       "conversation": {
//         "conversationId": "ABwppHH6ZiJlZEg5pdOTFKazUSsuZu2BOqVR3t3tATRqKdZGmWkDMZK25IUZcCvDDxc9VEzca5T-Kr-T-W8",
//         "type": "ACTIVE",
//         "conversationToken": "[]"
//       },
//       "inputs": [
//         {
//           "intent": "actions.intent.TEXT",
//           "rawInputs": [
//             {
//               "inputType": "VOICE",
//               "query": "123"
//             }
//           ],
//           "arguments": [
//             {
//               "name": "text",
//               "rawText": "123",
//               "textValue": "123"
//             }
//           ]
//         }
//       ],
//       "surface": {
//         "capabilities": [
//           {
//             "name": "actions.capability.ACCOUNT_LINKING"
//           },
//           {
//             "name": "actions.capability.MEDIA_RESPONSE_AUDIO"
//           },
//           {
//             "name": "actions.capability.AUDIO_OUTPUT"
//           },
//           {
//             "name": "actions.capability.SCREEN_OUTPUT"
//           }
//         ]
//       },
//       "isInSandbox": true,
//       "availableSurfaces": [
//         {
//           "capabilities": [
//             {
//               "name": "actions.capability.AUDIO_OUTPUT"
//             },
//             {
//               "name": "actions.capability.SCREEN_OUTPUT"
//             },
//             {
//               "name": "actions.capability.WEB_BROWSER"
//             }
//           ]
//         }
//       ],
//       "requestType": "SIMULATOR"
//     }
//   },
//   "session": "projects/tax-time-c294a/agent/sessions/ABwppHH6ZiJlZEg5pdOTFKazUSsuZu2BOqVR3t3tATRqKdZGmWkDMZK25IUZcCvDDxc9VEzca5T-Kr-T-W8"
// }
