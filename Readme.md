# This is a credit card validator which uses Luhn Algorithm
Author : [@Prabesh Bista]
- This algorithm is entirely based on Go Programming Language

Here is my Step By Step Process to Implement the algorithm: 
- Implement the Luhn algorithm
- Create an HTTP server
- Configure the server to respond to GET requests having a JSON payload
- Accept valid JSON requests and proceed to step 5, whilst rejecting invalid requests using an HTTP 400 status code
- Extract the credit card number from the JSON payload
- Run the Luhn algorithm on the number
- Wrap the result into a JSON response payload
- Return the payload back to the user through the HTTP server







    