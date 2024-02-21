# AuthStartRequestInstantAuthConsent


## Fields

| Field                                                                                                                                                                                                                                  | Type                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                            | Example                                                                                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `CollectedTimestamp`                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                     | CollectedTimestamp is the date/time when original consent was collected by the client in ISO 8601 format.<br/>Required if status = optedIn.                                                                                            | 2022-02-17 00:00:00 +0000 UTC                                                                                                                                                                                                          |
| `Description`                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                     | Description of the type of consent (electronic/paper), use case, and reference to the terms and conditions (T&C)<br/>version (if applicable). Required if status = optedIn.<br/>Acceptable characters are: alphanumeric with symbols '-._+=/'. | Customer Application Name                                                                                                                                                                                                              |
| `Status`                                                                                                                                                                                                                               | *string*                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                     | Status denotes whether consent has been provided by the mobile subscriber for this transaction. Acceptable values<br/>are: optedIn, optedOut, notCollected, unknown.                                                                   | optedIn                                                                                                                                                                                                                                |
| `TransactionID`                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                     | TransactionID uniquely identifies the consent collected by the client. Required if status = optedIn.<br/>Acceptable characters are: alphanumeric with symbols '-._+=/'.                                                                | eba12f3a-5555-47bc-b85d-21c0cbc4b973                                                                                                                                                                                                   |