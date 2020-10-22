--查询菜单是否已经存在 
SELECT * FROM `sys_base_menus` WHERE `sys_base_menus`.`deleted_at` IS NULL AND((name = 'orderManage')); 

-- 添加根菜单 
INSERT INTO `sys_base_menus` ( `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon`) VALUES ( CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP(), NULL, 0, '0', 'daigou/orderManage', 'orderManage', false, 'view/dgOrderManage/index', 11, TRUE, false, '订单管理', 's-operation' );

-- 添加子菜单 
INSERT INTO `sys_base_menus` ( `created_at`, `updated_at`, `deleted_at`, `menu_level`, `parent_id`, `path`, `name`, `hidden`, `component`, `sort`, `keep_alive`, `default_menu`, `title`, `icon` ) VALUES ( '2020-08-28 10:12:39', '2020-08-28 10:12:39', NULL, 0, '64', 'daigou/orderList', 'orderList', false, 'view/dgOrderManage', 1, TRUE, false, '订单列表', '' ) 

-- 给角色增加菜单权限 
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`,`sys_base_menu_id`) SELECT '888',64 FROM DUAL WHERE NOT EXISTS (SELECT * FROM `sys_authority_menus` WHERE `sys_authority_authority_id` = '888' AND `sys_base_menu_id` = 64)
INSERT INTO `sys_authority_menus` (`sys_authority_authority_id`,`sys_base_menu_id`) SELECT '888',65 FROM DUAL WHERE NOT EXISTS (SELECT * FROM `sys_authority_menus` WHERE `sys_authority_authority_id` = '888' AND `sys_base_menu_id` = 65) 

-- 增加api
INSERT INTO `sys_apis`(`created_at`, `updated_at`, `deleted_at`, `authority_id`, `path`, `description`, `api_group`, `method`) VALUES ('2020-08-24 14:42:01', '2020-08-28 09:49:33', NULL, NULL, '/daigou/shopManage/createShopInfo', '新增shopInfo表', 'daigou/shopManage', 'POST');
INSERT INTO `sys_apis`(`created_at`, `updated_at`, `deleted_at`, `authority_id`, `path`, `description`, `api_group`, `method`) VALUES ('2020-08-24 14:42:01', '2020-08-28 09:49:45', NULL, NULL, '/daigou/shopManage/deleteShopInfo', '删除shopInfo表', 'daigou/shopManage', 'DELETE');
INSERT INTO `sys_apis`(`created_at`, `updated_at`, `deleted_at`, `authority_id`, `path`, `description`, `api_group`, `method`) VALUES ('2020-08-24 14:42:02', '2020-08-28 09:50:01', NULL, NULL, '/daigou/shopManage/updateShopInfo', '更新shopInfo表', 'daigou/shopManage', 'PUT');
INSERT INTO `sys_apis`(`created_at`, `updated_at`, `deleted_at`, `authority_id`, `path`, `description`, `api_group`, `method`) VALUES ('2020-08-24 14:42:02', '2020-08-28 09:50:47', NULL, NULL, '/daigou/shopManage/findShopInfo', '查找ShopInfo', 'daigou/shopManage', 'GET');
INSERT INTO `sys_apis`(`created_at`, `updated_at`, `deleted_at`, `authority_id`, `path`, `description`, `api_group`, `method`) VALUES ('2020-08-24 14:42:02', '2020-08-28 09:50:14', NULL, NULL, '/daigou/shopManage/getShopInfoList', '获取shopInfo表列表', 'daigou/shopManage', 'GET');

-- 增加casbin权限
INSERT INTO `casbin_rule`(`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/daigou/shopManage/createShopInfo', 'POST', '', '', '');
INSERT INTO `casbin_rule`(`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/daigou/shopManage/deleteShopInfo', 'DELETE', '', '', '');
INSERT INTO `casbin_rule`(`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/daigou/shopManage/updateShopInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule`(`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/daigou/shopManage/findShopInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule`(`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', '888', '/daigou/shopManage/getShopInfoList', 'GET', '', '', '');


