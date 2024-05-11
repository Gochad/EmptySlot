# EmptySlot
A web application for booking any service and making payments for it through the Stripe.

## Run

1. prepare `.env` file (see `example.env`)

### frontend
1. `cd frontend`
1. make sure if you make `npm ci`
2. with `npm start`


### backend 
1. `cd backend`
1. with `docker compose up`

## Sequence diagram
```mermaid
sequenceDiagram
    participant U as User
    participant A as Admin
    participant Auth as Authentication Service
    participant SS as Slot Service
    participant DB as Database

    U->>+Auth: Provide credentials
    Auth-->>-U: Authentication Response (Success/Fail)

    alt Authentication Success
        U->>+SS: Request to view slots
        SS->>+DB: Fetch available slots
        DB-->>-SS: Return slots data
        SS-->>-U: Display slots

        U->>+SS: Select and book slot
        SS->>+DB: Update slot status
        DB-->>-SS: Confirm update
        SS-->>-U: Booking confirmation

        A->>+SS: Add new slot
        SS->>+DB: Insert new slot
        DB-->>-SS: Confirm insertion
        SS-->>-A: Slot added confirmation
    else Authentication Failed
        Auth-->>U: Show error
    end
```

## Database
```mermaid
erDiagram
    Category ||--o{ Merchandise : "has many"
    Merchandise ||--|| Category : "belongs to"
    Merchandise ||--o{ Reservation : "has many"
    Reservation }o--|| Merchandise : "includes"
    User ||--o{ Reservation : "makes"
    Reservation }o--|| User : "made by"
    User ||--o{ History : "logged in"
    History }o--|| User : "logs"
    Reservation ||--o{ History : "tracked by"
    History }o--|| Reservation : "tracks"

    Category {
        string ID PK "Unique Identifier"
        string Name "Category Name"
        string Color "Visual Identifier"
    }

    History {
        string ID PK "Unique Identifier"
        string Name "Event Name"
        string UserID FK "User Involved"
        string ReservationID FK "Reservation Involved"
    }

    Merchandise {
        string ID PK "Unique Identifier"
        string Name "Merchandise Name"
        string Description "Details"
        string CategoryID FK "Belongs to Category"
        string ReservationID FK "Linked Reservation"
        int64 Price "Cost"
        bool Confirmed "Status"
        string StartTime "Start Time"
        string EndTime "End Time"
    }

    Reservation {
        string ID PK "Unique Identifier"
        string Name "Reservation Name"
        string Description "Details"
        bool IsReserved "Reservation Status"
        int64 CalculatedPrice "Total Cost"
    }

    User {
        string ID PK "Unique Identifier"
        string Username "Username"
        string Email "Email Address"
        string Password "Password (hashed)"
        string Address "Physical Address"
        string Phone "Telephone Number"
        int Role "User Role"
        string ReservationID FK "User's Reservation"
    }
```
