# MEMO
## Firebase Set up
1. Create project
2. Create realtime database
3. Configure the Realtime Database rules as follows:
```
{
  "rules": {
    ".read": "auth != null && now < 1747234800000",
    ".write": "auth != null && now < 1747234800000"
  }
}
```
4. In the [Authentication] section, enable the email/password provider.
5. Create a user with email/password authentication.
6. Verify that the API key is properly set.

If the following API for obtaining the ID token succeeds, the setup is probably complete.
```
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "email": "{{ mail address}}",
    "password": "{{ password }}",
    "returnSecureToken": true
  }' \
  "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key={{API KEY}}"
```

## ID token cache
Saved in the format below.
```
{
  "idToken": "{{ ID token }}",
  "refreshToken": "{{ refresh token }}",
  "expiresAt": {{ Unix time }}
}
```