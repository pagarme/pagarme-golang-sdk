
# Get Checkout Payment Response

Resposta das configurações de pagamento do checkout

## Structure

`GetCheckoutPaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `types.Optional[string]` | Optional | - |
| `amount` | `types.Optional[int]` | Optional | Valor em centavos |
| `defaultPaymentMethod` | `types.Optional[string]` | Optional | Meio de pagamento padrão no checkout |
| `successUrl` | `types.Optional[string]` | Optional | Url de redirecionamento de sucesso após o checkou |
| `paymentUrl` | `types.Optional[string]` | Optional | Url para pagamento usando o checkout |
| `gatewayAffiliationId` | `types.Optional[string]` | Optional | Código da afiliação onde o pagamento será processado no gateway |
| `acceptedPaymentMethods` | `types.Optional[[]string]` | Optional | Meios de pagamento aceitos no checkout |
| `status` | `types.Optional[string]` | Optional | Status do checkout |
| `skipCheckoutSuccessPage` | `types.Optional[bool]` | Optional | Pular tela de sucesso pós-pagamento? |
| `createdAt` | `types.Optional[time.Time]` | Optional | Data de criação |
| `updatedAt` | `types.Optional[time.Time]` | Optional | Data de atualização |
| `canceledAt` | `types.Optional[time.Time]` | Optional | Data de cancelamento |
| `customerEditable` | `types.Optional[bool]` | Optional | Torna o objeto customer editável |
| `customer` | [`types.Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | Dados do comprador |
| `billingaddress` | [`types.Optional[models.GetAddressResponse]`](../../doc/models/get-address-response.md) | Optional | Dados do endereço de cobrança |
| `creditCard` | [`types.Optional[models.GetCheckoutCreditCardPaymentResponse]`](../../doc/models/get-checkout-credit-card-payment-response.md) | Optional | Configurações de cartão de crédito |
| `boleto` | [`types.Optional[models.GetCheckoutBoletoPaymentResponse]`](../../doc/models/get-checkout-boleto-payment-response.md) | Optional | Configurações de boleto |
| `billingAddressEditable` | `types.Optional[bool]` | Optional | Indica se o billing address poderá ser editado |
| `shipping` | [`types.Optional[models.GetShippingResponse]`](../../doc/models/get-shipping-response.md) | Optional | Configurações  de entrega |
| `shippable` | `types.Optional[bool]` | Optional | Indica se possui entrega |
| `closedAt` | `types.Optional[time.Time]` | Optional | Data de fechamento |
| `expiresAt` | `types.Optional[time.Time]` | Optional | Data de expiração |
| `currency` | `types.Optional[string]` | Optional | Moeda |
| `debitCard` | [`types.Optional[models.GetCheckoutDebitCardPaymentResponse]`](../../doc/models/get-checkout-debit-card-payment-response.md) | Optional | Configurações de cartão de débito |
| `bankTransfer` | [`types.Optional[models.GetCheckoutBankTransferPaymentResponse]`](../../doc/models/get-checkout-bank-transfer-payment-response.md) | Optional | Bank transfer payment response |
| `acceptedBrands` | `types.Optional[[]string]` | Optional | Accepted Brands |
| `pix` | [`types.Optional[models.GetCheckoutPixPaymentResponse]`](../../doc/models/get-checkout-pix-payment-response.md) | Optional | Pix payment response |

## Example (as JSON)

```json
{
  "id": "id0",
  "amount": 46,
  "default_payment_method": "default_payment_method0",
  "success_url": "success_url2",
  "payment_url": "payment_url6"
}
```

