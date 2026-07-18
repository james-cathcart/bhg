# Credential Harvester
One of the staples of social engineering is the credential-harvesting attack. This type of attack captures users’ login information to specific websites by getting them to enter their credentials in a cloned version of the original site. The attack is useful against organizations that expose a single-factor authentication interface to the internet. Once you have a user’s credentials, you can use them to access their account on the actual site. This often leads to an initial breach of the organization’s perimeter network.

## CURL
**REST Call**
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "Alice", "role": "admin"}' http://localhost:8080/login
```

**Form Submission**
```bash
curl -d "username=alice&password=secretpassword" http://localhost:8080/login

```

**Authorization Header**
```bash
curl -H "Authorization: Bearer YOUR_TOKEN_HERE" -d "status=active" http://localhost:8080/login
```
