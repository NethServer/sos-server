CREATE TABLE `session_history` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `session_id` varchar(64) NOT NULL,
  `started` int(13) NOT NULL,
  `lk` varchar(64) NOT NULL,
  `vpn_ip` varchar(64) NOT NULL,
  `bytes_received` varchar(64) NOT NULL,
  `bytes_send` varchar(64) NOT NULL,
  `common_name` varchar(64) NOT NULL,
  `dev` varchar(64) NOT NULL,
  `ifconfig_local` varchar(64) NOT NULL,
  `ifconfig_remote` varchar(64) NOT NULL,
  `ifconfig_netmask` varchar(64) NOT NULL,
  `ifconfig_pool_local_ip` varchar(64) NOT NULL,
  `ifconfig_pool_netmask2` varchar(64) NOT NULL,
  `ifconfig_pool_remote_ip` varchar(64) NOT NULL,
  `link_mtu` varchar(64) NOT NULL,
  `local` varchar(64) NOT NULL,
  `local_port` varchar(64) NOT NULL,
  `proto` varchar(64) NOT NULL,
  `remote_{n}` varchar(64) NOT NULL,
  `remote_port_{n}` varchar(64) NOT NULL,
  `route_net_gateway` varchar(64) NOT NULL,
  `route_vpn_gateway` varchar(64) NOT NULL,
  `route_network_{n}` varchar(64) NOT NULL,
  `route_netmask_{n}` varchar(64) NOT NULL,
  `route_gateway_{n}` varchar(64) NOT NULL,
  `peer_cert` varchar(64) NOT NULL,
  `signal` varchar(64) NOT NULL,
  `time_ascii` varchar(64) NOT NULL,
  `time_duration` varchar(64) NOT NULL,
  `time_unix` varchar(64) NOT NULL,
  `tls_digest_{n}` varchar(64) NOT NULL,
  `tls_digest_sha256_{n}` varchar(64) NOT NULL,
  `tls_id_{n}` varchar(64) NOT NULL,
  `tls_serial_{n}` varchar(64) NOT NULL,
  `tun_mtu` varchar(64) NOT NULL,
  `trusted_ip` varchar(64) NOT NULL,
  `trusted_port` varchar(64) NOT NULL,
  `username` varchar(64) NOT NULL,
  `X509_0_emailAddress` varchar(64) NOT NULL,
  `X509_0_CN` varchar(64) NOT NULL,
  `X509_0_O` varchar(64) NOT NULL,
  `X509_0_ST` varchar(64) NOT NULL,
  `X509_0_C` varchar(64) NOT NULL,
  `X509_1_emailAddress` varchar(64) NOT NULL,
  `X509_1_O` varchar(64) NOT NULL,
  `X509_1_L` varchar(64) NOT NULL,
  `X509_1_ST` varchar(64) NOT NULL,
  `X509_1_C` varchar(64) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `session_id` (`session_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;