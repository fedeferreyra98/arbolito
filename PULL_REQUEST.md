# Description

This PR introduces a new endpoint `/convert` that allows users to convert a specified amount between Argentine Pesos (ARS) and US Dollars (USD). The conversion is flexible and supports applying different historical dollar rate types (e.g., `blue`, `oficial`, `mep`, etc.) based on user input.

The functionality correctly uses the `Sell` price when converting from ARS to USD and the `Buy` price when converting from USD to ARS, accurately reflecting real-world exchange mechanics.

## Type of change

- [x] New feature (non-breaking change which adds functionality)
- [x] This change requires a documentation update

# How Has This Been Tested?

The conversion logic within the handler has been rigorously tested using Go's testing framework. Tests were added to mock the service layer to ensure it behaves deterministically for both ARS-to-USD and USD-to-ARS conversions. Edge cases involving bad requests were also tested.

- [x] **TestRateHandler_Convert/Valid_ARS_to_USD**: Verifies calculation logic using the `Sell` price.
- [x] **TestRateHandler_Convert/Valid_USD_to_ARS**: Verifies calculation logic using the `Buy` price.
- [x] **TestRateHandler_Convert/Missing_params**: Checks appropriate Bad Request response when required query params are missing.
- [x] **TestRateHandler_Convert/Invalid_amount**: Checks appropriate Bad Request response for negative or invalid amounts.
- [x] **TestRateHandler_Convert/Invalid_currency_pair**: Ensures only ARS to USD or USD to ARS requests are processed.

# Checklist:

- [x] My code follows the style guidelines of this project
- [x] I have performed a self-review of my own code
- [x] I have commented my code, particularly in hard-to-understand areas
- [x] I have made corresponding changes to the documentation
- [x] My changes generate no new warnings
- [x] I have added tests that prove my fix is effective or that my feature works
- [x] New and existing unit tests pass locally with my changes
- [x] Any dependent changes have been merged and published in downstream modules
