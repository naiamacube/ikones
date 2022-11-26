# ikones

* Dockerfiles for projects
* 1 image = 1 UK town name

1. `chelmsford` is the base docker-build image w/ git + golang. #build
2. `elland` contains Python 3.9.x. PyTorch 1.13.1 + Lightning 1.7.7. CPU-only. #dev
3. `longor` contains a vanilla Terraform 1.3.5 install. #dev

TODO: add a `dev` branch. Make sure all code in `main` pass the Go tests.
