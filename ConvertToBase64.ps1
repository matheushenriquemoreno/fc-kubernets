param (
    [string]$value
)

# Converter a string em bytes
$bytes = [System.Text.Encoding]::UTF8.GetBytes($value)

# Converter os bytes em Base64
$base64String = [Convert]::ToBase64String($bytes)

# Exibir o resultado
$base64String