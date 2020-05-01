package tfvm

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const defaultStateJson = `
{
 "lastUpdateTime": "2020-05-01T16:35:41+02:00",
 "terraformReleases": [
  {
   "version": "0.1.0"
  },
  {
   "version": "0.1.1"
  },
  {
   "version": "0.2.0"
  },
  {
   "version": "0.2.1"
  },
  {
   "version": "0.2.2"
  },
  {
   "version": "0.3.0"
  },
  {
   "version": "0.3.1"
  },
  {
   "version": "0.3.5"
  },
  {
   "version": "0.3.6"
  },
  {
   "version": "0.3.7"
  },
  {
   "version": "0.4.0"
  },
  {
   "version": "0.4.1"
  },
  {
   "version": "0.4.2"
  },
  {
   "version": "0.5.0"
  },
  {
   "version": "0.5.1"
  },
  {
   "version": "0.5.3"
  },
  {
   "version": "0.6.0"
  },
  {
   "version": "0.6.1"
  },
  {
   "version": "0.6.2"
  },
  {
   "version": "0.6.3"
  },
  {
   "version": "0.6.4"
  },
  {
   "version": "0.6.5"
  },
  {
   "version": "0.6.6"
  },
  {
   "version": "0.6.7"
  },
  {
   "version": "0.6.8"
  },
  {
   "version": "0.6.9"
  },
  {
   "version": "0.6.10"
  },
  {
   "version": "0.6.11"
  },
  {
   "version": "0.6.12"
  },
  {
   "version": "0.6.13"
  },
  {
   "version": "0.6.14"
  },
  {
   "version": "0.6.15"
  },
  {
   "version": "0.6.16"
  },
  {
   "version": "0.7.0"
  },
  {
   "version": "0.7.1"
  },
  {
   "version": "0.7.2"
  },
  {
   "version": "0.7.3"
  },
  {
   "version": "0.7.4"
  },
  {
   "version": "0.7.5"
  },
  {
   "version": "0.7.6"
  },
  {
   "version": "0.7.7"
  },
  {
   "version": "0.7.8"
  },
  {
   "version": "0.7.9"
  },
  {
   "version": "0.7.10"
  },
  {
   "version": "0.7.11"
  },
  {
   "version": "0.7.12"
  },
  {
   "version": "0.7.13"
  },
  {
   "version": "0.8.0"
  },
  {
   "version": "0.8.1"
  },
  {
   "version": "0.8.2"
  },
  {
   "version": "0.8.3"
  },
  {
   "version": "0.8.4"
  },
  {
   "version": "0.8.5"
  },
  {
   "version": "0.8.6"
  },
  {
   "version": "0.8.7"
  },
  {
   "version": "0.8.8"
  },
  {
   "version": "0.9.0"
  },
  {
   "version": "0.9.1"
  },
  {
   "version": "0.9.2"
  },
  {
   "version": "0.9.3"
  },
  {
   "version": "0.9.4"
  },
  {
   "version": "0.9.5"
  },
  {
   "version": "0.9.6"
  },
  {
   "version": "0.9.7"
  },
  {
   "version": "0.9.8"
  },
  {
   "version": "0.9.9"
  },
  {
   "version": "0.9.10"
  },
  {
   "version": "0.9.11"
  },
  {
   "version": "0.10.0-beta1"
  },
  {
   "version": "0.10.0-beta2"
  },
  {
   "version": "0.10.0-rc1"
  },
  {
   "version": "0.10.0"
  },
  {
   "version": "0.10.1"
  },
  {
   "version": "0.10.2"
  },
  {
   "version": "0.10.3"
  },
  {
   "version": "0.10.4"
  },
  {
   "version": "0.10.5"
  },
  {
   "version": "0.10.6"
  },
  {
   "version": "0.10.7"
  },
  {
   "version": "0.10.8"
  },
  {
   "version": "0.11.0-beta1"
  },
  {
   "version": "0.11.0-rc1"
  },
  {
   "version": "0.11.0"
  },
  {
   "version": "0.11.1"
  },
  {
   "version": "0.11.2"
  },
  {
   "version": "0.11.3"
  },
  {
   "version": "0.11.4"
  },
  {
   "version": "0.11.5"
  },
  {
   "version": "0.11.6"
  },
  {
   "version": "0.11.7"
  },
  {
   "version": "0.11.8"
  },
  {
   "version": "0.11.9-beta1"
  },
  {
   "version": "0.11.9"
  },
  {
   "version": "0.11.10"
  },
  {
   "version": "0.11.11"
  },
  {
   "version": "0.11.12-beta1"
  },
  {
   "version": "0.11.12"
  },
  {
   "version": "0.11.13"
  },
  {
   "version": "0.11.14"
  },
  {
   "version": "0.11.15-oci"
  },
  {
   "version": "0.12.0-alpha1"
  },
  {
   "version": "0.12.0-alpha2"
  },
  {
   "version": "0.12.0-alpha3"
  },
  {
   "version": "0.12.0-alpha4"
  },
  {
   "version": "0.12.0-beta1"
  },
  {
   "version": "0.12.0-beta2"
  },
  {
   "version": "0.12.0-rc1"
  },
  {
   "version": "0.12.0"
  },
  {
   "version": "0.12.1"
  },
  {
   "version": "0.12.2"
  },
  {
   "version": "0.12.3"
  },
  {
   "version": "0.12.4"
  },
  {
   "version": "0.12.5"
  },
  {
   "version": "0.12.6"
  },
  {
   "version": "0.12.7"
  },
  {
   "version": "0.12.8"
  },
  {
   "version": "0.12.9"
  },
  {
   "version": "0.12.10"
  },
  {
   "version": "0.12.11"
  },
  {
   "version": "0.12.12"
  },
  {
   "version": "0.12.13"
  },
  {
   "version": "0.12.14"
  },
  {
   "version": "0.12.15"
  },
  {
   "version": "0.12.16"
  },
  {
   "version": "0.12.17"
  },
  {
   "version": "0.12.18"
  },
  {
   "version": "0.12.19"
  },
  {
   "version": "0.12.20"
  },
  {
   "version": "0.12.21"
  },
  {
   "version": "0.12.22"
  },
  {
   "version": "0.12.23"
  },
  {
   "version": "0.12.24"
  }
 ]
}`

