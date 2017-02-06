package lib

import (
	"os"
	"path/filepath"

	"strconv"

	"github.com/BurntSushi/toml"
)

const (
	//CmdInsert insert mode 定数
	CmdInsert = "insert"
	//CmdUpdate update mode 定数
	CmdUpdate = "update"
	//CmdDelete delete mode 定数
	CmdDelete = "delete"
	//CmdOptionStore -store オプション
	CmdOptionStore = "store"
	//CmdOptionImage -image オプション
	CmdOptionImage = "image"
	//AssetsDirName assets dir name
	AssetsDirName = "assets"
	//ImageDirName dir name
	ImageDirName = "image"
	//LogoDirName logo dir name
	LogoDirName = "logo"
	//ImageMicroName micro kind name
	ImageMicroName = "micro"
	//ImageSmallName small kind name
	ImageSmallName = "small"
	//ImageMediumName medium kind name
	ImageMediumName = "medium"
	//ImageLargeName large kind name
	ImageLargeName = "large"
	//ImageOriginName origin kind name
	ImageOriginName = "origin"
	//ImageOriginNoLogoName origin no logo kind name
	ImageOriginNoLogoName = "origin_nologo"
	//TypeCookingName cooking type name
	TypeCookingName = "cooking"
	//TypeOtherName other type name
	TypeOtherName = "other"
	//SaveImageExt save image extension
	SaveImageExt = "jpeg"
	//SaveImageQuality is image quality
	SaveImageQuality100 = 100
	//SaveNoOriginImageQuality is image quality
	SaveImageQuality50 = 50
)

var (
	//Config 設定ファイルのグローバル変数
	Config Configs
	//CurrentBasePath current path
	CurrentBasePath string
)

//Configs 設定ファイル
type Configs struct {
	Database  DatabaseConfig
	Logo      LogoConfig
	Cooking   CookingConfig
	Other     OtherConfig
	ImageSize ImageSizeConfig
	Aws       AwsConfig
}

//DatabaseConfig database 設定ファイル
type DatabaseConfig struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
}

//LogoConfig logo画像の設定
type LogoConfig struct {
	LargeName  string
	OriginName string
}

//CookingConfig 料理画像の設定
type CookingConfig struct {
	OriginID       int
	LargeID        int
	MediumID       int
	SmallID        int
	MicroID        int
	OriginNoLogoID int
}

//OtherConfig その他画像の設定
type OtherConfig struct {
	OriginID       int
	LargeID        int
	MediumID       int
	SmallID        int
	MicroID        int
	OriginNoLogoID int
}

//AwsConfig awsの設定
type AwsConfig struct {
	BucketName      string
	AccessKeyID     string
	SecretAccessKey string
}

//ImageSizeConfig define image size
type ImageSizeConfig struct {
	LargeWidth   int
	LargeHeight  int
	MediumWidth  int
	MediumHeight int
	SmallWidth   int
	SmallHeight  int
	MicroWidth   int
	MicroHeight  int
}

//GetAwsBucketName is get aws bucket name by toml or env
func (con *AwsConfig) GetAwsBucketName() string {
	var result string
	if result = con.BucketName; result == "" {
		if result = os.Getenv("AwsBucketName"); result == "" {
			panic("awsのbucket nameが設定されていません")
		}
	}
	return result
}

//GetAwsAccessKeyID is get aws accesskeyID by toml or env
func (con *AwsConfig) GetAwsAccessKeyID() string {
	var result string
	if result = con.AccessKeyID; result == "" {
		if result = os.Getenv("AwsAccessKeyID"); result == "" {
			panic("awsのaccesskeyIDが設定されていません")
		}
	}
	return result
}

//GetAwsSecretAccessKey is get aws SecretAccesskey by toml or env
func (con *AwsConfig) GetAwsSecretAccessKey() string {
	var result string
	if result = con.SecretAccessKey; result == "" {
		if result = os.Getenv("AwsSecretAccessKey"); result == "" {
			panic("awsのsecret access keyが設定されていません")
		}
	}
	return result
}

//GetAssetsPath assets path
func (con *Configs) GetAssetsPath() string {
	return filepath.Join(CurrentBasePath, AssetsDirName)
}

