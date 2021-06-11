INSERT INTO core_user (
    use_id,
    use_name,
    use_middle_name,
    use_last_name,
    use_email,
    use_username,
    use_status
) VALUES (
    '00000001-0001-0001-0001-000000000001',
    'Phoenix',
    'J',
    '',
    'phoenix@example.com',
    'phoenix',
    2
),
(
    '00000001-0001-0001-0001-000000000002',
    'all',
    '',
    '',
    'all@example.com',
    'all',
    2
);

INSERT INTO core_password (
    pas_data,
    pas_type,
    pas_user_id,
    pas_status
) VALUES (
    '{"cleartext":"1", "password":"76bd4a1abf79b2ae1a0d90207c9fcd0cf279a3e0ab804717ac41dcee2fbe3beb01e04775478ffa1de1e1912721541647ce8d2faa21503829248acf7854a5d5e64d41653ef837d6bac141abe30f4f67a966f02e89f19bf524deb9ebbe1a47e89148946df9e1a640d5f3b543d3bf875fa673376f3c7ab264f5ddf3e6135fe28d6ee63842ad5d8a42e80f42cf473e0ccae94ad2cfa16138b7698acd5625cf649ad5eeab17c94c74f4e7f675c53c6387e65dc07223a7e5797ee673fe33bb450d336eed333e7aa9d748e34722f087d9e2c3a496becd63c6b51ae5668441258fcb2a506ee31704a797ffae9b1d5d68491cef71eab02278fc5203aa0509c425b43dac6b", "salt":"bd2feff895ac4aa15b2ee47203b75fa079da1a67f200f84d1580d54d8c6505daf8a2c68083c745b21e3db5cfd0639a9e30acc821d87a0951c8696533d7b350b6d1c86e55d6360f20fcf0b580b2ef5251f2f8ac8727dda2fee8cc5e48fb21df050ad700e5a05bbd7a62eb7319edd51040a799154104080e1103c2ccdf2e71e1c5"}',
    'Scrypt2017',
    '00000001-0001-0001-0001-000000000001',
    2
), (
    '{"cleartext":"1", "password":"76bd4a1abf79b2ae1a0d90207c9fcd0cf279a3e0ab804717ac41dcee2fbe3beb01e04775478ffa1de1e1912721541647ce8d2faa21503829248acf7854a5d5e64d41653ef837d6bac141abe30f4f67a966f02e89f19bf524deb9ebbe1a47e89148946df9e1a640d5f3b543d3bf875fa673376f3c7ab264f5ddf3e6135fe28d6ee63842ad5d8a42e80f42cf473e0ccae94ad2cfa16138b7698acd5625cf649ad5eeab17c94c74f4e7f675c53c6387e65dc07223a7e5797ee673fe33bb450d336eed333e7aa9d748e34722f087d9e2c3a496becd63c6b51ae5668441258fcb2a506ee31704a797ffae9b1d5d68491cef71eab02278fc5203aa0509c425b43dac6b", "salt":"bd2feff895ac4aa15b2ee47203b75fa079da1a67f200f84d1580d54d8c6505daf8a2c68083c745b21e3db5cfd0639a9e30acc821d87a0951c8696533d7b350b6d1c86e55d6360f20fcf0b580b2ef5251f2f8ac8727dda2fee8cc5e48fb21df050ad700e5a05bbd7a62eb7319edd51040a799154104080e1103c2ccdf2e71e1c5"}',
    'Scrypt2017',
    '00000001-0001-0001-0001-000000000002',
    2
);

INSERT INTO core_role(
    rol_id,
    rol_name,
    rol_description,
    rol_icon,
    rol_status
) VALUES
    ('SYS_ADM', 'System admin', 'the system admin has acceess to all the actions in the system', 'stars', 2);

INSERT INTO core_module(
    mod_id,
    mod_name,
    mod_description,
    mod_order,
    mod_parent_id,
    mod_status
) VALUES
    ('CORE_NAVIGATOR', 'Navigator', 'Administration for user navigator.', 1, NULL, 2),
    ('CORE_MODULE', 'Modules', 'Administration for system modules.', 2, NULL, 2),
    ('CORE_ACTION', 'Module Actions', 'Administration for system module actions.', 3, NULL, 2);


INSERT INTO core_action(
    act_module_id,
    act_action_id,
    act_name,
    act_description,
    act_order,
    act_status
) VALUES
    ('CORE_NAVIGATOR', 'EDIT_ALL',   'Create, Edit and delete the records information.', '', 1, 2),
    ('CORE_NAVIGATOR', 'QUERY_ONE',  'Query one record.', '', 1, 2),
    ('CORE_NAVIGATOR', 'QUERY_ALL',  'Query the records.', '', 1, 2),
    ('CORE_NAVIGATOR', 'ACTIVATE',   'Activate records.', '', 1, 2),
    ('CORE_NAVIGATOR', 'INACTIVATE', 'Inactivate records.', '', 1, 2),
    ('CORE_NAVIGATOR', 'ROLE',       'Associate and dissociate to roles.', '', 1, 2),

    ('CORE_MODULE', 'NEW',        'Create a new record.', '', 1, 2),
    ('CORE_MODULE', 'EDIT',       'Edit the records information.', '', 1, 2),
    ('CORE_MODULE', 'QUERY_ONE',  'Query one record.', '', 1, 2),
    ('CORE_MODULE', 'QUERY_ALL',  'Query all the records.', '', 1, 2),
    ('CORE_MODULE', 'QUERY_ACTV', 'Query only the active records.', '', 1, 2),
    ('CORE_MODULE', 'ACTIVATE',   'Activate records.', '', 1, 2),
    ('CORE_MODULE', 'INACTIVATE', 'Inactivate records.', '', 1, 2),
    ('CORE_MODULE', 'HDELETE',    'Hard delete the records.', '', 1, 2),

    ('CORE_ACTION', 'NEW',        'Create a new record.', '', 1, 2),
    ('CORE_ACTION', 'EDIT',       'Edit the records information.', '', 1, 2),
    ('CORE_ACTION', 'QUERY_ONE',  'Query one record.', '', 1, 2),
    ('CORE_ACTION', 'QUERY_ALL',  'Query all the records.', '', 1, 2),
    ('CORE_ACTION', 'QUERY_ACTV', 'Query only the active records.', '', 1, 2),
    ('CORE_ACTION', 'ACTIVATE',   'Activate records.', '', 1, 2),
    ('CORE_ACTION', 'INACTIVATE', 'Inactivate records.', '', 1, 2),
    ('CORE_ACTION', 'HDELETE',    'Hard delete the records.', '', 1, 2);