type TerraformRelease struct {
	Version string `json:"version"`
}

type StateStruct struct {
	LastUpdateTime    string             `json:"lastUpdateTime"`
	TerraformReleases []TerraformRelease `json:"terraformReleases"`
}

type Inventory struct {
	LastUpdateTime    time.Time
	TerraformReleases []TerraformRelease
}

func GetInventory() (*Inventory, error) {
	inventory := Inventory{LastUpdateTime: time.Now(), TerraformReleases: make([]TerraformRelease, 0)}
	err := inventory.initInventory()
	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

func (inventory *Inventory) GetTerraform(version string) (*Terraform, error) {
	tfPath, err := inventory.getTerraformPath(version)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(tfPath); err != nil {
		return nil, err
	}

	return &Terraform{version: version, path: tfPath}, nil
}

func (inventory *Inventory) GetTerraformBasePath(version string) (string, error) {
	inventoryDir, err := getInventoryDir()
	if err != nil {
		return "", err
	}

	versionedTfPath := filepath.Join(inventoryDir, "v1", "installed", version, fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH))
	return versionedTfPath, nil
}

func (inventory *Inventory) getTerraformPath(version string) (string, error) {
	basePath, err := inventory.GetTerraformBasePath(version)
	if err != nil {
		return "", err
	}

	terraformPath := filepath.Join(basePath, "terraform")
	return terraformPath, nil
}

func (inventory *Inventory) IsTerraformInstalled(version string) (bool, error) {
	versionedTfPath, err := inventory.getTerraformPath(version)
	if err != nil {
		return false, err
	}

	if _, err := os.Stat(versionedTfPath); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (inventory *Inventory) initInventory() error {
	inventoryDir, err := getInventoryDir()
	if err != nil {
		return err
	}

	err = os.MkdirAll(inventoryDir, 0755)
	if err != nil {
		return err
	}

	err = inventory.loadState()

	return err
}

func getInventoryDir() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".tfvm"), nil
}

func (inventory *Inventory) loadState() error {
	statefilepath, err := getStateFilePath()
	if err != nil {
		return err
	}

	file, err := ioutil.ReadFile(statefilepath)
	if err != nil {
		if os.IsNotExist(err) {
			file = []byte(defaultStateJson)
			err = nil
		} else {
			return err
		}
	}

	state := StateStruct{}
	err = json.Unmarshal(file, &state)
	if err != nil {
		return err
	}

	inventory.TerraformReleases = state.TerraformReleases
	inventory.LastUpdateTime, err = time.Parse(time.RFC3339, state.LastUpdateTime)

	return nil
}

func getStateFilePath() (string, error) {
	inventoryDir, err := getInventoryDir()
	if err != nil {
		return "", err
	}
	statefilepath := filepath.Join(inventoryDir, "v1", "state.json")
	return statefilepath, nil
}

func (inventory *Inventory) saveState() error {
	statefilepath, err := getStateFilePath()
	if err != nil {
		return err
	}

	state := StateStruct{}
	state.LastUpdateTime = inventory.LastUpdateTime.Format(time.RFC3339)
	state.TerraformReleases = inventory.TerraformReleases

	file, err := json.MarshalIndent(state, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(statefilepath, file, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (inventory *Inventory) Save() error {
	err := inventory.saveState()
	return err
}
