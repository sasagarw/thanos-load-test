package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	log.Println("Starting go application ...")
	for {
		runLoadTest()
		time.Sleep(time.Minute)
	}
}

func runLoadTest() {
	for j := 0; j < 60; j++ {
		log.Println("Printing ... ", j)
		time.Sleep(time.Second)
	}
	namespaces := []string{"addon-active-monitor-ns", "addon-alb-ingress-controller-ns", "addon-armador-ns", "addon-chaos-system", "addon-cluster-autoscaler-ns", "addon-efs-csi-ns", "addon-event-router-ns", "addon-external-dns-ns", "addon-ipa-ns", "addon-istio-installer-ns", "addon-ksm-forwarder-ns", "addon-lifecycle-manager-ns", "addon-manager-system", "addon-metricset-ns", "addon-nodereaper-ns", "addon-oil-ns", "addon-opa-gatekeeper-ns", "addon-opentelemetry-agent-ns", "addon-pdbreaper-ns", "addon-podreaper-ns", "addon-upgrademgr-ns", "addon-wavefront-s3-adapter-ns", "addon-wavefrontcollectors-ns", "admiral-sync", "ambassador", "api-management-apiexpress-usw2-e2e", "api-management-apisandboxext01-usw2-e2e", "api-management-apisandboxext01-usw2-qal", "api-management-gqltest0727-usw2-e2e", "api-management-gqltest0727-usw2-qal", "api-management-graphqlproxyserver-usw2-e2e", "api-management-graphqlproxyserver-usw2-qal", "argo", "argo-events", "argo-rollouts", "chaos-ns", "cloud-prnagonbdapi-usw2-e2e", "cloud-prnagonbdapi-usw2-qal", "cloud-security-mdscachersiks-usw2-qal", "default", "dev-asteriasapigateway-usw2-e2e", "dev-asteriasapigateway-usw2-qal", "dev-asteriasappfregistration-usw2-e2e", "dev-asteriasappfregistration-usw2-qal", "dev-asteriasargocd-usw2-e2e", "dev-asteriasargocd-usw2-qal", "dev-asteriasartifactory-usw2-e2e", "dev-asteriasartifactory-usw2-qal", "dev-asteriasasset-usw2-e2e", "dev-asteriasasset-usw2-qal", "dev-asteriasbridge-usw2-e2e", "dev-asteriasbridge-usw2-qal", "dev-asteriasbuilder-usw2-e2e", "dev-asteriasbuilder-usw2-qal", "dev-asteriascaas-usw2-e2e", "dev-asteriascaas-usw2-qal", "dev-asteriascloudresource-usw2-e2e", "dev-asteriascloudresource-usw2-qal", "dev-asteriascontent-usw2-e2e", "dev-asteriascontent-usw2-qal", "dev-asteriasconversationbot-usw2-e2e", "dev-asteriasconversationbot-usw2-qal", "dev-asteriasdataservice-usw2-e2e", "dev-asteriasdataservice-usw2-qal", "dev-asteriasdevicefarm-usw2-e2e", "dev-asteriasdevicefarm-usw2-qal", "dev-asteriasdevservicesync-usw2-e2e", "dev-asteriasdevservicesync-usw2-qal", "dev-asteriasgateway-usw2-e2e", "dev-asteriasgateway-usw2-qal", "dev-asteriasgithub-usw2-e2e", "dev-asteriasgithub-usw2-qal", "dev-asteriasidps-usw2-e2e", "dev-asteriasidps-usw2-qal", "dev-asteriasiksm-usw2-e2e", "dev-asteriasiksm-usw2-qal", "dev-asteriasjenkins-usw2-e2e", "dev-asteriasjenkins-usw2-qal", "dev-asteriasnetgenie-usw2-e2e", "dev-asteriasnetgenie-usw2-qal", "dev-asteriasnotification-usw2-e2e", "dev-asteriasnotification-usw2-qal", "dev-asteriaspaging-usw2-e2e", "dev-asteriaspaging-usw2-qal", "dev-asteriasrproxy-usw2-e2e", "dev-asteriasrproxy-usw2-qal", "dev-asteriassecretsmanager-usw2-e2e", "dev-asteriassecretsmanager-usw2-qal", "dev-asteriassplunk-usw2-e2e", "dev-asteriassplunk-usw2-qal", "dev-asteriasspringbootadmin-usw2-e2e", "dev-asteriasspringbootadmin-usw2-qal", "dev-asteriassslcert-usw2-e2e", "dev-asteriassslcert-usw2-qal", "dev-asteriaswavefrontalerts-usw2-e2e", "dev-asteriaswavefrontalerts-usw2-qal", "dev-build-mobilepavedroadeiamclient-usw2-e2e", "dev-build-mobilepavedroadeiamclient-usw2-qal", "dev-containers-fxudeleteme0902-usw2-e2e", "dev-containers-fxudeleteme0902-usw2-qal", "dev-containers-oicpcontainerimages-usw2-e2e", "dev-containers-oicpcontainerimages-usw2-qal", "dev-cpd2-usw2-qal", "dev-deploy-meshrolloutpoc11-usw2-e2e", "dev-deploy-meshrolloutpoc11-usw2-qal", "dev-devx-cia-usw2-e2e", "dev-devx-cia-usw2-qal", "dev-devx-datacollector-usw2-e2e", "dev-devx-developercli-usw2-e2e", "dev-devx-developercli-usw2-qal", "dev-devx-developerftest-usw2-e2e", "dev-devx-developerftest-usw2-qal", "dev-devx-dexter-usw2-e2e", "dev-devx-dexter-usw2-qal", "dev-devx-druidmanagementservices-usw2-qal", "dev-devx-druidreverseproxy-usw2-e2e", "dev-devx-druidreverseproxy-usw2-prf", "dev-devx-druidreverseproxy-usw2-qal", "dev-devx-getafix-usw2-qdat", "dev-devx-k8sodlforwarder-usw2-qal", "dev-devx-ntantrydelme0831-usw2-e2e", "dev-devx-ntantrydelme0831-usw2-qal", "dev-devx-o11yacp-usw2-qal", "dev-devx-o11yanalytics-usw2-e2e", "dev-devx-o11yanalytics-usw2-qal", "dev-devx-o11ydataplatformdemo-usw2-e2e", "dev-devx-o11ydataplatformdemo-usw2-qal", "dev-devx-o11ydocmgmtsvcdemo-usw2-e2e", "dev-devx-o11ydocmgmtsvcdemo-usw2-qal", "dev-devx-o11yfuzzygql-usw2-qal", "dev-devx-o11yfuzzygqlfederation-usw2-pprd", "dev-devx-o11yfuzzygqlfederation-usw2-qal", "dev-devx-o11ygsapigwharouter-usw2-e2e", "dev-devx-o11ygsapigwharouter-usw2-qal", "dev-devx-odldevelopervelocitygql-usw2-qal", "dev-devx-odldpingestionprocessor-usw2-e2e", "dev-devx-odldpingestionprocessor-usw2-qal", "dev-devx-odldruidlookupdatagen-usw2-qal", "dev-devx-odldveleventsforwarder-usw2-qal", "dev-devx-odlelasticconsumers-usw2-qal", "dev-devx-odlextelasticconsumer-usw2-qal", "dev-devx-odlgithubstats-usw2-qal", "dev-devx-odlgitwebhookservice-usw2-qal", "dev-devx-odlkafkaconnects3proc-usw2-qal", "dev-devx-odlosamdetector-usw2-qal", "dev-devx-odlosammltrainer-usw2-qal", "dev-devx-odlosampublisher-usw2-qal", "dev-devx-odlscheduledjobs-usw2-qal", "dev-devx-odlsupportservicetool-usw2-qal", "dev-devx-osamassetsinterface-usw2-qal", "dev-devx-osamgenericstreamer-usw2-qal", "dev-devx-osamgoldensignal-usw2-qal", "dev-devx-osamksmnormalizer-usw2-qal", "dev-devx-osammlpfci-usw2-e2e", "dev-devx-osamodmetrictrainer-usw2-qal", "dev-devx-osamondemandgql-usw2-qal", "dev-devx-osamsw1-usw2-qal", "dev-devx-osamsw1-usw2-qald", "dev-devx-rts-usw2-e2e", "dev-devx-rts-usw2-qal", "dev-devx-wavefrontcost-usw2-prf", "dev-devx-wavefrontcost-usw2-qal", "dev-devx-zookeeperexporter-usw2-e2e", "dev-druid-usw2-dev", "dev-druid-usw2-qal", "dev-druid-usw2-qala", "dev-druid-usw2-qdat", "dev-druid-usw2-qmst", "dev-integration-appdpolicygen-usw2-e2e", "dev-integration-appdpolicygen-usw2-qal", "dev-integration-msaasmetrics-usw2-e2e", "dev-integration-msaasmetrics-usw2-qal", "dev-integration-newmetrics2-usw2-e2e", "dev-integration-newmetrics2-usw2-qal", "dev-integration-o11yalerts-usw2-e2e", "dev-integration-o11yalerts-usw2-qal", "dev-integration-o11yterraform-usw2-e2e", "dev-integration-o11yterraform-usw2-qal", "dev-integration-te01stgithubhl-usw2-e2e", "dev-integration-te01stgithubhl-usw2-qal", "dev-integration-testgatewaycert-usw2-e2e", "dev-integration-testgatewaycert-usw2-qal", "dev-integration-testmetrics22-usw2-e2e", "dev-integration-testmetrics22-usw2-qal", "dev-integration-testtemplateerrors222-usw2-e2e", "dev-integration-testtemplateerrors222-usw2-qal", "dev-integration-testtemplatetype0125-usw2-e2e", "dev-integration-testtemplatetype0125-usw2-qal", "dev-ists-usw2-e2e", "dev-ko-usw2-e2e", "dev-mpaapigweventstp-usw2-qal", "dev-mpaassetidparity-usw2-qal", "dev-mpacostresanalysisargo-usw2-qal", "dev-mpaehifigraphql-usw2-e2e", "dev-mpaehifigraphql-usw2-qal", "dev-mpaehifitestclnt-usw2-qal", "dev-mpaehifitestsvc-usw2-qal", "dev-mpaiksmdatappfilter-usw2-qal", "dev-mpak8sevtsproc-usw2-qal", "dev-mpameetingstats-usw2-qal", "dev-mpaoimmetricsfilter-usw2-qal", "dev-msaaspolicyagent-usw2-qal", "dev-odlesmetrics-usw2-qal", "dev-odliksmsystemnsconsumer-usw2-qal", "dev-odlmanagementservice-usw2-qal", "dev-odltestconsumerservice-usw2-qal", "dev-odlworkdayservice-usw2-qal", "dev-patterns-artifactscan-usw2-e2e-dev", "dev-patterns-artifactscan-usw2-qal-dev", "dev-patterns-cpd2service-usw2-e2e", "dev-patterns-klir-usw2-e2e", "dev-patterns-klir-usw2-qal", "dev-patterns-platform360-usw2-qal", "dev-tools-chaostest-usw2-e2e", "dev-tools-deletemeinttest060122-usw2-e2e", "dev-tools-deletemeinttest060122-usw2-qal", "dev-tools-deletememvc1012-usw2-e2e", "dev-tools-deletememvc1012-usw2-qal", "dev-tools-deletemesigsci-usw2-e2e", "dev-tools-deletemesigsci-usw2-qal", "dev-tools-deletemesplunk040622-usw2-e2e", "dev-tools-deletemesplunk040622-usw2-qal", "dev-tools-deletemesplunk0927-usw2-e2e", "dev-tools-deletemesplunk0927-usw2-qal", "dev-tools-deletemetest0928-usw2-e2e", "dev-tools-deletemetest0928-usw2-qal", "dev-tools-deletemetestmvc0929-usw2-e2e", "dev-tools-deletemetestmvc0929-usw2-qal", "dev-tools-deletemetestsvc0928-usw2-e2e", "dev-tools-deletemetestsvc0928-usw2-qal", "dev-tools-gedfall22lesaint-usw2-e2e", "dev-tools-gedfall22lesaint-usw2-qal", "dev-tools-intuitzenhub-usw2-e2e", "dev-tools-intuitzenhub-usw2-qal", "dev-tools-springmvc1014-usw2-e2e", "dev-tools-springmvc1014-usw2-qal", "dev-tools-testawsincident0928-usw2-e2e", "dev-tools-testawsincident0928-usw2-qal", "dev-tools-testsvc0921221-usw2-e2e", "dev-tools-testsvc0921221-usw2-qal", "falco-ns", "iam-manager-system", "instance-manager", "istio-system", "khirani-test", "kube-node-lease", "kube-public", "kube-system", "kubernetes-clustermetadatasync-usw2-e2e", "kubernetes-clustermetadatasync-usw2-qal", "mesh-health-check", "numa-logic", "numaflow-system", "odl-odlonboardingselfserv-usw2-e2e", "odl-odlonboardingselfserv-usw2-qal", "odl-odlonboardingworkflow-usw2-e2e", "odl-odlonboardingworkflow-usw2-qal", "opa", "oss-analytics-sampledataflow-usw2-e2e", "oss-analytics-sampledataflow-usw2-qal", "platformexps-devx-test20222708prod-usw2-e2e", "platformexps-devx-test20222708prod-usw2-qal", "platformops-cloudops-netgenietest5-usw2-e2e", "platformops-cloudops-netgenietest5-usw2-qal", "s360-datapersistenceapi-usw2-e2e", "s360-opexapi-usw2-e2e", "s360-s360scheduler-usw2-e2e", "s360-s360scheduler-usw2-qal", "sandbox-meshdemoc-usw2-e2e", "sandbox-meshdemoc-usw2-qal", "sandbox-meshdemod-usw2-e2e", "sandbox-meshdemod-usw2-qal", "sandbox-meshdemoe-usw2-e2e", "sandbox-meshdemoe-usw2-qal", "sandbox-sandbox-adhasongqlservice2-usw2-e2e", "sandbox-sandbox-adhasongqlservice2-usw2-qal", "sandbox-sandbox-adhasongqlsvc1-usw2-e2e", "sandbox-sandbox-adhasongqlsvc1-usw2-qal", "sandbox-sandbox-beginerdeleteme-usw2-e2e", "sandbox-sandbox-beginerdeleteme-usw2-qal", "sandbox-sandbox-bing1014deleteme-usw2-e2e", "sandbox-sandbox-bing1014deleteme-usw2-qal", "sandbox-sandbox-ckumar11testservice-usw2-e2e", "sandbox-sandbox-ckumar11testservice-usw2-qal", "sandbox-sandbox-dasdadele-usw2-e2e", "sandbox-sandbox-dasdadele-usw2-qal", "sandbox-sandbox-deleteme0606a-usw2-e2e", "sandbox-sandbox-deleteme0606a-usw2-qal", "sandbox-sandbox-deleteme1011svc-usw2-e2e", "sandbox-sandbox-deleteme1011svc-usw2-qal", "sandbox-sandbox-deleteme1018aexp-usw2-e2e", "sandbox-sandbox-deleteme1018aexp-usw2-qal", "sandbox-sandbox-deletemedemo0526b-usw2-e2e", "sandbox-sandbox-deletemedemo0526b-usw2-qal", "sandbox-sandbox-deletemedemo0921-usw2-e2e", "sandbox-sandbox-deletemedemo0921-usw2-qal", "sandbox-sandbox-deletemeexpress0901a-usw2-e2e", "sandbox-sandbox-deletemeexpress0901a-usw2-qal", "sandbox-sandbox-deletemegena12303-usw2-e2e", "sandbox-sandbox-deletemegena12303-usw2-qal", "sandbox-sandbox-deletemegena123132-usw2-e2e", "sandbox-sandbox-deletemegena123132-usw2-qal", "sandbox-sandbox-deletemelwestcdd1011a-usw2-e2e", "sandbox-sandbox-deletemelwestcdd1011a-usw2-qal", "sandbox-sandbox-deletemelwestjavaver-usw2-e2e", "sandbox-sandbox-deletemelwestjavaver-usw2-qal", "sandbox-sandbox-deletemelwestkarate1011-usw2-e2e", "sandbox-sandbox-deletemelwestkarate1011-usw2-qal", "sandbox-sandbox-deletemetestlieu-usw2-e2e", "sandbox-sandbox-deletemetestlieu-usw2-qal", "sandbox-sandbox-demobackendtest1-usw2-e2e", "sandbox-sandbox-demobackendtest1-usw2-qal", "sandbox-sandbox-dep25deleteme-usw2-e2e", "sandbox-sandbox-dep25deleteme-usw2-qal", "sandbox-sandbox-eksadilnew-usw2-e2e", "sandbox-sandbox-eksadilnew-usw2-qal", "sandbox-sandbox-exampleikustomizedepr1-usw2-e2e", "sandbox-sandbox-exampleikustomizedepr1-usw2-qal", "sandbox-sandbox-exampleikustomizedepr2-usw2-e2e", "sandbox-sandbox-exampleikustomizedepr2-usw2-qal", "sandbox-sandbox-expressiksmprd-usw2-e2e", "sandbox-sandbox-expressiksmprd-usw2-qal", "sandbox-sandbox-gqltestsubgraph1-usw2-e2e", "sandbox-sandbox-gqltestsubgraph1-usw2-prf", "sandbox-sandbox-gqltestsubgraph1-usw2-qal", "sandbox-sandbox-gsdeletemeoct25-usw2-e2e", "sandbox-sandbox-gsdeletemeoct25-usw2-qal", "sandbox-sandbox-iksexpressdeleteme-usw2-e2e", "sandbox-sandbox-iksexpressdeleteme-usw2-qal", "sandbox-sandbox-inctdeleteme-usw2-e2e", "sandbox-sandbox-inctdeleteme-usw2-qal", "sandbox-sandbox-istslib4mvcdeleteme-usw2-e2e", "sandbox-sandbox-istslib4mvcdeleteme-usw2-qal", "sandbox-sandbox-justiceleaguesvc-usw2-e2e", "sandbox-sandbox-justiceleaguesvc-usw2-qal", "sandbox-sandbox-lchan8testservice6-usw2-e2e", "sandbox-sandbox-lchan8testservice6-usw2-qal", "sandbox-sandbox-luiselimon-usw2-e2e", "sandbox-sandbox-luiselimon-usw2-qal", "sandbox-sandbox-luistestplayground-usw2-e2e", "sandbox-sandbox-luistestplayground-usw2-qal", "sandbox-sandbox-mcdemo4-usw2-e2e", "sandbox-sandbox-mcdemo4-usw2-qal", "sandbox-sandbox-nbn01playground-usw2-e2e", "sandbox-sandbox-nbn01playground-usw2-qal", "sandbox-sandbox-nbn01playground1-usw2-e2e", "sandbox-sandbox-nbn01playground1-usw2-qal", "sandbox-sandbox-nexprdeleteme-usw2-e2e", "sandbox-sandbox-nexprdeleteme-usw2-qal", "sandbox-sandbox-ntantrydeletemegosvc-usw2-e2e", "sandbox-sandbox-ntantrydeletemegosvc-usw2-qal", "sandbox-sandbox-ntantrydeletemejavasvc-usw2-e2e", "sandbox-sandbox-ntantrydeletemejavasvc-usw2-qal", "sandbox-sandbox-recorddeleteme-usw2-e2e", "sandbox-sandbox-recorddeleteme-usw2-qal", "sandbox-sandbox-rolloutsconverterists-usw2-e2e", "sandbox-sandbox-rolloutsconverterists-usw2-qal", "sandbox-sandbox-rpdemo101-usw2-e2e", "sandbox-sandbox-rpdemo101-usw2-qal", "sandbox-sandbox-rpgqlsubgraphpavedroad1-usw2-e2e", "sandbox-sandbox-rpgqlsubgraphpavedroad1-usw2-qal", "sandbox-sandbox-rpr2deleteme-usw2-e2e", "sandbox-sandbox-rpr2deleteme-usw2-qal", "sandbox-sandbox-sringbootgpqlconnection-usw2-e2e", "sandbox-sandbox-sringbootgpqlconnection-usw2-qal", "sandbox-sandbox-svcstax-usw2-e2e", "sandbox-sandbox-svcstax-usw2-qal", "sandbox-sandbox-tes2deleteme-usw2-e2e", "sandbox-sandbox-tes2deleteme-usw2-qal", "sandbox-sandbox-testexpressservice001-usw2-e2e", "sandbox-sandbox-testexpressservice001-usw2-qal", "sandbox-sandbox-testservice004-usw2-e2e", "sandbox-sandbox-testservice004-usw2-qal", "sandbox-sandbox-testtemplateerror-usw2-e2e", "sandbox-sandbox-testtemplateerror-usw2-qal", "sandbox-sandbox-testtemplateerrors-usw2-e2e", "sandbox-sandbox-testtemplateerrors-usw2-qal", "services-config-apollosvc1-usw2-e2e", "services-config-apollosvc1-usw2-qal", "services-config-basnetchhetri000-usw2-e2e", "services-config-basnetchhetri000-usw2-qal", "services-config-basnetkaji03-usw2-e2e", "services-config-basnetkaji03-usw2-qal", "services-config-cleartripservice-usw2-e2e", "services-config-cleartripservice-usw2-qal", "services-config-mmtservice-usw2-e2e", "services-config-mmtservice-usw2-qal", "services-config-starwarsservice-usw2-e2e", "services-config-starwarsservice-usw2-qal", "services-config-uprendermanifest-usw2-e2e", "services-config-uprendermanifest-usw2-qal", "services-java-api1780testprod-usw2-e2e", "services-java-api1780testprod-usw2-qal", "services-java-jskasteriasclient-usw2-e2e", "services-java-jskasteriasclient-usw2-qal", "services-jskv4serviceinitializr-usw2-e2e-a"}
	endpoint := getEnv("ENDPOINT", "prometheus.addon-metricset-ns.svc.cluster.local")
	port := getEnv("PORT", "9090")
	noOfRoutines := getEnv("GOROUTINES", "5")
	noOfRoutinesInt, _ := strconv.Atoi(noOfRoutines)

	wg := new(sync.WaitGroup)
	wg.Add(noOfRoutinesInt)

	for i := 0; i < noOfRoutinesInt; i++ {
		go makeHttpCall(wg, endpoint, port, namespaces[i])
		time.Sleep(time.Second)
	}
	wg.Wait()
}

func makeHttpCall(wg *sync.WaitGroup, endpoint, port, namespace string) {
	defer wg.Done()
	metrics := []string{
		"namespace_app_pod_cpu_utilization",
		"namespace_pod_memory_utilization",
		"namespace_pod_cpu_utilization",
		"namespace_pod_http_server_requests_2xx",
		"namespace_app_pod_http_server_requests_2xx",
		"namespace_app_pod_count",
		"namespace_app_pod_http_server_requests_count",
		"namespace_app_pod_jvm_gc_pause_seconds_avg",
		"namespace_pod_go_goroutines",
		"namespace_pod_go_gc_duration_seconds_rate"}
	start := getEnv("START_TIME", "1668286200")
	end := getEnv("END_TIME", "1668415800")
	for _, metric := range metrics {
		url := fmt.Sprintf("http://%s:%s/api/v1/query_range?query=%s&namespace=%s&start=%s&end=%s&step=518", endpoint, port, metric, namespace, start, end)
		callEndpoint(url)
	}
}

func callEndpoint(url string) {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("Connection", "close")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}
	return value
}
