use MarketMeDb;

-- Init Script for MarketMe database
-- -----------------------------------------------------
-- Table `Company`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `Company` (
     `Company_Id` INT NOT NULL AUTO_INCREMENT,
     `CompanyName` VARCHAR(100) CHARACTER SET 'utf16' NOT NULL,
     `MyTest` FLOAT     (6, 2) NOT NULL DEFAULT 90,
     `Country` VARCHAR(100) CHARACTER SET 'utf16' COLLATE 'utf16_bin' NOT NULL,
     PRIMARY KEY (`Company_Id`),
     UNIQUE INDEX `idCompany_UNIQUE` (`Company_Id` ASC) )
     ENGINE = InnoDB
     AUTO_INCREMENT = 1
     DEFAULT CHARACTER SET = utf16
     COLLATE = utf16_bin;

-- -----------------------------------------------------
-- Table `Configuration`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `Configuration` (
                                               `Configuration_Id` INT NOT NULL AUTO_INCREMENT,
                                               `ConfigurationKey` VARCHAR(50) CHARACTER SET 'utf16' NOT NULL,
    `ConfigurationValue` VARCHAR(500) CHARACTER SET 'utf16' NULL DEFAULT NULL,
    PRIMARY KEY (`Configuration_Id`),
    UNIQUE INDEX `Configuration_Id_UNIQUE` (`Configuration_Id` ASC) )
    ENGINE = InnoDB
    AUTO_INCREMENT = 1
    DEFAULT CHARACTER SET = utf16
    COLLATE = utf16_bin;

-- -----------------------------------------------------
-- Table `FileRef`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `FileRef` (
                                         `FileRef_Id` INT NOT NULL AUTO_INCREMENT,
                                         `Filename` VARCHAR(100) CHARACTER SET 'utf16' NOT NULL,
    `Path` VARCHAR(255) CHARACTER SET 'utf16' NOT NULL,
    `ContentType` VARCHAR(255) CHARACTER SET 'utf16' NOT NULL,
    PRIMARY KEY (`FileRef_Id`),
    UNIQUE INDEX `FileRef_Id_UNIQUE` (`FileRef_Id` ASC) )
    ENGINE = InnoDB
    AUTO_INCREMENT = 1
    DEFAULT CHARACTER SET = utf16
    COLLATE = utf16_bin;


