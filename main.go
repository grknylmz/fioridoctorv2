package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

type resData struct {
	D struct {
		Results []struct {
			Metadata struct {
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"__metadata"`
			ID                               string      `json:"Id"`
			Productinstancebe                string      `json:"productInstanceBE"`
			Releasegrouptext                 string      `json:"releaseGroupText"`
			Appid                            string      `json:"appId"`
			Releaseid                        string      `json:"releaseId"`
			Releasename                      string      `json:"releaseName"`
			Otherreleases                    string      `json:"otherReleases"`
			Appnameall                       string      `json:"AppNameAll"`
			Rolenamecombined                 string      `json:"RoleNameCombined"`
			Solutionscapabilitycombined      string      `json:"SolutionsCapabilityCombined"`
			Solutionscapabilityidcombined    string      `json:"SolutionsCapabilityIDCombined"`
			Availabilityinfaascombined       string      `json:"availabilityInFaaSCombined"`
			Lobcombined                      string      `json:"LobCombined"`
			Industrycombined                 string      `json:"IndustryCombined"`
			Pcdcombined                      string      `json:"PCDCombined"`
			Achcombined                      string      `json:"ACHCombined"`
			Achlevel2Combined                string      `json:"ACHLevel2Combined"`
			Databasecombined                 string      `json:"DatabaseCombined"`
			Primarypvofficialnamecombined    string      `json:"PrimaryPVOfficialNameCombined"`
			Pvfrontendcombined               string      `json:"PVFrontendCombined"`
			Pvbackendcombined                string      `json:"PVBackendCombined"`
			Apptypecombined                  string      `json:"AppTypeCombined"`
			Technicalcatalognamecombined     string      `json:"TechnicalCatalogNameCombined"`
			Businesscatalognamecombined      string      `json:"BusinessCatalogNameCombined"`
			Frontendscvcombined              string      `json:"FrontendSCVCombined"`
			Hanascvcombined                  string      `json:"HANASCVCombined"`
			Backendscvcombined               string      `json:"BackendSCVCombined"`
			Releasegrouptextcombined         string      `json:"releaseGroupTextCombined"`
			Portfoliocategorycombined        string      `json:"PortfolioCategoryCombined"`
			Productcategorydetails           string      `json:"ProductCategoryDetails"`
			Roledescription                  string      `json:"RoleDescription"`
			Businesscatalog                  string      `json:"BusinessCatalog"`
			Technicalcatalog                 string      `json:"TechnicalCatalog"`
			Fitanalysisach                   string      `json:"FitAnalysisACH"`
			Pvbackend                        string      `json:"PVBackend"`
			Pvfrontend                       string      `json:"PVFrontend"`
			Apptypelevel                     string      `json:"AppTypeLevel"`
			Productinstancehana              string      `json:"productInstanceHANA"`
			Applicationcomponenttext         string      `json:"ApplicationComponentText"`
			Formfactors                      string      `json:"FormFactors"`
			Portfoliocategoryiv              interface{} `json:"PortfolioCategoryIV"`
			Portfoliocategoryimp             string      `json:"PortfolioCategoryImp"`
			Appname                          string      `json:"AppName"`
			Applicationtype                  string      `json:"ApplicationType"`
			Bspname                          string      `json:"BSPName"`
			Bspapplicationurl                string      `json:"BSPApplicationURL"`
			Newrolename                      string      `json:"NewRoleName"`
			Rolename                         string      `json:"RoleName"`
			Gtmlobname                       string      `json:"GTMLoBName"`
			Productcategory                  string      `json:"ProductCategory"`
			Applicationcomponent             string      `json:"ApplicationComponent"`
			Database                         string      `json:"Database"`
			Backendsoftwarecomponentversions string      `json:"BackendSoftwareComponentVersions"`
			Frontendsoftwarecomponent        string      `json:"FrontendSoftwareComponent"`
			Hanasoftwarecomponentversions    string      `json:"HANASoftwareComponentVersions"`
			Rolecombinedtooltipdescription   string      `json:"RoleCombinedToolTipDescription"`
			Businessroleoamname              string      `json:"BusinessRoleOAMName"`
			Innovationscombined              string      `json:"InnovationsCombined"`
			Intentscombined                  string      `json:"IntentsCombined"`
			Bspnamecombined                  string      `json:"BSPNameCombined"`
			Odataservicescombined            string      `json:"ODataServicesCombined"`
			Webdynprocomponentnamecombined   string      `json:"WebDynproComponentNameCombined"`
			Tcodescombined                   string      `json:"TCodesCombined"`
			Rolenamecombinedonlyname         string      `json:"RoleNameCombinedOnlyName"`
			Businessgroupnamecombined        string      `json:"BusinessGroupNameCombined"`
			Businessgroupdescriptioncombined string      `json:"BusinessGroupDescriptionCombined"`
			Bexquerynamecombined             string      `json:"BExQueryNameCombined"`
			Bexquerydescriptioncombined      string      `json:"BExQueryDescriptionCombined"`
			Sapui5Componentidcombined        string      `json:"SAPUI5ComponentIdCombined"`
			Businessrolenamecombined         string      `json:"BusinessRoleNameCombined"`
			Businessroledescriptioncombined  string      `json:"BusinessRoleDescriptionCombined"`
			Fitanalysisachcombined           string      `json:"FitAnalysisACHCombined"`
			Formfactorscombined              string      `json:"FormFactorsCombined"`
			Uitechnologycombined             string      `json:"UITechnologyCombined"`
			Appnamecombined                  string      `json:"AppNameCombined"`
			Productinstanceui                string      `json:"productInstanceUI"`
			Scopeitemid                      interface{} `json:"ScopeItemID"`
			Highlightedappscombined          string      `json:"HighlightedAppsCombined"`
			Highlightappssorterl1            string      `json:"HighlightAppsSorterL1"`
			Highlightappssorterl2            string      `json:"HighlightAppsSorterL2"`
			Scopeitemdetailscombined         string      `json:"ScopeItemDetailsCombined"`
			Solutionscapabilityguidcombined  string      `json:"SolutionsCapabilityGUIDCombined"`
			Scopeitemgroupcombined           interface{} `json:"ScopeItemGroupCombined"`
			Odatav4Servicegroupscombined     interface{} `json:"ODataV4ServiceGroupsCombined"`
			Applaunchertitlecombined         string      `json:"AppLauncherTitleCombined"`
			Highlightapps                    string      `json:"HighlightApps"`
			Bcmultilanguagecombined          interface{} `json:"BCMultiLanguageCombined"`
			Brmultilanguagecombined          interface{} `json:"BRMultiLanguageCombined"`
			Bgmultilanguagecombined          interface{} `json:"BGMultiLanguageCombined"`
			Tilemultilanguagecombined        interface{} `json:"TileMultiLanguageCombined"`
			Tcmultilanguagecombined          interface{} `json:"TCMultiLanguageCombined"`
		} `json:"results"`
	} `json:"d"`
}

func main() {
	// var echoTimes int
	getData()

	var cmdPrint = &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	var cmdEcho = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	var cmdTimes = &cobra.Command{
		Use:   "times [string to echo]",
		Short: "Echo anything to the screen more times",
		Long: `echo things multiple times back to the user by providing
a count and a string.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			// for i := 0; i < echoTimes; i++ {
			// 	fmt.Println("Echo: " + strings.Join(args, " "))
			// }
		},
	}

	var cmdSync = &cobra.Command{
		Use:   "sync",
		Short: "Synchronize DB",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			// go getAPI(0, 100)
		},
	}

	// cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPrint, cmdEcho, cmdSync)
	cmdEcho.AddCommand(cmdTimes)
	rootCmd.Execute()
}

func getData() {
	getAPI(0, 1)
}

func getAPI(skip int, top int) {

	var url string = fmt.Sprintf("https://fioriappslibrary.hana.ondemand.com/sap/fix/externalViewer/services/SingleApp.xsodata/InputFilterParam(InpFilterValue='1NA')/Results?$skip=%d&$top=%d&$orderby=tolower(HighlightAppsSorterL1)%%20asc,tolower(HighlightAppsSorterL2)%%20asc,tolower(AppNameAll)%%20asc&$format=json", skip, top)

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	body, _ := ioutil.ReadAll(response.Body)

	var f resData

	err = json.Unmarshal(body, &f)

	fmt.Println(f.D.Results[0].Appnameall)

}
