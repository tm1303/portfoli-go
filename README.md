# RUN

From the cli, inputs are [json input][root fund] optional ...[highlight funds]

```
go run cmd/main.go "fund-data.json" "Ethical Global Fund" "GoldenGadgets"
```

# NOTES

Assumtion
* the input file could be massive so lets be careful loading
* "Ethical Global Fund" might not be the first fund in the file
* "Ethical Global Fund" ~ "Global Ethical Fund" :)
* there might be funds listed which are not in our portfolio
* a fund is a holding with 1 or more holding
* a copmany is a holding with exactly zero holdings
* a holdings name won't be stupidly long, this is to tune our buffers
* holding names are CASE SENSTAIVE (madness, but I probably won't have the energy for it)
