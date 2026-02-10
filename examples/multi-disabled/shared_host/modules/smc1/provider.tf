terraform {
  required_providers {
    smc = {
      source = "forcepoint/fp-ngfw-smc"
      #      configuration_aliases = [ smc.smc1 ]
    }
  }
}