-- -----------------------------------------------------
-- Table `ProductRange`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ProductRange` (
                                              `ProductRange_Id` INT NOT NULL AUTO_INCREMENT,
                                              `ProductRangeName` VARCHAR(100) CHARACTER SET 'utf16' NOT NULL,
    PRIMARY KEY (`ProductRange_Id`),
    UNIQUE INDEX `idRange_UNIQUE` (`ProductRange_Id` ASC) )
    ENGINE = InnoDB
    AUTO_INCREMENT = 1
    DEFAULT CHARACTER SET = utf16
    COLLATE = utf16_bin;

-- -----------------------------------------------------
-- Table `Template`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `Template` (
                                          `Template_Id` INT NOT NULL AUTO_INCREMENT,
                                          `TemplateName` VARCHAR(200) CHARACTER SET 'utf16' NOT NULL,
    `Company_Id` INT NOT NULL,
    `ProductRange_Id` INT NOT NULL,
    `TemplateDescription` VARCHAR(500) CHARACTER SET 'utf16' NULL DEFAULT NULL,
    `CreatedAt` DATETIME NOT NULL,
    `UpdatedAt` DATETIME NULL,
    `PublishedAt` DATETIME NULL,
    `TemplateVersion` VARCHAR(20) CHARACTER SET 'utf16' NOT NULL,
    `Status` VARCHAR(20) CHARACTER SET 'utf16' NOT NULL,
    PRIMARY KEY (`Template_Id`),
    UNIQUE INDEX `idTemplate_UNIQUE` (`Template_Id` ASC) ,
    INDEX `Fk_Company_idx` (`Company_Id` ASC) ,
    INDEX `Fk_ProductRange_idx` (`ProductRange_Id` ASC) ,
    CONSTRAINT `Fk_CompanyId_Template` FOREIGN KEY (`Company_Id`) REFERENCES `Company` (`Company_Id`) ON UPDATE CASCADE,
    CONSTRAINT `Fk_ProductRange_Id_Template`
    FOREIGN KEY (`ProductRange_Id`)
    REFERENCES `ProductRange` (`ProductRange_Id`) ON DELETE RESTRICT
    )

    ENGINE = InnoDB
    AUTO_INCREMENT = 1
    DEFAULT CHARACTER SET = utf16
    COLLATE = utf16_bin;
#
# -- -----------------------------------------------------
# -- Table `OrderableContent`
# -- -----------------------------------------------------
# CREATE TABLE IF NOT EXISTS `OrderableContent` (
#                                                   `OrderableContent_Id` INT NOT NULL AUTO_INCREMENT,
#                                                   `Type` ENUM('Family_Template','Product_Family_Template') NOT NULL,
#     `Position` INT NULL,
#     `Template_Id` INT NOT NULL,
#     `ParentFamilyTemplate_Id` INT NULL,
#     `TileJson` JSON NOT NULL,
#     PRIMARY KEY (`OrderableContent_Id`),
#     UNIQUE INDEX `OrderableContent_Id_UNIQUE` (`OrderableContent_Id` ASC),
#     CONSTRAINT `Fk_TemplateId_FamTplt`
#     FOREIGN KEY (`Template_Id`)
#     REFERENCES `Template` (`Template_Id`)
#     )
#     ENGINE = InnoDB
#     AUTO_INCREMENT = 1
#     DEFAULT CHARACTER SET = utf16
#     COLLATE = utf16_bin;
#
#
# -- -----------------------------------------------------
# -- Table `Family_Template`
# -- -----------------------------------------------------
# CREATE TABLE IF NOT EXISTS `Family_Template` (
#                                                  `FamilyTemplate_Id` INT NOT NULL,
#                                                  `Family_TemplateName` VARCHAR(100) CHARACTER SET 'utf16' NOT NULL,
#     `ThumbnailFileRefId` INT NULL,
#     `TextFileRefId` INT NULL,
#     `IsEnabled` TINYINT(1) NOT NULL,
#     `IsAllProduct` TINYINT(1) NOT NULL,
#     PRIMARY KEY (`FamilyTemplate_Id`),
#     UNIQUE INDEX `FamilyTemplate_Id_UNIQUE` (`FamilyTemplate_Id` ASC) ,
#     INDEX `Fk_TemplateId_FamTplt_idx` (`FamilyTemplate_Id` ASC) ,
#     INDEX `Fk_ThumbnailFileRefId_idx` (`ThumbnailFileRefId` ASC) ,
#     INDEX `Fk_TextFileRefId_idx` (`TextFileRefId` ASC) ,
#     CONSTRAINT `Fk_FamilyTemplate_Id` FOREIGN KEY (`FamilyTemplate_Id`) REFERENCES `OrderableContent` (`OrderableContent_Id`),
#
#     CONSTRAINT `Fk_ThumbnailFileRefId_Family_Template` FOREIGN KEY (`ThumbnailFileRefId`) REFERENCES `FileRef` (`FileRef_Id`),
#     CONSTRAINT `Fk_TextFileRefId_Family_Template` FOREIGN KEY (`TextFileRefId`) REFERENCES `FileRef` (`FileRef_Id`))
#
#     ENGINE = InnoDB
#     AUTO_INCREMENT = 1
#     DEFAULT CHARACTER SET = utf16
#     COLLATE = utf16_bin;
#
#
#
# -- -----------------------------------------------------
# -- Table `ProductType`
# -- -----------------------------------------------------
# CREATE TABLE IF NOT EXISTS `ProductType` (
#                                              `ProductType_Id` INT NOT NULL AUTO_INCREMENT,
#                                              `ProductTypeName` VARCHAR(100) NOT NULL,
#     PRIMARY KEY (`ProductType_Id`),
#     UNIQUE INDEX `idType_UNIQUE` (`ProductType_Id` ASC))
#     ENGINE = InnoDB
#     AUTO_INCREMENT = 1
#     DEFAULT CHARACTER SET = utf16;
#
# -- -----------------------------------------------------
# -- Table `Product`
# -- -----------------------------------------------------
# CREATE TABLE IF NOT EXISTS `Product` (
#                                          `Product_Id` INT NOT NULL AUTO_INCREMENT,
#                                          `ProductName` VARCHAR(100) CHARACTER SET 'utf16' NOT NULL,
#     `ProductVersion` VARCHAR(20) CHARACTER SET 'utf16' NOT NULL,
#     `Company_Id` INT NOT NULL,
#     `ProductRange_Id` INT NOT NULL,
#     `ProductType_Id` INT NOT NULL,
#     PRIMARY KEY (`Product_Id`),
#     UNIQUE INDEX `ProductTemplate_Id_UNIQUE` (`Product_Id` ASC) ,
#     UNIQUE INDEX `Product_NameVersion_UNIQUE` (`ProductName`, `ProductVersion`, `Company_Id`, `ProductRange_Id` ASC) ,
#     INDEX `FK_Compagny_Id_idx` (`Company_Id` ASC) ,
#     INDEX `FK_ProductRange_Id_idx` (`ProductRange_Id` ASC) ,
#     INDEX `FK_ProductType_Id_idx` (`ProductType_Id` ASC) ,
#
#     CONSTRAINT `FK_Compagny_Id` FOREIGN KEY (`Company_Id`) REFERENCES `Company` (`Company_Id`),
#     CONSTRAINT `FK_ProductRange_Id` FOREIGN KEY (`ProductRange_Id`) REFERENCES `ProductRange` (`ProductRange_Id`),
#     CONSTRAINT `FK_ProductType_Id` FOREIGN KEY (`ProductType_Id`) REFERENCES `ProductType` (`ProductType_Id`))
#
#     ENGINE = InnoDB
#     AUTO_INCREMENT = 1
#     DEFAULT CHARACTER SET = utf16
#     COLLATE = utf16_bin;
#
# -- -----------------------------------------------------
# -- Table `ProductTemplate`
# -- -----------------------------------------------------
# CREATE TABLE IF NOT EXISTS `ProductTemplate` (
#                                                  `ProductTemplate_Id` INT NOT NULL AUTO_INCREMENT,
#                                                  `ProductTemplateName` VARCHAR(100) CHARACTER SET 'utf16' NOT NULL,
#     `ThumbnailFileRefId` INT NOT NULL,
#     `TextFileRefId` INT NOT NULL,
#     `Product_Id` INT NOT NULL,
#     PRIMARY KEY (`ProductTemplate_Id`),
#     UNIQUE INDEX `ProductTemplate_Id_UNIQUE` (`ProductTemplate_Id` ASC) ,
#     INDEX `FK_Product_Id_idx` (`Product_Id` ASC) ,
#     INDEX `Fk_ThumbnailFileRefId_idx` (`ThumbnailFileRefId` ASC) ,
#     INDEX `Fk_TextFileRefId_idx` (`TextFileRefId` ASC) ,
#
#     CONSTRAINT `FK_Product_Id` FOREIGN KEY (`Product_Id`) REFERENCES `Product` (`Product_Id`),
#     CONSTRAINT `Fk_ThumbnailFileRefId_Product_Template` FOREIGN KEY (`ThumbnailFileRefId`) REFERENCES `FileRef` (`FileRef_Id`),
#     CONSTRAINT `Fk_TextFileRefId_Product_Template` FOREIGN KEY (`TextFileRefId`) REFERENCES `FileRef` (`FileRef_Id`))
#
#     ENGINE = InnoDB
#     AUTO_INCREMENT = 1
#     DEFAULT CHARACTER SET = utf16
#     COLLATE = utf16_bin;
#
# -- -----------------------------------------------------
# -- Table `HardwareRequirement`
# -- -----------------------------------------------------
# CREATE TABLE IF NOT EXISTS `HardwareRequirement` (
#                                                      `HardwareRequirement_Id` INT NOT NULL AUTO_INCREMENT,
#                                                      `HardwareRequirementName` VARCHAR(200) CHARACTER SET 'utf16' NOT NULL,
#     `ProductTemplate_Id` INT NOT NULL,
#     PRIMARY KEY (`HardwareRequirement_Id`),
#     UNIQUE INDEX `HardwareRequirement_Id_UNIQUE` (`HardwareRequirement_Id` ASC),
#     CONSTRAINT `Fk_ProductTemplate_Id`
#     FOREIGN KEY (`ProductTemplate_Id`)
#     REFERENCES `ProductTemplate` (`ProductTemplate_Id`)
#     )
#     ENGINE = InnoDB
#     AUTO_INCREMENT = 1
#     DEFAULT CHARACTER SET = utf16
#     COLLATE = utf16_bin;
#
# -- -----------------------------------------------------
# -- Table `Product_Family_Template`
# -- -----------------------------------------------------
# CREATE TABLE IF NOT EXISTS `Product_Family_Template` (
#                                                          `Product_Family_Template_Id` INT NOT NULL,
#                                                          `ProductTemplate_Id` INT NOT NULL,
#                                                          PRIMARY KEY (`Product_Family_Template_Id`),
#     UNIQUE INDEX `Product_Family_Template_Id_UNIQUE` (`Product_Family_Template_Id` ASC) ,
#     INDEX `Fk_ProductTemplate_Id_idx` (`ProductTemplate_Id` ASC) ,
#     CONSTRAINT `Fk_Product_Family_Template_Id`
#     FOREIGN KEY (`Product_Family_Template_Id`)
#     REFERENCES `OrderableContent` (`OrderableContent_Id`),
#     CONSTRAINT `Fk_ProductId_ProdFamTplt`
#     FOREIGN KEY (`ProductTemplate_Id`)
#     REFERENCES `ProductTemplate` (`ProductTemplate_Id`))
#     ENGINE = InnoDB
#     AUTO_INCREMENT = 1
#     DEFAULT CHARACTER SET = utf16
#     COLLATE = utf16_bin;
#
# -- -----------------------------------------------------
# -- Table `SlideShowMedia`
# -- -----------------------------------------------------
# CREATE TABLE IF NOT EXISTS `SlideShowMedia` (
#                                                 `SlideShowMedia_Company_Id` INT NOT NULL AUTO_INCREMENT,
#                                                 `SlideShowMediaName` VARCHAR(100) CHARACTER SET 'utf16' NOT NULL,
#     `Company_Id` INT NOT NULL,
#     `ProductRange_Id` INT NOT NULL,
#     `MediaFileRefId` INT NULL,
#     PRIMARY KEY (`SlideShowMedia_Company_Id`),
#     UNIQUE INDEX `SlideShowMedia_Country_Id_UNIQUE` (`SlideShowMedia_Company_Id` ASC) ,
#     INDEX `Fk_CompanyId_idx` (`Company_Id` ASC) ,
#     INDEX `Fk_ProductRangeId_idx` (`ProductRange_Id` ASC) ,
#     INDEX `Fk_MediaFileRefId_idx` (`MediaFileRefId` ASC) ,
#     CONSTRAINT `Fk_CompanyId3`
#     FOREIGN KEY (`Company_Id`)
#     REFERENCES `Company` (`Company_Id`),
#     CONSTRAINT `Fk_ProductRangeId`
#     FOREIGN KEY (`ProductRange_Id`)
#     REFERENCES `ProductRange` (`ProductRange_Id`),
#     CONSTRAINT `Fk_MediaFileRefId`
#     FOREIGN KEY (`MediaFileRefId`)
#     REFERENCES `FileRef` (`FileRef_Id`))
#     ENGINE = InnoDB
#     DEFAULT CHARACTER SET = utf16
#     COLLATE = utf16_bin;
#
# -- -----------------------------------------------------
# -- Table `SlideShowMedia_Template`
# -- -----------------------------------------------------
# CREATE TABLE IF NOT EXISTS `SlideShowMedia_Template` (
#                                                          `SlideShowMedia_Template_Id` INT NOT NULL AUTO_INCREMENT,
#                                                          `Template_Id` INT NOT NULL,
#                                                          `MediaIndex` INT NOT NULL,
#                                                          `SlideShowMedia_Id` INT NOT NULL,
#                                                          PRIMARY KEY (`SlideShowMedia_Template_Id`),
#     UNIQUE INDEX `Id_UNIQUE` (`SlideShowMedia_Template_Id` ASC),
#     INDEX `Id` (`SlideShowMedia_Template_Id` ASC) ,
#     INDEX `Fk_SlideShowMediaId_idx` (`SlideShowMedia_Id` ASC) ,
#     INDEX `Fk_TemplateId_idx` (`Template_Id` ASC) ,
#     CONSTRAINT `Fk_SlideShowMedia_Id`
#     FOREIGN KEY (`SlideShowMedia_Id`)
#     REFERENCES `SlideShowMedia` (`SlideShowMedia_Company_Id`),
#     CONSTRAINT `Fk_Template_Id`
#     FOREIGN KEY (`Template_Id`)
#     REFERENCES `Template` (`Template_Id`))
#     ENGINE = InnoDB
#     DEFAULT CHARACTER SET = utf16
#     COLLATE = utf16_bin;