//GetLogoPath logo path
func (con *Configs) GetLogoPath() string {
	return filepath.Join(con.GetAssetsPath(), LogoDirName)
}

////GetMediumLogoPath mediumLogo path
//func (con *Configs) GetMediumLogoPath() string {
//	return filepath.Join(con.GetLogoPath(), con.Logo.MediumName)
//}

//GetLargeLogoPath largeLogo path
func (con *Configs) GetLargeLogoPath() string {
	return filepath.Join(con.GetLogoPath(), con.Logo.LargeName)
}

//GetOriginLogoPath originLogo path
func (con *Configs) GetOriginLogoPath() string {
	return filepath.Join(con.GetLogoPath(), con.Logo.OriginName)
}

//GetImagePath image dir path
func (con *Configs) GetImagePath() string {
	return filepath.Join(con.GetAssetsPath(), ImageDirName)
}

//GetStoreImagePath store image path
func (con *Configs) GetStoreImagePath(storeID int) string {
	return filepath.Join(con.GetImagePath(), strconv.Itoa(storeID))
}

//GetImageMicroPath micro image path
func (con *Configs) GetImageMicroPath(storeID int, typeName string) string {
	return filepath.Join(con.GetStoreImagePath(storeID), typeName, ImageMicroName)
}

//GetImageSmallPath small image path
func (con *Configs) GetImageSmallPath(storeID int, typeName string) string {
	return filepath.Join(con.GetStoreImagePath(storeID), typeName, ImageSmallName)
}

//GetImageMediumPath medium image path
func (con *Configs) GetImageMediumPath(storeID int, typeName string) string {
	return filepath.Join(con.GetStoreImagePath(storeID), typeName, ImageMediumName)
}

//GetImageLargePath large image path
func (con *Configs) GetImageLargePath(storeID int, typeName string) string {
	return filepath.Join(con.GetStoreImagePath(storeID), typeName, ImageLargeName)
}

//GetImagePath origin image path
func (con *Configs) GetImageOriginPath(storeID int, typeName string) string {
	return filepath.Join(con.GetStoreImagePath(storeID), typeName, ImageOriginName)
}

//GetImageOriginNoLogoPath originLogo image path
func (con *Configs) GetImageOriginNoLogoPath(storeID int, typeName string) string {
	return filepath.Join(con.GetStoreImagePath(storeID), typeName, ImageOriginNoLogoName)
}

//GetImageSrcPath cooking image source dir
func (con *Configs) GetImageSrcPath(storeID int, typeName string) string {
	return filepath.Join(con.GetStoreImagePath(storeID), typeName)
}

//GetImageCookingSrcPath cooking image source dir
func (con *Configs) GetImageCookingSrcPath(storeID int) string {
	return filepath.Join(con.GetStoreImagePath(storeID), TypeCookingName)
}

//GetImageCookingMicroPath cooking micro image path
func (con *Configs) GetImageCookingMicroPath(storeID int) string {
	return con.GetImageMicroPath(storeID, TypeCookingName)
}

//GetImageCookingSmallPath cooiking small image path
func (con *Configs) GetImageCookingSmallPath(storeID int) string {
	return con.GetImageSmallPath(storeID, TypeCookingName)
}

//GetImageCookingMediumPath cooking medium image path
func (con *Configs) GetImageCookingMediumPath(storeID int) string {
	return con.GetImageMediumPath(storeID, TypeCookingName)
}

//GetImageCookingLargePath cooking large image pa†h
func (con *Configs) GetImageCookingLargePath(storeID int) string {
	return con.GetImageLargePath(storeID, TypeCookingName)
}

//GetImageCookingOriginPath cooking origin image path
func (con *Configs) GetImageCookingOriginPath(storeID int) string {
	return con.GetImageOriginPath(storeID, TypeCookingName)
}

//GetImageCookingOriginLogoPath cooking origin logo path
func (con *Configs) GetImageCookingOriginLogoPath(storeID int) string {
	return con.GetImageOriginNoLogoPath(storeID, TypeCookingName)
}

//GetImageOtherSrcPath other image source dir
func (con *Configs) GetImageOtherSrcPath(storeID int) string {
	return filepath.Join(con.GetStoreImagePath(storeID), TypeCookingName)
}

