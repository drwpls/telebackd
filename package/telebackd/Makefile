include $(TOPDIR)/rules.mk

PKG_NAME:=telebackd
PKG_VERSION:=1.0.0
PKG_RELEASE:=1

PKG_BUILD_DIR:=$(BUILD_DIR)/$(PKG_NAME)-$(PKG_VERSION)

GO_FLAG:=-ldflags "-s -w"

include $(INCLUDE_DIR)/package.mk

define Package/$(PKG_NAME)    
	SECTION:=utils
	CATEGORY:=Utilities
	TITLE:=$(PKG_NAME)
	MAINTAINER:=dragino <myEmail.add>
endef

define Package/$(PKG_NAME)/description
	This is a sample program
endef

CONFIGURE_VARS+= \
	#CC="$(TOOLCHAIN_DIR)/bin/$(TARGET_CC)"
	CC=go

define Build/Prepare
	mkdir -p $(PKG_BUILD_DIR)
	$(CP) ./src/* $(PKG_BUILD_DIR)/
endef

define Build/Compile
	$(CC) build -C $(PKG_BUILD_DIR)/ -o $(PKG_BUILD_DIR)/$1 $(GO_FLAGS) 
endef

define Package/$(PKG_NAME)/install
	$(INSTALL_DIR) $(1)/usr/bin
	$(CP) $(PKG_BUILD_DIR)/$(PKG_NAME) $(1)/usr/bin
	$(INSTALL_DIR) $(1)/etc/config/
	$(INSTALL_CONF) ./files/telebackd.config $(1)/etc/config/telebackd
	$(INSTALL_DIR) $(1)/etc/init.d/
	$(INSTALL_BIN) ./files/telebackd.init $(1)/etc/init.d/telebackd
endef

export GOOS=linux
export GOARCH=mipsle
export GOMIPS=softfloat
export CGO_ENABLED=0
$(eval $(call BuildPackage,$(PKG_NAME)))