INSERT INTO core_action_route (
    acr_module_id,
    acr_action_id,
    acr_method,
    acr_route
) VALUES
    ('CORE_NAVIGATOR', 'EDIT_ALL',   'POST', '/api/core/navigator/upsert'),
    ('CORE_NAVIGATOR', 'QUERY_ONE',  'GET',  '/api/core/navigator/elements/element/:id'),
    ('CORE_NAVIGATOR', 'QUERY_ALL',  'GET',  '/api/core/navigator/elements'),
    ('CORE_NAVIGATOR', 'ACTIVATE',   'PUT',  '/api/core/navigator/elements/element/activate'),
    ('CORE_NAVIGATOR', 'INACTIVATE', 'PUT',  '/api/core/navigator/elements/element/inactivate'),
    ('CORE_NAVIGATOR', 'ROLE',       'GET',  '/api/core/navigator/elements/role/:roleID'),
    ('CORE_NAVIGATOR', 'ROLE',       'PUT',  '/api/core/navigator/elements/element/associate-role'),
    ('CORE_NAVIGATOR', 'ROLE',       'PUT',  '/api/core/navigator/elements/element/dissociate-role'),

    ('CORE_MODULE', 'NEW',        'POST', '/api/core/modules/module'),
    ('CORE_MODULE', 'EDIT',       'PUT',  '/api/core/modules/module'),
    ('CORE_MODULE', 'QUERY_ONE',  'GET',  '/api/core/modules/module/:moduleID'),
    ('CORE_MODULE', 'QUERY_ALL',  'GET',  '/api/core/modules'),
    ('CORE_MODULE', 'QUERY_ALL',  'POST', '/api/core/modules'),
    ('CORE_MODULE', 'QUERY_ACTV', 'GET',  '/api/core/modules/active-records'),
    ('CORE_MODULE', 'QUERY_ACTV', 'POST', '/api/core/modules/active-records'),
    ('CORE_MODULE', 'ACTIVATE',   'PUT',  '/api/core/modules/module/activate'),
    ('CORE_MODULE', 'INACTIVATE', 'PUT',  '/api/core/modules/module/inactivate'),
    ('CORE_MODULE', 'HDELETE',    'PUT',  '/api/core/modules/module/hard-delete'),

    ('CORE_ACTION', 'NEW',        'POST', '/api/core/modules/actions/action'),
    ('CORE_ACTION', 'EDIT',       'PUT',  '/api/core/modules/actions/action'),
    ('CORE_ACTION', 'QUERY_ONE',  'GET',  '/api/core/modules/module/:moduleID/actions/action/:actionID'),
    ('CORE_ACTION', 'QUERY_ALL',  'GET',  '/api/core/modules/module/:moduleID/actions'),
    ('CORE_ACTION', 'QUERY_ALL',  'POST', '/api/core/modules/module/:moduleID/actions'),
    ('CORE_ACTION', 'QUERY_ACTV', 'GET',  '/api/core/modules/module/:moduleID/actions/active-records'),
    ('CORE_ACTION', 'QUERY_ACTV', 'POST', '/api/core/modules/module/:moduleID/actions/active-records'),
    ('CORE_ACTION', 'ACTIVATE',   'PUT',  '/api/core/modules/actions/action/activate'),
    ('CORE_ACTION', 'INACTIVATE', 'PUT',  '/api/core/modules/actions/action/inactivate'),
    ('CORE_ACTION', 'HDELETE',    'PUT',  '/api/core/modules/actions/action/hard-delete');

INSERT INTO core_role(
    rol_id,
    rol_name,
    rol_description,
    rol_icon,
    rol_status
) VALUES
    ('CORE_NAVIGATOR', 'CORE_NAVIGATOR',  'CORE_NAVIGATOR', 'adjust', 2),
    ('CORE_MODULE',    'CORE_MODULE',     'CORE_MODULE', 'adjust', 2),
    ('CORE_ACTION',    'CORE_ACTION',     'CORE_ACTION', 'adjust', 2);

INSERT INTO core_privilege(
    pri_role_id,
    pri_module_id,
    pri_action_id
) SELECT rol_id, act_module_id, act_action_id
    FROM core_role
    INNER JOIN core_action
    ON rol_id = act_module_id;


INSERT INTO core_user_role(
    usr_user_id,
    usr_role_id
) VALUES
    ('00000001-0001-0001-0001-000000000001', 'SYS_ADM'),
    ('00000001-0001-0001-0001-000000000002', 'CORE_MODULE'),
    ('00000001-0001-0001-0001-000000000002', 'CORE_ACTION');

