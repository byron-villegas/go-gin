Feature: Products API

  Scenario: Get all products
    Given An endpoint "/products"
    When I send a GET request
    Then the response code should be 200

  Scenario Outline: Get product By SKU
    Given An endpoint "/products/<SKU>"
    When I send a GET request
    Then the response code should be <STATUS>

    Examples:
      | SKU      | STATUS |
      | 15207410 | 200    |
      | 15207411 | 404    |