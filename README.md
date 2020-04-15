 ``` 
    ___
   /. `\_._._._.
  / _/_ _ _ _ _|\
  |/ / \____/__/\\
    ///   |/|/   \_
```
# Pangolin
An operator for most common database-backups to s3

## Deployment
There will be a helm-cart in the future but for now we use a manifest:
```bash
git clone git@github.com:Netzlink/Pangolin.git
cd Pangolin
```
Now we will deploy the operator:
```bash
kubectl -n pangolin create -f pangolin-operator/deploy/crds/pangolin.netzlink.com_backupjobs_crd.yaml
kubectl -n pangolin create -f pangolin-operator/deploy/
```
Wait for it to be running:
```bash
kubectl -n pangolin get pods -w
```
## Usage
We give you two ways to configure backup-jobs:
- CustomRessource
- CLI (in making)
- Dashboard (future)
### CustomRessource Example
This is a fully configured example of our CRD
Change your schedule, set your DBMS to true (every other to false or remove it) and fill out the other information.
For the Secrets you have to create Kubernetes-secrets beforehand.
```yaml
apiVersion: pangolin.netzlink.com/v1alpha1
kind: BackupJob
metadata:
  name: example-backupjob
  namespace: example
spec:
  schedule: "5 4 * * *"
  type:
    mssql: false
    mariadb: false
    mysql: false
    mongodb: false
    postgres: true
    custom:
      enabled: false
      image: arangodb:2.8
  databaseConfig:
    endpoint: example-postgres # maybe port
    database: example-database
    user: root
    passwordSecret: root-example-database-secret
  s3Config:
    endpoint: s3.eu-west-1.amazonaws.com
    bucket:   db-backups
    secret:   aws-s3-secret
```
The Databaseconfig secret should look like this:
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: root-example-database-secret
  namespace: example
type: Opaque
data:
  password: MWYyZDFlMmU2N2Rm
```
The S3config secret like this:
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: aws-s3-secret
  namespace: example
type: Opaque
data:
  access-key: YWRtaW4=
  secret-key: MWYyZDFlMmU2N2Rm
```
### CLI
## Installation
*__optional__ Install golang if not already*
*Set the PATH*
```bash
export PATH="$PATH:$GOPATH/bin"
```
Install pangolin-cli
```bash
go install github.com/Netzlink/Pangolin/cli
```
Verify the installation
```bash
pangolin
```
You schould be greeted by a nice pangolin!
```text
	     ___
	    /. '\_._._._.
   	   / _/_ _ _ _ _|\
   	  |/  /\____/__/\\
	     ///   |/|/  \_ 

Pangolin holds your back!
Version:	v1alpha1
```
## Get your BackupJobs
```bash
pangolin get
```
if there are backup-jobs the output should roughly resemble this:
```text
Number	Name		Schedule	Type	Database	Bucket
0	nli-ipam	0 */12 * * *	MariaDB	mariadb-0	pangolin-backups
```
## Create a Backup
We are working on this bit :D

# Contribution
Although the license isnt specified yet you are welcome to contribute via _pulls_ and _issues_
