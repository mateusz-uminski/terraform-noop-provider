terraform {
  required_providers {
    noop = {
      source = "local/noop"
    }
  }
}

provider "noop" {
  dir = "/tmp"
}

resource "noop_tmp_file" "example" {
  filename = "test_file.txt"
  content  = "example content"
}

output "file_content" {
  value = noop_tmp_file.example
}
