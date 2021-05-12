package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/mohae/struct2csv"
	"github.com/spf13/cobra"
)

type resData struct {
	D struct {
		Results []struct {
			// Metadata struct {
			// 	Type string `json:"type"`
			// 	URI  string `json:"uri"`
			// } `json:"__metadata"`
			ID string `json:"Id"`
			// Productinstancebe                string      `json:"productInstanceBE"`
			// Releasegrouptext                 string      `json:"releaseGroupText"`
			Appid string `json:"appId"`
			// Releaseid                        string      `json:"releaseId"`
			// Releasename                      string      `json:"releaseName"`
			// Otherreleases                    string      `json:"otherReleases"`
			Appnameall string `json:"AppNameAll"`
			// Rolenamecombined                 string      `json:"RoleNameCombined"`
			// Solutionscapabilitycombined      string      `json:"SolutionsCapabilityCombined"`
			// Solutionscapabilityidcombined    string      `json:"SolutionsCapabilityIDCombined"`
			// Availabilityinfaascombined       string      `json:"availabilityInFaaSCombined"`
			// Lobcombined                      string      `json:"LobCombined"`
			// Industrycombined                 string      `json:"IndustryCombined"`
			// Pcdcombined                      string      `json:"PCDCombined"`
			// Achcombined                      string      `json:"ACHCombined"`
			// Achlevel2Combined                string      `json:"ACHLevel2Combined"`
			// Databasecombined                 string      `json:"DatabaseCombined"`
			// Primarypvofficialnamecombined    string      `json:"PrimaryPVOfficialNameCombined"`
			// Pvfrontendcombined               string      `json:"PVFrontendCombined"`
			// Pvbackendcombined                string      `json:"PVBackendCombined"`
			// Apptypecombined                  string      `json:"AppTypeCombined"`
			// Technicalcatalognamecombined     string      `json:"TechnicalCatalogNameCombined"`
			// Businesscatalognamecombined      string      `json:"BusinessCatalogNameCombined"`
			// Frontendscvcombined              string      `json:"FrontendSCVCombined"`
			// Hanascvcombined                  string      `json:"HANASCVCombined"`
			// Backendscvcombined               string      `json:"BackendSCVCombined"`
			// Releasegrouptextcombined         string      `json:"releaseGroupTextCombined"`
			// Portfoliocategorycombined        string      `json:"PortfolioCategoryCombined"`
			// Productcategorydetails           string      `json:"ProductCategoryDetails"`
			// Roledescription                  string      `json:"RoleDescription"`
			// Businesscatalog                  string      `json:"BusinessCatalog"`
			// Technicalcatalog                 string      `json:"TechnicalCatalog"`
			// Fitanalysisach                   string      `json:"FitAnalysisACH"`
			// Pvbackend                        string      `json:"PVBackend"`
			// Pvfrontend                       string      `json:"PVFrontend"`
			// Apptypelevel                     string      `json:"AppTypeLevel"`
			// Productinstancehana              string      `json:"productInstanceHANA"`
			// Applicationcomponenttext         string      `json:"ApplicationComponentText"`
			// Formfactors                      string      `json:"FormFactors"`
			// Portfoliocategoryiv              interface{} `json:"PortfolioCategoryIV"`
			// Portfoliocategoryimp             string      `json:"PortfolioCategoryImp"`
			// Appname                          string      `json:"AppName"`
			// Applicationtype                  string      `json:"ApplicationType"`
			// Bspname                          string      `json:"BSPName"`
			// Bspapplicationurl                string      `json:"BSPApplicationURL"`
			// Newrolename                      string      `json:"NewRoleName"`
			// Rolename                         string      `json:"RoleName"`
			// Gtmlobname                       string      `json:"GTMLoBName"`
			// Productcategory                  string      `json:"ProductCategory"`
			// Applicationcomponent             string      `json:"ApplicationComponent"`
			// Database                         string      `json:"Database"`
			// Backendsoftwarecomponentversions string      `json:"BackendSoftwareComponentVersions"`
			// Frontendsoftwarecomponent        string      `json:"FrontendSoftwareComponent"`
			// Hanasoftwarecomponentversions    string      `json:"HANASoftwareComponentVersions"`
			// Rolecombinedtooltipdescription   string      `json:"RoleCombinedToolTipDescription"`
			// Businessroleoamname              string      `json:"BusinessRoleOAMName"`
			// Innovationscombined              string      `json:"InnovationsCombined"`
			// Intentscombined                  string      `json:"IntentsCombined"`
			// Bspnamecombined                  string      `json:"BSPNameCombined"`
			// Odataservicescombined            string      `json:"ODataServicesCombined"`
			// Webdynprocomponentnamecombined   string      `json:"WebDynproComponentNameCombined"`
			// Tcodescombined                   string      `json:"TCodesCombined"`
			// Rolenamecombinedonlyname         string      `json:"RoleNameCombinedOnlyName"`
			// Businessgroupnamecombined        string      `json:"BusinessGroupNameCombined"`
			// Businessgroupdescriptioncombined string      `json:"BusinessGroupDescriptionCombined"`
			// Bexquerynamecombined             string      `json:"BExQueryNameCombined"`
			// Bexquerydescriptioncombined      string      `json:"BExQueryDescriptionCombined"`
			// Sapui5Componentidcombined        string      `json:"SAPUI5ComponentIdCombined"`
			// Businessrolenamecombined         string      `json:"BusinessRoleNameCombined"`
			// Businessroledescriptioncombined  string      `json:"BusinessRoleDescriptionCombined"`
			// Fitanalysisachcombined           string      `json:"FitAnalysisACHCombined"`
			// Formfactorscombined              string      `json:"FormFactorsCombined"`
			// Uitechnologycombined             string      `json:"UITechnologyCombined"`
			// Appnamecombined                  string      `json:"AppNameCombined"`
			// Productinstanceui                string      `json:"productInstanceUI"`
			// Scopeitemid                      interface{} `json:"ScopeItemID"`
			// Highlightedappscombined          string      `json:"HighlightedAppsCombined"`
			// Highlightappssorterl1            string      `json:"HighlightAppsSorterL1"`
			// Highlightappssorterl2            string      `json:"HighlightAppsSorterL2"`
			// Scopeitemdetailscombined         string      `json:"ScopeItemDetailsCombined"`
			// Solutionscapabilityguidcombined  string      `json:"SolutionsCapabilityGUIDCombined"`
			// Scopeitemgroupcombined           interface{} `json:"ScopeItemGroupCombined"`
			// Odatav4Servicegroupscombined     interface{} `json:"ODataV4ServiceGroupsCombined"`
			// Applaunchertitlecombined         string      `json:"AppLauncherTitleCombined"`
			// Highlightapps                    string      `json:"HighlightApps"`
			// Bcmultilanguagecombined          interface{} `json:"BCMultiLanguageCombined"`
			// Brmultilanguagecombined          interface{} `json:"BRMultiLanguageCombined"`
			// Bgmultilanguagecombined          interface{} `json:"BGMultiLanguageCombined"`
			// Tilemultilanguagecombined        interface{} `json:"TileMultiLanguageCombined"`
			// Tcmultilanguagecombined          interface{} `json:"TCMultiLanguageCombined"`
		} `json:"results"`
	} `json:"d"`
}

func main() {
	var cmdSync = &cobra.Command{
		Use:   "sync",
		Short: "Synchronize DB",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			var start int = 0
			var end int = 100
  // It should loop 130 times
			for i := 0; i < 1; i++ {
				getAPI(start, end)
				start += 100
				end += 100
			}

			fmt.Print("data fetched")
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdSync)
	rootCmd.Execute()
}

func getData() {

}

func getAPI(skip int, top int) {
	var data resData

	var url string = fmt.Sprintf("https://fioriappslibrary.hana.ondemand.com/sap/fix/externalViewer/services/SingleApp.xsodata/InputFilterParam(InpFilterValue='1NA')/Results?$skip=%d&$top=%d&$orderby=tolower(HighlightAppsSorterL1)%%20asc,tolower(HighlightAppsSorterL2)%%20asc,tolower(AppNameAll)%%20asc&$format=json", skip, top)

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(body, &data)

	exportCSV(data)
}

func exportCSV(appData resData) {
	data := appData.D.Results
	enc := struct2csv.New()
	rows, err := enc.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	csvFile, err := os.Create("appData.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, singleRow := range rows {
		_ = csvwriter.Write(singleRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}
