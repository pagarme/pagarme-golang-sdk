
# Paging Response

Object used for returning lists of objects with pagination

## Structure

`PagingResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `total` | `types.Optional[int]` | Optional | Total number of pages |
| `previous` | `types.Optional[string]` | Optional | Previous page |
| `next` | `types.Optional[string]` | Optional | Next page |

## Example (as JSON)

```json
{
  "total": 10,
  "previous": "previous8",
  "next": "next2"
}
```

