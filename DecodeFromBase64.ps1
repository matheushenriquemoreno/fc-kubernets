param (
    [string]$base64
)

# Converter a string Base64 de volta em bytes
$bytes = [Convert]::FromBase64String($base64)

# Converter os bytes em uma string
$decodedString = [System.Text.Encoding]::UTF8.GetString($bytes)

# Exibir o resultado
$decodedString