//GetImageOtherMicroPath other micro image path
func (con *Configs) GetImageOtherMicroPath(storeID int) string {
	return con.GetImageMicroPath(storeID, TypeOtherName)
}

//GetImageOtherSmallPath other small image path
func (con *Configs) GetImageOtherSmallPath(storeID int) string {
	return con.GetImageSmallPath(storeID, TypeOtherName)
}

//GetImageOtherMediumPath other medium image path
func (con *Configs) GetImageOtherMediumPath(storeID int) string {
	return con.GetImageMediumPath(storeID, TypeOtherName)
}

//GetImageOtherLargePath other large image pa†h
func (con *Configs) GetImageOtherLargePath(storeID int) string {
	return con.GetImageLargePath(storeID, TypeOtherName)
}

//GetImageOtherOriginPath other origin image path
func (con *Configs) GetImageOtherOriginPath(storeID int) string {
	return con.GetImageOriginPath(storeID, TypeOtherName)
}

//GetImageOtherOriginLogoPath other origin logo path
func (con *Configs) GetImageOtherOriginLogoPath(storeID int) string {
	return con.GetImageOriginNoLogoPath(storeID, TypeOtherName)
}

//GetImageKindNameByKind
func (con *Configs) GetImageKindNameByKind(kind int) string {
	var kindName string

	switch kind {
	case con.Cooking.OriginNoLogoID, con.Other.OriginNoLogoID:
		kindName = ImageOriginNoLogoName
	case con.Cooking.OriginID, con.Other.OriginID:
		kindName = ImageOriginName
	case con.Cooking.LargeID, con.Other.LargeID:
		kindName = ImageLargeName
	case con.Cooking.MediumID, con.Other.MediumID:
		kindName = ImageMediumName
	case con.Cooking.SmallID, con.Other.SmallID:
		kindName = ImageSmallName
	case con.Cooking.MicroID, con.Other.MicroID:
		kindName = ImageMicroName
	default:
		panic("画像種類名が取得できませんでした。")
	}

	return kindName
}

//getKindByName get kind from config
func (cook *CookingConfig) getKindByName(name string) int {
	var kind int

	switch name {
	case ImageOriginNoLogoName:
		kind = cook.OriginNoLogoID
	case ImageOriginName:
		kind = cook.OriginID
	case ImageLargeName:
		kind = cook.LargeID
	case ImageMediumName:
		kind = cook.MediumID
	case ImageSmallName:
		kind = cook.SmallID
	case ImageMicroName:
		kind = cook.MicroID
	}
	return kind
}

//getKindByName get kind from config
func (other *OtherConfig) getKingByName(name string) int {
	var kind int

	switch name {
	case ImageOriginNoLogoName:
		kind = other.OriginNoLogoID
	case ImageOriginName:
		kind = other.OriginID
	case ImageLargeName:
		kind = other.LargeID
	case ImageMediumName:
		kind = other.MediumID
	case ImageSmallName:
		kind = other.SmallID
	case ImageMicroName:
		kind = other.MicroID
	}
	return kind
}

//GetImageKindByKindNameAndTypeName
func (con *Configs) GetKindByKindNameAndTypeName(kindN string, typeN string) int {
	var kind int

	if typeN == TypeCookingName {
		kind = con.Cooking.getKindByName(kindN)
	} else {
		kind = con.Other.getKingByName(kindN)
	}

	return kind
}

//GetImageTypeByKind return kind name
func (con *Configs) GetImageTypeByKind(kind int) string {

	var typeName string

	switch kind {
	case con.Cooking.OriginNoLogoID, con.Cooking.OriginID, con.Cooking.LargeID,
		con.Cooking.MediumID, con.Cooking.MicroID, con.Cooking.SmallID:
		typeName = TypeCookingName
	case con.Other.OriginNoLogoID, con.Other.OriginID, con.Other.LargeID,
		con.Other.MediumID, con.Other.MicroID, con.Other.SmallID:
		typeName = TypeOtherName
	default:
		panic("logo kind が一致しませんでした。")
	}
	return typeName
}

func init() {
	_, err := toml.DecodeFile("./config.toml", &Config)
	if err != nil {
		panic(err)
	}

	basePath, _ := os.Getwd()
	CurrentBasePath = basePath
}
