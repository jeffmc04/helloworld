package ako

var AkoDeployFile = `apiVersion: v1
kind: Namespace
metadata:
  labels:
    control-plane: controller-manager
  name: aerospike
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.4.1
  creationTimestamp: null
  name: aerospikeclusters.asdb.aerospike.com
spec:
  group: asdb.aerospike.com
  names:
    kind: AerospikeCluster
    listKind: AerospikeClusterList
    plural: aerospikeclusters
    singular: aerospikecluster
  scope: Namespaced
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: AerospikeCluster is the schema for the AerospikeCluster API
        properties:
          apiVersion:
            description: APIVersion defines the versioned schema of this representation of an object.
            type: string
          kind:
            description: Kind is a string value representing the REST resource this object represents.
            type: string
          metadata:
            type: object
          spec:
            description: AerospikeClusterSpec defines the desired state of AerospikeCluster
            properties:
              aerospikeAccessControl:
                description: AerospikeAccessControl has the Aerospike roles and users definitions.
                properties:
                  adminPolicy:
                    description: 'AerospikeClientAdminPolicy specify the aerospike client admin policy for access '
                    properties:
                      timeout:
                        description: Timeout for admin client policy in milliseconds.
                        type: integer
                    required:
                    - timeout
                    type: object
                  roles:
                    description: Roles is the set of roles to allow on the Aerospike cluster.
                    items:
                      description: AerospikeRoleSpec specifies an Aerospike database role and its associated privil
                      properties:
                        name:
                          description: Name of this role.
                          type: string
                        privileges:
                          description: Privileges granted to this role.
                          items:
                            type: string
                          type: array
                          x-kubernetes-list-type: set
                        readQuota:
                          description: ReadQuota specifies permitted rate of read records for current role (the value i
                          format: int32
                          type: integer
                        whitelist:
                          description: Whitelist of host address allowed for this role.
                          items:
                            type: string
                          type: array
                          x-kubernetes-list-type: set
                        writeQuota:
                          description: WriteQuota specifies permitted rate of write records for current role (the value
                          format: int32
                          type: integer
                      required:
                      - name
                      - privileges
                      type: object
                    type: array
                    x-kubernetes-list-map-keys:
                    - name
                    x-kubernetes-list-type: map
                  users:
                    description: Users is the set of users to allow on the Aerospike cluster.
                    items:
                      description: 'AerospikeUserSpec specifies an Aerospike database user, the secret name for the '
                      properties:
                        name:
                          description: Name is the user's username.
                          type: string
                        roles:
                          description: Roles is the list of roles granted to the user.
                          items:
                            type: string
                          type: array
                          x-kubernetes-list-type: set
                        secretName:
                          description: SecretName has secret info created by user.
                          type: string
                      required:
                      - name
                      - roles
                      - secretName
                      type: object
                    type: array
                    x-kubernetes-list-map-keys:
                    - name
                    x-kubernetes-list-type: map
                required:
                - users
                type: object
              aerospikeConfig:
                description: AerospikeConfig sets config in aerospike.conf file.
                type: object
                x-kubernetes-preserve-unknown-fields: true
              aerospikeNetworkPolicy:
                description: AerospikeNetworkPolicy specifies how clients and tools access the Aerospike clus
                properties:
                  access:
                    description: AccessType is the type of network address to use for Aerospike access address.
                    enum:
                    - pod
                    - hostInternal
                    - hostExternal
                    type: string
                  alternateAccess:
                    description: AlternateAccessType is the type of network address to use for Aerospike alternat
                    enum:
                    - pod
                    - hostInternal
                    - hostExternal
                    type: string
                  tlsAccess:
                    description: TLSAccessType is the type of network address to use for Aerospike TLS access add
                    enum:
                    - pod
                    - hostInternal
                    - hostExternal
                    type: string
                  tlsAlternateAccess:
                    description: TLSAlternateAccessType is the type of network address to use for Aerospike TLS a
                    enum:
                    - pod
                    - hostInternal
                    - hostExternal
                    type: string
                type: object
              image:
                description: Aerospike server image
                type: string
              operatorClientCert:
                description: Certificates to connect to Aerospike.
                properties:
                  certPathInOperator:
                    description: AerospikeCertPathInOperatorSource contain configuration for certificates used by
                    properties:
                      caCertsPath:
                        type: string
                      clientCertPath:
                        type: string
                      clientKeyPath:
                        type: string
                    type: object
                  secretCertSource:
                    properties:
                      caCertsFilename:
                        type: string
                      clientCertFilename:
                        type: string
                      clientKeyFilename:
                        type: string
                      secretName:
                        type: string
                      secretNamespace:
                        type: string
                    required:
                    - secretName
                    type: object
                  tlsClientName:
                    description: If specified, this name will be added to tls-authenticate-client list by the ope
                    type: string
                type: object
              podSpec:
                description: Additional configuration for create Aerospike pods.
                properties:
                  aerospikeContainer:
                    description: 'AerospikeContainerSpec contains settings for aerospike-server container created '
                    properties:
                      resources:
                        description: Define resources requests and limits for Aerospike Server Container.
                        properties:
                          limits:
                            additionalProperties:
                              anyOf:
                              - type: integer
                              - type: string
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            description: Limits describes the maximum amount of compute resources allowed.
                            type: object
                          requests:
                            additionalProperties:
                              anyOf:
                              - type: integer
                              - type: string
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            description: Requests describes the minimum amount of compute resources required.
                            type: object
                        type: object
                      securityContext:
                        description: SecurityContext that will be added to aerospike-server container created by oper
                        properties:
                          allowPrivilegeEscalation:
                            description: AllowPrivilegeEscalation controls whether a process can gain more privileges tha
                            type: boolean
                          capabilities:
                            description: The capabilities to add/drop when running containers.
                            properties:
                              add:
                                description: Added capabilities
                                items:
                                  description: Capability represent POSIX capabilities type
                                  type: string
                                type: array
                              drop:
                                description: Removed capabilities
                                items:
                                  description: Capability represent POSIX capabilities type
                                  type: string
                                type: array
                            type: object
                          privileged:
                            description: Run container in privileged mode.
                            type: boolean
                          procMount:
                            description: procMount denotes the type of proc mount to use for the containers.
                            type: string
                          readOnlyRootFilesystem:
                            description: Whether this container has a read-only root filesystem. Default is false.
                            type: boolean
                          runAsGroup:
                            description: The GID to run the entrypoint of the container process.
                            format: int64
                            type: integer
                          runAsNonRoot:
                            description: Indicates that the container must run as a non-root user.
                            type: boolean
                          runAsUser:
                            description: The UID to run the entrypoint of the container process.
                            format: int64
                            type: integer
                          seLinuxOptions:
                            description: The SELinux context to be applied to the container.
                            properties:
                              level:
                                description: Level is SELinux level label that applies to the container.
                                type: string
                              role:
                                description: Role is a SELinux role label that applies to the container.
                                type: string
                              type:
                                description: Type is a SELinux type label that applies to the container.
                                type: string
                              user:
                                description: User is a SELinux user label that applies to the container.
                                type: string
                            type: object
                          seccompProfile:
                            description: The seccomp options to use by this container.
                            properties:
                              localhostProfile:
                                description: localhostProfile indicates a profile defined in a file on the node should be use
                                type: string
                              type:
                                description: type indicates which kind of seccomp profile will be applied.
                                type: string
                            required:
                            - type
                            type: object
                          windowsOptions:
                            description: The Windows specific settings applied to all containers.
                            properties:
                              gmsaCredentialSpec:
                                description: GMSACredentialSpec is where the GMSA admission webhook (https://github.
                                type: string
                              gmsaCredentialSpecName:
                                description: GMSACredentialSpecName is the name of the GMSA credential spec to use.
                                type: string
                              runAsUserName:
                                description: The UserName in Windows to run the entrypoint of the container process.
                                type: string
                            type: object
                        type: object
                    type: object
                  affinity:
                    description: Affinity rules for pod placement.
                    properties:
                      nodeAffinity:
                        description: Describes node affinity scheduling rules for the pod.
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods to nodes that satisfy the affinity ex
                            items:
                              description: An empty preferred scheduling term matches all objects with implicit weight 0 (i
                              properties:
                                preference:
                                  description: A node selector term, associated with the corresponding weight.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements by node's labels.
                                      items:
                                        description: A node selector requirement is a selector that contains values, a key, and an op
                                        properties:
                                          key:
                                            description: The label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: An array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements by node's fields.
                                      items:
                                        description: A node selector requirement is a selector that contains values, a key, and an op
                                        properties:
                                          key:
                                            description: The label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: An array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                weight:
                                  description: Weight associated with matching the corresponding nodeSelectorTerm, in the range
                                  format: int32
                                  type: integer
                              required:
                              - preference
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by this field are not met at scheduling t
                            properties:
                              nodeSelectorTerms:
                                description: Required. A list of node selector terms. The terms are ORed.
                                items:
                                  description: A null or empty node selector term matches no objects.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements by node's labels.
                                      items:
                                        description: A node selector requirement is a selector that contains values, a key, and an op
                                        properties:
                                          key:
                                            description: The label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: An array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements by node's fields.
                                      items:
                                        description: A node selector requirement is a selector that contains values, a key, and an op
                                        properties:
                                          key:
                                            description: The label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: An array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                type: array
                            required:
                            - nodeSelectorTerms
                            type: object
                        type: object
                      podAffinity:
                        description: Describes pod affinity scheduling rules (e.g.
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods to nodes that satisfy the affinity ex
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm fields are added per-n
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources, in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list of label selector requirements.
                                          items:
                                            description: A label selector requirement is a selector that contains values, a key, and an o
                                            properties:
                                              key:
                                                description: key is the label key that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a key's relationship to a set of values.
                                                type: string
                                              values:
                                                description: values is an array of string values.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value} pairs.
                                          type: object
                                      type: object
                                    namespaceSelector:
                                      description: A label query over the set of namespaces that the term applies to.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list of label selector requirements.
                                          items:
                                            description: A label selector requirement is a selector that contains values, a key, and an o
                                            properties:
                                              key:
                                                description: key is the label key that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a key's relationship to a set of values.
                                                type: string
                                              values:
                                                description: values is an array of string values.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value} pairs.
                                          type: object
                                      type: object
                                    namespaces:
                                      description: namespaces specifies a static list of namespace names that the term applies to.
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: 'weight associated with matching the corresponding podAffinityTerm, in the range '
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by this field are not met at scheduling t
                            items:
                              description: Defines a set of pods (namely those matching the labelSelector relative to the g
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources, in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label selector requirements.
                                      items:
                                        description: A label selector requirement is a selector that contains values, a key, and an o
                                        properties:
                                          key:
                                            description: key is the label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: values is an array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value} pairs.
                                      type: object
                                  type: object
                                namespaceSelector:
                                  description: A label query over the set of namespaces that the term applies to.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label selector requirements.
                                      items:
                                        description: A label selector requirement is a selector that contains values, a key, and an o
                                        properties:
                                          key:
                                            description: key is the label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: values is an array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value} pairs.
                                      type: object
                                  type: object
                                namespaces:
                                  description: namespaces specifies a static list of namespace names that the term applies to.
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                      podAntiAffinity:
                        description: Describes pod anti-affinity scheduling rules (e.g.
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods to nodes that satisfy the anti-affini
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm fields are added per-n
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources, in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list of label selector requirements.
                                          items:
                                            description: A label selector requirement is a selector that contains values, a key, and an o
                                            properties:
                                              key:
                                                description: key is the label key that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a key's relationship to a set of values.
                                                type: string
                                              values:
                                                description: values is an array of string values.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value} pairs.
                                          type: object
                                      type: object
                                    namespaceSelector:
                                      description: A label query over the set of namespaces that the term applies to.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list of label selector requirements.
                                          items:
                                            description: A label selector requirement is a selector that contains values, a key, and an o
                                            properties:
                                              key:
                                                description: key is the label key that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a key's relationship to a set of values.
                                                type: string
                                              values:
                                                description: values is an array of string values.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value} pairs.
                                          type: object
                                      type: object
                                    namespaces:
                                      description: namespaces specifies a static list of namespace names that the term applies to.
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: 'weight associated with matching the corresponding podAffinityTerm, in the range '
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the anti-affinity requirements specified by this field are not met at schedul
                            items:
                              description: Defines a set of pods (namely those matching the labelSelector relative to the g
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources, in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label selector requirements.
                                      items:
                                        description: A label selector requirement is a selector that contains values, a key, and an o
                                        properties:
                                          key:
                                            description: key is the label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: values is an array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value} pairs.
                                      type: object
                                  type: object
                                namespaceSelector:
                                  description: A label query over the set of namespaces that the term applies to.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label selector requirements.
                                      items:
                                        description: A label selector requirement is a selector that contains values, a key, and an o
                                        properties:
                                          key:
                                            description: key is the label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: values is an array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value} pairs.
                                      type: object
                                  type: object
                                namespaces:
                                  description: namespaces specifies a static list of namespace names that the term applies to.
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                    type: object
                  dnsPolicy:
                    description: DnsPolicy same as https://kubernetes.
                    type: string
                  effectiveDNSPolicy:
                    description: Effective value of the DNSPolicy
                    type: string
                  hostNetwork:
                    description: HostNetwork enables host networking for the pod.
                    type: boolean
                  initContainers:
                    description: InitContainers to add to the pods.
                    items:
                      description: A single application container that you want to run within a pod.
                      properties:
                        args:
                          description: Arguments to the entrypoint.
                          items:
                            type: string
                          type: array
                        command:
                          description: Entrypoint array. Not executed within a shell.
                          items:
                            type: string
                          type: array
                        env:
                          description: List of environment variables to set in the container. Cannot be updated.
                          items:
                            description: EnvVar represents an environment variable present in a Container.
                            properties:
                              name:
                                description: Name of the environment variable. Must be a C_IDENTIFIER.
                                type: string
                              value:
                                description: Variable references $(VAR_NAME) are expanded using the previous defined environm
                                type: string
                              valueFrom:
                                description: Source for the environment variable's value.
                                properties:
                                  configMapKeyRef:
                                    description: Selects a key of a ConfigMap.
                                    properties:
                                      key:
                                        description: The key to select.
                                        type: string
                                      name:
                                        description: 'Name of the referent. More info: https://kubernetes.'
                                        type: string
                                      optional:
                                        description: Specify whether the ConfigMap or its key must be defined
                                        type: boolean
                                    required:
                                    - key
                                    type: object
                                  fieldRef:
                                    description: 'Selects a field of the pod: supports metadata.name, metadata.'
                                    properties:
                                      apiVersion:
                                        description: Version of the schema the FieldPath is written in terms of, defaults to "v1".
                                        type: string
                                      fieldPath:
                                        description: Path of the field to select in the specified API version.
                                        type: string
                                    required:
                                    - fieldPath
                                    type: object
                                  resourceFieldRef:
                                    description: 'Selects a resource of the container: only resources limits and requests (limits.'
                                    properties:
                                      containerName:
                                        description: 'Container name: required for volumes, optional for env vars'
                                        type: string
                                      divisor:
                                        anyOf:
                                        - type: integer
                                        - type: string
                                        description: Specifies the output format of the exposed resources, defaults to "1"
                                        pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                        x-kubernetes-int-or-string: true
                                      resource:
                                        description: 'Required: resource to select'
                                        type: string
                                    required:
                                    - resource
                                    type: object
                                  secretKeyRef:
                                    description: Selects a key of a secret in the pod's namespace
                                    properties:
                                      key:
                                        description: The key of the secret to select from.  Must be a valid secret key.
                                        type: string
                                      name:
                                        description: 'Name of the referent. More info: https://kubernetes.'
                                        type: string
                                      optional:
                                        description: Specify whether the Secret or its key must be defined
                                        type: boolean
                                    required:
                                    - key
                                    type: object
                                type: object
                            required:
                            - name
                            type: object
                          type: array
                        envFrom:
                          description: List of sources to populate environment variables in the container.
                          items:
                            description: EnvFromSource represents the source of a set of ConfigMaps
                            properties:
                              configMapRef:
                                description: The ConfigMap to select from
                                properties:
                                  name:
                                    description: 'Name of the referent. More info: https://kubernetes.'
                                    type: string
                                  optional:
                                    description: Specify whether the ConfigMap must be defined
                                    type: boolean
                                type: object
                              prefix:
                                description: An optional identifier to prepend to each key in the ConfigMap.
                                type: string
                              secretRef:
                                description: The Secret to select from
                                properties:
                                  name:
                                    description: 'Name of the referent. More info: https://kubernetes.'
                                    type: string
                                  optional:
                                    description: Specify whether the Secret must be defined
                                    type: boolean
                                type: object
                            type: object
                          type: array
                        image:
                          description: 'Docker image name. More info: https://kubernetes.'
                          type: string
                        imagePullPolicy:
                          description: Image pull policy. One of Always, Never, IfNotPresent.
                          type: string
                        lifecycle:
                          description: Actions that the management system should take in response to container lifecycl
                          properties:
                            postStart:
                              description: PostStart is called immediately after a container is created.
                              properties:
                                exec:
                                  description: One and only one of the following should be specified.
                                  properties:
                                    command:
                                      description: Command is the command line to execute inside the container, the working directo
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                httpGet:
                                  description: HTTPGet specifies the http request to perform.
                                  properties:
                                    host:
                                      description: Host name to connect to, defaults to the pod IP.
                                      type: string
                                    httpHeaders:
                                      description: Custom headers to set in the request. HTTP allows repeated headers.
                                      items:
                                        description: HTTPHeader describes a custom header to be used in HTTP probes
                                        properties:
                                          name:
                                            description: The header field name
                                            type: string
                                          value:
                                            description: The header field value
                                            type: string
                                        required:
                                        - name
                                        - value
                                        type: object
                                      type: array
                                    path:
                                      description: Path to access on the HTTP server.
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Name or number of the port to access on the container.
                                      x-kubernetes-int-or-string: true
                                    scheme:
                                      description: Scheme to use for connecting to the host. Defaults to HTTP.
                                      type: string
                                  required:
                                  - port
                                  type: object
                                tcpSocket:
                                  description: TCPSocket specifies an action involving a TCP port.
                                  properties:
                                    host:
                                      description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Number or name of the port to access on the container.
                                      x-kubernetes-int-or-string: true
                                  required:
                                  - port
                                  type: object
                              type: object
                            preStop:
                              description: PreStop is called immediately before a container is terminated due to an API req
                              properties:
                                exec:
                                  description: One and only one of the following should be specified.
                                  properties:
                                    command:
                                      description: Command is the command line to execute inside the container, the working directo
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                httpGet:
                                  description: HTTPGet specifies the http request to perform.
                                  properties:
                                    host:
                                      description: Host name to connect to, defaults to the pod IP.
                                      type: string
                                    httpHeaders:
                                      description: Custom headers to set in the request. HTTP allows repeated headers.
                                      items:
                                        description: HTTPHeader describes a custom header to be used in HTTP probes
                                        properties:
                                          name:
                                            description: The header field name
                                            type: string
                                          value:
                                            description: The header field value
                                            type: string
                                        required:
                                        - name
                                        - value
                                        type: object
                                      type: array
                                    path:
                                      description: Path to access on the HTTP server.
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Name or number of the port to access on the container.
                                      x-kubernetes-int-or-string: true
                                    scheme:
                                      description: Scheme to use for connecting to the host. Defaults to HTTP.
                                      type: string
                                  required:
                                  - port
                                  type: object
                                tcpSocket:
                                  description: TCPSocket specifies an action involving a TCP port.
                                  properties:
                                    host:
                                      description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Number or name of the port to access on the container.
                                      x-kubernetes-int-or-string: true
                                  required:
                                  - port
                                  type: object
                              type: object
                          type: object
                        livenessProbe:
                          description: Periodic probe of container liveness.
                          properties:
                            exec:
                              description: One and only one of the following should be specified.
                              properties:
                                command:
                                  description: Command is the command line to execute inside the container, the working directo
                                  items:
                                    type: string
                                  type: array
                              type: object
                            failureThreshold:
                              description: 'Minimum consecutive failures for the probe to be considered failed after having '
                              format: int32
                              type: integer
                            httpGet:
                              description: HTTPGet specifies the http request to perform.
                              properties:
                                host:
                                  description: Host name to connect to, defaults to the pod IP.
                                  type: string
                                httpHeaders:
                                  description: Custom headers to set in the request. HTTP allows repeated headers.
                                  items:
                                    description: HTTPHeader describes a custom header to be used in HTTP probes
                                    properties:
                                      name:
                                        description: The header field name
                                        type: string
                                      value:
                                        description: The header field value
                                        type: string
                                    required:
                                    - name
                                    - value
                                    type: object
                                  type: array
                                path:
                                  description: Path to access on the HTTP server.
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Name or number of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                                scheme:
                                  description: Scheme to use for connecting to the host. Defaults to HTTP.
                                  type: string
                              required:
                              - port
                              type: object
                            initialDelaySeconds:
                              description: Number of seconds after the container has started before liveness probes are ini
                              format: int32
                              type: integer
                            periodSeconds:
                              description: How often (in seconds) to perform the probe. Default to 10 seconds.
                              format: int32
                              type: integer
                            successThreshold:
                              description: Minimum consecutive successes for the probe to be considered successful after ha
                              format: int32
                              type: integer
                            tcpSocket:
                              description: TCPSocket specifies an action involving a TCP port.
                              properties:
                                host:
                                  description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Number or name of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                              required:
                              - port
                              type: object
                            terminationGracePeriodSeconds:
                              description: Optional duration in seconds the pod needs to terminate gracefully upon probe fa
                              format: int64
                              type: integer
                            timeoutSeconds:
                              description: Number of seconds after which the probe times out. Defaults to 1 second.
                              format: int32
                              type: integer
                          type: object
                        name:
                          description: Name of the container specified as a DNS_LABEL.
                          type: string
                        ports:
                          description: List of ports to expose from the container.
                          items:
                            description: ContainerPort represents a network port in a single container.
                            properties:
                              containerPort:
                                description: Number of port to expose on the pod's IP address.
                                format: int32
                                type: integer
                              hostIP:
                                description: What host IP to bind the external port to.
                                type: string
                              hostPort:
                                description: Number of port to expose on the host.
                                format: int32
                                type: integer
                              name:
                                description: If specified, this must be an IANA_SVC_NAME and unique within the pod.
                                type: string
                              protocol:
                                default: TCP
                                description: Protocol for port. Must be UDP, TCP, or SCTP. Defaults to "TCP".
                                type: string
                            required:
                            - containerPort
                            type: object
                          type: array
                          x-kubernetes-list-map-keys:
                          - containerPort
                          - protocol
                          x-kubernetes-list-type: map
                        readinessProbe:
                          description: Periodic probe of container service readiness.
                          properties:
                            exec:
                              description: One and only one of the following should be specified.
                              properties:
                                command:
                                  description: Command is the command line to execute inside the container, the working directo
                                  items:
                                    type: string
                                  type: array
                              type: object
                            failureThreshold:
                              description: 'Minimum consecutive failures for the probe to be considered failed after having '
                              format: int32
                              type: integer
                            httpGet:
                              description: HTTPGet specifies the http request to perform.
                              properties:
                                host:
                                  description: Host name to connect to, defaults to the pod IP.
                                  type: string
                                httpHeaders:
                                  description: Custom headers to set in the request. HTTP allows repeated headers.
                                  items:
                                    description: HTTPHeader describes a custom header to be used in HTTP probes
                                    properties:
                                      name:
                                        description: The header field name
                                        type: string
                                      value:
                                        description: The header field value
                                        type: string
                                    required:
                                    - name
                                    - value
                                    type: object
                                  type: array
                                path:
                                  description: Path to access on the HTTP server.
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Name or number of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                                scheme:
                                  description: Scheme to use for connecting to the host. Defaults to HTTP.
                                  type: string
                              required:
                              - port
                              type: object
                            initialDelaySeconds:
                              description: Number of seconds after the container has started before liveness probes are ini
                              format: int32
                              type: integer
                            periodSeconds:
                              description: How often (in seconds) to perform the probe. Default to 10 seconds.
                              format: int32
                              type: integer
                            successThreshold:
                              description: Minimum consecutive successes for the probe to be considered successful after ha
                              format: int32
                              type: integer
                            tcpSocket:
                              description: TCPSocket specifies an action involving a TCP port.
                              properties:
                                host:
                                  description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Number or name of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                              required:
                              - port
                              type: object
                            terminationGracePeriodSeconds:
                              description: Optional duration in seconds the pod needs to terminate gracefully upon probe fa
                              format: int64
                              type: integer
                            timeoutSeconds:
                              description: Number of seconds after which the probe times out. Defaults to 1 second.
                              format: int32
                              type: integer
                          type: object
                        resources:
                          description: Compute Resources required by this container. Cannot be updated.
                          properties:
                            limits:
                              additionalProperties:
                                anyOf:
                                - type: integer
                                - type: string
                                pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                x-kubernetes-int-or-string: true
                              description: Limits describes the maximum amount of compute resources allowed.
                              type: object
                            requests:
                              additionalProperties:
                                anyOf:
                                - type: integer
                                - type: string
                                pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                x-kubernetes-int-or-string: true
                              description: Requests describes the minimum amount of compute resources required.
                              type: object
                          type: object
                        securityContext:
                          description: 'Security options the pod should run with. More info: https://kubernetes.'
                          properties:
                            allowPrivilegeEscalation:
                              description: AllowPrivilegeEscalation controls whether a process can gain more privileges tha
                              type: boolean
                            capabilities:
                              description: The capabilities to add/drop when running containers.
                              properties:
                                add:
                                  description: Added capabilities
                                  items:
                                    description: Capability represent POSIX capabilities type
                                    type: string
                                  type: array
                                drop:
                                  description: Removed capabilities
                                  items:
                                    description: Capability represent POSIX capabilities type
                                    type: string
                                  type: array
                              type: object
                            privileged:
                              description: Run container in privileged mode.
                              type: boolean
                            procMount:
                              description: procMount denotes the type of proc mount to use for the containers.
                              type: string
                            readOnlyRootFilesystem:
                              description: Whether this container has a read-only root filesystem. Default is false.
                              type: boolean
                            runAsGroup:
                              description: The GID to run the entrypoint of the container process.
                              format: int64
                              type: integer
                            runAsNonRoot:
                              description: Indicates that the container must run as a non-root user.
                              type: boolean
                            runAsUser:
                              description: The UID to run the entrypoint of the container process.
                              format: int64
                              type: integer
                            seLinuxOptions:
                              description: The SELinux context to be applied to the container.
                              properties:
                                level:
                                  description: Level is SELinux level label that applies to the container.
                                  type: string
                                role:
                                  description: Role is a SELinux role label that applies to the container.
                                  type: string
                                type:
                                  description: Type is a SELinux type label that applies to the container.
                                  type: string
                                user:
                                  description: User is a SELinux user label that applies to the container.
                                  type: string
                              type: object
                            seccompProfile:
                              description: The seccomp options to use by this container.
                              properties:
                                localhostProfile:
                                  description: localhostProfile indicates a profile defined in a file on the node should be use
                                  type: string
                                type:
                                  description: type indicates which kind of seccomp profile will be applied.
                                  type: string
                              required:
                              - type
                              type: object
                            windowsOptions:
                              description: The Windows specific settings applied to all containers.
                              properties:
                                gmsaCredentialSpec:
                                  description: GMSACredentialSpec is where the GMSA admission webhook (https://github.
                                  type: string
                                gmsaCredentialSpecName:
                                  description: GMSACredentialSpecName is the name of the GMSA credential spec to use.
                                  type: string
                                runAsUserName:
                                  description: The UserName in Windows to run the entrypoint of the container process.
                                  type: string
                              type: object
                          type: object
                        startupProbe:
                          description: StartupProbe indicates that the Pod has successfully initialized.
                          properties:
                            exec:
                              description: One and only one of the following should be specified.
                              properties:
                                command:
                                  description: Command is the command line to execute inside the container, the working directo
                                  items:
                                    type: string
                                  type: array
                              type: object
                            failureThreshold:
                              description: 'Minimum consecutive failures for the probe to be considered failed after having '
                              format: int32
                              type: integer
                            httpGet:
                              description: HTTPGet specifies the http request to perform.
                              properties:
                                host:
                                  description: Host name to connect to, defaults to the pod IP.
                                  type: string
                                httpHeaders:
                                  description: Custom headers to set in the request. HTTP allows repeated headers.
                                  items:
                                    description: HTTPHeader describes a custom header to be used in HTTP probes
                                    properties:
                                      name:
                                        description: The header field name
                                        type: string
                                      value:
                                        description: The header field value
                                        type: string
                                    required:
                                    - name
                                    - value
                                    type: object
                                  type: array
                                path:
                                  description: Path to access on the HTTP server.
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Name or number of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                                scheme:
                                  description: Scheme to use for connecting to the host. Defaults to HTTP.
                                  type: string
                              required:
                              - port
                              type: object
                            initialDelaySeconds:
                              description: Number of seconds after the container has started before liveness probes are ini
                              format: int32
                              type: integer
                            periodSeconds:
                              description: How often (in seconds) to perform the probe. Default to 10 seconds.
                              format: int32
                              type: integer
                            successThreshold:
                              description: Minimum consecutive successes for the probe to be considered successful after ha
                              format: int32
                              type: integer
                            tcpSocket:
                              description: TCPSocket specifies an action involving a TCP port.
                              properties:
                                host:
                                  description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Number or name of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                              required:
                              - port
                              type: object
                            terminationGracePeriodSeconds:
                              description: Optional duration in seconds the pod needs to terminate gracefully upon probe fa
                              format: int64
                              type: integer
                            timeoutSeconds:
                              description: Number of seconds after which the probe times out. Defaults to 1 second.
                              format: int32
                              type: integer
                          type: object
                        stdin:
                          description: Whether this container should allocate a buffer for stdin in the container runti
                          type: boolean
                        stdinOnce:
                          description: Whether the container runtime should close the stdin channel after it has been o
                          type: boolean
                        terminationMessagePath:
                          description: 'Optional: Path at which the file to which the container''s termination message wi'
                          type: string
                        terminationMessagePolicy:
                          description: Indicate how the termination message should be populated.
                          type: string
                        tty:
                          description: Whether this container should allocate a TTY for itself, also requires 'stdin' t
                          type: boolean
                        volumeDevices:
                          description: volumeDevices is the list of block devices to be used by the container.
                          items:
                            description: volumeDevice describes a mapping of a raw block device within a container.
                            properties:
                              devicePath:
                                description: devicePath is the path inside of the container that the device will be mapped to
                                type: string
                              name:
                                description: name must match the name of a persistentVolumeClaim in the pod
                                type: string
                            required:
                            - devicePath
                            - name
                            type: object
                          type: array
                        volumeMounts:
                          description: Pod volumes to mount into the container's filesystem. Cannot be updated.
                          items:
                            description: VolumeMount describes a mounting of a Volume within a container.
                            properties:
                              mountPath:
                                description: Path within the container at which the volume should be mounted.
                                type: string
                              mountPropagation:
                                description: mountPropagation determines how mounts are propagated from the host to container
                                type: string
                              name:
                                description: This must match the Name of a Volume.
                                type: string
                              readOnly:
                                description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                type: boolean
                              subPath:
                                description: Path within the volume from which the container's volume should be mounted.
                                type: string
                              subPathExpr:
                                description: Expanded path within the volume from which the container's volume should be moun
                                type: string
                            required:
                            - mountPath
                            - name
                            type: object
                          type: array
                        workingDir:
                          description: Container's working directory.
                          type: string
                      required:
                      - name
                      type: object
                    type: array
                  multiPodPerHost:
                    description: If set true then multiple pods can be created per Kubernetes Node.
                    type: boolean
                  nodeSelector:
                    additionalProperties:
                      type: string
                    description: NodeSelector constraints for this pod.
                    type: object
                  sidecars:
                    description: Sidecars to add to pods.
                    items:
                      description: A single application container that you want to run within a pod.
                      properties:
                        args:
                          description: Arguments to the entrypoint.
                          items:
                            type: string
                          type: array
                        command:
                          description: Entrypoint array. Not executed within a shell.
                          items:
                            type: string
                          type: array
                        env:
                          description: List of environment variables to set in the container. Cannot be updated.
                          items:
                            description: EnvVar represents an environment variable present in a Container.
                            properties:
                              name:
                                description: Name of the environment variable. Must be a C_IDENTIFIER.
                                type: string
                              value:
                                description: Variable references $(VAR_NAME) are expanded using the previous defined environm
                                type: string
                              valueFrom:
                                description: Source for the environment variable's value.
                                properties:
                                  configMapKeyRef:
                                    description: Selects a key of a ConfigMap.
                                    properties:
                                      key:
                                        description: The key to select.
                                        type: string
                                      name:
                                        description: 'Name of the referent. More info: https://kubernetes.'
                                        type: string
                                      optional:
                                        description: Specify whether the ConfigMap or its key must be defined
                                        type: boolean
                                    required:
                                    - key
                                    type: object
                                  fieldRef:
                                    description: 'Selects a field of the pod: supports metadata.name, metadata.'
                                    properties:
                                      apiVersion:
                                        description: Version of the schema the FieldPath is written in terms of, defaults to "v1".
                                        type: string
                                      fieldPath:
                                        description: Path of the field to select in the specified API version.
                                        type: string
                                    required:
                                    - fieldPath
                                    type: object
                                  resourceFieldRef:
                                    description: 'Selects a resource of the container: only resources limits and requests (limits.'
                                    properties:
                                      containerName:
                                        description: 'Container name: required for volumes, optional for env vars'
                                        type: string
                                      divisor:
                                        anyOf:
                                        - type: integer
                                        - type: string
                                        description: Specifies the output format of the exposed resources, defaults to "1"
                                        pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                        x-kubernetes-int-or-string: true
                                      resource:
                                        description: 'Required: resource to select'
                                        type: string
                                    required:
                                    - resource
                                    type: object
                                  secretKeyRef:
                                    description: Selects a key of a secret in the pod's namespace
                                    properties:
                                      key:
                                        description: The key of the secret to select from.  Must be a valid secret key.
                                        type: string
                                      name:
                                        description: 'Name of the referent. More info: https://kubernetes.'
                                        type: string
                                      optional:
                                        description: Specify whether the Secret or its key must be defined
                                        type: boolean
                                    required:
                                    - key
                                    type: object
                                type: object
                            required:
                            - name
                            type: object
                          type: array
                        envFrom:
                          description: List of sources to populate environment variables in the container.
                          items:
                            description: EnvFromSource represents the source of a set of ConfigMaps
                            properties:
                              configMapRef:
                                description: The ConfigMap to select from
                                properties:
                                  name:
                                    description: 'Name of the referent. More info: https://kubernetes.'
                                    type: string
                                  optional:
                                    description: Specify whether the ConfigMap must be defined
                                    type: boolean
                                type: object
                              prefix:
                                description: An optional identifier to prepend to each key in the ConfigMap.
                                type: string
                              secretRef:
                                description: The Secret to select from
                                properties:
                                  name:
                                    description: 'Name of the referent. More info: https://kubernetes.'
                                    type: string
                                  optional:
                                    description: Specify whether the Secret must be defined
                                    type: boolean
                                type: object
                            type: object
                          type: array
                        image:
                          description: 'Docker image name. More info: https://kubernetes.'
                          type: string
                        imagePullPolicy:
                          description: Image pull policy. One of Always, Never, IfNotPresent.
                          type: string
                        lifecycle:
                          description: Actions that the management system should take in response to container lifecycl
                          properties:
                            postStart:
                              description: PostStart is called immediately after a container is created.
                              properties:
                                exec:
                                  description: One and only one of the following should be specified.
                                  properties:
                                    command:
                                      description: Command is the command line to execute inside the container, the working directo
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                httpGet:
                                  description: HTTPGet specifies the http request to perform.
                                  properties:
                                    host:
                                      description: Host name to connect to, defaults to the pod IP.
                                      type: string
                                    httpHeaders:
                                      description: Custom headers to set in the request. HTTP allows repeated headers.
                                      items:
                                        description: HTTPHeader describes a custom header to be used in HTTP probes
                                        properties:
                                          name:
                                            description: The header field name
                                            type: string
                                          value:
                                            description: The header field value
                                            type: string
                                        required:
                                        - name
                                        - value
                                        type: object
                                      type: array
                                    path:
                                      description: Path to access on the HTTP server.
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Name or number of the port to access on the container.
                                      x-kubernetes-int-or-string: true
                                    scheme:
                                      description: Scheme to use for connecting to the host. Defaults to HTTP.
                                      type: string
                                  required:
                                  - port
                                  type: object
                                tcpSocket:
                                  description: TCPSocket specifies an action involving a TCP port.
                                  properties:
                                    host:
                                      description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Number or name of the port to access on the container.
                                      x-kubernetes-int-or-string: true
                                  required:
                                  - port
                                  type: object
                              type: object
                            preStop:
                              description: PreStop is called immediately before a container is terminated due to an API req
                              properties:
                                exec:
                                  description: One and only one of the following should be specified.
                                  properties:
                                    command:
                                      description: Command is the command line to execute inside the container, the working directo
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                httpGet:
                                  description: HTTPGet specifies the http request to perform.
                                  properties:
                                    host:
                                      description: Host name to connect to, defaults to the pod IP.
                                      type: string
                                    httpHeaders:
                                      description: Custom headers to set in the request. HTTP allows repeated headers.
                                      items:
                                        description: HTTPHeader describes a custom header to be used in HTTP probes
                                        properties:
                                          name:
                                            description: The header field name
                                            type: string
                                          value:
                                            description: The header field value
                                            type: string
                                        required:
                                        - name
                                        - value
                                        type: object
                                      type: array
                                    path:
                                      description: Path to access on the HTTP server.
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Name or number of the port to access on the container.
                                      x-kubernetes-int-or-string: true
                                    scheme:
                                      description: Scheme to use for connecting to the host. Defaults to HTTP.
                                      type: string
                                  required:
                                  - port
                                  type: object
                                tcpSocket:
                                  description: TCPSocket specifies an action involving a TCP port.
                                  properties:
                                    host:
                                      description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Number or name of the port to access on the container.
                                      x-kubernetes-int-or-string: true
                                  required:
                                  - port
                                  type: object
                              type: object
                          type: object
                        livenessProbe:
                          description: Periodic probe of container liveness.
                          properties:
                            exec:
                              description: One and only one of the following should be specified.
                              properties:
                                command:
                                  description: Command is the command line to execute inside the container, the working directo
                                  items:
                                    type: string
                                  type: array
                              type: object
                            failureThreshold:
                              description: 'Minimum consecutive failures for the probe to be considered failed after having '
                              format: int32
                              type: integer
                            httpGet:
                              description: HTTPGet specifies the http request to perform.
                              properties:
                                host:
                                  description: Host name to connect to, defaults to the pod IP.
                                  type: string
                                httpHeaders:
                                  description: Custom headers to set in the request. HTTP allows repeated headers.
                                  items:
                                    description: HTTPHeader describes a custom header to be used in HTTP probes
                                    properties:
                                      name:
                                        description: The header field name
                                        type: string
                                      value:
                                        description: The header field value
                                        type: string
                                    required:
                                    - name
                                    - value
                                    type: object
                                  type: array
                                path:
                                  description: Path to access on the HTTP server.
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Name or number of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                                scheme:
                                  description: Scheme to use for connecting to the host. Defaults to HTTP.
                                  type: string
                              required:
                              - port
                              type: object
                            initialDelaySeconds:
                              description: Number of seconds after the container has started before liveness probes are ini
                              format: int32
                              type: integer
                            periodSeconds:
                              description: How often (in seconds) to perform the probe. Default to 10 seconds.
                              format: int32
                              type: integer
                            successThreshold:
                              description: Minimum consecutive successes for the probe to be considered successful after ha
                              format: int32
                              type: integer
                            tcpSocket:
                              description: TCPSocket specifies an action involving a TCP port.
                              properties:
                                host:
                                  description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Number or name of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                              required:
                              - port
                              type: object
                            terminationGracePeriodSeconds:
                              description: Optional duration in seconds the pod needs to terminate gracefully upon probe fa
                              format: int64
                              type: integer
                            timeoutSeconds:
                              description: Number of seconds after which the probe times out. Defaults to 1 second.
                              format: int32
                              type: integer
                          type: object
                        name:
                          description: Name of the container specified as a DNS_LABEL.
                          type: string
                        ports:
                          description: List of ports to expose from the container.
                          items:
                            description: ContainerPort represents a network port in a single container.
                            properties:
                              containerPort:
                                description: Number of port to expose on the pod's IP address.
                                format: int32
                                type: integer
                              hostIP:
                                description: What host IP to bind the external port to.
                                type: string
                              hostPort:
                                description: Number of port to expose on the host.
                                format: int32
                                type: integer
                              name:
                                description: If specified, this must be an IANA_SVC_NAME and unique within the pod.
                                type: string
                              protocol:
                                default: TCP
                                description: Protocol for port. Must be UDP, TCP, or SCTP. Defaults to "TCP".
                                type: string
                            required:
                            - containerPort
                            type: object
                          type: array
                          x-kubernetes-list-map-keys:
                          - containerPort
                          - protocol
                          x-kubernetes-list-type: map
                        readinessProbe:
                          description: Periodic probe of container service readiness.
                          properties:
                            exec:
                              description: One and only one of the following should be specified.
                              properties:
                                command:
                                  description: Command is the command line to execute inside the container, the working directo
                                  items:
                                    type: string
                                  type: array
                              type: object
                            failureThreshold:
                              description: 'Minimum consecutive failures for the probe to be considered failed after having '
                              format: int32
                              type: integer
                            httpGet:
                              description: HTTPGet specifies the http request to perform.
                              properties:
                                host:
                                  description: Host name to connect to, defaults to the pod IP.
                                  type: string
                                httpHeaders:
                                  description: Custom headers to set in the request. HTTP allows repeated headers.
                                  items:
                                    description: HTTPHeader describes a custom header to be used in HTTP probes
                                    properties:
                                      name:
                                        description: The header field name
                                        type: string
                                      value:
                                        description: The header field value
                                        type: string
                                    required:
                                    - name
                                    - value
                                    type: object
                                  type: array
                                path:
                                  description: Path to access on the HTTP server.
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Name or number of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                                scheme:
                                  description: Scheme to use for connecting to the host. Defaults to HTTP.
                                  type: string
                              required:
                              - port
                              type: object
                            initialDelaySeconds:
                              description: Number of seconds after the container has started before liveness probes are ini
                              format: int32
                              type: integer
                            periodSeconds:
                              description: How often (in seconds) to perform the probe. Default to 10 seconds.
                              format: int32
                              type: integer
                            successThreshold:
                              description: Minimum consecutive successes for the probe to be considered successful after ha
                              format: int32
                              type: integer
                            tcpSocket:
                              description: TCPSocket specifies an action involving a TCP port.
                              properties:
                                host:
                                  description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Number or name of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                              required:
                              - port
                              type: object
                            terminationGracePeriodSeconds:
                              description: Optional duration in seconds the pod needs to terminate gracefully upon probe fa
                              format: int64
                              type: integer
                            timeoutSeconds:
                              description: Number of seconds after which the probe times out. Defaults to 1 second.
                              format: int32
                              type: integer
                          type: object
                        resources:
                          description: Compute Resources required by this container. Cannot be updated.
                          properties:
                            limits:
                              additionalProperties:
                                anyOf:
                                - type: integer
                                - type: string
                                pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                x-kubernetes-int-or-string: true
                              description: Limits describes the maximum amount of compute resources allowed.
                              type: object
                            requests:
                              additionalProperties:
                                anyOf:
                                - type: integer
                                - type: string
                                pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                x-kubernetes-int-or-string: true
                              description: Requests describes the minimum amount of compute resources required.
                              type: object
                          type: object
                        securityContext:
                          description: 'Security options the pod should run with. More info: https://kubernetes.'
                          properties:
                            allowPrivilegeEscalation:
                              description: AllowPrivilegeEscalation controls whether a process can gain more privileges tha
                              type: boolean
                            capabilities:
                              description: The capabilities to add/drop when running containers.
                              properties:
                                add:
                                  description: Added capabilities
                                  items:
                                    description: Capability represent POSIX capabilities type
                                    type: string
                                  type: array
                                drop:
                                  description: Removed capabilities
                                  items:
                                    description: Capability represent POSIX capabilities type
                                    type: string
                                  type: array
                              type: object
                            privileged:
                              description: Run container in privileged mode.
                              type: boolean
                            procMount:
                              description: procMount denotes the type of proc mount to use for the containers.
                              type: string
                            readOnlyRootFilesystem:
                              description: Whether this container has a read-only root filesystem. Default is false.
                              type: boolean
                            runAsGroup:
                              description: The GID to run the entrypoint of the container process.
                              format: int64
                              type: integer
                            runAsNonRoot:
                              description: Indicates that the container must run as a non-root user.
                              type: boolean
                            runAsUser:
                              description: The UID to run the entrypoint of the container process.
                              format: int64
                              type: integer
                            seLinuxOptions:
                              description: The SELinux context to be applied to the container.
                              properties:
                                level:
                                  description: Level is SELinux level label that applies to the container.
                                  type: string
                                role:
                                  description: Role is a SELinux role label that applies to the container.
                                  type: string
                                type:
                                  description: Type is a SELinux type label that applies to the container.
                                  type: string
                                user:
                                  description: User is a SELinux user label that applies to the container.
                                  type: string
                              type: object
                            seccompProfile:
                              description: The seccomp options to use by this container.
                              properties:
                                localhostProfile:
                                  description: localhostProfile indicates a profile defined in a file on the node should be use
                                  type: string
                                type:
                                  description: type indicates which kind of seccomp profile will be applied.
                                  type: string
                              required:
                              - type
                              type: object
                            windowsOptions:
                              description: The Windows specific settings applied to all containers.
                              properties:
                                gmsaCredentialSpec:
                                  description: GMSACredentialSpec is where the GMSA admission webhook (https://github.
                                  type: string
                                gmsaCredentialSpecName:
                                  description: GMSACredentialSpecName is the name of the GMSA credential spec to use.
                                  type: string
                                runAsUserName:
                                  description: The UserName in Windows to run the entrypoint of the container process.
                                  type: string
                              type: object
                          type: object
                        startupProbe:
                          description: StartupProbe indicates that the Pod has successfully initialized.
                          properties:
                            exec:
                              description: One and only one of the following should be specified.
                              properties:
                                command:
                                  description: Command is the command line to execute inside the container, the working directo
                                  items:
                                    type: string
                                  type: array
                              type: object
                            failureThreshold:
                              description: 'Minimum consecutive failures for the probe to be considered failed after having '
                              format: int32
                              type: integer
                            httpGet:
                              description: HTTPGet specifies the http request to perform.
                              properties:
                                host:
                                  description: Host name to connect to, defaults to the pod IP.
                                  type: string
                                httpHeaders:
                                  description: Custom headers to set in the request. HTTP allows repeated headers.
                                  items:
                                    description: HTTPHeader describes a custom header to be used in HTTP probes
                                    properties:
                                      name:
                                        description: The header field name
                                        type: string
                                      value:
                                        description: The header field value
                                        type: string
                                    required:
                                    - name
                                    - value
                                    type: object
                                  type: array
                                path:
                                  description: Path to access on the HTTP server.
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Name or number of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                                scheme:
                                  description: Scheme to use for connecting to the host. Defaults to HTTP.
                                  type: string
                              required:
                              - port
                              type: object
                            initialDelaySeconds:
                              description: Number of seconds after the container has started before liveness probes are ini
                              format: int32
                              type: integer
                            periodSeconds:
                              description: How often (in seconds) to perform the probe. Default to 10 seconds.
                              format: int32
                              type: integer
                            successThreshold:
                              description: Minimum consecutive successes for the probe to be considered successful after ha
                              format: int32
                              type: integer
                            tcpSocket:
                              description: TCPSocket specifies an action involving a TCP port.
                              properties:
                                host:
                                  description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Number or name of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                              required:
                              - port
                              type: object
                            terminationGracePeriodSeconds:
                              description: Optional duration in seconds the pod needs to terminate gracefully upon probe fa
                              format: int64
                              type: integer
                            timeoutSeconds:
                              description: Number of seconds after which the probe times out. Defaults to 1 second.
                              format: int32
                              type: integer
                          type: object
                        stdin:
                          description: Whether this container should allocate a buffer for stdin in the container runti
                          type: boolean
                        stdinOnce:
                          description: Whether the container runtime should close the stdin channel after it has been o
                          type: boolean
                        terminationMessagePath:
                          description: 'Optional: Path at which the file to which the container''s termination message wi'
                          type: string
                        terminationMessagePolicy:
                          description: Indicate how the termination message should be populated.
                          type: string
                        tty:
                          description: Whether this container should allocate a TTY for itself, also requires 'stdin' t
                          type: boolean
                        volumeDevices:
                          description: volumeDevices is the list of block devices to be used by the container.
                          items:
                            description: volumeDevice describes a mapping of a raw block device within a container.
                            properties:
                              devicePath:
                                description: devicePath is the path inside of the container that the device will be mapped to
                                type: string
                              name:
                                description: name must match the name of a persistentVolumeClaim in the pod
                                type: string
                            required:
                            - devicePath
                            - name
                            type: object
                          type: array
                        volumeMounts:
                          description: Pod volumes to mount into the container's filesystem. Cannot be updated.
                          items:
                            description: VolumeMount describes a mounting of a Volume within a container.
                            properties:
                              mountPath:
                                description: Path within the container at which the volume should be mounted.
                                type: string
                              mountPropagation:
                                description: mountPropagation determines how mounts are propagated from the host to container
                                type: string
                              name:
                                description: This must match the Name of a Volume.
                                type: string
                              readOnly:
                                description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                type: boolean
                              subPath:
                                description: Path within the volume from which the container's volume should be mounted.
                                type: string
                              subPathExpr:
                                description: Expanded path within the volume from which the container's volume should be moun
                                type: string
                            required:
                            - mountPath
                            - name
                            type: object
                          type: array
                        workingDir:
                          description: Container's working directory.
                          type: string
                      required:
                      - name
                      type: object
                    type: array
                  tolerations:
                    description: Tolerations for this pod.
                    items:
                      description: The pod this Toleration is attached to tolerates any taint that matches the trip
                      properties:
                        effect:
                          description: Effect indicates the taint effect to match. Empty means match all taint effects.
                          type: string
                        key:
                          description: Key is the taint key that the toleration applies to.
                          type: string
                        operator:
                          description: Operator represents a key's relationship to the value.
                          type: string
                        tolerationSeconds:
                          description: TolerationSeconds represents the period of time the toleration (which must be of
                          format: int64
                          type: integer
                        value:
                          description: Value is the taint value the toleration matches to.
                          type: string
                      type: object
                    type: array
                type: object
              rackConfig:
                description: RackConfig Configures the operator to deploy rack aware Aerospike cluster.
                properties:
                  namespaces:
                    description: List of Aerospike namespaces for which rack feature will be enabled
                    items:
                      type: string
                    type: array
                  racks:
                    description: Racks is the list of all racks
                    items:
                      description: Rack specifies single rack config
                      properties:
                        aerospikeConfig:
                          description: AerospikeConfig overrides the common AerospikeConfig for this Rack.
                          type: object
                          x-kubernetes-preserve-unknown-fields: true
                        effectiveAerospikeConfig:
                          description: Effective/operative Aerospike config.
                          type: object
                          x-kubernetes-preserve-unknown-fields: true
                        effectivePodSpec:
                          description: Effective/operative PodSpec.
                          properties:
                            affinity:
                              description: Affinity rules for pod placement.
                              properties:
                                nodeAffinity:
                                  description: Describes node affinity scheduling rules for the pod.
                                  properties:
                                    preferredDuringSchedulingIgnoredDuringExecution:
                                      description: The scheduler will prefer to schedule pods to nodes that satisfy the affinity ex
                                      items:
                                        description: An empty preferred scheduling term matches all objects with implicit weight 0 (i
                                        properties:
                                          preference:
                                            description: A node selector term, associated with the corresponding weight.
                                            properties:
                                              matchExpressions:
                                                description: A list of node selector requirements by node's labels.
                                                items:
                                                  description: A node selector requirement is a selector that contains values, a key, and an op
                                                  properties:
                                                    key:
                                                      description: The label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: Represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: An array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchFields:
                                                description: A list of node selector requirements by node's fields.
                                                items:
                                                  description: A node selector requirement is a selector that contains values, a key, and an op
                                                  properties:
                                                    key:
                                                      description: The label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: Represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: An array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                            type: object
                                          weight:
                                            description: Weight associated with matching the corresponding nodeSelectorTerm, in the range
                                            format: int32
                                            type: integer
                                        required:
                                        - preference
                                        - weight
                                        type: object
                                      type: array
                                    requiredDuringSchedulingIgnoredDuringExecution:
                                      description: If the affinity requirements specified by this field are not met at scheduling t
                                      properties:
                                        nodeSelectorTerms:
                                          description: Required. A list of node selector terms. The terms are ORed.
                                          items:
                                            description: A null or empty node selector term matches no objects.
                                            properties:
                                              matchExpressions:
                                                description: A list of node selector requirements by node's labels.
                                                items:
                                                  description: A node selector requirement is a selector that contains values, a key, and an op
                                                  properties:
                                                    key:
                                                      description: The label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: Represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: An array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchFields:
                                                description: A list of node selector requirements by node's fields.
                                                items:
                                                  description: A node selector requirement is a selector that contains values, a key, and an op
                                                  properties:
                                                    key:
                                                      description: The label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: Represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: An array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                            type: object
                                          type: array
                                      required:
                                      - nodeSelectorTerms
                                      type: object
                                  type: object
                                podAffinity:
                                  description: Describes pod affinity scheduling rules (e.g.
                                  properties:
                                    preferredDuringSchedulingIgnoredDuringExecution:
                                      description: The scheduler will prefer to schedule pods to nodes that satisfy the affinity ex
                                      items:
                                        description: The weights of all of the matched WeightedPodAffinityTerm fields are added per-n
                                        properties:
                                          podAffinityTerm:
                                            description: Required. A pod affinity term, associated with the corresponding weight.
                                            properties:
                                              labelSelector:
                                                description: A label query over a set of resources, in this case pods.
                                                properties:
                                                  matchExpressions:
                                                    description: matchExpressions is a list of label selector requirements.
                                                    items:
                                                      description: A label selector requirement is a selector that contains values, a key, and an o
                                                      properties:
                                                        key:
                                                          description: key is the label key that the selector applies to.
                                                          type: string
                                                        operator:
                                                          description: operator represents a key's relationship to a set of values.
                                                          type: string
                                                        values:
                                                          description: values is an array of string values.
                                                          items:
                                                            type: string
                                                          type: array
                                                      required:
                                                      - key
                                                      - operator
                                                      type: object
                                                    type: array
                                                  matchLabels:
                                                    additionalProperties:
                                                      type: string
                                                    description: matchLabels is a map of {key,value} pairs.
                                                    type: object
                                                type: object
                                              namespaceSelector:
                                                description: A label query over the set of namespaces that the term applies to.
                                                properties:
                                                  matchExpressions:
                                                    description: matchExpressions is a list of label selector requirements.
                                                    items:
                                                      description: A label selector requirement is a selector that contains values, a key, and an o
                                                      properties:
                                                        key:
                                                          description: key is the label key that the selector applies to.
                                                          type: string
                                                        operator:
                                                          description: operator represents a key's relationship to a set of values.
                                                          type: string
                                                        values:
                                                          description: values is an array of string values.
                                                          items:
                                                            type: string
                                                          type: array
                                                      required:
                                                      - key
                                                      - operator
                                                      type: object
                                                    type: array
                                                  matchLabels:
                                                    additionalProperties:
                                                      type: string
                                                    description: matchLabels is a map of {key,value} pairs.
                                                    type: object
                                                type: object
                                              namespaces:
                                                description: namespaces specifies a static list of namespace names that the term applies to.
                                                items:
                                                  type: string
                                                type: array
                                              topologyKey:
                                                description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                                type: string
                                            required:
                                            - topologyKey
                                            type: object
                                          weight:
                                            description: 'weight associated with matching the corresponding podAffinityTerm, in the range '
                                            format: int32
                                            type: integer
                                        required:
                                        - podAffinityTerm
                                        - weight
                                        type: object
                                      type: array
                                    requiredDuringSchedulingIgnoredDuringExecution:
                                      description: If the affinity requirements specified by this field are not met at scheduling t
                                      items:
                                        description: Defines a set of pods (namely those matching the labelSelector relative to the g
                                        properties:
                                          labelSelector:
                                            description: A label query over a set of resources, in this case pods.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          namespaceSelector:
                                            description: A label query over the set of namespaces that the term applies to.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          namespaces:
                                            description: namespaces specifies a static list of namespace names that the term applies to.
                                            items:
                                              type: string
                                            type: array
                                          topologyKey:
                                            description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                            type: string
                                        required:
                                        - topologyKey
                                        type: object
                                      type: array
                                  type: object
                                podAntiAffinity:
                                  description: Describes pod anti-affinity scheduling rules (e.g.
                                  properties:
                                    preferredDuringSchedulingIgnoredDuringExecution:
                                      description: The scheduler will prefer to schedule pods to nodes that satisfy the anti-affini
                                      items:
                                        description: The weights of all of the matched WeightedPodAffinityTerm fields are added per-n
                                        properties:
                                          podAffinityTerm:
                                            description: Required. A pod affinity term, associated with the corresponding weight.
                                            properties:
                                              labelSelector:
                                                description: A label query over a set of resources, in this case pods.
                                                properties:
                                                  matchExpressions:
                                                    description: matchExpressions is a list of label selector requirements.
                                                    items:
                                                      description: A label selector requirement is a selector that contains values, a key, and an o
                                                      properties:
                                                        key:
                                                          description: key is the label key that the selector applies to.
                                                          type: string
                                                        operator:
                                                          description: operator represents a key's relationship to a set of values.
                                                          type: string
                                                        values:
                                                          description: values is an array of string values.
                                                          items:
                                                            type: string
                                                          type: array
                                                      required:
                                                      - key
                                                      - operator
                                                      type: object
                                                    type: array
                                                  matchLabels:
                                                    additionalProperties:
                                                      type: string
                                                    description: matchLabels is a map of {key,value} pairs.
                                                    type: object
                                                type: object
                                              namespaceSelector:
                                                description: A label query over the set of namespaces that the term applies to.
                                                properties:
                                                  matchExpressions:
                                                    description: matchExpressions is a list of label selector requirements.
                                                    items:
                                                      description: A label selector requirement is a selector that contains values, a key, and an o
                                                      properties:
                                                        key:
                                                          description: key is the label key that the selector applies to.
                                                          type: string
                                                        operator:
                                                          description: operator represents a key's relationship to a set of values.
                                                          type: string
                                                        values:
                                                          description: values is an array of string values.
                                                          items:
                                                            type: string
                                                          type: array
                                                      required:
                                                      - key
                                                      - operator
                                                      type: object
                                                    type: array
                                                  matchLabels:
                                                    additionalProperties:
                                                      type: string
                                                    description: matchLabels is a map of {key,value} pairs.
                                                    type: object
                                                type: object
                                              namespaces:
                                                description: namespaces specifies a static list of namespace names that the term applies to.
                                                items:
                                                  type: string
                                                type: array
                                              topologyKey:
                                                description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                                type: string
                                            required:
                                            - topologyKey
                                            type: object
                                          weight:
                                            description: 'weight associated with matching the corresponding podAffinityTerm, in the range '
                                            format: int32
                                            type: integer
                                        required:
                                        - podAffinityTerm
                                        - weight
                                        type: object
                                      type: array
                                    requiredDuringSchedulingIgnoredDuringExecution:
                                      description: If the anti-affinity requirements specified by this field are not met at schedul
                                      items:
                                        description: Defines a set of pods (namely those matching the labelSelector relative to the g
                                        properties:
                                          labelSelector:
                                            description: A label query over a set of resources, in this case pods.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          namespaceSelector:
                                            description: A label query over the set of namespaces that the term applies to.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          namespaces:
                                            description: namespaces specifies a static list of namespace names that the term applies to.
                                            items:
                                              type: string
                                            type: array
                                          topologyKey:
                                            description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                            type: string
                                        required:
                                        - topologyKey
                                        type: object
                                      type: array
                                  type: object
                              type: object
                            nodeSelector:
                              additionalProperties:
                                type: string
                              description: NodeSelector constraints for this pod.
                              type: object
                            tolerations:
                              description: Tolerations for this pod.
                              items:
                                description: The pod this Toleration is attached to tolerates any taint that matches the trip
                                properties:
                                  effect:
                                    description: Effect indicates the taint effect to match. Empty means match all taint effects.
                                    type: string
                                  key:
                                    description: Key is the taint key that the toleration applies to.
                                    type: string
                                  operator:
                                    description: Operator represents a key's relationship to the value.
                                    type: string
                                  tolerationSeconds:
                                    description: TolerationSeconds represents the period of time the toleration (which must be of
                                    format: int64
                                    type: integer
                                  value:
                                    description: Value is the taint value the toleration matches to.
                                    type: string
                                type: object
                              type: array
                          type: object
                        effectiveStorage:
                          description: Effective/operative storage.
                          properties:
                            blockVolumePolicy:
                              description: BlockVolumePolicy contains default policies for block volumes.
                              properties:
                                cascadeDelete:
                                  description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                                  type: boolean
                                effectiveCascadeDelete:
                                  description: Effective/operative value to use for cascade delete after applying defaults.
                                  type: boolean
                                effectiveInitMethod:
                                  description: Effective/operative value to use as the volume init method after applying defaul
                                  enum:
                                  - none
                                  - dd
                                  - blkdiscard
                                  - deleteFiles
                                  type: string
                                initMethod:
                                  description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                                  enum:
                                  - none
                                  - dd
                                  - blkdiscard
                                  - deleteFiles
                                  type: string
                              type: object
                            filesystemVolumePolicy:
                              description: FileSystemVolumePolicy contains default policies for filesystem volumes.
                              properties:
                                cascadeDelete:
                                  description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                                  type: boolean
                                effectiveCascadeDelete:
                                  description: Effective/operative value to use for cascade delete after applying defaults.
                                  type: boolean
                                effectiveInitMethod:
                                  description: Effective/operative value to use as the volume init method after applying defaul
                                  enum:
                                  - none
                                  - dd
                                  - blkdiscard
                                  - deleteFiles
                                  type: string
                                initMethod:
                                  description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                                  enum:
                                  - none
                                  - dd
                                  - blkdiscard
                                  - deleteFiles
                                  type: string
                              type: object
                            volumes:
                              description: Volumes list to attach to created pods.
                              items:
                                properties:
                                  aerospike:
                                    description: Aerospike attachment of this volume on Aerospike server container.
                                    properties:
                                      mountOptions:
                                        properties:
                                          mountPropagation:
                                            description: mountPropagation determines how mounts are propagated from the host to container
                                            type: string
                                          readOnly:
                                            description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                            type: boolean
                                          subPath:
                                            description: Path within the volume from which the container's volume should be mounted.
                                            type: string
                                          subPathExpr:
                                            description: Expanded path within the volume from which the container's volume should be moun
                                            type: string
                                        type: object
                                      path:
                                        description: Path to attach the volume on the Aerospike server container.
                                        type: string
                                    required:
                                    - path
                                    type: object
                                  cascadeDelete:
                                    description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                                    type: boolean
                                  effectiveCascadeDelete:
                                    description: Effective/operative value to use for cascade delete after applying defaults.
                                    type: boolean
                                  effectiveInitMethod:
                                    description: Effective/operative value to use as the volume init method after applying defaul
                                    enum:
                                    - none
                                    - dd
                                    - blkdiscard
                                    - deleteFiles
                                    type: string
                                  initContainers:
                                    description: InitContainers are additional init containers where this volume will be mounted
                                    items:
                                      description: VolumeAttachment specifies volume attachment to a container.
                                      properties:
                                        containerName:
                                          description: ContainerName is the name of the container to attach this volume to.
                                          type: string
                                        mountOptions:
                                          properties:
                                            mountPropagation:
                                              description: mountPropagation determines how mounts are propagated from the host to container
                                              type: string
                                            readOnly:
                                              description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                              type: boolean
                                            subPath:
                                              description: Path within the volume from which the container's volume should be mounted.
                                              type: string
                                            subPathExpr:
                                              description: Expanded path within the volume from which the container's volume should be moun
                                              type: string
                                          type: object
                                        path:
                                          description: Path to attache the volume on the container.
                                          type: string
                                      required:
                                      - containerName
                                      - path
                                      type: object
                                    type: array
                                  initMethod:
                                    description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                                    enum:
                                    - none
                                    - dd
                                    - blkdiscard
                                    - deleteFiles
                                    type: string
                                  name:
                                    description: Name for this volume, Name or path should be given.
                                    type: string
                                  sidecars:
                                    description: Sidecars are side containers where this volume will be mounted
                                    items:
                                      description: VolumeAttachment specifies volume attachment to a container.
                                      properties:
                                        containerName:
                                          description: ContainerName is the name of the container to attach this volume to.
                                          type: string
                                        mountOptions:
                                          properties:
                                            mountPropagation:
                                              description: mountPropagation determines how mounts are propagated from the host to container
                                              type: string
                                            readOnly:
                                              description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                              type: boolean
                                            subPath:
                                              description: Path within the volume from which the container's volume should be mounted.
                                              type: string
                                            subPathExpr:
                                              description: Expanded path within the volume from which the container's volume should be moun
                                              type: string
                                          type: object
                                        path:
                                          description: Path to attache the volume on the container.
                                          type: string
                                      required:
                                      - containerName
                                      - path
                                      type: object
                                    type: array
                                  source:
                                    description: Source of this volume.
                                    properties:
                                      configMap:
                                        description: ConfigMap represents a configMap that should populate this volume
                                        properties:
                                          defaultMode:
                                            description: 'Optional: mode bits used to set permissions on created files by default.'
                                            format: int32
                                            type: integer
                                          items:
                                            description: If unspecified, each key-value pair in the Data field of the referenced ConfigMa
                                            items:
                                              description: Maps a string key to a path within a volume.
                                              properties:
                                                key:
                                                  description: The key to project.
                                                  type: string
                                                mode:
                                                  description: 'Optional: mode bits used to set permissions on this file.'
                                                  format: int32
                                                  type: integer
                                                path:
                                                  description: The relative path of the file to map the key to. May not be an absolute path.
                                                  type: string
                                              required:
                                              - key
                                              - path
                                              type: object
                                            type: array
                                          name:
                                            description: 'Name of the referent. More info: https://kubernetes.'
                                            type: string
                                          optional:
                                            description: Specify whether the ConfigMap or its keys must be defined
                                            type: boolean
                                        type: object
                                      emptyDir:
                                        description: EmptyDir represents a temporary directory that shares a pod's lifetime.
                                        properties:
                                          medium:
                                            description: What type of storage medium should back this directory.
                                            type: string
                                          sizeLimit:
                                            anyOf:
                                            - type: integer
                                            - type: string
                                            description: Total amount of local storage required for this EmptyDir volume.
                                            pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                            x-kubernetes-int-or-string: true
                                        type: object
                                      persistentVolume:
                                        description: PersistentVolumeSpec describes a persistent volume to claim and attach to Aerosp
                                        properties:
                                          accessModes:
                                            items:
                                              type: string
                                            type: array
                                          selector:
                                            description: A label query over volumes to consider for binding.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          size:
                                            anyOf:
                                            - type: integer
                                            - type: string
                                            description: Size of volume.
                                            pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                            x-kubernetes-int-or-string: true
                                          storageClass:
                                            description: StorageClass should be pre-created by user.
                                            type: string
                                          volumeMode:
                                            description: VolumeMode specifies if the volume is block/raw or a filesystem.
                                            type: string
                                        required:
                                        - size
                                        - storageClass
                                        - volumeMode
                                        type: object
                                      secret:
                                        description: Adapts a Secret into a volume.
                                        properties:
                                          defaultMode:
                                            description: 'Optional: mode bits used to set permissions on created files by default.'
                                            format: int32
                                            type: integer
                                          items:
                                            description: If unspecified, each key-value pair in the Data field of the referenced Secret w
                                            items:
                                              description: Maps a string key to a path within a volume.
                                              properties:
                                                key:
                                                  description: The key to project.
                                                  type: string
                                                mode:
                                                  description: 'Optional: mode bits used to set permissions on this file.'
                                                  format: int32
                                                  type: integer
                                                path:
                                                  description: The relative path of the file to map the key to. May not be an absolute path.
                                                  type: string
                                              required:
                                              - key
                                              - path
                                              type: object
                                            type: array
                                          optional:
                                            description: Specify whether the Secret or its keys must be defined
                                            type: boolean
                                          secretName:
                                            description: 'Name of the secret in the pod''s namespace to use. More info: https://kubernetes.'
                                            type: string
                                        type: object
                                    type: object
                                required:
                                - name
                                type: object
                              type: array
                              x-kubernetes-list-map-keys:
                              - name
                              x-kubernetes-list-type: map
                          type: object
                        id:
                          description: Identifier for the rack
                          type: integer
                        nodeName:
                          description: K8s Node name for setting rack affinity.
                          type: string
                        podSpec:
                          description: PodSpec to use for the pods in this rack.
                          properties:
                            affinity:
                              description: Affinity rules for pod placement.
                              properties:
                                nodeAffinity:
                                  description: Describes node affinity scheduling rules for the pod.
                                  properties:
                                    preferredDuringSchedulingIgnoredDuringExecution:
                                      description: The scheduler will prefer to schedule pods to nodes that satisfy the affinity ex
                                      items:
                                        description: An empty preferred scheduling term matches all objects with implicit weight 0 (i
                                        properties:
                                          preference:
                                            description: A node selector term, associated with the corresponding weight.
                                            properties:
                                              matchExpressions:
                                                description: A list of node selector requirements by node's labels.
                                                items:
                                                  description: A node selector requirement is a selector that contains values, a key, and an op
                                                  properties:
                                                    key:
                                                      description: The label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: Represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: An array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchFields:
                                                description: A list of node selector requirements by node's fields.
                                                items:
                                                  description: A node selector requirement is a selector that contains values, a key, and an op
                                                  properties:
                                                    key:
                                                      description: The label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: Represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: An array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                            type: object
                                          weight:
                                            description: Weight associated with matching the corresponding nodeSelectorTerm, in the range
                                            format: int32
                                            type: integer
                                        required:
                                        - preference
                                        - weight
                                        type: object
                                      type: array
                                    requiredDuringSchedulingIgnoredDuringExecution:
                                      description: If the affinity requirements specified by this field are not met at scheduling t
                                      properties:
                                        nodeSelectorTerms:
                                          description: Required. A list of node selector terms. The terms are ORed.
                                          items:
                                            description: A null or empty node selector term matches no objects.
                                            properties:
                                              matchExpressions:
                                                description: A list of node selector requirements by node's labels.
                                                items:
                                                  description: A node selector requirement is a selector that contains values, a key, and an op
                                                  properties:
                                                    key:
                                                      description: The label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: Represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: An array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchFields:
                                                description: A list of node selector requirements by node's fields.
                                                items:
                                                  description: A node selector requirement is a selector that contains values, a key, and an op
                                                  properties:
                                                    key:
                                                      description: The label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: Represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: An array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                            type: object
                                          type: array
                                      required:
                                      - nodeSelectorTerms
                                      type: object
                                  type: object
                                podAffinity:
                                  description: Describes pod affinity scheduling rules (e.g.
                                  properties:
                                    preferredDuringSchedulingIgnoredDuringExecution:
                                      description: The scheduler will prefer to schedule pods to nodes that satisfy the affinity ex
                                      items:
                                        description: The weights of all of the matched WeightedPodAffinityTerm fields are added per-n
                                        properties:
                                          podAffinityTerm:
                                            description: Required. A pod affinity term, associated with the corresponding weight.
                                            properties:
                                              labelSelector:
                                                description: A label query over a set of resources, in this case pods.
                                                properties:
                                                  matchExpressions:
                                                    description: matchExpressions is a list of label selector requirements.
                                                    items:
                                                      description: A label selector requirement is a selector that contains values, a key, and an o
                                                      properties:
                                                        key:
                                                          description: key is the label key that the selector applies to.
                                                          type: string
                                                        operator:
                                                          description: operator represents a key's relationship to a set of values.
                                                          type: string
                                                        values:
                                                          description: values is an array of string values.
                                                          items:
                                                            type: string
                                                          type: array
                                                      required:
                                                      - key
                                                      - operator
                                                      type: object
                                                    type: array
                                                  matchLabels:
                                                    additionalProperties:
                                                      type: string
                                                    description: matchLabels is a map of {key,value} pairs.
                                                    type: object
                                                type: object
                                              namespaceSelector:
                                                description: A label query over the set of namespaces that the term applies to.
                                                properties:
                                                  matchExpressions:
                                                    description: matchExpressions is a list of label selector requirements.
                                                    items:
                                                      description: A label selector requirement is a selector that contains values, a key, and an o
                                                      properties:
                                                        key:
                                                          description: key is the label key that the selector applies to.
                                                          type: string
                                                        operator:
                                                          description: operator represents a key's relationship to a set of values.
                                                          type: string
                                                        values:
                                                          description: values is an array of string values.
                                                          items:
                                                            type: string
                                                          type: array
                                                      required:
                                                      - key
                                                      - operator
                                                      type: object
                                                    type: array
                                                  matchLabels:
                                                    additionalProperties:
                                                      type: string
                                                    description: matchLabels is a map of {key,value} pairs.
                                                    type: object
                                                type: object
                                              namespaces:
                                                description: namespaces specifies a static list of namespace names that the term applies to.
                                                items:
                                                  type: string
                                                type: array
                                              topologyKey:
                                                description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                                type: string
                                            required:
                                            - topologyKey
                                            type: object
                                          weight:
                                            description: 'weight associated with matching the corresponding podAffinityTerm, in the range '
                                            format: int32
                                            type: integer
                                        required:
                                        - podAffinityTerm
                                        - weight
                                        type: object
                                      type: array
                                    requiredDuringSchedulingIgnoredDuringExecution:
                                      description: If the affinity requirements specified by this field are not met at scheduling t
                                      items:
                                        description: Defines a set of pods (namely those matching the labelSelector relative to the g
                                        properties:
                                          labelSelector:
                                            description: A label query over a set of resources, in this case pods.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          namespaceSelector:
                                            description: A label query over the set of namespaces that the term applies to.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          namespaces:
                                            description: namespaces specifies a static list of namespace names that the term applies to.
                                            items:
                                              type: string
                                            type: array
                                          topologyKey:
                                            description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                            type: string
                                        required:
                                        - topologyKey
                                        type: object
                                      type: array
                                  type: object
                                podAntiAffinity:
                                  description: Describes pod anti-affinity scheduling rules (e.g.
                                  properties:
                                    preferredDuringSchedulingIgnoredDuringExecution:
                                      description: The scheduler will prefer to schedule pods to nodes that satisfy the anti-affini
                                      items:
                                        description: The weights of all of the matched WeightedPodAffinityTerm fields are added per-n
                                        properties:
                                          podAffinityTerm:
                                            description: Required. A pod affinity term, associated with the corresponding weight.
                                            properties:
                                              labelSelector:
                                                description: A label query over a set of resources, in this case pods.
                                                properties:
                                                  matchExpressions:
                                                    description: matchExpressions is a list of label selector requirements.
                                                    items:
                                                      description: A label selector requirement is a selector that contains values, a key, and an o
                                                      properties:
                                                        key:
                                                          description: key is the label key that the selector applies to.
                                                          type: string
                                                        operator:
                                                          description: operator represents a key's relationship to a set of values.
                                                          type: string
                                                        values:
                                                          description: values is an array of string values.
                                                          items:
                                                            type: string
                                                          type: array
                                                      required:
                                                      - key
                                                      - operator
                                                      type: object
                                                    type: array
                                                  matchLabels:
                                                    additionalProperties:
                                                      type: string
                                                    description: matchLabels is a map of {key,value} pairs.
                                                    type: object
                                                type: object
                                              namespaceSelector:
                                                description: A label query over the set of namespaces that the term applies to.
                                                properties:
                                                  matchExpressions:
                                                    description: matchExpressions is a list of label selector requirements.
                                                    items:
                                                      description: A label selector requirement is a selector that contains values, a key, and an o
                                                      properties:
                                                        key:
                                                          description: key is the label key that the selector applies to.
                                                          type: string
                                                        operator:
                                                          description: operator represents a key's relationship to a set of values.
                                                          type: string
                                                        values:
                                                          description: values is an array of string values.
                                                          items:
                                                            type: string
                                                          type: array
                                                      required:
                                                      - key
                                                      - operator
                                                      type: object
                                                    type: array
                                                  matchLabels:
                                                    additionalProperties:
                                                      type: string
                                                    description: matchLabels is a map of {key,value} pairs.
                                                    type: object
                                                type: object
                                              namespaces:
                                                description: namespaces specifies a static list of namespace names that the term applies to.
                                                items:
                                                  type: string
                                                type: array
                                              topologyKey:
                                                description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                                type: string
                                            required:
                                            - topologyKey
                                            type: object
                                          weight:
                                            description: 'weight associated with matching the corresponding podAffinityTerm, in the range '
                                            format: int32
                                            type: integer
                                        required:
                                        - podAffinityTerm
                                        - weight
                                        type: object
                                      type: array
                                    requiredDuringSchedulingIgnoredDuringExecution:
                                      description: If the anti-affinity requirements specified by this field are not met at schedul
                                      items:
                                        description: Defines a set of pods (namely those matching the labelSelector relative to the g
                                        properties:
                                          labelSelector:
                                            description: A label query over a set of resources, in this case pods.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          namespaceSelector:
                                            description: A label query over the set of namespaces that the term applies to.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          namespaces:
                                            description: namespaces specifies a static list of namespace names that the term applies to.
                                            items:
                                              type: string
                                            type: array
                                          topologyKey:
                                            description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                            type: string
                                        required:
                                        - topologyKey
                                        type: object
                                      type: array
                                  type: object
                              type: object
                            nodeSelector:
                              additionalProperties:
                                type: string
                              description: NodeSelector constraints for this pod.
                              type: object
                            tolerations:
                              description: Tolerations for this pod.
                              items:
                                description: The pod this Toleration is attached to tolerates any taint that matches the trip
                                properties:
                                  effect:
                                    description: Effect indicates the taint effect to match. Empty means match all taint effects.
                                    type: string
                                  key:
                                    description: Key is the taint key that the toleration applies to.
                                    type: string
                                  operator:
                                    description: Operator represents a key's relationship to the value.
                                    type: string
                                  tolerationSeconds:
                                    description: TolerationSeconds represents the period of time the toleration (which must be of
                                    format: int64
                                    type: integer
                                  value:
                                    description: Value is the taint value the toleration matches to.
                                    type: string
                                type: object
                              type: array
                          type: object
                        rackLabel:
                          description: RackLabel for setting rack affinity.
                          type: string
                        region:
                          description: Region name for setting rack affinity.
                          type: string
                        storage:
                          description: Storage specify persistent storage to use for the pods in this rack.
                          properties:
                            blockVolumePolicy:
                              description: BlockVolumePolicy contains default policies for block volumes.
                              properties:
                                cascadeDelete:
                                  description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                                  type: boolean
                                effectiveCascadeDelete:
                                  description: Effective/operative value to use for cascade delete after applying defaults.
                                  type: boolean
                                effectiveInitMethod:
                                  description: Effective/operative value to use as the volume init method after applying defaul
                                  enum:
                                  - none
                                  - dd
                                  - blkdiscard
                                  - deleteFiles
                                  type: string
                                initMethod:
                                  description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                                  enum:
                                  - none
                                  - dd
                                  - blkdiscard
                                  - deleteFiles
                                  type: string
                              type: object
                            filesystemVolumePolicy:
                              description: FileSystemVolumePolicy contains default policies for filesystem volumes.
                              properties:
                                cascadeDelete:
                                  description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                                  type: boolean
                                effectiveCascadeDelete:
                                  description: Effective/operative value to use for cascade delete after applying defaults.
                                  type: boolean
                                effectiveInitMethod:
                                  description: Effective/operative value to use as the volume init method after applying defaul
                                  enum:
                                  - none
                                  - dd
                                  - blkdiscard
                                  - deleteFiles
                                  type: string
                                initMethod:
                                  description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                                  enum:
                                  - none
                                  - dd
                                  - blkdiscard
                                  - deleteFiles
                                  type: string
                              type: object
                            volumes:
                              description: Volumes list to attach to created pods.
                              items:
                                properties:
                                  aerospike:
                                    description: Aerospike attachment of this volume on Aerospike server container.
                                    properties:
                                      mountOptions:
                                        properties:
                                          mountPropagation:
                                            description: mountPropagation determines how mounts are propagated from the host to container
                                            type: string
                                          readOnly:
                                            description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                            type: boolean
                                          subPath:
                                            description: Path within the volume from which the container's volume should be mounted.
                                            type: string
                                          subPathExpr:
                                            description: Expanded path within the volume from which the container's volume should be moun
                                            type: string
                                        type: object
                                      path:
                                        description: Path to attach the volume on the Aerospike server container.
                                        type: string
                                    required:
                                    - path
                                    type: object
                                  cascadeDelete:
                                    description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                                    type: boolean
                                  effectiveCascadeDelete:
                                    description: Effective/operative value to use for cascade delete after applying defaults.
                                    type: boolean
                                  effectiveInitMethod:
                                    description: Effective/operative value to use as the volume init method after applying defaul
                                    enum:
                                    - none
                                    - dd
                                    - blkdiscard
                                    - deleteFiles
                                    type: string
                                  initContainers:
                                    description: InitContainers are additional init containers where this volume will be mounted
                                    items:
                                      description: VolumeAttachment specifies volume attachment to a container.
                                      properties:
                                        containerName:
                                          description: ContainerName is the name of the container to attach this volume to.
                                          type: string
                                        mountOptions:
                                          properties:
                                            mountPropagation:
                                              description: mountPropagation determines how mounts are propagated from the host to container
                                              type: string
                                            readOnly:
                                              description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                              type: boolean
                                            subPath:
                                              description: Path within the volume from which the container's volume should be mounted.
                                              type: string
                                            subPathExpr:
                                              description: Expanded path within the volume from which the container's volume should be moun
                                              type: string
                                          type: object
                                        path:
                                          description: Path to attache the volume on the container.
                                          type: string
                                      required:
                                      - containerName
                                      - path
                                      type: object
                                    type: array
                                  initMethod:
                                    description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                                    enum:
                                    - none
                                    - dd
                                    - blkdiscard
                                    - deleteFiles
                                    type: string
                                  name:
                                    description: Name for this volume, Name or path should be given.
                                    type: string
                                  sidecars:
                                    description: Sidecars are side containers where this volume will be mounted
                                    items:
                                      description: VolumeAttachment specifies volume attachment to a container.
                                      properties:
                                        containerName:
                                          description: ContainerName is the name of the container to attach this volume to.
                                          type: string
                                        mountOptions:
                                          properties:
                                            mountPropagation:
                                              description: mountPropagation determines how mounts are propagated from the host to container
                                              type: string
                                            readOnly:
                                              description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                              type: boolean
                                            subPath:
                                              description: Path within the volume from which the container's volume should be mounted.
                                              type: string
                                            subPathExpr:
                                              description: Expanded path within the volume from which the container's volume should be moun
                                              type: string
                                          type: object
                                        path:
                                          description: Path to attache the volume on the container.
                                          type: string
                                      required:
                                      - containerName
                                      - path
                                      type: object
                                    type: array
                                  source:
                                    description: Source of this volume.
                                    properties:
                                      configMap:
                                        description: ConfigMap represents a configMap that should populate this volume
                                        properties:
                                          defaultMode:
                                            description: 'Optional: mode bits used to set permissions on created files by default.'
                                            format: int32
                                            type: integer
                                          items:
                                            description: If unspecified, each key-value pair in the Data field of the referenced ConfigMa
                                            items:
                                              description: Maps a string key to a path within a volume.
                                              properties:
                                                key:
                                                  description: The key to project.
                                                  type: string
                                                mode:
                                                  description: 'Optional: mode bits used to set permissions on this file.'
                                                  format: int32
                                                  type: integer
                                                path:
                                                  description: The relative path of the file to map the key to. May not be an absolute path.
                                                  type: string
                                              required:
                                              - key
                                              - path
                                              type: object
                                            type: array
                                          name:
                                            description: 'Name of the referent. More info: https://kubernetes.'
                                            type: string
                                          optional:
                                            description: Specify whether the ConfigMap or its keys must be defined
                                            type: boolean
                                        type: object
                                      emptyDir:
                                        description: EmptyDir represents a temporary directory that shares a pod's lifetime.
                                        properties:
                                          medium:
                                            description: What type of storage medium should back this directory.
                                            type: string
                                          sizeLimit:
                                            anyOf:
                                            - type: integer
                                            - type: string
                                            description: Total amount of local storage required for this EmptyDir volume.
                                            pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                            x-kubernetes-int-or-string: true
                                        type: object
                                      persistentVolume:
                                        description: PersistentVolumeSpec describes a persistent volume to claim and attach to Aerosp
                                        properties:
                                          accessModes:
                                            items:
                                              type: string
                                            type: array
                                          selector:
                                            description: A label query over volumes to consider for binding.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          size:
                                            anyOf:
                                            - type: integer
                                            - type: string
                                            description: Size of volume.
                                            pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                            x-kubernetes-int-or-string: true
                                          storageClass:
                                            description: StorageClass should be pre-created by user.
                                            type: string
                                          volumeMode:
                                            description: VolumeMode specifies if the volume is block/raw or a filesystem.
                                            type: string
                                        required:
                                        - size
                                        - storageClass
                                        - volumeMode
                                        type: object
                                      secret:
                                        description: Adapts a Secret into a volume.
                                        properties:
                                          defaultMode:
                                            description: 'Optional: mode bits used to set permissions on created files by default.'
                                            format: int32
                                            type: integer
                                          items:
                                            description: If unspecified, each key-value pair in the Data field of the referenced Secret w
                                            items:
                                              description: Maps a string key to a path within a volume.
                                              properties:
                                                key:
                                                  description: The key to project.
                                                  type: string
                                                mode:
                                                  description: 'Optional: mode bits used to set permissions on this file.'
                                                  format: int32
                                                  type: integer
                                                path:
                                                  description: The relative path of the file to map the key to. May not be an absolute path.
                                                  type: string
                                              required:
                                              - key
                                              - path
                                              type: object
                                            type: array
                                          optional:
                                            description: Specify whether the Secret or its keys must be defined
                                            type: boolean
                                          secretName:
                                            description: 'Name of the secret in the pod''s namespace to use. More info: https://kubernetes.'
                                            type: string
                                        type: object
                                    type: object
                                required:
                                - name
                                type: object
                              type: array
                              x-kubernetes-list-map-keys:
                              - name
                              x-kubernetes-list-type: map
                          type: object
                        zone:
                          description: Zone name for setting rack affinity. Rack pods will be deployed to given Zone
                          type: string
                      required:
                      - id
                      type: object
                    nullable: true
                    type: array
                type: object
              seedsFinderServices:
                description: SeedsFinderServices describes services which are used for seeding Aerospike node
                properties:
                  loadBalancer:
                    description: LoadBalancer configuration.
                    properties:
                      annotations:
                        additionalProperties:
                          type: string
                        type: object
                      externalTrafficPolicy:
                        description: Service External Traffic Policy Type string
                        enum:
                        - Local
                        - Cluster
                        type: string
                      loadBalancerSourceRanges:
                        items:
                          type: string
                        type: array
                      port:
                        description: Port Exposed port on load balancer. If not specified TargetPort is used.
                        format: int32
                        maximum: 65535
                        minimum: 1024
                        type: integer
                      targetPort:
                        description: TargetPort Target port. If not specified the tls-port of network.
                        format: int32
                        maximum: 65535
                        minimum: 1024
                        type: integer
                    type: object
                type: object
              size:
                description: Aerospike cluster size
                format: int32
                type: integer
              storage:
                description: Storage specify persistent storage to use for the Aerospike pods.
                properties:
                  blockVolumePolicy:
                    description: BlockVolumePolicy contains default policies for block volumes.
                    properties:
                      cascadeDelete:
                        description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                        type: boolean
                      effectiveCascadeDelete:
                        description: Effective/operative value to use for cascade delete after applying defaults.
                        type: boolean
                      effectiveInitMethod:
                        description: Effective/operative value to use as the volume init method after applying defaul
                        enum:
                        - none
                        - dd
                        - blkdiscard
                        - deleteFiles
                        type: string
                      initMethod:
                        description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                        enum:
                        - none
                        - dd
                        - blkdiscard
                        - deleteFiles
                        type: string
                    type: object
                  filesystemVolumePolicy:
                    description: FileSystemVolumePolicy contains default policies for filesystem volumes.
                    properties:
                      cascadeDelete:
                        description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                        type: boolean
                      effectiveCascadeDelete:
                        description: Effective/operative value to use for cascade delete after applying defaults.
                        type: boolean
                      effectiveInitMethod:
                        description: Effective/operative value to use as the volume init method after applying defaul
                        enum:
                        - none
                        - dd
                        - blkdiscard
                        - deleteFiles
                        type: string
                      initMethod:
                        description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                        enum:
                        - none
                        - dd
                        - blkdiscard
                        - deleteFiles
                        type: string
                    type: object
                  volumes:
                    description: Volumes list to attach to created pods.
                    items:
                      properties:
                        aerospike:
                          description: Aerospike attachment of this volume on Aerospike server container.
                          properties:
                            mountOptions:
                              properties:
                                mountPropagation:
                                  description: mountPropagation determines how mounts are propagated from the host to container
                                  type: string
                                readOnly:
                                  description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                  type: boolean
                                subPath:
                                  description: Path within the volume from which the container's volume should be mounted.
                                  type: string
                                subPathExpr:
                                  description: Expanded path within the volume from which the container's volume should be moun
                                  type: string
                              type: object
                            path:
                              description: Path to attach the volume on the Aerospike server container.
                              type: string
                          required:
                          - path
                          type: object
                        cascadeDelete:
                          description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                          type: boolean
                        effectiveCascadeDelete:
                          description: Effective/operative value to use for cascade delete after applying defaults.
                          type: boolean
                        effectiveInitMethod:
                          description: Effective/operative value to use as the volume init method after applying defaul
                          enum:
                          - none
                          - dd
                          - blkdiscard
                          - deleteFiles
                          type: string
                        initContainers:
                          description: InitContainers are additional init containers where this volume will be mounted
                          items:
                            description: VolumeAttachment specifies volume attachment to a container.
                            properties:
                              containerName:
                                description: ContainerName is the name of the container to attach this volume to.
                                type: string
                              mountOptions:
                                properties:
                                  mountPropagation:
                                    description: mountPropagation determines how mounts are propagated from the host to container
                                    type: string
                                  readOnly:
                                    description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                    type: boolean
                                  subPath:
                                    description: Path within the volume from which the container's volume should be mounted.
                                    type: string
                                  subPathExpr:
                                    description: Expanded path within the volume from which the container's volume should be moun
                                    type: string
                                type: object
                              path:
                                description: Path to attache the volume on the container.
                                type: string
                            required:
                            - containerName
                            - path
                            type: object
                          type: array
                        initMethod:
                          description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                          enum:
                          - none
                          - dd
                          - blkdiscard
                          - deleteFiles
                          type: string
                        name:
                          description: Name for this volume, Name or path should be given.
                          type: string
                        sidecars:
                          description: Sidecars are side containers where this volume will be mounted
                          items:
                            description: VolumeAttachment specifies volume attachment to a container.
                            properties:
                              containerName:
                                description: ContainerName is the name of the container to attach this volume to.
                                type: string
                              mountOptions:
                                properties:
                                  mountPropagation:
                                    description: mountPropagation determines how mounts are propagated from the host to container
                                    type: string
                                  readOnly:
                                    description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                    type: boolean
                                  subPath:
                                    description: Path within the volume from which the container's volume should be mounted.
                                    type: string
                                  subPathExpr:
                                    description: Expanded path within the volume from which the container's volume should be moun
                                    type: string
                                type: object
                              path:
                                description: Path to attache the volume on the container.
                                type: string
                            required:
                            - containerName
                            - path
                            type: object
                          type: array
                        source:
                          description: Source of this volume.
                          properties:
                            configMap:
                              description: ConfigMap represents a configMap that should populate this volume
                              properties:
                                defaultMode:
                                  description: 'Optional: mode bits used to set permissions on created files by default.'
                                  format: int32
                                  type: integer
                                items:
                                  description: If unspecified, each key-value pair in the Data field of the referenced ConfigMa
                                  items:
                                    description: Maps a string key to a path within a volume.
                                    properties:
                                      key:
                                        description: The key to project.
                                        type: string
                                      mode:
                                        description: 'Optional: mode bits used to set permissions on this file.'
                                        format: int32
                                        type: integer
                                      path:
                                        description: The relative path of the file to map the key to. May not be an absolute path.
                                        type: string
                                    required:
                                    - key
                                    - path
                                    type: object
                                  type: array
                                name:
                                  description: 'Name of the referent. More info: https://kubernetes.'
                                  type: string
                                optional:
                                  description: Specify whether the ConfigMap or its keys must be defined
                                  type: boolean
                              type: object
                            emptyDir:
                              description: EmptyDir represents a temporary directory that shares a pod's lifetime.
                              properties:
                                medium:
                                  description: What type of storage medium should back this directory.
                                  type: string
                                sizeLimit:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Total amount of local storage required for this EmptyDir volume.
                                  pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                  x-kubernetes-int-or-string: true
                              type: object
                            persistentVolume:
                              description: PersistentVolumeSpec describes a persistent volume to claim and attach to Aerosp
                              properties:
                                accessModes:
                                  items:
                                    type: string
                                  type: array
                                selector:
                                  description: A label query over volumes to consider for binding.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label selector requirements.
                                      items:
                                        description: A label selector requirement is a selector that contains values, a key, and an o
                                        properties:
                                          key:
                                            description: key is the label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: values is an array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value} pairs.
                                      type: object
                                  type: object
                                size:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Size of volume.
                                  pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                  x-kubernetes-int-or-string: true
                                storageClass:
                                  description: StorageClass should be pre-created by user.
                                  type: string
                                volumeMode:
                                  description: VolumeMode specifies if the volume is block/raw or a filesystem.
                                  type: string
                              required:
                              - size
                              - storageClass
                              - volumeMode
                              type: object
                            secret:
                              description: Adapts a Secret into a volume.
                              properties:
                                defaultMode:
                                  description: 'Optional: mode bits used to set permissions on created files by default.'
                                  format: int32
                                  type: integer
                                items:
                                  description: If unspecified, each key-value pair in the Data field of the referenced Secret w
                                  items:
                                    description: Maps a string key to a path within a volume.
                                    properties:
                                      key:
                                        description: The key to project.
                                        type: string
                                      mode:
                                        description: 'Optional: mode bits used to set permissions on this file.'
                                        format: int32
                                        type: integer
                                      path:
                                        description: The relative path of the file to map the key to. May not be an absolute path.
                                        type: string
                                    required:
                                    - key
                                    - path
                                    type: object
                                  type: array
                                optional:
                                  description: Specify whether the Secret or its keys must be defined
                                  type: boolean
                                secretName:
                                  description: 'Name of the secret in the pod''s namespace to use. More info: https://kubernetes.'
                                  type: string
                              type: object
                          type: object
                      required:
                      - name
                      type: object
                    type: array
                    x-kubernetes-list-map-keys:
                    - name
                    x-kubernetes-list-type: map
                type: object
              validationPolicy:
                description: ValidationPolicy controls validation of the Aerospike cluster resource.
                properties:
                  skipWorkDirValidate:
                    description: skipWorkDirValidate validates that Aerospike work directory is mounted on a pers
                    type: boolean
                  skipXdrDlogFileValidate:
                    description: ValidateXdrDigestLogFile validates that xdr digest log file is mounted on a pers
                    type: boolean
                required:
                - skipWorkDirValidate
                - skipXdrDlogFileValidate
                type: object
            required:
            - aerospikeConfig
            - image
            - size
            type: object
          status:
            description: AerospikeClusterStatus defines the observed state of AerospikeCluster
            nullable: true
            properties:
              aerospikeAccessControl:
                description: AerospikeAccessControl has the Aerospike roles and users definitions.
                properties:
                  adminPolicy:
                    description: 'AerospikeClientAdminPolicy specify the aerospike client admin policy for access '
                    properties:
                      timeout:
                        description: Timeout for admin client policy in milliseconds.
                        type: integer
                    required:
                    - timeout
                    type: object
                  roles:
                    description: Roles is the set of roles to allow on the Aerospike cluster.
                    items:
                      description: AerospikeRoleSpec specifies an Aerospike database role and its associated privil
                      properties:
                        name:
                          description: Name of this role.
                          type: string
                        privileges:
                          description: Privileges granted to this role.
                          items:
                            type: string
                          type: array
                          x-kubernetes-list-type: set
                        readQuota:
                          description: ReadQuota specifies permitted rate of read records for current role (the value i
                          format: int32
                          type: integer
                        whitelist:
                          description: Whitelist of host address allowed for this role.
                          items:
                            type: string
                          type: array
                          x-kubernetes-list-type: set
                        writeQuota:
                          description: WriteQuota specifies permitted rate of write records for current role (the value
                          format: int32
                          type: integer
                      required:
                      - name
                      - privileges
                      type: object
                    type: array
                    x-kubernetes-list-map-keys:
                    - name
                    x-kubernetes-list-type: map
                  users:
                    description: Users is the set of users to allow on the Aerospike cluster.
                    items:
                      description: 'AerospikeUserSpec specifies an Aerospike database user, the secret name for the '
                      properties:
                        name:
                          description: Name is the user's username.
                          type: string
                        roles:
                          description: Roles is the list of roles granted to the user.
                          items:
                            type: string
                          type: array
                          x-kubernetes-list-type: set
                        secretName:
                          description: SecretName has secret info created by user.
                          type: string
                      required:
                      - name
                      - roles
                      - secretName
                      type: object
                    type: array
                    x-kubernetes-list-map-keys:
                    - name
                    x-kubernetes-list-type: map
                required:
                - users
                type: object
              aerospikeConfig:
                description: AerospikeConfig sets config in aerospike.conf file.
                nullable: true
                type: object
                x-kubernetes-preserve-unknown-fields: true
              aerospikeNetworkPolicy:
                description: AerospikeNetworkPolicy specifies how clients and tools access the Aerospike clus
                properties:
                  access:
                    description: AccessType is the type of network address to use for Aerospike access address.
                    enum:
                    - pod
                    - hostInternal
                    - hostExternal
                    type: string
                  alternateAccess:
                    description: AlternateAccessType is the type of network address to use for Aerospike alternat
                    enum:
                    - pod
                    - hostInternal
                    - hostExternal
                    type: string
                  tlsAccess:
                    description: TLSAccessType is the type of network address to use for Aerospike TLS access add
                    enum:
                    - pod
                    - hostInternal
                    - hostExternal
                    type: string
                  tlsAlternateAccess:
                    description: TLSAlternateAccessType is the type of network address to use for Aerospike TLS a
                    enum:
                    - pod
                    - hostInternal
                    - hostExternal
                    type: string
                type: object
              image:
                description: Aerospike server image
                type: string
              multiPodPerHost:
                description: If set true then multiple pods can be created per Kubernetes Node.
                type: boolean
              operatorClientCertSpec:
                description: Certificates to connect to Aerospike.
                properties:
                  certPathInOperator:
                    description: AerospikeCertPathInOperatorSource contain configuration for certificates used by
                    properties:
                      caCertsPath:
                        type: string
                      clientCertPath:
                        type: string
                      clientKeyPath:
                        type: string
                    type: object
                  secretCertSource:
                    properties:
                      caCertsFilename:
                        type: string
                      clientCertFilename:
                        type: string
                      clientKeyFilename:
                        type: string
                      secretName:
                        type: string
                      secretNamespace:
                        type: string
                    required:
                    - secretName
                    type: object
                  tlsClientName:
                    description: If specified, this name will be added to tls-authenticate-client list by the ope
                    type: string
                type: object
              podSpec:
                description: Additional configuration for create Aerospike pods.
                properties:
                  aerospikeContainer:
                    description: 'AerospikeContainerSpec contains settings for aerospike-server container created '
                    properties:
                      resources:
                        description: Define resources requests and limits for Aerospike Server Container.
                        properties:
                          limits:
                            additionalProperties:
                              anyOf:
                              - type: integer
                              - type: string
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            description: Limits describes the maximum amount of compute resources allowed.
                            type: object
                          requests:
                            additionalProperties:
                              anyOf:
                              - type: integer
                              - type: string
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            description: Requests describes the minimum amount of compute resources required.
                            type: object
                        type: object
                      securityContext:
                        description: SecurityContext that will be added to aerospike-server container created by oper
                        properties:
                          allowPrivilegeEscalation:
                            description: AllowPrivilegeEscalation controls whether a process can gain more privileges tha
                            type: boolean
                          capabilities:
                            description: The capabilities to add/drop when running containers.
                            properties:
                              add:
                                description: Added capabilities
                                items:
                                  description: Capability represent POSIX capabilities type
                                  type: string
                                type: array
                              drop:
                                description: Removed capabilities
                                items:
                                  description: Capability represent POSIX capabilities type
                                  type: string
                                type: array
                            type: object
                          privileged:
                            description: Run container in privileged mode.
                            type: boolean
                          procMount:
                            description: procMount denotes the type of proc mount to use for the containers.
                            type: string
                          readOnlyRootFilesystem:
                            description: Whether this container has a read-only root filesystem. Default is false.
                            type: boolean
                          runAsGroup:
                            description: The GID to run the entrypoint of the container process.
                            format: int64
                            type: integer
                          runAsNonRoot:
                            description: Indicates that the container must run as a non-root user.
                            type: boolean
                          runAsUser:
                            description: The UID to run the entrypoint of the container process.
                            format: int64
                            type: integer
                          seLinuxOptions:
                            description: The SELinux context to be applied to the container.
                            properties:
                              level:
                                description: Level is SELinux level label that applies to the container.
                                type: string
                              role:
                                description: Role is a SELinux role label that applies to the container.
                                type: string
                              type:
                                description: Type is a SELinux type label that applies to the container.
                                type: string
                              user:
                                description: User is a SELinux user label that applies to the container.
                                type: string
                            type: object
                          seccompProfile:
                            description: The seccomp options to use by this container.
                            properties:
                              localhostProfile:
                                description: localhostProfile indicates a profile defined in a file on the node should be use
                                type: string
                              type:
                                description: type indicates which kind of seccomp profile will be applied.
                                type: string
                            required:
                            - type
                            type: object
                          windowsOptions:
                            description: The Windows specific settings applied to all containers.
                            properties:
                              gmsaCredentialSpec:
                                description: GMSACredentialSpec is where the GMSA admission webhook (https://github.
                                type: string
                              gmsaCredentialSpecName:
                                description: GMSACredentialSpecName is the name of the GMSA credential spec to use.
                                type: string
                              runAsUserName:
                                description: The UserName in Windows to run the entrypoint of the container process.
                                type: string
                            type: object
                        type: object
                    type: object
                  affinity:
                    description: Affinity rules for pod placement.
                    properties:
                      nodeAffinity:
                        description: Describes node affinity scheduling rules for the pod.
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods to nodes that satisfy the affinity ex
                            items:
                              description: An empty preferred scheduling term matches all objects with implicit weight 0 (i
                              properties:
                                preference:
                                  description: A node selector term, associated with the corresponding weight.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements by node's labels.
                                      items:
                                        description: A node selector requirement is a selector that contains values, a key, and an op
                                        properties:
                                          key:
                                            description: The label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: An array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements by node's fields.
                                      items:
                                        description: A node selector requirement is a selector that contains values, a key, and an op
                                        properties:
                                          key:
                                            description: The label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: An array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                weight:
                                  description: Weight associated with matching the corresponding nodeSelectorTerm, in the range
                                  format: int32
                                  type: integer
                              required:
                              - preference
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by this field are not met at scheduling t
                            properties:
                              nodeSelectorTerms:
                                description: Required. A list of node selector terms. The terms are ORed.
                                items:
                                  description: A null or empty node selector term matches no objects.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements by node's labels.
                                      items:
                                        description: A node selector requirement is a selector that contains values, a key, and an op
                                        properties:
                                          key:
                                            description: The label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: An array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements by node's fields.
                                      items:
                                        description: A node selector requirement is a selector that contains values, a key, and an op
                                        properties:
                                          key:
                                            description: The label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: An array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                type: array
                            required:
                            - nodeSelectorTerms
                            type: object
                        type: object
                      podAffinity:
                        description: Describes pod affinity scheduling rules (e.g.
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods to nodes that satisfy the affinity ex
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm fields are added per-n
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources, in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list of label selector requirements.
                                          items:
                                            description: A label selector requirement is a selector that contains values, a key, and an o
                                            properties:
                                              key:
                                                description: key is the label key that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a key's relationship to a set of values.
                                                type: string
                                              values:
                                                description: values is an array of string values.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value} pairs.
                                          type: object
                                      type: object
                                    namespaceSelector:
                                      description: A label query over the set of namespaces that the term applies to.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list of label selector requirements.
                                          items:
                                            description: A label selector requirement is a selector that contains values, a key, and an o
                                            properties:
                                              key:
                                                description: key is the label key that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a key's relationship to a set of values.
                                                type: string
                                              values:
                                                description: values is an array of string values.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value} pairs.
                                          type: object
                                      type: object
                                    namespaces:
                                      description: namespaces specifies a static list of namespace names that the term applies to.
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: 'weight associated with matching the corresponding podAffinityTerm, in the range '
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by this field are not met at scheduling t
                            items:
                              description: Defines a set of pods (namely those matching the labelSelector relative to the g
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources, in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label selector requirements.
                                      items:
                                        description: A label selector requirement is a selector that contains values, a key, and an o
                                        properties:
                                          key:
                                            description: key is the label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: values is an array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value} pairs.
                                      type: object
                                  type: object
                                namespaceSelector:
                                  description: A label query over the set of namespaces that the term applies to.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label selector requirements.
                                      items:
                                        description: A label selector requirement is a selector that contains values, a key, and an o
                                        properties:
                                          key:
                                            description: key is the label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: values is an array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value} pairs.
                                      type: object
                                  type: object
                                namespaces:
                                  description: namespaces specifies a static list of namespace names that the term applies to.
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                      podAntiAffinity:
                        description: Describes pod anti-affinity scheduling rules (e.g.
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods to nodes that satisfy the anti-affini
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm fields are added per-n
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources, in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list of label selector requirements.
                                          items:
                                            description: A label selector requirement is a selector that contains values, a key, and an o
                                            properties:
                                              key:
                                                description: key is the label key that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a key's relationship to a set of values.
                                                type: string
                                              values:
                                                description: values is an array of string values.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value} pairs.
                                          type: object
                                      type: object
                                    namespaceSelector:
                                      description: A label query over the set of namespaces that the term applies to.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list of label selector requirements.
                                          items:
                                            description: A label selector requirement is a selector that contains values, a key, and an o
                                            properties:
                                              key:
                                                description: key is the label key that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a key's relationship to a set of values.
                                                type: string
                                              values:
                                                description: values is an array of string values.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value} pairs.
                                          type: object
                                      type: object
                                    namespaces:
                                      description: namespaces specifies a static list of namespace names that the term applies to.
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: 'weight associated with matching the corresponding podAffinityTerm, in the range '
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the anti-affinity requirements specified by this field are not met at schedul
                            items:
                              description: Defines a set of pods (namely those matching the labelSelector relative to the g
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources, in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label selector requirements.
                                      items:
                                        description: A label selector requirement is a selector that contains values, a key, and an o
                                        properties:
                                          key:
                                            description: key is the label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: values is an array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value} pairs.
                                      type: object
                                  type: object
                                namespaceSelector:
                                  description: A label query over the set of namespaces that the term applies to.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label selector requirements.
                                      items:
                                        description: A label selector requirement is a selector that contains values, a key, and an o
                                        properties:
                                          key:
                                            description: key is the label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: values is an array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value} pairs.
                                      type: object
                                  type: object
                                namespaces:
                                  description: namespaces specifies a static list of namespace names that the term applies to.
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                    type: object
                  dnsPolicy:
                    description: DnsPolicy same as https://kubernetes.
                    type: string
                  effectiveDNSPolicy:
                    description: Effective value of the DNSPolicy
                    type: string
                  hostNetwork:
                    description: HostNetwork enables host networking for the pod.
                    type: boolean
                  initContainers:
                    description: InitContainers to add to the pods.
                    items:
                      description: A single application container that you want to run within a pod.
                      properties:
                        args:
                          description: Arguments to the entrypoint.
                          items:
                            type: string
                          type: array
                        command:
                          description: Entrypoint array. Not executed within a shell.
                          items:
                            type: string
                          type: array
                        env:
                          description: List of environment variables to set in the container. Cannot be updated.
                          items:
                            description: EnvVar represents an environment variable present in a Container.
                            properties:
                              name:
                                description: Name of the environment variable. Must be a C_IDENTIFIER.
                                type: string
                              value:
                                description: Variable references $(VAR_NAME) are expanded using the previous defined environm
                                type: string
                              valueFrom:
                                description: Source for the environment variable's value.
                                properties:
                                  configMapKeyRef:
                                    description: Selects a key of a ConfigMap.
                                    properties:
                                      key:
                                        description: The key to select.
                                        type: string
                                      name:
                                        description: 'Name of the referent. More info: https://kubernetes.'
                                        type: string
                                      optional:
                                        description: Specify whether the ConfigMap or its key must be defined
                                        type: boolean
                                    required:
                                    - key
                                    type: object
                                  fieldRef:
                                    description: 'Selects a field of the pod: supports metadata.name, metadata.'
                                    properties:
                                      apiVersion:
                                        description: Version of the schema the FieldPath is written in terms of, defaults to "v1".
                                        type: string
                                      fieldPath:
                                        description: Path of the field to select in the specified API version.
                                        type: string
                                    required:
                                    - fieldPath
                                    type: object
                                  resourceFieldRef:
                                    description: 'Selects a resource of the container: only resources limits and requests (limits.'
                                    properties:
                                      containerName:
                                        description: 'Container name: required for volumes, optional for env vars'
                                        type: string
                                      divisor:
                                        anyOf:
                                        - type: integer
                                        - type: string
                                        description: Specifies the output format of the exposed resources, defaults to "1"
                                        pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                        x-kubernetes-int-or-string: true
                                      resource:
                                        description: 'Required: resource to select'
                                        type: string
                                    required:
                                    - resource
                                    type: object
                                  secretKeyRef:
                                    description: Selects a key of a secret in the pod's namespace
                                    properties:
                                      key:
                                        description: The key of the secret to select from.  Must be a valid secret key.
                                        type: string
                                      name:
                                        description: 'Name of the referent. More info: https://kubernetes.'
                                        type: string
                                      optional:
                                        description: Specify whether the Secret or its key must be defined
                                        type: boolean
                                    required:
                                    - key
                                    type: object
                                type: object
                            required:
                            - name
                            type: object
                          type: array
                        envFrom:
                          description: List of sources to populate environment variables in the container.
                          items:
                            description: EnvFromSource represents the source of a set of ConfigMaps
                            properties:
                              configMapRef:
                                description: The ConfigMap to select from
                                properties:
                                  name:
                                    description: 'Name of the referent. More info: https://kubernetes.'
                                    type: string
                                  optional:
                                    description: Specify whether the ConfigMap must be defined
                                    type: boolean
                                type: object
                              prefix:
                                description: An optional identifier to prepend to each key in the ConfigMap.
                                type: string
                              secretRef:
                                description: The Secret to select from
                                properties:
                                  name:
                                    description: 'Name of the referent. More info: https://kubernetes.'
                                    type: string
                                  optional:
                                    description: Specify whether the Secret must be defined
                                    type: boolean
                                type: object
                            type: object
                          type: array
                        image:
                          description: 'Docker image name. More info: https://kubernetes.'
                          type: string
                        imagePullPolicy:
                          description: Image pull policy. One of Always, Never, IfNotPresent.
                          type: string
                        lifecycle:
                          description: Actions that the management system should take in response to container lifecycl
                          properties:
                            postStart:
                              description: PostStart is called immediately after a container is created.
                              properties:
                                exec:
                                  description: One and only one of the following should be specified.
                                  properties:
                                    command:
                                      description: Command is the command line to execute inside the container, the working directo
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                httpGet:
                                  description: HTTPGet specifies the http request to perform.
                                  properties:
                                    host:
                                      description: Host name to connect to, defaults to the pod IP.
                                      type: string
                                    httpHeaders:
                                      description: Custom headers to set in the request. HTTP allows repeated headers.
                                      items:
                                        description: HTTPHeader describes a custom header to be used in HTTP probes
                                        properties:
                                          name:
                                            description: The header field name
                                            type: string
                                          value:
                                            description: The header field value
                                            type: string
                                        required:
                                        - name
                                        - value
                                        type: object
                                      type: array
                                    path:
                                      description: Path to access on the HTTP server.
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Name or number of the port to access on the container.
                                      x-kubernetes-int-or-string: true
                                    scheme:
                                      description: Scheme to use for connecting to the host. Defaults to HTTP.
                                      type: string
                                  required:
                                  - port
                                  type: object
                                tcpSocket:
                                  description: TCPSocket specifies an action involving a TCP port.
                                  properties:
                                    host:
                                      description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Number or name of the port to access on the container.
                                      x-kubernetes-int-or-string: true
                                  required:
                                  - port
                                  type: object
                              type: object
                            preStop:
                              description: PreStop is called immediately before a container is terminated due to an API req
                              properties:
                                exec:
                                  description: One and only one of the following should be specified.
                                  properties:
                                    command:
                                      description: Command is the command line to execute inside the container, the working directo
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                httpGet:
                                  description: HTTPGet specifies the http request to perform.
                                  properties:
                                    host:
                                      description: Host name to connect to, defaults to the pod IP.
                                      type: string
                                    httpHeaders:
                                      description: Custom headers to set in the request. HTTP allows repeated headers.
                                      items:
                                        description: HTTPHeader describes a custom header to be used in HTTP probes
                                        properties:
                                          name:
                                            description: The header field name
                                            type: string
                                          value:
                                            description: The header field value
                                            type: string
                                        required:
                                        - name
                                        - value
                                        type: object
                                      type: array
                                    path:
                                      description: Path to access on the HTTP server.
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Name or number of the port to access on the container.
                                      x-kubernetes-int-or-string: true
                                    scheme:
                                      description: Scheme to use for connecting to the host. Defaults to HTTP.
                                      type: string
                                  required:
                                  - port
                                  type: object
                                tcpSocket:
                                  description: TCPSocket specifies an action involving a TCP port.
                                  properties:
                                    host:
                                      description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Number or name of the port to access on the container.
                                      x-kubernetes-int-or-string: true
                                  required:
                                  - port
                                  type: object
                              type: object
                          type: object
                        livenessProbe:
                          description: Periodic probe of container liveness.
                          properties:
                            exec:
                              description: One and only one of the following should be specified.
                              properties:
                                command:
                                  description: Command is the command line to execute inside the container, the working directo
                                  items:
                                    type: string
                                  type: array
                              type: object
                            failureThreshold:
                              description: 'Minimum consecutive failures for the probe to be considered failed after having '
                              format: int32
                              type: integer
                            httpGet:
                              description: HTTPGet specifies the http request to perform.
                              properties:
                                host:
                                  description: Host name to connect to, defaults to the pod IP.
                                  type: string
                                httpHeaders:
                                  description: Custom headers to set in the request. HTTP allows repeated headers.
                                  items:
                                    description: HTTPHeader describes a custom header to be used in HTTP probes
                                    properties:
                                      name:
                                        description: The header field name
                                        type: string
                                      value:
                                        description: The header field value
                                        type: string
                                    required:
                                    - name
                                    - value
                                    type: object
                                  type: array
                                path:
                                  description: Path to access on the HTTP server.
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Name or number of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                                scheme:
                                  description: Scheme to use for connecting to the host. Defaults to HTTP.
                                  type: string
                              required:
                              - port
                              type: object
                            initialDelaySeconds:
                              description: Number of seconds after the container has started before liveness probes are ini
                              format: int32
                              type: integer
                            periodSeconds:
                              description: How often (in seconds) to perform the probe. Default to 10 seconds.
                              format: int32
                              type: integer
                            successThreshold:
                              description: Minimum consecutive successes for the probe to be considered successful after ha
                              format: int32
                              type: integer
                            tcpSocket:
                              description: TCPSocket specifies an action involving a TCP port.
                              properties:
                                host:
                                  description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Number or name of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                              required:
                              - port
                              type: object
                            terminationGracePeriodSeconds:
                              description: Optional duration in seconds the pod needs to terminate gracefully upon probe fa
                              format: int64
                              type: integer
                            timeoutSeconds:
                              description: Number of seconds after which the probe times out. Defaults to 1 second.
                              format: int32
                              type: integer
                          type: object
                        name:
                          description: Name of the container specified as a DNS_LABEL.
                          type: string
                        ports:
                          description: List of ports to expose from the container.
                          items:
                            description: ContainerPort represents a network port in a single container.
                            properties:
                              containerPort:
                                description: Number of port to expose on the pod's IP address.
                                format: int32
                                type: integer
                              hostIP:
                                description: What host IP to bind the external port to.
                                type: string
                              hostPort:
                                description: Number of port to expose on the host.
                                format: int32
                                type: integer
                              name:
                                description: If specified, this must be an IANA_SVC_NAME and unique within the pod.
                                type: string
                              protocol:
                                default: TCP
                                description: Protocol for port. Must be UDP, TCP, or SCTP. Defaults to "TCP".
                                type: string
                            required:
                            - containerPort
                            type: object
                          type: array
                          x-kubernetes-list-map-keys:
                          - containerPort
                          - protocol
                          x-kubernetes-list-type: map
                        readinessProbe:
                          description: Periodic probe of container service readiness.
                          properties:
                            exec:
                              description: One and only one of the following should be specified.
                              properties:
                                command:
                                  description: Command is the command line to execute inside the container, the working directo
                                  items:
                                    type: string
                                  type: array
                              type: object
                            failureThreshold:
                              description: 'Minimum consecutive failures for the probe to be considered failed after having '
                              format: int32
                              type: integer
                            httpGet:
                              description: HTTPGet specifies the http request to perform.
                              properties:
                                host:
                                  description: Host name to connect to, defaults to the pod IP.
                                  type: string
                                httpHeaders:
                                  description: Custom headers to set in the request. HTTP allows repeated headers.
                                  items:
                                    description: HTTPHeader describes a custom header to be used in HTTP probes
                                    properties:
                                      name:
                                        description: The header field name
                                        type: string
                                      value:
                                        description: The header field value
                                        type: string
                                    required:
                                    - name
                                    - value
                                    type: object
                                  type: array
                                path:
                                  description: Path to access on the HTTP server.
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Name or number of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                                scheme:
                                  description: Scheme to use for connecting to the host. Defaults to HTTP.
                                  type: string
                              required:
                              - port
                              type: object
                            initialDelaySeconds:
                              description: Number of seconds after the container has started before liveness probes are ini
                              format: int32
                              type: integer
                            periodSeconds:
                              description: How often (in seconds) to perform the probe. Default to 10 seconds.
                              format: int32
                              type: integer
                            successThreshold:
                              description: Minimum consecutive successes for the probe to be considered successful after ha
                              format: int32
                              type: integer
                            tcpSocket:
                              description: TCPSocket specifies an action involving a TCP port.
                              properties:
                                host:
                                  description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Number or name of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                              required:
                              - port
                              type: object
                            terminationGracePeriodSeconds:
                              description: Optional duration in seconds the pod needs to terminate gracefully upon probe fa
                              format: int64
                              type: integer
                            timeoutSeconds:
                              description: Number of seconds after which the probe times out. Defaults to 1 second.
                              format: int32
                              type: integer
                          type: object
                        resources:
                          description: Compute Resources required by this container. Cannot be updated.
                          properties:
                            limits:
                              additionalProperties:
                                anyOf:
                                - type: integer
                                - type: string
                                pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                x-kubernetes-int-or-string: true
                              description: Limits describes the maximum amount of compute resources allowed.
                              type: object
                            requests:
                              additionalProperties:
                                anyOf:
                                - type: integer
                                - type: string
                                pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                x-kubernetes-int-or-string: true
                              description: Requests describes the minimum amount of compute resources required.
                              type: object
                          type: object
                        securityContext:
                          description: 'Security options the pod should run with. More info: https://kubernetes.'
                          properties:
                            allowPrivilegeEscalation:
                              description: AllowPrivilegeEscalation controls whether a process can gain more privileges tha
                              type: boolean
                            capabilities:
                              description: The capabilities to add/drop when running containers.
                              properties:
                                add:
                                  description: Added capabilities
                                  items:
                                    description: Capability represent POSIX capabilities type
                                    type: string
                                  type: array
                                drop:
                                  description: Removed capabilities
                                  items:
                                    description: Capability represent POSIX capabilities type
                                    type: string
                                  type: array
                              type: object
                            privileged:
                              description: Run container in privileged mode.
                              type: boolean
                            procMount:
                              description: procMount denotes the type of proc mount to use for the containers.
                              type: string
                            readOnlyRootFilesystem:
                              description: Whether this container has a read-only root filesystem. Default is false.
                              type: boolean
                            runAsGroup:
                              description: The GID to run the entrypoint of the container process.
                              format: int64
                              type: integer
                            runAsNonRoot:
                              description: Indicates that the container must run as a non-root user.
                              type: boolean
                            runAsUser:
                              description: The UID to run the entrypoint of the container process.
                              format: int64
                              type: integer
                            seLinuxOptions:
                              description: The SELinux context to be applied to the container.
                              properties:
                                level:
                                  description: Level is SELinux level label that applies to the container.
                                  type: string
                                role:
                                  description: Role is a SELinux role label that applies to the container.
                                  type: string
                                type:
                                  description: Type is a SELinux type label that applies to the container.
                                  type: string
                                user:
                                  description: User is a SELinux user label that applies to the container.
                                  type: string
                              type: object
                            seccompProfile:
                              description: The seccomp options to use by this container.
                              properties:
                                localhostProfile:
                                  description: localhostProfile indicates a profile defined in a file on the node should be use
                                  type: string
                                type:
                                  description: type indicates which kind of seccomp profile will be applied.
                                  type: string
                              required:
                              - type
                              type: object
                            windowsOptions:
                              description: The Windows specific settings applied to all containers.
                              properties:
                                gmsaCredentialSpec:
                                  description: GMSACredentialSpec is where the GMSA admission webhook (https://github.
                                  type: string
                                gmsaCredentialSpecName:
                                  description: GMSACredentialSpecName is the name of the GMSA credential spec to use.
                                  type: string
                                runAsUserName:
                                  description: The UserName in Windows to run the entrypoint of the container process.
                                  type: string
                              type: object
                          type: object
                        startupProbe:
                          description: StartupProbe indicates that the Pod has successfully initialized.
                          properties:
                            exec:
                              description: One and only one of the following should be specified.
                              properties:
                                command:
                                  description: Command is the command line to execute inside the container, the working directo
                                  items:
                                    type: string
                                  type: array
                              type: object
                            failureThreshold:
                              description: 'Minimum consecutive failures for the probe to be considered failed after having '
                              format: int32
                              type: integer
                            httpGet:
                              description: HTTPGet specifies the http request to perform.
                              properties:
                                host:
                                  description: Host name to connect to, defaults to the pod IP.
                                  type: string
                                httpHeaders:
                                  description: Custom headers to set in the request. HTTP allows repeated headers.
                                  items:
                                    description: HTTPHeader describes a custom header to be used in HTTP probes
                                    properties:
                                      name:
                                        description: The header field name
                                        type: string
                                      value:
                                        description: The header field value
                                        type: string
                                    required:
                                    - name
                                    - value
                                    type: object
                                  type: array
                                path:
                                  description: Path to access on the HTTP server.
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Name or number of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                                scheme:
                                  description: Scheme to use for connecting to the host. Defaults to HTTP.
                                  type: string
                              required:
                              - port
                              type: object
                            initialDelaySeconds:
                              description: Number of seconds after the container has started before liveness probes are ini
                              format: int32
                              type: integer
                            periodSeconds:
                              description: How often (in seconds) to perform the probe. Default to 10 seconds.
                              format: int32
                              type: integer
                            successThreshold:
                              description: Minimum consecutive successes for the probe to be considered successful after ha
                              format: int32
                              type: integer
                            tcpSocket:
                              description: TCPSocket specifies an action involving a TCP port.
                              properties:
                                host:
                                  description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Number or name of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                              required:
                              - port
                              type: object
                            terminationGracePeriodSeconds:
                              description: Optional duration in seconds the pod needs to terminate gracefully upon probe fa
                              format: int64
                              type: integer
                            timeoutSeconds:
                              description: Number of seconds after which the probe times out. Defaults to 1 second.
                              format: int32
                              type: integer
                          type: object
                        stdin:
                          description: Whether this container should allocate a buffer for stdin in the container runti
                          type: boolean
                        stdinOnce:
                          description: Whether the container runtime should close the stdin channel after it has been o
                          type: boolean
                        terminationMessagePath:
                          description: 'Optional: Path at which the file to which the container''s termination message wi'
                          type: string
                        terminationMessagePolicy:
                          description: Indicate how the termination message should be populated.
                          type: string
                        tty:
                          description: Whether this container should allocate a TTY for itself, also requires 'stdin' t
                          type: boolean
                        volumeDevices:
                          description: volumeDevices is the list of block devices to be used by the container.
                          items:
                            description: volumeDevice describes a mapping of a raw block device within a container.
                            properties:
                              devicePath:
                                description: devicePath is the path inside of the container that the device will be mapped to
                                type: string
                              name:
                                description: name must match the name of a persistentVolumeClaim in the pod
                                type: string
                            required:
                            - devicePath
                            - name
                            type: object
                          type: array
                        volumeMounts:
                          description: Pod volumes to mount into the container's filesystem. Cannot be updated.
                          items:
                            description: VolumeMount describes a mounting of a Volume within a container.
                            properties:
                              mountPath:
                                description: Path within the container at which the volume should be mounted.
                                type: string
                              mountPropagation:
                                description: mountPropagation determines how mounts are propagated from the host to container
                                type: string
                              name:
                                description: This must match the Name of a Volume.
                                type: string
                              readOnly:
                                description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                type: boolean
                              subPath:
                                description: Path within the volume from which the container's volume should be mounted.
                                type: string
                              subPathExpr:
                                description: Expanded path within the volume from which the container's volume should be moun
                                type: string
                            required:
                            - mountPath
                            - name
                            type: object
                          type: array
                        workingDir:
                          description: Container's working directory.
                          type: string
                      required:
                      - name
                      type: object
                    type: array
                  multiPodPerHost:
                    description: If set true then multiple pods can be created per Kubernetes Node.
                    type: boolean
                  nodeSelector:
                    additionalProperties:
                      type: string
                    description: NodeSelector constraints for this pod.
                    type: object
                  sidecars:
                    description: Sidecars to add to pods.
                    items:
                      description: A single application container that you want to run within a pod.
                      properties:
                        args:
                          description: Arguments to the entrypoint.
                          items:
                            type: string
                          type: array
                        command:
                          description: Entrypoint array. Not executed within a shell.
                          items:
                            type: string
                          type: array
                        env:
                          description: List of environment variables to set in the container. Cannot be updated.
                          items:
                            description: EnvVar represents an environment variable present in a Container.
                            properties:
                              name:
                                description: Name of the environment variable. Must be a C_IDENTIFIER.
                                type: string
                              value:
                                description: Variable references $(VAR_NAME) are expanded using the previous defined environm
                                type: string
                              valueFrom:
                                description: Source for the environment variable's value.
                                properties:
                                  configMapKeyRef:
                                    description: Selects a key of a ConfigMap.
                                    properties:
                                      key:
                                        description: The key to select.
                                        type: string
                                      name:
                                        description: 'Name of the referent. More info: https://kubernetes.'
                                        type: string
                                      optional:
                                        description: Specify whether the ConfigMap or its key must be defined
                                        type: boolean
                                    required:
                                    - key
                                    type: object
                                  fieldRef:
                                    description: 'Selects a field of the pod: supports metadata.name, metadata.'
                                    properties:
                                      apiVersion:
                                        description: Version of the schema the FieldPath is written in terms of, defaults to "v1".
                                        type: string
                                      fieldPath:
                                        description: Path of the field to select in the specified API version.
                                        type: string
                                    required:
                                    - fieldPath
                                    type: object
                                  resourceFieldRef:
                                    description: 'Selects a resource of the container: only resources limits and requests (limits.'
                                    properties:
                                      containerName:
                                        description: 'Container name: required for volumes, optional for env vars'
                                        type: string
                                      divisor:
                                        anyOf:
                                        - type: integer
                                        - type: string
                                        description: Specifies the output format of the exposed resources, defaults to "1"
                                        pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                        x-kubernetes-int-or-string: true
                                      resource:
                                        description: 'Required: resource to select'
                                        type: string
                                    required:
                                    - resource
                                    type: object
                                  secretKeyRef:
                                    description: Selects a key of a secret in the pod's namespace
                                    properties:
                                      key:
                                        description: The key of the secret to select from.  Must be a valid secret key.
                                        type: string
                                      name:
                                        description: 'Name of the referent. More info: https://kubernetes.'
                                        type: string
                                      optional:
                                        description: Specify whether the Secret or its key must be defined
                                        type: boolean
                                    required:
                                    - key
                                    type: object
                                type: object
                            required:
                            - name
                            type: object
                          type: array
                        envFrom:
                          description: List of sources to populate environment variables in the container.
                          items:
                            description: EnvFromSource represents the source of a set of ConfigMaps
                            properties:
                              configMapRef:
                                description: The ConfigMap to select from
                                properties:
                                  name:
                                    description: 'Name of the referent. More info: https://kubernetes.'
                                    type: string
                                  optional:
                                    description: Specify whether the ConfigMap must be defined
                                    type: boolean
                                type: object
                              prefix:
                                description: An optional identifier to prepend to each key in the ConfigMap.
                                type: string
                              secretRef:
                                description: The Secret to select from
                                properties:
                                  name:
                                    description: 'Name of the referent. More info: https://kubernetes.'
                                    type: string
                                  optional:
                                    description: Specify whether the Secret must be defined
                                    type: boolean
                                type: object
                            type: object
                          type: array
                        image:
                          description: 'Docker image name. More info: https://kubernetes.'
                          type: string
                        imagePullPolicy:
                          description: Image pull policy. One of Always, Never, IfNotPresent.
                          type: string
                        lifecycle:
                          description: Actions that the management system should take in response to container lifecycl
                          properties:
                            postStart:
                              description: PostStart is called immediately after a container is created.
                              properties:
                                exec:
                                  description: One and only one of the following should be specified.
                                  properties:
                                    command:
                                      description: Command is the command line to execute inside the container, the working directo
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                httpGet:
                                  description: HTTPGet specifies the http request to perform.
                                  properties:
                                    host:
                                      description: Host name to connect to, defaults to the pod IP.
                                      type: string
                                    httpHeaders:
                                      description: Custom headers to set in the request. HTTP allows repeated headers.
                                      items:
                                        description: HTTPHeader describes a custom header to be used in HTTP probes
                                        properties:
                                          name:
                                            description: The header field name
                                            type: string
                                          value:
                                            description: The header field value
                                            type: string
                                        required:
                                        - name
                                        - value
                                        type: object
                                      type: array
                                    path:
                                      description: Path to access on the HTTP server.
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Name or number of the port to access on the container.
                                      x-kubernetes-int-or-string: true
                                    scheme:
                                      description: Scheme to use for connecting to the host. Defaults to HTTP.
                                      type: string
                                  required:
                                  - port
                                  type: object
                                tcpSocket:
                                  description: TCPSocket specifies an action involving a TCP port.
                                  properties:
                                    host:
                                      description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Number or name of the port to access on the container.
                                      x-kubernetes-int-or-string: true
                                  required:
                                  - port
                                  type: object
                              type: object
                            preStop:
                              description: PreStop is called immediately before a container is terminated due to an API req
                              properties:
                                exec:
                                  description: One and only one of the following should be specified.
                                  properties:
                                    command:
                                      description: Command is the command line to execute inside the container, the working directo
                                      items:
                                        type: string
                                      type: array
                                  type: object
                                httpGet:
                                  description: HTTPGet specifies the http request to perform.
                                  properties:
                                    host:
                                      description: Host name to connect to, defaults to the pod IP.
                                      type: string
                                    httpHeaders:
                                      description: Custom headers to set in the request. HTTP allows repeated headers.
                                      items:
                                        description: HTTPHeader describes a custom header to be used in HTTP probes
                                        properties:
                                          name:
                                            description: The header field name
                                            type: string
                                          value:
                                            description: The header field value
                                            type: string
                                        required:
                                        - name
                                        - value
                                        type: object
                                      type: array
                                    path:
                                      description: Path to access on the HTTP server.
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Name or number of the port to access on the container.
                                      x-kubernetes-int-or-string: true
                                    scheme:
                                      description: Scheme to use for connecting to the host. Defaults to HTTP.
                                      type: string
                                  required:
                                  - port
                                  type: object
                                tcpSocket:
                                  description: TCPSocket specifies an action involving a TCP port.
                                  properties:
                                    host:
                                      description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                      type: string
                                    port:
                                      anyOf:
                                      - type: integer
                                      - type: string
                                      description: Number or name of the port to access on the container.
                                      x-kubernetes-int-or-string: true
                                  required:
                                  - port
                                  type: object
                              type: object
                          type: object
                        livenessProbe:
                          description: Periodic probe of container liveness.
                          properties:
                            exec:
                              description: One and only one of the following should be specified.
                              properties:
                                command:
                                  description: Command is the command line to execute inside the container, the working directo
                                  items:
                                    type: string
                                  type: array
                              type: object
                            failureThreshold:
                              description: 'Minimum consecutive failures for the probe to be considered failed after having '
                              format: int32
                              type: integer
                            httpGet:
                              description: HTTPGet specifies the http request to perform.
                              properties:
                                host:
                                  description: Host name to connect to, defaults to the pod IP.
                                  type: string
                                httpHeaders:
                                  description: Custom headers to set in the request. HTTP allows repeated headers.
                                  items:
                                    description: HTTPHeader describes a custom header to be used in HTTP probes
                                    properties:
                                      name:
                                        description: The header field name
                                        type: string
                                      value:
                                        description: The header field value
                                        type: string
                                    required:
                                    - name
                                    - value
                                    type: object
                                  type: array
                                path:
                                  description: Path to access on the HTTP server.
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Name or number of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                                scheme:
                                  description: Scheme to use for connecting to the host. Defaults to HTTP.
                                  type: string
                              required:
                              - port
                              type: object
                            initialDelaySeconds:
                              description: Number of seconds after the container has started before liveness probes are ini
                              format: int32
                              type: integer
                            periodSeconds:
                              description: How often (in seconds) to perform the probe. Default to 10 seconds.
                              format: int32
                              type: integer
                            successThreshold:
                              description: Minimum consecutive successes for the probe to be considered successful after ha
                              format: int32
                              type: integer
                            tcpSocket:
                              description: TCPSocket specifies an action involving a TCP port.
                              properties:
                                host:
                                  description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Number or name of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                              required:
                              - port
                              type: object
                            terminationGracePeriodSeconds:
                              description: Optional duration in seconds the pod needs to terminate gracefully upon probe fa
                              format: int64
                              type: integer
                            timeoutSeconds:
                              description: Number of seconds after which the probe times out. Defaults to 1 second.
                              format: int32
                              type: integer
                          type: object
                        name:
                          description: Name of the container specified as a DNS_LABEL.
                          type: string
                        ports:
                          description: List of ports to expose from the container.
                          items:
                            description: ContainerPort represents a network port in a single container.
                            properties:
                              containerPort:
                                description: Number of port to expose on the pod's IP address.
                                format: int32
                                type: integer
                              hostIP:
                                description: What host IP to bind the external port to.
                                type: string
                              hostPort:
                                description: Number of port to expose on the host.
                                format: int32
                                type: integer
                              name:
                                description: If specified, this must be an IANA_SVC_NAME and unique within the pod.
                                type: string
                              protocol:
                                default: TCP
                                description: Protocol for port. Must be UDP, TCP, or SCTP. Defaults to "TCP".
                                type: string
                            required:
                            - containerPort
                            type: object
                          type: array
                          x-kubernetes-list-map-keys:
                          - containerPort
                          - protocol
                          x-kubernetes-list-type: map
                        readinessProbe:
                          description: Periodic probe of container service readiness.
                          properties:
                            exec:
                              description: One and only one of the following should be specified.
                              properties:
                                command:
                                  description: Command is the command line to execute inside the container, the working directo
                                  items:
                                    type: string
                                  type: array
                              type: object
                            failureThreshold:
                              description: 'Minimum consecutive failures for the probe to be considered failed after having '
                              format: int32
                              type: integer
                            httpGet:
                              description: HTTPGet specifies the http request to perform.
                              properties:
                                host:
                                  description: Host name to connect to, defaults to the pod IP.
                                  type: string
                                httpHeaders:
                                  description: Custom headers to set in the request. HTTP allows repeated headers.
                                  items:
                                    description: HTTPHeader describes a custom header to be used in HTTP probes
                                    properties:
                                      name:
                                        description: The header field name
                                        type: string
                                      value:
                                        description: The header field value
                                        type: string
                                    required:
                                    - name
                                    - value
                                    type: object
                                  type: array
                                path:
                                  description: Path to access on the HTTP server.
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Name or number of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                                scheme:
                                  description: Scheme to use for connecting to the host. Defaults to HTTP.
                                  type: string
                              required:
                              - port
                              type: object
                            initialDelaySeconds:
                              description: Number of seconds after the container has started before liveness probes are ini
                              format: int32
                              type: integer
                            periodSeconds:
                              description: How often (in seconds) to perform the probe. Default to 10 seconds.
                              format: int32
                              type: integer
                            successThreshold:
                              description: Minimum consecutive successes for the probe to be considered successful after ha
                              format: int32
                              type: integer
                            tcpSocket:
                              description: TCPSocket specifies an action involving a TCP port.
                              properties:
                                host:
                                  description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Number or name of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                              required:
                              - port
                              type: object
                            terminationGracePeriodSeconds:
                              description: Optional duration in seconds the pod needs to terminate gracefully upon probe fa
                              format: int64
                              type: integer
                            timeoutSeconds:
                              description: Number of seconds after which the probe times out. Defaults to 1 second.
                              format: int32
                              type: integer
                          type: object
                        resources:
                          description: Compute Resources required by this container. Cannot be updated.
                          properties:
                            limits:
                              additionalProperties:
                                anyOf:
                                - type: integer
                                - type: string
                                pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                x-kubernetes-int-or-string: true
                              description: Limits describes the maximum amount of compute resources allowed.
                              type: object
                            requests:
                              additionalProperties:
                                anyOf:
                                - type: integer
                                - type: string
                                pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                x-kubernetes-int-or-string: true
                              description: Requests describes the minimum amount of compute resources required.
                              type: object
                          type: object
                        securityContext:
                          description: 'Security options the pod should run with. More info: https://kubernetes.'
                          properties:
                            allowPrivilegeEscalation:
                              description: AllowPrivilegeEscalation controls whether a process can gain more privileges tha
                              type: boolean
                            capabilities:
                              description: The capabilities to add/drop when running containers.
                              properties:
                                add:
                                  description: Added capabilities
                                  items:
                                    description: Capability represent POSIX capabilities type
                                    type: string
                                  type: array
                                drop:
                                  description: Removed capabilities
                                  items:
                                    description: Capability represent POSIX capabilities type
                                    type: string
                                  type: array
                              type: object
                            privileged:
                              description: Run container in privileged mode.
                              type: boolean
                            procMount:
                              description: procMount denotes the type of proc mount to use for the containers.
                              type: string
                            readOnlyRootFilesystem:
                              description: Whether this container has a read-only root filesystem. Default is false.
                              type: boolean
                            runAsGroup:
                              description: The GID to run the entrypoint of the container process.
                              format: int64
                              type: integer
                            runAsNonRoot:
                              description: Indicates that the container must run as a non-root user.
                              type: boolean
                            runAsUser:
                              description: The UID to run the entrypoint of the container process.
                              format: int64
                              type: integer
                            seLinuxOptions:
                              description: The SELinux context to be applied to the container.
                              properties:
                                level:
                                  description: Level is SELinux level label that applies to the container.
                                  type: string
                                role:
                                  description: Role is a SELinux role label that applies to the container.
                                  type: string
                                type:
                                  description: Type is a SELinux type label that applies to the container.
                                  type: string
                                user:
                                  description: User is a SELinux user label that applies to the container.
                                  type: string
                              type: object
                            seccompProfile:
                              description: The seccomp options to use by this container.
                              properties:
                                localhostProfile:
                                  description: localhostProfile indicates a profile defined in a file on the node should be use
                                  type: string
                                type:
                                  description: type indicates which kind of seccomp profile will be applied.
                                  type: string
                              required:
                              - type
                              type: object
                            windowsOptions:
                              description: The Windows specific settings applied to all containers.
                              properties:
                                gmsaCredentialSpec:
                                  description: GMSACredentialSpec is where the GMSA admission webhook (https://github.
                                  type: string
                                gmsaCredentialSpecName:
                                  description: GMSACredentialSpecName is the name of the GMSA credential spec to use.
                                  type: string
                                runAsUserName:
                                  description: The UserName in Windows to run the entrypoint of the container process.
                                  type: string
                              type: object
                          type: object
                        startupProbe:
                          description: StartupProbe indicates that the Pod has successfully initialized.
                          properties:
                            exec:
                              description: One and only one of the following should be specified.
                              properties:
                                command:
                                  description: Command is the command line to execute inside the container, the working directo
                                  items:
                                    type: string
                                  type: array
                              type: object
                            failureThreshold:
                              description: 'Minimum consecutive failures for the probe to be considered failed after having '
                              format: int32
                              type: integer
                            httpGet:
                              description: HTTPGet specifies the http request to perform.
                              properties:
                                host:
                                  description: Host name to connect to, defaults to the pod IP.
                                  type: string
                                httpHeaders:
                                  description: Custom headers to set in the request. HTTP allows repeated headers.
                                  items:
                                    description: HTTPHeader describes a custom header to be used in HTTP probes
                                    properties:
                                      name:
                                        description: The header field name
                                        type: string
                                      value:
                                        description: The header field value
                                        type: string
                                    required:
                                    - name
                                    - value
                                    type: object
                                  type: array
                                path:
                                  description: Path to access on the HTTP server.
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Name or number of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                                scheme:
                                  description: Scheme to use for connecting to the host. Defaults to HTTP.
                                  type: string
                              required:
                              - port
                              type: object
                            initialDelaySeconds:
                              description: Number of seconds after the container has started before liveness probes are ini
                              format: int32
                              type: integer
                            periodSeconds:
                              description: How often (in seconds) to perform the probe. Default to 10 seconds.
                              format: int32
                              type: integer
                            successThreshold:
                              description: Minimum consecutive successes for the probe to be considered successful after ha
                              format: int32
                              type: integer
                            tcpSocket:
                              description: TCPSocket specifies an action involving a TCP port.
                              properties:
                                host:
                                  description: 'Optional: Host name to connect to, defaults to the pod IP.'
                                  type: string
                                port:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Number or name of the port to access on the container.
                                  x-kubernetes-int-or-string: true
                              required:
                              - port
                              type: object
                            terminationGracePeriodSeconds:
                              description: Optional duration in seconds the pod needs to terminate gracefully upon probe fa
                              format: int64
                              type: integer
                            timeoutSeconds:
                              description: Number of seconds after which the probe times out. Defaults to 1 second.
                              format: int32
                              type: integer
                          type: object
                        stdin:
                          description: Whether this container should allocate a buffer for stdin in the container runti
                          type: boolean
                        stdinOnce:
                          description: Whether the container runtime should close the stdin channel after it has been o
                          type: boolean
                        terminationMessagePath:
                          description: 'Optional: Path at which the file to which the container''s termination message wi'
                          type: string
                        terminationMessagePolicy:
                          description: Indicate how the termination message should be populated.
                          type: string
                        tty:
                          description: Whether this container should allocate a TTY for itself, also requires 'stdin' t
                          type: boolean
                        volumeDevices:
                          description: volumeDevices is the list of block devices to be used by the container.
                          items:
                            description: volumeDevice describes a mapping of a raw block device within a container.
                            properties:
                              devicePath:
                                description: devicePath is the path inside of the container that the device will be mapped to
                                type: string
                              name:
                                description: name must match the name of a persistentVolumeClaim in the pod
                                type: string
                            required:
                            - devicePath
                            - name
                            type: object
                          type: array
                        volumeMounts:
                          description: Pod volumes to mount into the container's filesystem. Cannot be updated.
                          items:
                            description: VolumeMount describes a mounting of a Volume within a container.
                            properties:
                              mountPath:
                                description: Path within the container at which the volume should be mounted.
                                type: string
                              mountPropagation:
                                description: mountPropagation determines how mounts are propagated from the host to container
                                type: string
                              name:
                                description: This must match the Name of a Volume.
                                type: string
                              readOnly:
                                description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                type: boolean
                              subPath:
                                description: Path within the volume from which the container's volume should be mounted.
                                type: string
                              subPathExpr:
                                description: Expanded path within the volume from which the container's volume should be moun
                                type: string
                            required:
                            - mountPath
                            - name
                            type: object
                          type: array
                        workingDir:
                          description: Container's working directory.
                          type: string
                      required:
                      - name
                      type: object
                    type: array
                  tolerations:
                    description: Tolerations for this pod.
                    items:
                      description: The pod this Toleration is attached to tolerates any taint that matches the trip
                      properties:
                        effect:
                          description: Effect indicates the taint effect to match. Empty means match all taint effects.
                          type: string
                        key:
                          description: Key is the taint key that the toleration applies to.
                          type: string
                        operator:
                          description: Operator represents a key's relationship to the value.
                          type: string
                        tolerationSeconds:
                          description: TolerationSeconds represents the period of time the toleration (which must be of
                          format: int64
                          type: integer
                        value:
                          description: Value is the taint value the toleration matches to.
                          type: string
                      type: object
                    type: array
                type: object
              pods:
                additionalProperties:
                  description: AerospikePodStatus contains the Aerospike specific status of the Aerospike serve
                  properties:
                    aerospike:
                      description: Aerospike server instance summary for this pod.
                      properties:
                        accessEndpoints:
                          description: AccessEndpoints are the access endpoints for this pod.
                          items:
                            type: string
                          type: array
                        alternateAccessEndpoints:
                          description: AlternateAccessEndpoints are the alternate access endpoints for this pod.
                          items:
                            type: string
                          type: array
                        clusterName:
                          description: ClusterName is the name of the Aerospike cluster this pod belongs to.
                          type: string
                        nodeID:
                          description: NodeID is the unique Aerospike ID for this pod.
                          type: string
                        rackID:
                          description: RackID of rack to which this node belongs
                          type: integer
                        tlsAccessEndpoints:
                          description: TLSAccessEndpoints are the TLS access endpoints for this pod.
                          items:
                            type: string
                          type: array
                        tlsAlternateAccessEndpoints:
                          description: TLSAlternateAccessEndpoints are the alternate TLS access endpoints for this pod.
                          items:
                            type: string
                          type: array
                        tlsName:
                          description: TLSName is the TLS name of this pod in the Aerospike cluster.
                          type: string
                      required:
                      - clusterName
                      - nodeID
                      type: object
                    aerospikeConfigHash:
                      description: AerospikeConfigHash is ripemd160 hash of aerospikeConfig used by this pod
                      type: string
                    hostExternalIP:
                      description: HostExternalIP of the K8s host this pod is scheduled on.
                      type: string
                    hostInternalIP:
                      description: HostInternalIP of the K8s host this pod is scheduled on.
                      type: string
                    image:
                      description: Image is the Aerospike image this pod is running.
                      type: string
                    initializedVolumePaths:
                      description: InitializedVolumePaths is the list of device path that have already been initial
                      items:
                        type: string
                      type: array
                    networkPolicyHash:
                      description: NetworkPolicyHash is ripemd160 hash of NetworkPolicy used by this pod
                      type: string
                    podIP:
                      description: PodIP in the K8s network.
                      type: string
                    podPort:
                      description: PodPort is the port K8s intenral Aerospike clients can connect to.
                      type: integer
                    podSpecHash:
                      description: PodSpecHash is ripemd160 hash of PodSpec used by this pod
                      type: string
                    servicePort:
                      description: ServicePort is the port Aerospike clients outside K8s can connect to.
                      format: int32
                      type: integer
                  required:
                  - aerospikeConfigHash
                  - image
                  - initializedVolumePaths
                  - networkPolicyHash
                  - podIP
                  - podPort
                  - podSpecHash
                  - servicePort
                  type: object
                description: Pods has Aerospike specific status of the pods.
                type: object
              rackConfig:
                description: RackConfig Configures the operator to deploy rack aware Aerospike cluster.
                nullable: true
                properties:
                  namespaces:
                    description: List of Aerospike namespaces for which rack feature will be enabled
                    items:
                      type: string
                    type: array
                  racks:
                    description: Racks is the list of all racks
                    items:
                      description: Rack specifies single rack config
                      properties:
                        aerospikeConfig:
                          description: AerospikeConfig overrides the common AerospikeConfig for this Rack.
                          type: object
                          x-kubernetes-preserve-unknown-fields: true
                        effectiveAerospikeConfig:
                          description: Effective/operative Aerospike config.
                          type: object
                          x-kubernetes-preserve-unknown-fields: true
                        effectivePodSpec:
                          description: Effective/operative PodSpec.
                          properties:
                            affinity:
                              description: Affinity rules for pod placement.
                              properties:
                                nodeAffinity:
                                  description: Describes node affinity scheduling rules for the pod.
                                  properties:
                                    preferredDuringSchedulingIgnoredDuringExecution:
                                      description: The scheduler will prefer to schedule pods to nodes that satisfy the affinity ex
                                      items:
                                        description: An empty preferred scheduling term matches all objects with implicit weight 0 (i
                                        properties:
                                          preference:
                                            description: A node selector term, associated with the corresponding weight.
                                            properties:
                                              matchExpressions:
                                                description: A list of node selector requirements by node's labels.
                                                items:
                                                  description: A node selector requirement is a selector that contains values, a key, and an op
                                                  properties:
                                                    key:
                                                      description: The label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: Represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: An array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchFields:
                                                description: A list of node selector requirements by node's fields.
                                                items:
                                                  description: A node selector requirement is a selector that contains values, a key, and an op
                                                  properties:
                                                    key:
                                                      description: The label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: Represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: An array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                            type: object
                                          weight:
                                            description: Weight associated with matching the corresponding nodeSelectorTerm, in the range
                                            format: int32
                                            type: integer
                                        required:
                                        - preference
                                        - weight
                                        type: object
                                      type: array
                                    requiredDuringSchedulingIgnoredDuringExecution:
                                      description: If the affinity requirements specified by this field are not met at scheduling t
                                      properties:
                                        nodeSelectorTerms:
                                          description: Required. A list of node selector terms. The terms are ORed.
                                          items:
                                            description: A null or empty node selector term matches no objects.
                                            properties:
                                              matchExpressions:
                                                description: A list of node selector requirements by node's labels.
                                                items:
                                                  description: A node selector requirement is a selector that contains values, a key, and an op
                                                  properties:
                                                    key:
                                                      description: The label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: Represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: An array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchFields:
                                                description: A list of node selector requirements by node's fields.
                                                items:
                                                  description: A node selector requirement is a selector that contains values, a key, and an op
                                                  properties:
                                                    key:
                                                      description: The label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: Represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: An array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                            type: object
                                          type: array
                                      required:
                                      - nodeSelectorTerms
                                      type: object
                                  type: object
                                podAffinity:
                                  description: Describes pod affinity scheduling rules (e.g.
                                  properties:
                                    preferredDuringSchedulingIgnoredDuringExecution:
                                      description: The scheduler will prefer to schedule pods to nodes that satisfy the affinity ex
                                      items:
                                        description: The weights of all of the matched WeightedPodAffinityTerm fields are added per-n
                                        properties:
                                          podAffinityTerm:
                                            description: Required. A pod affinity term, associated with the corresponding weight.
                                            properties:
                                              labelSelector:
                                                description: A label query over a set of resources, in this case pods.
                                                properties:
                                                  matchExpressions:
                                                    description: matchExpressions is a list of label selector requirements.
                                                    items:
                                                      description: A label selector requirement is a selector that contains values, a key, and an o
                                                      properties:
                                                        key:
                                                          description: key is the label key that the selector applies to.
                                                          type: string
                                                        operator:
                                                          description: operator represents a key's relationship to a set of values.
                                                          type: string
                                                        values:
                                                          description: values is an array of string values.
                                                          items:
                                                            type: string
                                                          type: array
                                                      required:
                                                      - key
                                                      - operator
                                                      type: object
                                                    type: array
                                                  matchLabels:
                                                    additionalProperties:
                                                      type: string
                                                    description: matchLabels is a map of {key,value} pairs.
                                                    type: object
                                                type: object
                                              namespaceSelector:
                                                description: A label query over the set of namespaces that the term applies to.
                                                properties:
                                                  matchExpressions:
                                                    description: matchExpressions is a list of label selector requirements.
                                                    items:
                                                      description: A label selector requirement is a selector that contains values, a key, and an o
                                                      properties:
                                                        key:
                                                          description: key is the label key that the selector applies to.
                                                          type: string
                                                        operator:
                                                          description: operator represents a key's relationship to a set of values.
                                                          type: string
                                                        values:
                                                          description: values is an array of string values.
                                                          items:
                                                            type: string
                                                          type: array
                                                      required:
                                                      - key
                                                      - operator
                                                      type: object
                                                    type: array
                                                  matchLabels:
                                                    additionalProperties:
                                                      type: string
                                                    description: matchLabels is a map of {key,value} pairs.
                                                    type: object
                                                type: object
                                              namespaces:
                                                description: namespaces specifies a static list of namespace names that the term applies to.
                                                items:
                                                  type: string
                                                type: array
                                              topologyKey:
                                                description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                                type: string
                                            required:
                                            - topologyKey
                                            type: object
                                          weight:
                                            description: 'weight associated with matching the corresponding podAffinityTerm, in the range '
                                            format: int32
                                            type: integer
                                        required:
                                        - podAffinityTerm
                                        - weight
                                        type: object
                                      type: array
                                    requiredDuringSchedulingIgnoredDuringExecution:
                                      description: If the affinity requirements specified by this field are not met at scheduling t
                                      items:
                                        description: Defines a set of pods (namely those matching the labelSelector relative to the g
                                        properties:
                                          labelSelector:
                                            description: A label query over a set of resources, in this case pods.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          namespaceSelector:
                                            description: A label query over the set of namespaces that the term applies to.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          namespaces:
                                            description: namespaces specifies a static list of namespace names that the term applies to.
                                            items:
                                              type: string
                                            type: array
                                          topologyKey:
                                            description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                            type: string
                                        required:
                                        - topologyKey
                                        type: object
                                      type: array
                                  type: object
                                podAntiAffinity:
                                  description: Describes pod anti-affinity scheduling rules (e.g.
                                  properties:
                                    preferredDuringSchedulingIgnoredDuringExecution:
                                      description: The scheduler will prefer to schedule pods to nodes that satisfy the anti-affini
                                      items:
                                        description: The weights of all of the matched WeightedPodAffinityTerm fields are added per-n
                                        properties:
                                          podAffinityTerm:
                                            description: Required. A pod affinity term, associated with the corresponding weight.
                                            properties:
                                              labelSelector:
                                                description: A label query over a set of resources, in this case pods.
                                                properties:
                                                  matchExpressions:
                                                    description: matchExpressions is a list of label selector requirements.
                                                    items:
                                                      description: A label selector requirement is a selector that contains values, a key, and an o
                                                      properties:
                                                        key:
                                                          description: key is the label key that the selector applies to.
                                                          type: string
                                                        operator:
                                                          description: operator represents a key's relationship to a set of values.
                                                          type: string
                                                        values:
                                                          description: values is an array of string values.
                                                          items:
                                                            type: string
                                                          type: array
                                                      required:
                                                      - key
                                                      - operator
                                                      type: object
                                                    type: array
                                                  matchLabels:
                                                    additionalProperties:
                                                      type: string
                                                    description: matchLabels is a map of {key,value} pairs.
                                                    type: object
                                                type: object
                                              namespaceSelector:
                                                description: A label query over the set of namespaces that the term applies to.
                                                properties:
                                                  matchExpressions:
                                                    description: matchExpressions is a list of label selector requirements.
                                                    items:
                                                      description: A label selector requirement is a selector that contains values, a key, and an o
                                                      properties:
                                                        key:
                                                          description: key is the label key that the selector applies to.
                                                          type: string
                                                        operator:
                                                          description: operator represents a key's relationship to a set of values.
                                                          type: string
                                                        values:
                                                          description: values is an array of string values.
                                                          items:
                                                            type: string
                                                          type: array
                                                      required:
                                                      - key
                                                      - operator
                                                      type: object
                                                    type: array
                                                  matchLabels:
                                                    additionalProperties:
                                                      type: string
                                                    description: matchLabels is a map of {key,value} pairs.
                                                    type: object
                                                type: object
                                              namespaces:
                                                description: namespaces specifies a static list of namespace names that the term applies to.
                                                items:
                                                  type: string
                                                type: array
                                              topologyKey:
                                                description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                                type: string
                                            required:
                                            - topologyKey
                                            type: object
                                          weight:
                                            description: 'weight associated with matching the corresponding podAffinityTerm, in the range '
                                            format: int32
                                            type: integer
                                        required:
                                        - podAffinityTerm
                                        - weight
                                        type: object
                                      type: array
                                    requiredDuringSchedulingIgnoredDuringExecution:
                                      description: If the anti-affinity requirements specified by this field are not met at schedul
                                      items:
                                        description: Defines a set of pods (namely those matching the labelSelector relative to the g
                                        properties:
                                          labelSelector:
                                            description: A label query over a set of resources, in this case pods.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          namespaceSelector:
                                            description: A label query over the set of namespaces that the term applies to.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          namespaces:
                                            description: namespaces specifies a static list of namespace names that the term applies to.
                                            items:
                                              type: string
                                            type: array
                                          topologyKey:
                                            description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                            type: string
                                        required:
                                        - topologyKey
                                        type: object
                                      type: array
                                  type: object
                              type: object
                            nodeSelector:
                              additionalProperties:
                                type: string
                              description: NodeSelector constraints for this pod.
                              type: object
                            tolerations:
                              description: Tolerations for this pod.
                              items:
                                description: The pod this Toleration is attached to tolerates any taint that matches the trip
                                properties:
                                  effect:
                                    description: Effect indicates the taint effect to match. Empty means match all taint effects.
                                    type: string
                                  key:
                                    description: Key is the taint key that the toleration applies to.
                                    type: string
                                  operator:
                                    description: Operator represents a key's relationship to the value.
                                    type: string
                                  tolerationSeconds:
                                    description: TolerationSeconds represents the period of time the toleration (which must be of
                                    format: int64
                                    type: integer
                                  value:
                                    description: Value is the taint value the toleration matches to.
                                    type: string
                                type: object
                              type: array
                          type: object
                        effectiveStorage:
                          description: Effective/operative storage.
                          properties:
                            blockVolumePolicy:
                              description: BlockVolumePolicy contains default policies for block volumes.
                              properties:
                                cascadeDelete:
                                  description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                                  type: boolean
                                effectiveCascadeDelete:
                                  description: Effective/operative value to use for cascade delete after applying defaults.
                                  type: boolean
                                effectiveInitMethod:
                                  description: Effective/operative value to use as the volume init method after applying defaul
                                  enum:
                                  - none
                                  - dd
                                  - blkdiscard
                                  - deleteFiles
                                  type: string
                                initMethod:
                                  description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                                  enum:
                                  - none
                                  - dd
                                  - blkdiscard
                                  - deleteFiles
                                  type: string
                              type: object
                            filesystemVolumePolicy:
                              description: FileSystemVolumePolicy contains default policies for filesystem volumes.
                              properties:
                                cascadeDelete:
                                  description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                                  type: boolean
                                effectiveCascadeDelete:
                                  description: Effective/operative value to use for cascade delete after applying defaults.
                                  type: boolean
                                effectiveInitMethod:
                                  description: Effective/operative value to use as the volume init method after applying defaul
                                  enum:
                                  - none
                                  - dd
                                  - blkdiscard
                                  - deleteFiles
                                  type: string
                                initMethod:
                                  description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                                  enum:
                                  - none
                                  - dd
                                  - blkdiscard
                                  - deleteFiles
                                  type: string
                              type: object
                            volumes:
                              description: Volumes list to attach to created pods.
                              items:
                                properties:
                                  aerospike:
                                    description: Aerospike attachment of this volume on Aerospike server container.
                                    properties:
                                      mountOptions:
                                        properties:
                                          mountPropagation:
                                            description: mountPropagation determines how mounts are propagated from the host to container
                                            type: string
                                          readOnly:
                                            description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                            type: boolean
                                          subPath:
                                            description: Path within the volume from which the container's volume should be mounted.
                                            type: string
                                          subPathExpr:
                                            description: Expanded path within the volume from which the container's volume should be moun
                                            type: string
                                        type: object
                                      path:
                                        description: Path to attach the volume on the Aerospike server container.
                                        type: string
                                    required:
                                    - path
                                    type: object
                                  cascadeDelete:
                                    description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                                    type: boolean
                                  effectiveCascadeDelete:
                                    description: Effective/operative value to use for cascade delete after applying defaults.
                                    type: boolean
                                  effectiveInitMethod:
                                    description: Effective/operative value to use as the volume init method after applying defaul
                                    enum:
                                    - none
                                    - dd
                                    - blkdiscard
                                    - deleteFiles
                                    type: string
                                  initContainers:
                                    description: InitContainers are additional init containers where this volume will be mounted
                                    items:
                                      description: VolumeAttachment specifies volume attachment to a container.
                                      properties:
                                        containerName:
                                          description: ContainerName is the name of the container to attach this volume to.
                                          type: string
                                        mountOptions:
                                          properties:
                                            mountPropagation:
                                              description: mountPropagation determines how mounts are propagated from the host to container
                                              type: string
                                            readOnly:
                                              description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                              type: boolean
                                            subPath:
                                              description: Path within the volume from which the container's volume should be mounted.
                                              type: string
                                            subPathExpr:
                                              description: Expanded path within the volume from which the container's volume should be moun
                                              type: string
                                          type: object
                                        path:
                                          description: Path to attache the volume on the container.
                                          type: string
                                      required:
                                      - containerName
                                      - path
                                      type: object
                                    type: array
                                  initMethod:
                                    description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                                    enum:
                                    - none
                                    - dd
                                    - blkdiscard
                                    - deleteFiles
                                    type: string
                                  name:
                                    description: Name for this volume, Name or path should be given.
                                    type: string
                                  sidecars:
                                    description: Sidecars are side containers where this volume will be mounted
                                    items:
                                      description: VolumeAttachment specifies volume attachment to a container.
                                      properties:
                                        containerName:
                                          description: ContainerName is the name of the container to attach this volume to.
                                          type: string
                                        mountOptions:
                                          properties:
                                            mountPropagation:
                                              description: mountPropagation determines how mounts are propagated from the host to container
                                              type: string
                                            readOnly:
                                              description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                              type: boolean
                                            subPath:
                                              description: Path within the volume from which the container's volume should be mounted.
                                              type: string
                                            subPathExpr:
                                              description: Expanded path within the volume from which the container's volume should be moun
                                              type: string
                                          type: object
                                        path:
                                          description: Path to attache the volume on the container.
                                          type: string
                                      required:
                                      - containerName
                                      - path
                                      type: object
                                    type: array
                                  source:
                                    description: Source of this volume.
                                    properties:
                                      configMap:
                                        description: ConfigMap represents a configMap that should populate this volume
                                        properties:
                                          defaultMode:
                                            description: 'Optional: mode bits used to set permissions on created files by default.'
                                            format: int32
                                            type: integer
                                          items:
                                            description: If unspecified, each key-value pair in the Data field of the referenced ConfigMa
                                            items:
                                              description: Maps a string key to a path within a volume.
                                              properties:
                                                key:
                                                  description: The key to project.
                                                  type: string
                                                mode:
                                                  description: 'Optional: mode bits used to set permissions on this file.'
                                                  format: int32
                                                  type: integer
                                                path:
                                                  description: The relative path of the file to map the key to. May not be an absolute path.
                                                  type: string
                                              required:
                                              - key
                                              - path
                                              type: object
                                            type: array
                                          name:
                                            description: 'Name of the referent. More info: https://kubernetes.'
                                            type: string
                                          optional:
                                            description: Specify whether the ConfigMap or its keys must be defined
                                            type: boolean
                                        type: object
                                      emptyDir:
                                        description: EmptyDir represents a temporary directory that shares a pod's lifetime.
                                        properties:
                                          medium:
                                            description: What type of storage medium should back this directory.
                                            type: string
                                          sizeLimit:
                                            anyOf:
                                            - type: integer
                                            - type: string
                                            description: Total amount of local storage required for this EmptyDir volume.
                                            pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                            x-kubernetes-int-or-string: true
                                        type: object
                                      persistentVolume:
                                        description: PersistentVolumeSpec describes a persistent volume to claim and attach to Aerosp
                                        properties:
                                          accessModes:
                                            items:
                                              type: string
                                            type: array
                                          selector:
                                            description: A label query over volumes to consider for binding.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          size:
                                            anyOf:
                                            - type: integer
                                            - type: string
                                            description: Size of volume.
                                            pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                            x-kubernetes-int-or-string: true
                                          storageClass:
                                            description: StorageClass should be pre-created by user.
                                            type: string
                                          volumeMode:
                                            description: VolumeMode specifies if the volume is block/raw or a filesystem.
                                            type: string
                                        required:
                                        - size
                                        - storageClass
                                        - volumeMode
                                        type: object
                                      secret:
                                        description: Adapts a Secret into a volume.
                                        properties:
                                          defaultMode:
                                            description: 'Optional: mode bits used to set permissions on created files by default.'
                                            format: int32
                                            type: integer
                                          items:
                                            description: If unspecified, each key-value pair in the Data field of the referenced Secret w
                                            items:
                                              description: Maps a string key to a path within a volume.
                                              properties:
                                                key:
                                                  description: The key to project.
                                                  type: string
                                                mode:
                                                  description: 'Optional: mode bits used to set permissions on this file.'
                                                  format: int32
                                                  type: integer
                                                path:
                                                  description: The relative path of the file to map the key to. May not be an absolute path.
                                                  type: string
                                              required:
                                              - key
                                              - path
                                              type: object
                                            type: array
                                          optional:
                                            description: Specify whether the Secret or its keys must be defined
                                            type: boolean
                                          secretName:
                                            description: 'Name of the secret in the pod''s namespace to use. More info: https://kubernetes.'
                                            type: string
                                        type: object
                                    type: object
                                required:
                                - name
                                type: object
                              type: array
                              x-kubernetes-list-map-keys:
                              - name
                              x-kubernetes-list-type: map
                          type: object
                        id:
                          description: Identifier for the rack
                          type: integer
                        nodeName:
                          description: K8s Node name for setting rack affinity.
                          type: string
                        podSpec:
                          description: PodSpec to use for the pods in this rack.
                          properties:
                            affinity:
                              description: Affinity rules for pod placement.
                              properties:
                                nodeAffinity:
                                  description: Describes node affinity scheduling rules for the pod.
                                  properties:
                                    preferredDuringSchedulingIgnoredDuringExecution:
                                      description: The scheduler will prefer to schedule pods to nodes that satisfy the affinity ex
                                      items:
                                        description: An empty preferred scheduling term matches all objects with implicit weight 0 (i
                                        properties:
                                          preference:
                                            description: A node selector term, associated with the corresponding weight.
                                            properties:
                                              matchExpressions:
                                                description: A list of node selector requirements by node's labels.
                                                items:
                                                  description: A node selector requirement is a selector that contains values, a key, and an op
                                                  properties:
                                                    key:
                                                      description: The label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: Represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: An array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchFields:
                                                description: A list of node selector requirements by node's fields.
                                                items:
                                                  description: A node selector requirement is a selector that contains values, a key, and an op
                                                  properties:
                                                    key:
                                                      description: The label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: Represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: An array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                            type: object
                                          weight:
                                            description: Weight associated with matching the corresponding nodeSelectorTerm, in the range
                                            format: int32
                                            type: integer
                                        required:
                                        - preference
                                        - weight
                                        type: object
                                      type: array
                                    requiredDuringSchedulingIgnoredDuringExecution:
                                      description: If the affinity requirements specified by this field are not met at scheduling t
                                      properties:
                                        nodeSelectorTerms:
                                          description: Required. A list of node selector terms. The terms are ORed.
                                          items:
                                            description: A null or empty node selector term matches no objects.
                                            properties:
                                              matchExpressions:
                                                description: A list of node selector requirements by node's labels.
                                                items:
                                                  description: A node selector requirement is a selector that contains values, a key, and an op
                                                  properties:
                                                    key:
                                                      description: The label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: Represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: An array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchFields:
                                                description: A list of node selector requirements by node's fields.
                                                items:
                                                  description: A node selector requirement is a selector that contains values, a key, and an op
                                                  properties:
                                                    key:
                                                      description: The label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: Represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: An array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                            type: object
                                          type: array
                                      required:
                                      - nodeSelectorTerms
                                      type: object
                                  type: object
                                podAffinity:
                                  description: Describes pod affinity scheduling rules (e.g.
                                  properties:
                                    preferredDuringSchedulingIgnoredDuringExecution:
                                      description: The scheduler will prefer to schedule pods to nodes that satisfy the affinity ex
                                      items:
                                        description: The weights of all of the matched WeightedPodAffinityTerm fields are added per-n
                                        properties:
                                          podAffinityTerm:
                                            description: Required. A pod affinity term, associated with the corresponding weight.
                                            properties:
                                              labelSelector:
                                                description: A label query over a set of resources, in this case pods.
                                                properties:
                                                  matchExpressions:
                                                    description: matchExpressions is a list of label selector requirements.
                                                    items:
                                                      description: A label selector requirement is a selector that contains values, a key, and an o
                                                      properties:
                                                        key:
                                                          description: key is the label key that the selector applies to.
                                                          type: string
                                                        operator:
                                                          description: operator represents a key's relationship to a set of values.
                                                          type: string
                                                        values:
                                                          description: values is an array of string values.
                                                          items:
                                                            type: string
                                                          type: array
                                                      required:
                                                      - key
                                                      - operator
                                                      type: object
                                                    type: array
                                                  matchLabels:
                                                    additionalProperties:
                                                      type: string
                                                    description: matchLabels is a map of {key,value} pairs.
                                                    type: object
                                                type: object
                                              namespaceSelector:
                                                description: A label query over the set of namespaces that the term applies to.
                                                properties:
                                                  matchExpressions:
                                                    description: matchExpressions is a list of label selector requirements.
                                                    items:
                                                      description: A label selector requirement is a selector that contains values, a key, and an o
                                                      properties:
                                                        key:
                                                          description: key is the label key that the selector applies to.
                                                          type: string
                                                        operator:
                                                          description: operator represents a key's relationship to a set of values.
                                                          type: string
                                                        values:
                                                          description: values is an array of string values.
                                                          items:
                                                            type: string
                                                          type: array
                                                      required:
                                                      - key
                                                      - operator
                                                      type: object
                                                    type: array
                                                  matchLabels:
                                                    additionalProperties:
                                                      type: string
                                                    description: matchLabels is a map of {key,value} pairs.
                                                    type: object
                                                type: object
                                              namespaces:
                                                description: namespaces specifies a static list of namespace names that the term applies to.
                                                items:
                                                  type: string
                                                type: array
                                              topologyKey:
                                                description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                                type: string
                                            required:
                                            - topologyKey
                                            type: object
                                          weight:
                                            description: 'weight associated with matching the corresponding podAffinityTerm, in the range '
                                            format: int32
                                            type: integer
                                        required:
                                        - podAffinityTerm
                                        - weight
                                        type: object
                                      type: array
                                    requiredDuringSchedulingIgnoredDuringExecution:
                                      description: If the affinity requirements specified by this field are not met at scheduling t
                                      items:
                                        description: Defines a set of pods (namely those matching the labelSelector relative to the g
                                        properties:
                                          labelSelector:
                                            description: A label query over a set of resources, in this case pods.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          namespaceSelector:
                                            description: A label query over the set of namespaces that the term applies to.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          namespaces:
                                            description: namespaces specifies a static list of namespace names that the term applies to.
                                            items:
                                              type: string
                                            type: array
                                          topologyKey:
                                            description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                            type: string
                                        required:
                                        - topologyKey
                                        type: object
                                      type: array
                                  type: object
                                podAntiAffinity:
                                  description: Describes pod anti-affinity scheduling rules (e.g.
                                  properties:
                                    preferredDuringSchedulingIgnoredDuringExecution:
                                      description: The scheduler will prefer to schedule pods to nodes that satisfy the anti-affini
                                      items:
                                        description: The weights of all of the matched WeightedPodAffinityTerm fields are added per-n
                                        properties:
                                          podAffinityTerm:
                                            description: Required. A pod affinity term, associated with the corresponding weight.
                                            properties:
                                              labelSelector:
                                                description: A label query over a set of resources, in this case pods.
                                                properties:
                                                  matchExpressions:
                                                    description: matchExpressions is a list of label selector requirements.
                                                    items:
                                                      description: A label selector requirement is a selector that contains values, a key, and an o
                                                      properties:
                                                        key:
                                                          description: key is the label key that the selector applies to.
                                                          type: string
                                                        operator:
                                                          description: operator represents a key's relationship to a set of values.
                                                          type: string
                                                        values:
                                                          description: values is an array of string values.
                                                          items:
                                                            type: string
                                                          type: array
                                                      required:
                                                      - key
                                                      - operator
                                                      type: object
                                                    type: array
                                                  matchLabels:
                                                    additionalProperties:
                                                      type: string
                                                    description: matchLabels is a map of {key,value} pairs.
                                                    type: object
                                                type: object
                                              namespaceSelector:
                                                description: A label query over the set of namespaces that the term applies to.
                                                properties:
                                                  matchExpressions:
                                                    description: matchExpressions is a list of label selector requirements.
                                                    items:
                                                      description: A label selector requirement is a selector that contains values, a key, and an o
                                                      properties:
                                                        key:
                                                          description: key is the label key that the selector applies to.
                                                          type: string
                                                        operator:
                                                          description: operator represents a key's relationship to a set of values.
                                                          type: string
                                                        values:
                                                          description: values is an array of string values.
                                                          items:
                                                            type: string
                                                          type: array
                                                      required:
                                                      - key
                                                      - operator
                                                      type: object
                                                    type: array
                                                  matchLabels:
                                                    additionalProperties:
                                                      type: string
                                                    description: matchLabels is a map of {key,value} pairs.
                                                    type: object
                                                type: object
                                              namespaces:
                                                description: namespaces specifies a static list of namespace names that the term applies to.
                                                items:
                                                  type: string
                                                type: array
                                              topologyKey:
                                                description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                                type: string
                                            required:
                                            - topologyKey
                                            type: object
                                          weight:
                                            description: 'weight associated with matching the corresponding podAffinityTerm, in the range '
                                            format: int32
                                            type: integer
                                        required:
                                        - podAffinityTerm
                                        - weight
                                        type: object
                                      type: array
                                    requiredDuringSchedulingIgnoredDuringExecution:
                                      description: If the anti-affinity requirements specified by this field are not met at schedul
                                      items:
                                        description: Defines a set of pods (namely those matching the labelSelector relative to the g
                                        properties:
                                          labelSelector:
                                            description: A label query over a set of resources, in this case pods.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          namespaceSelector:
                                            description: A label query over the set of namespaces that the term applies to.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          namespaces:
                                            description: namespaces specifies a static list of namespace names that the term applies to.
                                            items:
                                              type: string
                                            type: array
                                          topologyKey:
                                            description: 'This pod should be co-located (affinity) or not co-located (anti-affinity) with '
                                            type: string
                                        required:
                                        - topologyKey
                                        type: object
                                      type: array
                                  type: object
                              type: object
                            nodeSelector:
                              additionalProperties:
                                type: string
                              description: NodeSelector constraints for this pod.
                              type: object
                            tolerations:
                              description: Tolerations for this pod.
                              items:
                                description: The pod this Toleration is attached to tolerates any taint that matches the trip
                                properties:
                                  effect:
                                    description: Effect indicates the taint effect to match. Empty means match all taint effects.
                                    type: string
                                  key:
                                    description: Key is the taint key that the toleration applies to.
                                    type: string
                                  operator:
                                    description: Operator represents a key's relationship to the value.
                                    type: string
                                  tolerationSeconds:
                                    description: TolerationSeconds represents the period of time the toleration (which must be of
                                    format: int64
                                    type: integer
                                  value:
                                    description: Value is the taint value the toleration matches to.
                                    type: string
                                type: object
                              type: array
                          type: object
                        rackLabel:
                          description: RackLabel for setting rack affinity.
                          type: string
                        region:
                          description: Region name for setting rack affinity.
                          type: string
                        storage:
                          description: Storage specify persistent storage to use for the pods in this rack.
                          properties:
                            blockVolumePolicy:
                              description: BlockVolumePolicy contains default policies for block volumes.
                              properties:
                                cascadeDelete:
                                  description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                                  type: boolean
                                effectiveCascadeDelete:
                                  description: Effective/operative value to use for cascade delete after applying defaults.
                                  type: boolean
                                effectiveInitMethod:
                                  description: Effective/operative value to use as the volume init method after applying defaul
                                  enum:
                                  - none
                                  - dd
                                  - blkdiscard
                                  - deleteFiles
                                  type: string
                                initMethod:
                                  description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                                  enum:
                                  - none
                                  - dd
                                  - blkdiscard
                                  - deleteFiles
                                  type: string
                              type: object
                            filesystemVolumePolicy:
                              description: FileSystemVolumePolicy contains default policies for filesystem volumes.
                              properties:
                                cascadeDelete:
                                  description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                                  type: boolean
                                effectiveCascadeDelete:
                                  description: Effective/operative value to use for cascade delete after applying defaults.
                                  type: boolean
                                effectiveInitMethod:
                                  description: Effective/operative value to use as the volume init method after applying defaul
                                  enum:
                                  - none
                                  - dd
                                  - blkdiscard
                                  - deleteFiles
                                  type: string
                                initMethod:
                                  description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                                  enum:
                                  - none
                                  - dd
                                  - blkdiscard
                                  - deleteFiles
                                  type: string
                              type: object
                            volumes:
                              description: Volumes list to attach to created pods.
                              items:
                                properties:
                                  aerospike:
                                    description: Aerospike attachment of this volume on Aerospike server container.
                                    properties:
                                      mountOptions:
                                        properties:
                                          mountPropagation:
                                            description: mountPropagation determines how mounts are propagated from the host to container
                                            type: string
                                          readOnly:
                                            description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                            type: boolean
                                          subPath:
                                            description: Path within the volume from which the container's volume should be mounted.
                                            type: string
                                          subPathExpr:
                                            description: Expanded path within the volume from which the container's volume should be moun
                                            type: string
                                        type: object
                                      path:
                                        description: Path to attach the volume on the Aerospike server container.
                                        type: string
                                    required:
                                    - path
                                    type: object
                                  cascadeDelete:
                                    description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                                    type: boolean
                                  effectiveCascadeDelete:
                                    description: Effective/operative value to use for cascade delete after applying defaults.
                                    type: boolean
                                  effectiveInitMethod:
                                    description: Effective/operative value to use as the volume init method after applying defaul
                                    enum:
                                    - none
                                    - dd
                                    - blkdiscard
                                    - deleteFiles
                                    type: string
                                  initContainers:
                                    description: InitContainers are additional init containers where this volume will be mounted
                                    items:
                                      description: VolumeAttachment specifies volume attachment to a container.
                                      properties:
                                        containerName:
                                          description: ContainerName is the name of the container to attach this volume to.
                                          type: string
                                        mountOptions:
                                          properties:
                                            mountPropagation:
                                              description: mountPropagation determines how mounts are propagated from the host to container
                                              type: string
                                            readOnly:
                                              description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                              type: boolean
                                            subPath:
                                              description: Path within the volume from which the container's volume should be mounted.
                                              type: string
                                            subPathExpr:
                                              description: Expanded path within the volume from which the container's volume should be moun
                                              type: string
                                          type: object
                                        path:
                                          description: Path to attache the volume on the container.
                                          type: string
                                      required:
                                      - containerName
                                      - path
                                      type: object
                                    type: array
                                  initMethod:
                                    description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                                    enum:
                                    - none
                                    - dd
                                    - blkdiscard
                                    - deleteFiles
                                    type: string
                                  name:
                                    description: Name for this volume, Name or path should be given.
                                    type: string
                                  sidecars:
                                    description: Sidecars are side containers where this volume will be mounted
                                    items:
                                      description: VolumeAttachment specifies volume attachment to a container.
                                      properties:
                                        containerName:
                                          description: ContainerName is the name of the container to attach this volume to.
                                          type: string
                                        mountOptions:
                                          properties:
                                            mountPropagation:
                                              description: mountPropagation determines how mounts are propagated from the host to container
                                              type: string
                                            readOnly:
                                              description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                              type: boolean
                                            subPath:
                                              description: Path within the volume from which the container's volume should be mounted.
                                              type: string
                                            subPathExpr:
                                              description: Expanded path within the volume from which the container's volume should be moun
                                              type: string
                                          type: object
                                        path:
                                          description: Path to attache the volume on the container.
                                          type: string
                                      required:
                                      - containerName
                                      - path
                                      type: object
                                    type: array
                                  source:
                                    description: Source of this volume.
                                    properties:
                                      configMap:
                                        description: ConfigMap represents a configMap that should populate this volume
                                        properties:
                                          defaultMode:
                                            description: 'Optional: mode bits used to set permissions on created files by default.'
                                            format: int32
                                            type: integer
                                          items:
                                            description: If unspecified, each key-value pair in the Data field of the referenced ConfigMa
                                            items:
                                              description: Maps a string key to a path within a volume.
                                              properties:
                                                key:
                                                  description: The key to project.
                                                  type: string
                                                mode:
                                                  description: 'Optional: mode bits used to set permissions on this file.'
                                                  format: int32
                                                  type: integer
                                                path:
                                                  description: The relative path of the file to map the key to. May not be an absolute path.
                                                  type: string
                                              required:
                                              - key
                                              - path
                                              type: object
                                            type: array
                                          name:
                                            description: 'Name of the referent. More info: https://kubernetes.'
                                            type: string
                                          optional:
                                            description: Specify whether the ConfigMap or its keys must be defined
                                            type: boolean
                                        type: object
                                      emptyDir:
                                        description: EmptyDir represents a temporary directory that shares a pod's lifetime.
                                        properties:
                                          medium:
                                            description: What type of storage medium should back this directory.
                                            type: string
                                          sizeLimit:
                                            anyOf:
                                            - type: integer
                                            - type: string
                                            description: Total amount of local storage required for this EmptyDir volume.
                                            pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                            x-kubernetes-int-or-string: true
                                        type: object
                                      persistentVolume:
                                        description: PersistentVolumeSpec describes a persistent volume to claim and attach to Aerosp
                                        properties:
                                          accessModes:
                                            items:
                                              type: string
                                            type: array
                                          selector:
                                            description: A label query over volumes to consider for binding.
                                            properties:
                                              matchExpressions:
                                                description: matchExpressions is a list of label selector requirements.
                                                items:
                                                  description: A label selector requirement is a selector that contains values, a key, and an o
                                                  properties:
                                                    key:
                                                      description: key is the label key that the selector applies to.
                                                      type: string
                                                    operator:
                                                      description: operator represents a key's relationship to a set of values.
                                                      type: string
                                                    values:
                                                      description: values is an array of string values.
                                                      items:
                                                        type: string
                                                      type: array
                                                  required:
                                                  - key
                                                  - operator
                                                  type: object
                                                type: array
                                              matchLabels:
                                                additionalProperties:
                                                  type: string
                                                description: matchLabels is a map of {key,value} pairs.
                                                type: object
                                            type: object
                                          size:
                                            anyOf:
                                            - type: integer
                                            - type: string
                                            description: Size of volume.
                                            pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                            x-kubernetes-int-or-string: true
                                          storageClass:
                                            description: StorageClass should be pre-created by user.
                                            type: string
                                          volumeMode:
                                            description: VolumeMode specifies if the volume is block/raw or a filesystem.
                                            type: string
                                        required:
                                        - size
                                        - storageClass
                                        - volumeMode
                                        type: object
                                      secret:
                                        description: Adapts a Secret into a volume.
                                        properties:
                                          defaultMode:
                                            description: 'Optional: mode bits used to set permissions on created files by default.'
                                            format: int32
                                            type: integer
                                          items:
                                            description: If unspecified, each key-value pair in the Data field of the referenced Secret w
                                            items:
                                              description: Maps a string key to a path within a volume.
                                              properties:
                                                key:
                                                  description: The key to project.
                                                  type: string
                                                mode:
                                                  description: 'Optional: mode bits used to set permissions on this file.'
                                                  format: int32
                                                  type: integer
                                                path:
                                                  description: The relative path of the file to map the key to. May not be an absolute path.
                                                  type: string
                                              required:
                                              - key
                                              - path
                                              type: object
                                            type: array
                                          optional:
                                            description: Specify whether the Secret or its keys must be defined
                                            type: boolean
                                          secretName:
                                            description: 'Name of the secret in the pod''s namespace to use. More info: https://kubernetes.'
                                            type: string
                                        type: object
                                    type: object
                                required:
                                - name
                                type: object
                              type: array
                              x-kubernetes-list-map-keys:
                              - name
                              x-kubernetes-list-type: map
                          type: object
                        zone:
                          description: Zone name for setting rack affinity. Rack pods will be deployed to given Zone
                          type: string
                      required:
                      - id
                      type: object
                    nullable: true
                    type: array
                type: object
              resources:
                description: Define resources requests and limits for Aerospike Server Container.
                properties:
                  limits:
                    additionalProperties:
                      anyOf:
                      - type: integer
                      - type: string
                      pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                      x-kubernetes-int-or-string: true
                    description: Limits describes the maximum amount of compute resources allowed.
                    type: object
                  requests:
                    additionalProperties:
                      anyOf:
                      - type: integer
                      - type: string
                      pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                      x-kubernetes-int-or-string: true
                    description: Requests describes the minimum amount of compute resources required.
                    type: object
                type: object
              seedsFinderServices:
                description: SeedsFinderServices describes services which are used for seeding Aerospike node
                properties:
                  loadBalancer:
                    description: LoadBalancer configuration.
                    properties:
                      annotations:
                        additionalProperties:
                          type: string
                        type: object
                      externalTrafficPolicy:
                        description: Service External Traffic Policy Type string
                        enum:
                        - Local
                        - Cluster
                        type: string
                      loadBalancerSourceRanges:
                        items:
                          type: string
                        type: array
                      port:
                        description: Port Exposed port on load balancer. If not specified TargetPort is used.
                        format: int32
                        maximum: 65535
                        minimum: 1024
                        type: integer
                      targetPort:
                        description: TargetPort Target port. If not specified the tls-port of network.
                        format: int32
                        maximum: 65535
                        minimum: 1024
                        type: integer
                    type: object
                type: object
              size:
                description: Aerospike cluster size
                format: int32
                type: integer
              storage:
                description: Storage specify persistent storage to use for the Aerospike pods.
                properties:
                  blockVolumePolicy:
                    description: BlockVolumePolicy contains default policies for block volumes.
                    properties:
                      cascadeDelete:
                        description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                        type: boolean
                      effectiveCascadeDelete:
                        description: Effective/operative value to use for cascade delete after applying defaults.
                        type: boolean
                      effectiveInitMethod:
                        description: Effective/operative value to use as the volume init method after applying defaul
                        enum:
                        - none
                        - dd
                        - blkdiscard
                        - deleteFiles
                        type: string
                      initMethod:
                        description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                        enum:
                        - none
                        - dd
                        - blkdiscard
                        - deleteFiles
                        type: string
                    type: object
                  filesystemVolumePolicy:
                    description: FileSystemVolumePolicy contains default policies for filesystem volumes.
                    properties:
                      cascadeDelete:
                        description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                        type: boolean
                      effectiveCascadeDelete:
                        description: Effective/operative value to use for cascade delete after applying defaults.
                        type: boolean
                      effectiveInitMethod:
                        description: Effective/operative value to use as the volume init method after applying defaul
                        enum:
                        - none
                        - dd
                        - blkdiscard
                        - deleteFiles
                        type: string
                      initMethod:
                        description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                        enum:
                        - none
                        - dd
                        - blkdiscard
                        - deleteFiles
                        type: string
                    type: object
                  volumes:
                    description: Volumes list to attach to created pods.
                    items:
                      properties:
                        aerospike:
                          description: Aerospike attachment of this volume on Aerospike server container.
                          properties:
                            mountOptions:
                              properties:
                                mountPropagation:
                                  description: mountPropagation determines how mounts are propagated from the host to container
                                  type: string
                                readOnly:
                                  description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                  type: boolean
                                subPath:
                                  description: Path within the volume from which the container's volume should be mounted.
                                  type: string
                                subPathExpr:
                                  description: Expanded path within the volume from which the container's volume should be moun
                                  type: string
                              type: object
                            path:
                              description: Path to attach the volume on the Aerospike server container.
                              type: string
                          required:
                          - path
                          type: object
                        cascadeDelete:
                          description: CascadeDelete determines if the persistent volumes are deleted after the pod thi
                          type: boolean
                        effectiveCascadeDelete:
                          description: Effective/operative value to use for cascade delete after applying defaults.
                          type: boolean
                        effectiveInitMethod:
                          description: Effective/operative value to use as the volume init method after applying defaul
                          enum:
                          - none
                          - dd
                          - blkdiscard
                          - deleteFiles
                          type: string
                        initContainers:
                          description: InitContainers are additional init containers where this volume will be mounted
                          items:
                            description: VolumeAttachment specifies volume attachment to a container.
                            properties:
                              containerName:
                                description: ContainerName is the name of the container to attach this volume to.
                                type: string
                              mountOptions:
                                properties:
                                  mountPropagation:
                                    description: mountPropagation determines how mounts are propagated from the host to container
                                    type: string
                                  readOnly:
                                    description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                    type: boolean
                                  subPath:
                                    description: Path within the volume from which the container's volume should be mounted.
                                    type: string
                                  subPathExpr:
                                    description: Expanded path within the volume from which the container's volume should be moun
                                    type: string
                                type: object
                              path:
                                description: Path to attache the volume on the container.
                                type: string
                            required:
                            - containerName
                            - path
                            type: object
                          type: array
                        initMethod:
                          description: InitMethod determines how volumes attached to Aerospike server pods are initiali
                          enum:
                          - none
                          - dd
                          - blkdiscard
                          - deleteFiles
                          type: string
                        name:
                          description: Name for this volume, Name or path should be given.
                          type: string
                        sidecars:
                          description: Sidecars are side containers where this volume will be mounted
                          items:
                            description: VolumeAttachment specifies volume attachment to a container.
                            properties:
                              containerName:
                                description: ContainerName is the name of the container to attach this volume to.
                                type: string
                              mountOptions:
                                properties:
                                  mountPropagation:
                                    description: mountPropagation determines how mounts are propagated from the host to container
                                    type: string
                                  readOnly:
                                    description: Mounted read-only if true, read-write otherwise (false or unspecified).
                                    type: boolean
                                  subPath:
                                    description: Path within the volume from which the container's volume should be mounted.
                                    type: string
                                  subPathExpr:
                                    description: Expanded path within the volume from which the container's volume should be moun
                                    type: string
                                type: object
                              path:
                                description: Path to attache the volume on the container.
                                type: string
                            required:
                            - containerName
                            - path
                            type: object
                          type: array
                        source:
                          description: Source of this volume.
                          properties:
                            configMap:
                              description: ConfigMap represents a configMap that should populate this volume
                              properties:
                                defaultMode:
                                  description: 'Optional: mode bits used to set permissions on created files by default.'
                                  format: int32
                                  type: integer
                                items:
                                  description: If unspecified, each key-value pair in the Data field of the referenced ConfigMa
                                  items:
                                    description: Maps a string key to a path within a volume.
                                    properties:
                                      key:
                                        description: The key to project.
                                        type: string
                                      mode:
                                        description: 'Optional: mode bits used to set permissions on this file.'
                                        format: int32
                                        type: integer
                                      path:
                                        description: The relative path of the file to map the key to. May not be an absolute path.
                                        type: string
                                    required:
                                    - key
                                    - path
                                    type: object
                                  type: array
                                name:
                                  description: 'Name of the referent. More info: https://kubernetes.'
                                  type: string
                                optional:
                                  description: Specify whether the ConfigMap or its keys must be defined
                                  type: boolean
                              type: object
                            emptyDir:
                              description: EmptyDir represents a temporary directory that shares a pod's lifetime.
                              properties:
                                medium:
                                  description: What type of storage medium should back this directory.
                                  type: string
                                sizeLimit:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Total amount of local storage required for this EmptyDir volume.
                                  pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                  x-kubernetes-int-or-string: true
                              type: object
                            persistentVolume:
                              description: PersistentVolumeSpec describes a persistent volume to claim and attach to Aerosp
                              properties:
                                accessModes:
                                  items:
                                    type: string
                                  type: array
                                selector:
                                  description: A label query over volumes to consider for binding.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label selector requirements.
                                      items:
                                        description: A label selector requirement is a selector that contains values, a key, and an o
                                        properties:
                                          key:
                                            description: key is the label key that the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's relationship to a set of values.
                                            type: string
                                          values:
                                            description: values is an array of string values.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value} pairs.
                                      type: object
                                  type: object
                                size:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  description: Size of volume.
                                  pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                  x-kubernetes-int-or-string: true
                                storageClass:
                                  description: StorageClass should be pre-created by user.
                                  type: string
                                volumeMode:
                                  description: VolumeMode specifies if the volume is block/raw or a filesystem.
                                  type: string
                              required:
                              - size
                              - storageClass
                              - volumeMode
                              type: object
                            secret:
                              description: Adapts a Secret into a volume.
                              properties:
                                defaultMode:
                                  description: 'Optional: mode bits used to set permissions on created files by default.'
                                  format: int32
                                  type: integer
                                items:
                                  description: If unspecified, each key-value pair in the Data field of the referenced Secret w
                                  items:
                                    description: Maps a string key to a path within a volume.
                                    properties:
                                      key:
                                        description: The key to project.
                                        type: string
                                      mode:
                                        description: 'Optional: mode bits used to set permissions on this file.'
                                        format: int32
                                        type: integer
                                      path:
                                        description: The relative path of the file to map the key to. May not be an absolute path.
                                        type: string
                                    required:
                                    - key
                                    - path
                                    type: object
                                  type: array
                                optional:
                                  description: Specify whether the Secret or its keys must be defined
                                  type: boolean
                                secretName:
                                  description: 'Name of the secret in the pod''s namespace to use. More info: https://kubernetes.'
                                  type: string
                              type: object
                          type: object
                      required:
                      - name
                      type: object
                    type: array
                    x-kubernetes-list-map-keys:
                    - name
                    x-kubernetes-list-type: map
                type: object
              validationPolicy:
                description: ValidationPolicy controls validation of the Aerospike cluster resource.
                properties:
                  skipWorkDirValidate:
                    description: skipWorkDirValidate validates that Aerospike work directory is mounted on a pers
                    type: boolean
                  skipXdrDlogFileValidate:
                    description: ValidateXdrDigestLogFile validates that xdr digest log file is mounted on a pers
                    type: boolean
                required:
                - skipWorkDirValidate
                - skipXdrDlogFileValidate
                type: object
            required:
            - pods
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: aerospike-operator-controller-manager
  namespace: aerospike
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: aerospike-operator-leader-election-role
  namespace: aerospike
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - coordination.k8s.io
  resources:
  - leases
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
  - patch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: aerospike-operator-aerospikecluster-editor-role
rules:
- apiGroups:
  - asdb.aerospike.com
  resources:
  - aerospikeclusters
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - asdb.aerospike.com
  resources:
  - aerospikeclusters/status
  verbs:
  - get
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: aerospike-operator-aerospikecluster-viewer-role
rules:
- apiGroups:
  - asdb.aerospike.com
  resources:
  - aerospikeclusters
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - asdb.aerospike.com
  resources:
  - aerospikeclusters/status
  verbs:
  - get
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  creationTimestamp: null
  name: aerospike-operator-manager-role
rules:
- apiGroups:
  - apps
  resources:
  - statefulsets
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - asdb.aerospike.com
  resources:
  - aerospikeclusters
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - asdb.aerospike.com
  resources:
  - aerospikeclusters/finalizers
  verbs:
  - update
- apiGroups:
  - asdb.aerospike.com
  resources:
  - aerospikeclusters/status
  verbs:
  - get
  - patch
  - update
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - ""
  resources:
  - nodes
  verbs:
  - get
  - list
- apiGroups:
  - ""
  resources:
  - persistentvolumeclaims
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - ""
  resources:
  - pods/exec
  verbs:
  - create
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - services
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: aerospike-operator-metrics-reader
rules:
- nonResourceURLs:
  - /metrics
  verbs:
  - get
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: aerospike-operator-proxy-role
rules:
- apiGroups:
  - authentication.k8s.io
  resources:
  - tokenreviews
  verbs:
  - create
- apiGroups:
  - authorization.k8s.io
  resources:
  - subjectaccessreviews
  verbs:
  - create
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: aerospike-operator-leader-election-rolebinding
  namespace: aerospike
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: aerospike-operator-leader-election-role
subjects:
- kind: ServiceAccount
  name: aerospike-operator-controller-manager
  namespace: aerospike
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: aerospike-operator-aerospikecluster-editor-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: aerospike-operator-aerospikecluster-editor-role
subjects:
- kind: ServiceAccount
  name: aerospike-operator-controller-manager
  namespace: aerospike
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: aerospike-operator-aerospikecluster-viewer-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: aerospike-operator-aerospikecluster-viewer-role
subjects:
- kind: ServiceAccount
  name: aerospike-operator-controller-manager
  namespace: aerospike
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: aerospike-operator-manager-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: aerospike-operator-manager-role
subjects:
- kind: ServiceAccount
  name: aerospike-operator-controller-manager
  namespace: aerospike
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: aerospike-operator-proxy-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: aerospike-operator-proxy-role
subjects:
- kind: ServiceAccount
  name: aerospike-operator-controller-manager
  namespace: aerospike
---
apiVersion: v1
data:
  controller_manager_config.yaml: |
    apiVersion: controller-runtime.sigs.k8s.io/v1alpha1
    kind: ControllerManagerConfig
    health:
      healthProbeBindAddress: :8081
    metrics:
      bindAddress: 127.0.0.1:8080
    webhook:
      port: 9443
    leaderElection:
      leaderElect: true
      resourceName: 96242fdf.aerospike.com
kind: ConfigMap
metadata:
  name: aerospike-operator-manager-config
  namespace: aerospike
---
apiVersion: v1
kind: Service
metadata:
  labels:
    control-plane: controller-manager
  name: aerospike-operator-controller-manager-metrics-service
  namespace: aerospike
spec:
  ports:
  - name: https
    port: 8443
    protocol: TCP
    targetPort: https
  selector:
    control-plane: controller-manager
---
apiVersion: v1
kind: Service
metadata:
  name: aerospike-operator-webhook-service
  namespace: aerospike
spec:
  ports:
  - port: 443
    protocol: TCP
    targetPort: 9443
  selector:
    control-plane: controller-manager
---
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: controller-manager
  name: aerospike-operator-controller-manager
  namespace: aerospike
spec:
  replicas: 1
  selector:
    matchLabels:
      control-plane: controller-manager
  template:
    metadata:
      labels:
        control-plane: controller-manager
    spec:
      containers:
      - args:
        - --secure-listen-address=0.0.0.0:8443
        - --upstream=http://127.0.0.1:8080/
        - --logtostderr=true
        - --v=10
        image: gcr.io/kubebuilder/kube-rbac-proxy:v0.5.0
        name: kube-rbac-proxy
        ports:
        - containerPort: 8443
          name: https
          protocol: TCP
      - args:
        - --health-probe-bind-address=:8081
        - --metrics-bind-address=127.0.0.1:8080
        - --leader-elect
        command:
        - /manager
        env:
        - name: WATCH_NAMESPACE
          value: aerospike
        image: public.ecr.aws/t4x6l9m5/aerospike-kubernetes-operator:2.0.0
        imagePullPolicy: Always
        livenessProbe:
          httpGet:
            path: /healthz
            port: 8081
          initialDelaySeconds: 15
          periodSeconds: 20
        name: manager
        ports:
        - containerPort: 9443
          name: webhook-server
          protocol: TCP
        readinessProbe:
          httpGet:
            path: /readyz
            port: 8081
          initialDelaySeconds: 5
          periodSeconds: 10
        securityContext:
          allowPrivilegeEscalation: false
        volumeMounts:
        - mountPath: /tmp/k8s-webhook-server/serving-certs
          name: cert
          readOnly: true
      securityContext:
        runAsUser: 65532
      serviceAccountName: aerospike-operator-controller-manager
      terminationGracePeriodSeconds: 10
      volumes:
      - name: cert
        secret:
          defaultMode: 420
          secretName: webhook-server-cert
---
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: aerospike-operator-serving-cert
  namespace: aerospike
spec:
  dnsNames:
  - aerospike-operator-webhook-service.aerospike.svc
  - aerospike-operator-webhook-service.aerospike.svc.cluster.local
  issuerRef:
    kind: Issuer
    name: aerospike-operator-selfsigned-issuer
  secretName: webhook-server-cert
---
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: aerospike-operator-selfsigned-issuer
  namespace: aerospike
spec:
  selfSigned: {}
---
apiVersion: admissionregistration.k8s.io/v1
kind: MutatingWebhookConfiguration
metadata:
  annotations:
    cert-manager.io/inject-ca-from: aerospike/aerospike-operator-serving-cert
  name: aerospike-operator-mutating-webhook-configuration
webhooks:
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: aerospike-operator-webhook-service
      namespace: aerospike
      path: /mutate-asdb-aerospike-com-v1beta1-aerospikecluster
  failurePolicy: Fail
  name: maerospikecluster.kb.io
  rules:
  - apiGroups:
    - asdb.aerospike.com
    apiVersions:
    - v1beta1
    operations:
    - CREATE
    - UPDATE
    resources:
    - aerospikeclusters
  sideEffects: None
---
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingWebhookConfiguration
metadata:
  annotations:
    cert-manager.io/inject-ca-from: aerospike/aerospike-operator-serving-cert
  name: aerospike-operator-validating-webhook-configuration
webhooks:
- admissionReviewVersions:
  - v1
  - v1beta1
  clientConfig:
    service:
      name: aerospike-operator-webhook-service
      namespace: aerospike
      path: /validate-asdb-aerospike-com-v1beta1-aerospikecluster
  failurePolicy: Fail
  name: vaerospikecluster.kb.io
  rules:
  - apiGroups:
    - asdb.aerospike.com
    apiVersions:
    - v1beta1
    operations:
    - CREATE
    - UPDATE
    resources:
    - aerospikeclusters
  sideEffects: None
`
