# booking-site

Practise project - Contains placeholders.

Can be used in real-life scenario, but will need a bit of refactoring
to remove placeholders and add some actual working logic for different places, mostly in front-end. This will require update for back-end as well, which mostly is caused by missing functions due to
the front-end being incomplete.

# Works without registration.

To controll a booking, I used token system instead of registration to ease the process of booking.
Once the payment is successful, a token can be generated and sent on a provided email.
This token is used for refunds and booking. Can be refunded 5-10 days before in-date, depending on the duration booked.
Can add a review after out-date. Once token is used for any of the purposes, it disables and can't be used for anymore.
Refunding causes token to go inactive and free up the booked space.
