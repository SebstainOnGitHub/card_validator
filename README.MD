# A basic api that uses the luhn algorithm to check if the PAN of a card is valid, and or determines the provider

Accepts a json get request with a structure of 

# {
    number: "example-number"
# }
, and returns either

# {
    message: "valid card"
# }
,
# {
    error: "invalid card number"
# }
or
# {
    provider: "(Provider/industry)"
# }

Provider also checks whether the card is valid or not.

To find out the whether a card is valid, request from 
# {
    (localhost):8080/check
# }

To find out the provider, request from 
# {
    (localhost):8080/check_suggest
# }