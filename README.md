# vpn-portal

The portal handles generating OpenVPN configs with embeded client certificates.  This allows you to:
- Generates certificates.
- Hand out expiring configs.
- Push routes to clients.
- Open firewall rules on server when connecting.

![alt text][logo]

[logo]: assets/example.png "Example Portal Image"

Requires 
- A reverse proxy that handles SSO.
- An OpenVPN Server configured to pick-up rules.

Examples:
- [Example Config](configs/conf.yaml)
- [Helm Chart](deployment/helm)


