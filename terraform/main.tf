terraform {
  required_version = ">=1.5.0"

  required_providers {
    github = {
      source  = "integrations/github"
      version = "5.31.0"
    }
  }

  cloud {
    organization = "mokmok"
    hostname     = "app.terraform.io"

    workspaces {
      name = "golang-template"
    }
  }
}

provider "github" {
  owner = var.github.owner
}
