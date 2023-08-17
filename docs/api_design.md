# API Design

## Objects

- user
  - name: String
  - email: String
  - id (pk): String - generated
  - password: String
- room
  - floor: Int
  - door_number (pk): Int
  - description: String
- reservation
  - id (pk): String - generated
  - user (fk): String
  - room (fk): Int
  - date: Int - epoch on date, no time

## Endpoints

- /api/v1/users/{id}
  - GET -> returns id, name
  - PATCH, DELETE is protected to user
- /api/v1/users
  - GET
  - POST -> returns email, id, name
- /api/v1/rooms/{door_number}
  - GET
- /api/v1/rooms
  - GET
    - ?free={date} -> epoch
    - ?oneWeekStatusFrom={date} -> epoch
- /api/v1/reservations/{id}
  - GET -> returns id, date in "2006-01-02", user's name, room's door num
  - PATCH, DELETE is protected to user
- /api/v1/reservations
  - GET
  - POST is protected to user -> give "doorNumber" and "date" in timestamp (same as GET)
- /login
  - POST -> returns token

## JWT

- Duration: 10 min
  - Authentication/Bearer {token}
- PW: root-pw
- Claims: UserID = uuid

## Return Codes

- Reach nonexistent route: 404
- Use non supported method: 405
- Not logged in when editing: 401
- Editing something that's not yours: 403
- Something is wrong with the request (i.e Tried to update a generated field): 400
- Create user/reservation: 201 (return the created object)
- Updated user/reservation: 200 (return the created object)
- Delete user/reservation: 204 (no return)

## Error messages

- {"message": "lower cased message here"}
