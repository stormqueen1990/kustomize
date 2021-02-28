// Code generated by "mdtogo"; DO NOT EDIT.
package tutorials

var ConfigurationBasicsShort = `### Synopsis`
var ConfigurationBasicsLong = `
` + "`" + `kustomize cfg` + "`" + ` provides tools for working with local configuration directories.

  First fetch a bundle of configuration to your local file system from the
  Kubernetes examples repository.

	git clone https://github.com/kubernetes/examples/
	cd examples/

### ` + "`" + `tree` + "`" + ` -- view Resources and directory structure

  ` + "`" + `tree` + "`" + ` can be used to summarize the collection of Resources in a directory:

	$ kustomize cfg tree mysql-wordpress-pd/
	mysql-wordpress-pd
	├── [gce-volumes.yaml]  v1.PersistentVolume wordpress-pv-1
	├── [gce-volumes.yaml]  v1.PersistentVolume wordpress-pv-2
	├── [local-volumes.yaml]  v1.PersistentVolume local-pv-1
	├── [local-volumes.yaml]  v1.PersistentVolume local-pv-2
	├── [mysql-deployment.yaml]  v1.PersistentVolumeClaim mysql-pv-claim
	├── [mysql-deployment.yaml]  apps/v1.Deployment wordpress-mysql
	├── [mysql-deployment.yaml]  v1.Service wordpress-mysql
	├── [wordpress-deployment.yaml]  apps/v1.Deployment wordpress
	├── [wordpress-deployment.yaml]  v1.Service wordpress
	└── [wordpress-deployment.yaml]  v1.PersistentVolumeClaim wp-pv-claim

  ` + "`" + `tree` + "`" + ` may be provided flags to print the Resource field values.  ` + "`" + `tree` + "`" + ` has a number of built-in
  supported fields, and may also print arbitrary values using the ` + "`" + `--field` + "`" + ` flag to specify a field
  path.

    $  kustomize cfg tree mysql-wordpress-pd/ --name --image --replicas --ports
    mysql-wordpress-pd
    ├── [gce-volumes.yaml]  PersistentVolume wordpress-pv-1
    ├── [gce-volumes.yaml]  PersistentVolume wordpress-pv-2
    ├── [local-volumes.yaml]  PersistentVolume local-pv-1
    ├── [local-volumes.yaml]  PersistentVolume local-pv-2
    ├── [mysql-deployment.yaml]  PersistentVolumeClaim mysql-pv-claim
    ├── [mysql-deployment.yaml]  Deployment wordpress-mysql
    │   └── spec.template.spec.containers
    │       └── 0
    │           ├── name: mysql
    │           ├── image: mysql:5.6
    │           └── ports: [{name: mysql, containerPort: 3306}]
    ├── [mysql-deployment.yaml]  Service wordpress-mysql
    │   └── spec.ports: [{port: 3306}]
    ├── [wordpress-deployment.yaml]  Deployment wordpress
    │   └── spec.template.spec.containers
    │       └── 0
    │           ├── name: wordpress
    │           ├── image: wordpress:4.8-apache
    │           └── ports: [{name: wordpress, containerPort: 80}]
    ├── [wordpress-deployment.yaml]  Service wordpress
    │   └── spec.ports: [{port: 80}]
    └── [wordpress-deployment.yaml]  PersistentVolumeClaim wp-pv-claim

  ` + "`" + `tree` + "`" + ` can also be used with ` + "`" + `kubectl get` + "`" + ` to print cluster Resources using OwnersReferences
  to build the tree structure.

    kubectl apply -R -f cockroachdb/
    kubectl get all -o yaml | kustomize cfg tree --graph-structure owners --name --image --replicas
    .
    ├── [Resource]  Deployment wp/wordpress
    │   ├── spec.replicas: 1
    │   ├── spec.template.spec.containers
    │   │   └── 0
    │   │       ├── name: wordpress
    │   │       └── image: wordpress:4.8-apache
    │   └── [Resource]  ReplicaSet wp/wordpress-76b5d9f5c8
    │       ├── spec.replicas: 1
    │       ├── spec.template.spec.containers
    │       │   └── 0
    │       │       ├── name: wordpress
    │       │       └── image: wordpress:4.8-apache
    │       └── [Resource]  Pod wp/wordpress-76b5d9f5c8-g656w
    │           └── spec.containers
    │               └── 0
    │                   ├── name: wordpress
    │                   └── image: wordpress:4.8-apache
    ├── [Resource]  Service wp/wordpress
    ...

### ` + "`" + `cat` + "`" + ` -- view the full collection of Resources

	$ kustomize cfg cat mysql-wordpress-pd/
	apiVersion: v1
	kind: PersistentVolume
	metadata:
	  name: wordpress-pv-1
	  annotations:
		config.kubernetes.io/path: gce-volumes.yaml
	spec:
	  accessModes:
	  - ReadWriteOnce
	  capacity:
		storage: 20Gi
	  gcePersistentDisk:
		fsType: ext4
		pdName: wordpress-1
	---
	apiVersion: v1
	...

  ` + "`" + `cat` + "`" + ` prints the raw package Resources.  This may be used to pipe them to other tools
  such as ` + "`" + `kubectl apply -f -` + "`" + `.

## ` + "`" + `fmt` + "`" + ` -- format the Resources for a directory (like go fmt for Kubernetes Resources)

  ` + "`" + `fmt` + "`" + ` formats the Resource Configuration by applying a consistent style, including
  ordering of fields and indentation.

	$ kustomize cfg fmt mysql-wordpress-pd/

  Run ` + "`" + `git diff` + "`" + ` and see the changes that have been applied.

### ` + "`" + `grep` + "`" + ` -- search for Resources by field values

  ` + "`" + `grep` + "`" + ` prints Resources matching some field value.  The Resources are annotated with their
  file source so they can be piped to other commands without losing this information.

	$ kustomize cfg grep "metadata.name=wordpress" wordpress/
	apiVersion: v1
	kind: Service
	metadata:
	  name: wordpress
	  labels:
		app: wordpress
	  annotations:
		config.kubernetes.io/path: wordpress-deployment.yaml
	spec:
	  ports:
	  - port: 80
	  selector:
		app: wordpress
		tier: frontend
	  type: LoadBalancer
	---
	...

  - list elements may be indexed by a field value using list[field=value]
  - '.' as part of a key or value may be escaped as '\.'

	$ kustomize cfg grep "spec.status.spec.containers[name=nginx].image=mysql:5\.6" wordpress/
	apiVersion: apps/v1 # for k8s versions before 1.9.0 use apps/v1beta2  and before 1.8.0 use extensions/v1beta1
	kind: Deployment
	metadata:
	  name: wordpress-mysql
	  labels:
		app: wordpress
	spec:
	  selector:
		matchLabels:
		  app: wordpress
		  tier: mysql
	  template:
		metadata:
		  labels:
			app: wordpress
			tier: mysql
	...

  ` + "`" + `grep` + "`" + ` may be used with kubectl to search for Resources in a cluster matching a value.

    kubectl get all -o yaml | kustomize cfg grep "spec.replicas>0" | kustomize cfg tree --replicas
    .
    └──
        ├── [.]  Deployment wp/wordpress
        │   └── spec.replicas: 1
        ├── [.]  ReplicaSet wp/wordpress-76b5d9f5c8
        │   └── spec.replicas: 1
        ├── [.]  Deployment wp/wordpress-mysql
        │   └── spec.replicas: 1
        └── [.]  ReplicaSet wp/wordpress-mysql-f9447f458
            └── spec.replicas: 1

### Error handling

  If there is an error parsing the Resource configuration, kustomize will print an error with the file.

    $ kustomize cfg grep "spec.template.spec.containers[name=\.*].resources.limits.cpu>1.0" ./staging/ | kustomize cfg tree --name --resources
    Error: staging/persistent-volume-provisioning/quobyte/quobyte-admin-secret.yaml: [0]: yaml: unmarshal errors:
      line 13: mapping key "type" already defined at line 9

  Here the ` + "`" + `staging/persistent-volume-provisioning/quobyte/quobyte-admin-secret.yaml` + "`" + ` has a malformed
  Resource.  Remove the malformed Resources:

    rm staging/persistent-volume-provisioning/quobyte/quobyte-admin-secret.yaml
    rm staging/storage/vitess/etcd-service-template.yaml

  When developing -- to get a stack trace for where an error was encountered,
  use the ` + "`" + `--stack-trace` + "`" + ` flag:

    $ kustomize cfg grep "spec.template.spec.containers[name=\.*].resources.limits.cpu>1.0" ./staging/ --stack-trace
    go/src/sigs.k8s.io/kustomize/kyaml/yaml/types.go:260 (0x4d35c86)
            (*RNode).GetMeta: return m, errors.Wrap(err)
    go/src/sigs.k8s.io/kustomize/kyaml/kio/byteio_reader.go:130 (0x4d3e099)
            (*ByteReader).Read: meta, err := node.GetMeta()
    ...


### Combine ` + "`" + `grep` + "`" + ` and ` + "`" + `tree` + "`" + `

  ` + "`" + `grep` + "`" + ` and ` + "`" + `tree` + "`" + ` may be combined to perform queries against configuration.

  Query for ` + "`" + `replicas` + "`" + `:

    $ kustomize cfg grep "spec.replicas>5" ./ | kustomize cfg tree --replicas
      .
      ├── staging/sysdig-cloud
      │   └── [sysdig-rc.yaml]  ReplicationController sysdig-agent
      │       └── spec.replicas: 100
      └── staging/volumes/vsphere
          └── [simple-statefulset.yaml]  StatefulSet web
              └── spec.replicas: 14

  Query for ` + "`" + `resource.limits` + "`" + `

	$ kustomize cfg grep "spec.template.spec.containers[name=\.*].resources.limits.memory>0" ./ | kustomize cfg tree --resources
	.
    ├── cassandra
    │   └── [cassandra-statefulset.yaml]  StatefulSet cassandra
    │       └── spec.template.spec.containers
    │           └── 0
    │               └── resources: {limits: {cpu: "500m", memory: 1Gi}, requests: {cpu: "500m", memory: 1Gi}}
    ├── staging/selenium
    │   ├── [selenium-hub-deployment.yaml]  Deployment selenium-hub
    │   │   └── spec.template.spec.containers
    │   │       └── 0
    │   │           └── resources: {limits: {memory: 1000Mi, cpu: ".5"}}
    │   ├── [selenium-node-chrome-deployment.yaml]  Deployment selenium-node-chrome
    │   │   └── spec.template.spec.containers
    │   │       └── 0
    │   │           └── resources: {limits: {memory: 1000Mi, cpu: ".5"}}
    │   └── [selenium-node-firefox-deployment.yaml]  Deployment selenium-node-firefox
    │       └── spec.template.spec.containers
    │           └── 0
    │               └── resources: {limits: {memory: 1000Mi, cpu: ".5"}}
    ...

### Inverting ` + "`" + `grep` + "`" + `

  The ` + "`" + `grep` + "`" + ` results may be inverted with the ` + "`" + `-v` + "`" + ` flag and used to find Resources that don't
  match a query.

  Find Resources that have an image specified, but the image doesn't have a tag:

    $ kustomize cfg grep "spec.template.spec.containers[name=\.*].name=\.*" ./ |  kustomize cfg grep "spec.template.spec.containers[name=\.*].image=\.*:\.*" -v | kustomize cfg tree --image --name
    .
    ├── staging/newrelic
    │   ├── [newrelic-daemonset.yaml]  DaemonSet newrelic-agent
    │   │   └── spec.template.spec.containers
    │   │       └── 0
    │   │           ├── name: newrelic
    │   │           └── image: newrelic/nrsysmond
    │   └── staging/newrelic-infrastructure
    │       └── [newrelic-infra-daemonset.yaml]  DaemonSet newrelic-infra-agent
    │           └── spec.template.spec.containers
    │               └── 0
    │                   ├── name: newrelic
    │                   └── image: newrelic/infrastructure
    ├── staging/nodesjs-mongodb
    │   ├── [mongo-controller.yaml]  ReplicationController mongo-controller
    │   │   └── spec.template.spec.containers
    │   │       └── 0
    │   │           ├── name: mongo
    │   │           └── image: mongo
    │   └── [web-controller.yaml]  ReplicationController web-controller
    │       └── spec.template.spec.containers
    │           └── 0
    │               ├── name: web
    │               └── image: <YOUR-CONTAINER>
    ...`

