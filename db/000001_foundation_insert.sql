INSERT INTO fnd_user (
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
);

INSERT INTO fnd_password (
    pas_data,
    pas_type,
    pas_user_id,
    pas_status
) VALUES (
    '{"cleartext":"1", "password":"76bd4a1abf79b2ae1a0d90207c9fcd0cf279a3e0ab804717ac41dcee2fbe3beb01e04775478ffa1de1e1912721541647ce8d2faa21503829248acf7854a5d5e64d41653ef837d6bac141abe30f4f67a966f02e89f19bf524deb9ebbe1a47e89148946df9e1a640d5f3b543d3bf875fa673376f3c7ab264f5ddf3e6135fe28d6ee63842ad5d8a42e80f42cf473e0ccae94ad2cfa16138b7698acd5625cf649ad5eeab17c94c74f4e7f675c53c6387e65dc07223a7e5797ee673fe33bb450d336eed333e7aa9d748e34722f087d9e2c3a496becd63c6b51ae5668441258fcb2a506ee31704a797ffae9b1d5d68491cef71eab02278fc5203aa0509c425b43dac6b", "salt":"bd2feff895ac4aa15b2ee47203b75fa079da1a67f200f84d1580d54d8c6505daf8a2c68083c745b21e3db5cfd0639a9e30acc821d87a0951c8696533d7b350b6d1c86e55d6360f20fcf0b580b2ef5251f2f8ac8727dda2fee8cc5e48fb21df050ad700e5a05bbd7a62eb7319edd51040a799154104080e1103c2ccdf2e71e1c5"}',
    'Scrypt2017',
    '00000001-0001-0001-0001-000000000001',
    2
);