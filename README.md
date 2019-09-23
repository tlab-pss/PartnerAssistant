# PartnerAssistant Module

This repository is Partner Assistant Module in PSS.

## Usage

### send request

`<base_url>/:moduls`

* modules list

  * defualt(sample bot)
  * LINE BOT

* request params 

| params  | type     | details                  |
| :------ | :------- | :----------------------- |
| message | `string` | Message sent by the user |

* response types (supposition)

  * image
  * location
  * text
  * carousel

* response params 

| params  | type            | details                          |
| :------ | :-------------- | :------------------------------- |
| type    | `string`        | `image`,`location`,`text` etc... |
| message | `string | null` | Message returned by AI           |

## Format

```
$ docker-compose exec app gofmt -w .
```
