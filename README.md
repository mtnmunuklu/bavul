# Bavul <sub><small><small>(Be aware of vulnerabilities)</small></small></sub>
Bavul is a project that collects and stores vulnerability information. It utilizes **Golang**, **Docker**, **Kubernetes**, **Hugo-site**, **Traefik**, **grpc-go**, and **Protocol Buffers** technologies.

## Table of Contents

* [Features](#features)
* [Setup](#setup)
* [Usage](#usage)
* [Sponsors](#sponsors)
* [License](#license)

## Features

Bavul offers the following features:

- Collect and store vulnerability information from various sources
- API service for retrieving vulnerability information
- Authentication service for user access and management

## Setup

To set up bavul, follow these steps:

1. Download the latest version:

    ```
    LATEST_VERSION=$(wget -qO - https://api.github.com/repos/mtnmunuklu/bavul/releases/latest \
    | grep tag_name \
    | cut -d  '"' -f 4)

    curl -LJO https://github.com/mtnmunuklu/bavul/archive/refs/tags/$LATEST_VERSION.tar.gz
    ```

2. Extract the downloaded file:

    ```
    FILE_NAME=bavul-$(echo $LATEST_VERSION | cut -d 'v' -f 2)
    tar -xvf $FILE_NAME.tar.gz
    ```

3. Execute the setup scripts:

    ```
    cd $FILE_NAME/scripts
    # Execute on worker and control plane servers.
    bash tools/setup_tools.sh
    bash k8s/setup_k8s.sh
    # Execute only on the first control plane server.
    # It will create setup_k8s_control_plane.sh and setup_k8s_worker.sh files.
    # Control plane and worker scripts are for joining the Kubernetes cluster.
    # You can use these scripts on new nodes when you add new nodes as control plane or worker.
    bash k8s/setup_k8s_first_control_plane.sh
    # Execute only on first control plane server.
    bash setup_bavul.sh
    ```

## Usage

Bavul consists of 2 different services: [authentication](authentication) and [api](api). All incoming requests are first forwarded to the API service. Afterwards, the API service decides to which service the incoming request will be forwarded. The requested URL plays a role in the decision-making process.

To understand the features of each service, the available endpoints, how to make requests, and the expected responses, refer to the [swagger.yml](docs/api/swagger.yml) file under the `docs` folder.

You can also access the documents describing the software structure of each service under the `docs` folder.

## Sponsors

We would like to express our gratitude to the following sponsors for their generous support:

<div align="center">
  <a href="https://github.com/alisentas">
    <img src="https://github.com/alisentas.png" alt="tolgaakkapulu" width="50" height="50" style="border-radius: 50%">
  </a>
  <a href="https://github.com/furkansekerci">
    <img src="https://github.com/furkansekerci.png" alt="mkdemir" width="50" height="50" style="border-radius: 50%">
  </a>
</div>

If you are interested in becoming a sponsor, please visit our [GitHub Sponsors](https://github.com/sponsors) page.

## License

Bavul is licensed under the MIT License. See [LICENSE](LICENSE) for the full text of the license.