#!/bin/bash

# Set the API key
API_KEY="apitest"
INVALID_API_KEY="invalid-key"
API_PORT=9090

# Function to make API requests
make_request() {
    local method=$1
    local endpoint=$2
    local data=$3
    local api_key=$4

    echo "Testing $endpoint..."
    if [ -z "$data" ]; then
        curl -X "$method" "http://localhost:$API_PORT$endpoint" \
            -H "api_key: ${api_key:-$API_KEY}" \
            -H "Content-Type: application/json"
    else
        curl -X "$method" "http://localhost:$API_PORT$endpoint" \
            -H "api_key: ${api_key:-$API_KEY}" \
            -H "Content-Type: application/json" \
            -d "$data"
    fi
    echo -e "\n"
}

# Test List Products
echo "=== Testing List Products ==="
make_request "GET" "/product"

# Test Get Product by ID
echo "=== Testing Get Product by ID ==="
make_request "GET" "/product/1"
make_request "GET" "/product/999" # Test non-existent product

# Test List Orders
echo "=== Testing List Orders ==="
make_request "GET" "/orders"

# Test Get Order by ID
echo "=== Testing Get Order by ID ==="
make_request "GET" "/order/69d3b889-a685-447a-bf29-0f5488c29b52"
make_request "GET" "/order/invalid-id" # Test invalid order ID
make_request "GET" "/order/00000000-0000-0000-0000-000000000000" # Test non-existent order

# Test Create Order
echo "=== Testing Create Order ==="
# Valid order
make_request "POST" "/order" '{
    "items": [
        {
            "productId": "1",
            "quantity": 2
        }
    ]
}'

# Order with invalid product ID
make_request "POST" "/order" '{
    "items": [
        {
            "productId": "999",
            "quantity": 2
        }
    ]
}'

# Order with invalid quantity
make_request "POST" "/order" '{
    "items": [
        {
            "productId": "1",
            "quantity": 0
        }
    ]
}'

# Order with empty items
make_request "POST" "/order" '{
    "items": []
}'

# Test Validate Promo Code
echo "=== Testing Validate Promo Code ==="
make_request "GET" "/validate-promo/HAPPYHRS"
make_request "GET" "/validate-promo/SUPER100"
make_request "GET" "/validate-promo/INVALID" # Test non-existent promo code

# Test Authentication
echo "=== Testing Authentication ==="
make_request "GET" "/product" "" "$INVALID_API_KEY"
make_request "GET" "/orders" "" "$INVALID_API_KEY"
make_request "POST" "/order" '{"items": [{"productId": "1", "quantity": 2}]}' "$INVALID_API_KEY"

# Test Invalid Methods
echo "=== Testing Invalid Methods ==="
make_request "PUT" "/product/1"
make_request "DELETE" "/order/1"
make_request "PATCH" "/validate-promo/HAPPYHRS" 