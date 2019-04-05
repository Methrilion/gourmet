# gourmet

## Сreate Currency

    curl -X POST \
      http://localhost:9092/currency \
      -H 'Content-Type: application/json' \
      -d '{
        "name": "BALUTA",
        "code": "TOP"
    }'

## Сreate Rate Of Exchange

    curl -X POST \
      http://localhost:9092/rates_of_exchange \
      -H 'Content-Type: application/json' \
      -d '{
        "from_id": 1,
        "to_id": 2,
        "price": 49.99
    }'

