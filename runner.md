## Gitlab CI/CD Pipelines

https://docs.gitlab.com/ee/ci/triggers/
Trigger pipeline on creation of Release branch (Deploy stage)

- Trigger token to trigger branch
- CI/CD job token to trigger multi-project pipeline

Using Directed Acyclic Graphs (DAG)

https://docs.gitlab.com/ee/ci/pipelines/

---

## What is Pipelines

### Jobs:

Jobs are the basic configuration component.

### Stages:

define stages from : build - test - deploy

Gitlab Docker

```
docker run -d --name gitlab-runner --restart always \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v /srv/gitlab-runner/config:/etc/gitlab-runner \
  gitlab/gitlab-runner:latest
```