var FunctionBasicsShort = `### Synopsis`
var FunctionBasicsLong = `
  ` + "`" + `kustomize config` + "`" + ` enables encapsulating function for manipulating Resource
  configuration inside containers, which are run using ` + "`" + `run` + "`" + `.

  First fetch the kustomize repository, which contains a collection of example
  functions

	git clone https://github.com/kubernetes-sigs/kustomize
	cd kustomize/functions/examples/

### Templating -- CockroachDB

  This section demonstrates how to leverage templating based solutions from
  ` + "`" + `kustomize config` + "`" + `.  The templating function is implemented as a ` + "`" + `bash` + "`" + ` script
  using a ` + "`" + `heredoc` + "`" + `.

  #### 1: Generate the Resources

  ` + "`" + `cd` + "`" + ` into the ` + "`" + `kustomize/functions/examples/template-heredoc-cockroachdb/` + "`" + `
  directory, and invoke ` + "`" + `run` + "`" + ` on the ` + "`" + `local-resource/` + "`" + ` directory.

    cd template-heredoc-cockroachdb/

    # view the Resources
    kustomize cfg tree local-resource/ --name --image --replicas

    # run the function
    kustomize fn run local-resource/

    # view the generated Resources
    kustomize cfg tree local-resource/ --name --image --replicas

  ` + "`" + `run` + "`" + ` generated the directory ` + "`" + ` local-resource/config` + "`" + ` containing the generated
  Resources.

  #### 2. Modify the Generated Resources

  - modify the generated Resources by adding an annotation, sidecar container, etc.
  - modify the ` + "`" + `local-resource/example-use.yaml` + "`" + ` by changing the replicas

  re-run ` + "`" + `run` + "`" + `.  this will apply the updated replicas to the generated Resources,
  but keep the fields that you manually added to the generated Resource configuration.

    # run the function
    kustomize fn run local-resource/

  ` + "`" + `run` + "`" + ` facilitates a non-destructive *smart templating* approach that allows templating
  to be composed with manual modifications directly to the template output, as well as
  composition with other functions which may appy validation or injection of values.

  #### 3. Function Implementation

  the function implementation is located under the ` + "`" + `image/` + "`" + ` directory as a ` + "`" + `Dockerfile` + "`" + `
  and a ` + "`" + `bash` + "`" + ` script.

### Templating -- Nginx

  The steps in this section are identical to the CockroachDB templating example,
  but the function implementation is very different, and implemented as a ` + "`" + `go` + "`" + `
  program rather than a ` + "`" + `bash` + "`" + ` script.

  #### 1: Generate the Resources

  ` + "`" + `cd` + "`" + ` into the ` + "`" + `kustomize/functions/examples/template-go-nginx/` + "`" + `
  directory, and invoke ` + "`" + `run` + "`" + ` on the ` + "`" + `local-resource/` + "`" + ` directory.

    cd template-go-nginx/

    # view the Resources
    kustomize cfg tree local-resource/ --name --image --replicas

    # run the function
    kustomize fn run local-resource/

    # view the generated Resources
    kustomize cfg tree local-resource/ --name --image --replicas

  ` + "`" + `run` + "`" + ` generated the directory ` + "`" + ` local-resource/config` + "`" + ` containing the generated
  Resources.  this time it put the configuration in a single file rather than multiple
  files.  The mapping of Resources to files is controlled by the function itself through
  annotations on the generated Resources.

  #### 2. Modify the Generated Resources

  - modify the generated Resources by adding an annotation, sidecar container, etc.
  - modify the ` + "`" + `local-resource/example-use.yaml` + "`" + ` by changing the replicas

  re-run ` + "`" + `run` + "`" + `.  this will apply the updated replicas to the generated Resources,
  but keep the fields that you manually added to the generated Resource configuration.

    # run the function
    kustomize fn run local-resource/

  Just like in the preceding section, the function is implemented using a non-destructive
  approach which merges the generated Resources into previously generated instances.

  #### 3. Function Implementation

  the function implementation is located under the ` + "`" + `image/` + "`" + ` directory as a ` + "`" + `Dockerfile` + "`" + `
  and a ` + "`" + `go` + "`" + ` program.

### Validation -- resource reservations

  This section uses ` + "`" + `run` + "`" + ` to perform validation rather than generate Resources.

  #### 1: Run the Validator

  ` + "`" + `cd` + "`" + ` into the ` + "`" + `kustomize/functions/examples/validator-resource-requests` + "`" + `
  directory, and invoke ` + "`" + `run` + "`" + ` on the ` + "`" + `local-resource/` + "`" + ` directory.

    # run the function
    kustomize fn run local-resource/
    cpu-requests missing for a container in Deployment nginx (example-use.yaml [1])
    Error: exit status 1
    Usage:
    ...

  #### 2: Fix the validation issue

  The command will fail complaining that the nginx Deployment is missing ` + "`" + `cpu-requests` + "`" + `,
  and print the name of the file + Resource index.  Edit the file and uncomment the resources,
  then re-run the functions.

    kustomize fn run local-resource/

  The validation now passes.

### Injection -- resource reservations

  This section uses ` + "`" + `run` + "`" + ` to perform injection of field values based off annotations
  on the Resource.

  #### 1: Run the Injector

  ` + "`" + `cd` + "`" + ` into the ` + "`" + `kustomize/functions/examples/inject-tshirt-sizes` + "`" + `
  directory, and invoke ` + "`" + `run` + "`" + ` on the ` + "`" + `local-resource/` + "`" + ` directory.

    # print the resources
    kustomize cfg tree local-resource --resources --name
    local-resource
    ├── [example-use.yaml]  Validator
    └── [example-use.yaml]  Deployment nginx
        └── spec.template.spec.containers
            └── 0
                └── name: nginx

    # run the functions
    kustomize fn run local-resource/

    # print the new resources
    kustomize cfg tree local-resource --resources --name
    ├── [example-use.yaml]  Validator
    └── [example-use.yaml]  Deployment nginx
        └── spec.template.spec.containers
            └── 0
                ├── name: nginx
                └── resources: {requests: {cpu: 4, memory: 1GiB}}

  #### 2: Change the tshirt-size

  Change the ` + "`" + `tshirt-size` + "`" + ` annotation from ` + "`" + `medium` + "`" + ` to ` + "`" + `small` + "`" + ` and re-run the functions.

    kustomize fn run local-resource/
    kustomize cfg tree local-resource/
    local-resource
    ├── [example-use.yaml]  Validator
    └── [example-use.yaml]  Deployment nginx
        └── spec.template.spec.containers
            └── 0
                ├── name: nginx
                └── resources: {requests: {cpu: 200m, memory: 50MiB}}

  The function has applied the reservations for the new tshirt-size

### Function Composition

Functions may be composed together.  Try putting the Injection (tshirt-size) and
Validation functions together in the same .yaml file (separated by ` + "`" + `---` + "`" + `).  Run
` + "`" + `run` + "`" + ` and observe that the first function in the file is applied to the Resources,
and then the second function in the file is applied.`
