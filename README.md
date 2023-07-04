# Online-Learning-Platform

<a name="readme-top"></a>



[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]



<!-- PROJECT LOGO -->
<div align="center">
  <a href="https://github.com/mdanys/be-online-store">
    <img src="utils/logo.png" alt="Logo" width="200px">
  </a>

<h3 align="center">Let's shopping!</h3>

  <p align="center">
    This project is a API for Online Store.
    <br />
    <br />
    <br />
    <a href="https://app.swaggerhub.com/apis-docs/mdanys/be-online-store/1.0.0#/">View Demo</a>
    ·
    <a href="https://github.com/mdanys/be-online-store/issues">Report Bug</a>
    ·
    <a href="https://github.com/mdanys/be-online-store/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#features">Features</a>
      <ul>
        <li><a href="#entity-relationship-diagram">Entity Relationship Diagram</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## Features

- Customer
    - Customer can view product list by product category
    - Customer can add product to shopping cart
    - Customer can see a list of products that have been added to the shopping cart
    - Customer can delete product list in shopping cart
    - Customer can checkout and make payment transactions
    - Login and register customers
- Admin
    - Admin can add category
    - Admin can add product
    - Admin can get list all customer

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Entity Relationship Diagram



[![Store-ERD][erd-screenshot]](https://github.com/mdanys/be-online-store/blob/main/utils/ERD.drawio.png)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Installation local

1. Clone the repo
   ```bash
   git clone https://github.com/mdanys/be-online-store
   ```
2. Get env at [config.env](https://drive.google.com/file/d/13wLy-4LO1EPOmMTaaCZFWr7fsc2_uNYz/view?usp=sharing)
3. Enter your config in `config.env`
   ```bash
   AWS_ACCESS_KEY_ID = "ENTER YOUR AWS ACCESS KEY ID"
   AWS_SECRET_ACCESS_KEY = "ENTER YOUR AWS SECRET KEY ID"
   MIDTRANS_CLIENT = "ENTER YOUR MIDTRANS CLIENT"
   MIDTRANS_SERVER = "ENTER YOUR MIDTRANS SERVER"
   ```
4. Run project
   ```bash
   cd be-online-store
   paste your config.env
   go run .
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Documentation

```bash
Login as admin :

email : admin@admin.com
password : admin123
```

_For more examples, please refer to the [OPEN API](https://app.swaggerhub.com/apis-docs/mdanys/be-online-store/1.0.0#/)_

[![Online-API][product-screenshot]](https://github.com/mdanys/be-online-store/blob/main/utils/online-learning-platform.png)



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
[contributors-shield]: https://img.shields.io/github/contributors/mdanys/be-online-store?style=for-the-badge
[contributors-url]: https://github.com/mdanys/be-online-store/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/mdanys/be-online-store?style=for-the-badge
[forks-url]: https://github.com/mdanys/be-online-store/network/members
[stars-shield]: https://img.shields.io/github/stars/mdanys/be-online-store?style=for-the-badge
[stars-url]: https://github.com/mdanys/be-online-store/stargazers
[issues-shield]: https://img.shields.io/github/issues/mdanys/be-online-store?style=for-the-badge
[issues-url]: https://github.com/mdanys/be-online-store/issues
[license-shield]: https://img.shields.io/github/license/mdanys/be-online-store?style=for-the-badge
[license-url]: https://github.com/mdanys/be-online-store/blob/main/LICENSE
[product-screenshot]: utils/online-learning-platform.png
[erd-screenshot]: utils/ERD.drawio.png