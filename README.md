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
