package templates

const CMSInitScriptTemplate = `
set -eu

echo Console-Config-Root.txt
cat /opt/kikimr/cfg/Console-Config-Root.txt
/opt/kikimr/bin/kikimr admin console execute --domain=root --retry=10 /opt/kikimr/cfg/Console-Config-Root.txt

echo Configure-Root.txt
cat /opt/kikimr/cfg/Configure-Root.txt
/opt/kikimr/bin/kikimr admin console execute --domain=root --retry=10 /opt/kikimr/cfg/Configure-Root.txt
`