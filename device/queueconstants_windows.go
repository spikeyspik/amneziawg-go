/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2025 WireGuard LLC. All Rights Reserved.
 */

package device

const (
	QueueStagedSize            = 128
	QueueOutboundSize          = 1024
	QueueInboundSize           = 1024
	QueueHandshakeSize         = 1024
	MaxSegmentSize             = 65535 // Match with WINTUN_MAX_IP_PACKET_SIZE macro definition
	PreallocatedBuffersPerPool = 0     // Disable and allow for infinite memory growth
)
