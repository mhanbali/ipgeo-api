This is a simple api that allows you to query an ip address for geo data. It uses httprouter, and the data source is from IPinfo.io (links below, and in the source code).

Since you should be using the MMDB file I'm using oschwald's reader.

This may be useful for things like web analytics software or broadly in cyber security (ip allow/deny).

Source for the IP database:
https://github.com/abdullahdevrel/kaggle_data_source
https://www.kaggle.com/datasets/ipinfo/ipinfo-country-asn

Here is an example output after running a query on http://localhost:8080/query/39.213.149.53

OUTPUT:

{"as_domain":"telkomsel.com","as_name":"PT. Telekomunikasi Selular","asn":"AS23693","continent":"AS","continent_name":"Asia","country":"ID","country_name":"Indonesia"}
