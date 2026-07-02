# validation

Shared validation module built on top of [go-playground/validator](https://github.com/go-playground/validator). Provides struct validation, single-value validation, and Thai error messages out of the box.

## Installation

```go
import "github.com/Siravitt/validation"
```

No setup required — `init()` runs automatically on import.

## Usage

### Validate a struct

```go
type CreateUserReq struct {
    Name  string `json:"name"  validate:"required,max=100"`
    Email string `json:"email" validate:"required,email"`
    Phone string `json:"phone" validate:"required,thai_mobile"`
}

err := validation.Validate(req)
if err != nil {
    // err is ValidationErrors — marshals to JSON array automatically
    return c.JSON(400, err)
}
```

Error response:
```json
[
  {"field": "name",  "message": "name ไม่ได้ระบุค่า"},
  {"field": "email", "message": "email รูปแบบอีเมลไม่ถูกต้อง"},
  {"field": "phone", "message": "phone เบอร์โทรไม่ถูกต้อง"}
]
```

### Validate a single value

```go
// validate path/query param without a struct
err := validation.ValidateVar(id, "required", "uuid4")

// with a key name in the error message
err := validation.ValidateWithKey("user_id", id, "required", "uuid4")
// → "'user_id' failed on 'uuid4' tag"
```

### Type-assert errors for field-level access

```go
if errs, ok := err.(validation.ValidationErrors); ok {
    for _, e := range errs {
        fmt.Println(e.Field)   // "email"
        fmt.Println(e.Message) // "email รูปแบบอีเมลไม่ถูกต้อง"
    }
}
```

## Built-in Thai messages

| Tag        | Message                              |
|------------|--------------------------------------|
| `required` | `{field} ไม่ได้ระบุค่า`              |
| `email`    | `{field} รูปแบบอีเมลไม่ถูกต้อง`     |
| `min`      | `{field} ต้องมีความยาวอย่างน้อย {n}` |
| `max`      | `{field} ต้องมีความยาวไม่เกิน {n}`   |
| `len`      | `{field} ต้องมีความยาว {n}`          |
| `numeric`  | `{field} ต้องเป็นตัวเลขเท่านั้น`     |
| `oneof`    | `{field} ต้องเป็นค่าใดค่าหนึ่งใน {n}`|
| `url`      | `{field} รูปแบบ URL ไม่ถูกต้อง`      |
| `uuid4`    | `{field} รูปแบบ UUID ไม่ถูกต้อง`     |
| `gt`       | `{field} ต้องมีค่ามากกว่า {n}`        |
| `gte`      | `{field} ต้องมีค่ามากกว่าหรือเท่ากับ {n}` |
| `lt`       | `{field} ต้องมีค่าน้อยกว่า {n}`       |
| `lte`      | `{field} ต้องมีค่าน้อยกว่าหรือเท่ากับ {n}` |

## Custom validators

| Tag           | Description               | Example                              |
|---------------|---------------------------|--------------------------------------|
| `thai_mobile` | Thai mobile number        | `0812345678`, `0912345678`           |
| `no_script`   | Reject `<script>` tags    | blocks `<script>`, `<SCRIPT>`        |
| `no_emoji`    | Reject emoji characters   | blocks `😀`, `🔥`                   |

## Override error messages

Override any tag message at app startup (e.g., `main.go` or `bootstrap.go`):

```go
// override a single tag
validation.RegisterTranslation(validation.TagRequired, "{0} จำเป็นต้องกรอก")
validation.RegisterTranslation(validation.TagEmail,    "{0} ต้องเป็น email บริษัทเท่านั้น")
```

Placeholders: `{0}` = field name, `{1}` = param value (e.g., `5` for `min=5`)

## Add custom validators

Register at app startup before serving:

```go
validation.RegisterValidation(
    "glo_email",
    func(fl validator.FieldLevel) bool {
        return strings.HasSuffix(fl.Field().String(), "@glo.or.th")
    },
    "{0} ต้องเป็น email @glo.or.th เท่านั้น",
)
```

Then use in struct tags:

```go
type Req struct {
    Email string `json:"email" validate:"required,glo_email"`
}
```

## Tag constants

Use constants instead of raw strings to avoid typos:

```go
// built-in
validation.TagRequired  // "required"
validation.TagEmail     // "email"
validation.TagMin       // "min"
// ...

// custom
validation.TagThaiMobile // "thai_mobile"
validation.TagNoScript   // "no_script"
validation.TagNoEmoji    // "no_emoji"
```
