#!/bin/bash

helm template ../chart --name-template=hockeytrainer > ../manifests/hockeytrainer-chart.yaml