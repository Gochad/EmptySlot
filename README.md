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


%%{init: {'theme': 'base', 'themeVariables': { 'primaryColor': '#FFCC00', 'primaryBorderColor': '#333', 'lineColor': '#333', 'textColor': '#333' }}}%%
usecaseDiagram
actor User
actor Admin

%% User Use Cases
User --> (Login/Logout)
User --> (Register)
User --> (View Slots)
User --> (Book Slot)
User --> (Cancel Booking)

%% Admin Use Cases
Admin --> (Add Slot)
Admin --> (Remove Slot)
Admin --> (Modify Slot)
Admin --> (View Slots)

%% Overlapping Use Cases
(View Slots) <.. (Modify Slot) : "<<extends>>"
(View Slots) <.. (Add Slot) : "<<extends>>"
(View Slots) <.. (Remove Slot) : "<<extends>>"
