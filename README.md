# kustomize-job-name-generator

Kustomize does not allow, by default, hashing the contents of jobs in their names
or the use of generateName:

- https://github.com/kubernetes-sigs/kustomize/issues/168
- https://github.com/kubernetes-sigs/kustomize/issues/641

This poses problems because Jobs in Kubernetes are largely immutable so if you
wanted to use Jobs for things like DB migrations or other kinds of preflight
tasks you basically can't use Kustomize.

The proposed solutions from Kustomize are to use the TTL controller to delete
the job immediately after success but this conflicts with popular open source
GitOps tools like ArgoCD which will keep syncing to recreate the job infinitely.

This plugin allows you to write a Job which will then have a name whose suffix
is a hash of its contents. This means that if the job changes you get a new job
and if it doesn't it simply applies cleanly (i.e. unchanged) to the existing
job preventing duplicate runs.