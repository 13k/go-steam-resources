// Generated code. DO NOT EDIT
package steamlang

import (
	"encoding/binary"
	"fmt"
	"github.com/13k/go-steam-resources/protobuf/steam"
	"github.com/golang/protobuf/proto"
	"io"
	"sort"
	"strings"
)

// Code included from snippet "emsgmask"

const EMsgMask uint32 = ^uint32(0x80000000)

func MakeEMsg(e uint32) EMsg {
	return EMsg(e & EMsgMask)
}

// ---

// Code included from snippet "protomask"

const ProtoMask uint32 = 0x80000000

func MaskProto(e uint32) uint32 {
	return e | ProtoMask
}

func IsProto(e uint32) bool {
	return e&ProtoMask > 0
}

// ---

type EMsg int32

const (
	EMsg_Invalid                                                  EMsg = 0
	EMsg_Multi                                                    EMsg = 1
	EMsg_ProtobufWrapped                                          EMsg = 2
	EMsg_BaseGeneral                                              EMsg = 100
	EMsg_GenericReply                                             EMsg = 100
	EMsg_DestJobFailed                                            EMsg = 113
	EMsg_Alert                                                    EMsg = 115
	EMsg_SCIDRequest                                              EMsg = 120
	EMsg_SCIDResponse                                             EMsg = 121
	EMsg_JobHeartbeat                                             EMsg = 123
	EMsg_HubConnect                                               EMsg = 124
	EMsg_Subscribe                                                EMsg = 126
	EMsg_RouteMessage                                             EMsg = 127 // removed
	EMsg_RemoteSysID                                              EMsg = 128 // removed
	EMsg_AMCreateAccountResponse                                  EMsg = 129
	EMsg_WGRequest                                                EMsg = 130
	EMsg_WGResponse                                               EMsg = 131
	EMsg_KeepAlive                                                EMsg = 132
	EMsg_WebAPIJobRequest                                         EMsg = 133
	EMsg_WebAPIJobResponse                                        EMsg = 134
	EMsg_ClientSessionStart                                       EMsg = 135
	EMsg_ClientSessionEnd                                         EMsg = 136
	EMsg_ClientSessionUpdateAuthTicket                            EMsg = 137 // removed: renamed to ClientSessionUpdate
	EMsg_ClientSessionUpdate                                      EMsg = 137
	EMsg_StatsDeprecated                                          EMsg = 138 // removed
	EMsg_Ping                                                     EMsg = 139
	EMsg_PingResponse                                             EMsg = 140
	EMsg_Stats                                                    EMsg = 141
	EMsg_RequestFullStatsBlock                                    EMsg = 142
	EMsg_LoadDBOCacheItem                                         EMsg = 143
	EMsg_LoadDBOCacheItemResponse                                 EMsg = 144
	EMsg_InvalidateDBOCacheItems                                  EMsg = 145
	EMsg_ServiceMethod                                            EMsg = 146
	EMsg_ServiceMethodResponse                                    EMsg = 147
	EMsg_ClientPackageVersions                                    EMsg = 148
	EMsg_TimestampRequest                                         EMsg = 149
	EMsg_TimestampResponse                                        EMsg = 150
	EMsg_ServiceMethodCallFromClient                              EMsg = 151
	EMsg_ServiceMethodSendToClient                                EMsg = 152
	EMsg_BaseShell                                                EMsg = 200
	EMsg_AssignSysID                                              EMsg = 200
	EMsg_Exit                                                     EMsg = 201
	EMsg_DirRequest                                               EMsg = 202
	EMsg_DirResponse                                              EMsg = 203
	EMsg_ZipRequest                                               EMsg = 204
	EMsg_ZipResponse                                              EMsg = 205
	EMsg_UpdateRecordResponse                                     EMsg = 215
	EMsg_UpdateCreditCardRequest                                  EMsg = 221
	EMsg_UpdateUserBanResponse                                    EMsg = 225
	EMsg_PrepareToExit                                            EMsg = 226
	EMsg_ContentDescriptionUpdate                                 EMsg = 227
	EMsg_TestResetServer                                          EMsg = 228
	EMsg_UniverseChanged                                          EMsg = 229
	EMsg_ShellConfigInfoUpdate                                    EMsg = 230
	EMsg_RequestWindowsEventLogEntries                            EMsg = 233
	EMsg_ProvideWindowsEventLogEntries                            EMsg = 234
	EMsg_ShellSearchLogs                                          EMsg = 235
	EMsg_ShellSearchLogsResponse                                  EMsg = 236
	EMsg_ShellCheckWindowsUpdates                                 EMsg = 237
	EMsg_ShellCheckWindowsUpdatesResponse                         EMsg = 238
	EMsg_ShellFlushUserLicenseCache                               EMsg = 239 // removed
	EMsg_BaseGM                                                   EMsg = 300
	EMsg_Heartbeat                                                EMsg = 300
	EMsg_ShellFailed                                              EMsg = 301
	EMsg_ExitShells                                               EMsg = 307
	EMsg_ExitShell                                                EMsg = 308
	EMsg_GracefulExitShell                                        EMsg = 309
	EMsg_NotifyWatchdog                                           EMsg = 314
	EMsg_LicenseProcessingComplete                                EMsg = 316
	EMsg_SetTestFlag                                              EMsg = 317
	EMsg_QueuedEmailsComplete                                     EMsg = 318
	EMsg_GMReportPHPError                                         EMsg = 319
	EMsg_GMDRMSync                                                EMsg = 320
	EMsg_PhysicalBoxInventory                                     EMsg = 321
	EMsg_UpdateConfigFile                                         EMsg = 322
	EMsg_TestInitDB                                               EMsg = 323
	EMsg_GMWriteConfigToSQL                                       EMsg = 324
	EMsg_GMLoadActivationCodes                                    EMsg = 325
	EMsg_GMQueueForFBS                                            EMsg = 326
	EMsg_GMSchemaConversionResults                                EMsg = 327
	EMsg_GMSchemaConversionResultsResponse                        EMsg = 328 // removed
	EMsg_GMWriteShellFailureToSQL                                 EMsg = 329
	EMsg_GMWriteStatsToSOS                                        EMsg = 330
	EMsg_GMGetServiceMethodRouting                                EMsg = 331
	EMsg_GMGetServiceMethodRoutingResponse                        EMsg = 332
	EMsg_GMConvertUserWallets                                     EMsg = 333
	EMsg_BaseAIS                                                  EMsg = 400
	EMsg_AISRefreshContentDescription                             EMsg = 401 // removed
	EMsg_AISRequestContentDescription                             EMsg = 402
	EMsg_AISUpdateAppInfo                                         EMsg = 403
	EMsg_AISUpdatePackageInfo                                     EMsg = 404 // removed: renamed to AISUpdatePackageCosts
	EMsg_AISUpdatePackageCosts                                    EMsg = 404
	EMsg_AISGetPackageChangeNumber                                EMsg = 405
	EMsg_AISGetPackageChangeNumberResponse                        EMsg = 406
	EMsg_AISAppInfoTableChanged                                   EMsg = 407 // removed
	EMsg_AISUpdatePackageCostsResponse                            EMsg = 408
	EMsg_AISCreateMarketingMessage                                EMsg = 409
	EMsg_AISCreateMarketingMessageResponse                        EMsg = 410
	EMsg_AISGetMarketingMessage                                   EMsg = 411
	EMsg_AISGetMarketingMessageResponse                           EMsg = 412
	EMsg_AISUpdateMarketingMessage                                EMsg = 413
	EMsg_AISUpdateMarketingMessageResponse                        EMsg = 414
	EMsg_AISRequestMarketingMessageUpdate                         EMsg = 415
	EMsg_AISDeleteMarketingMessage                                EMsg = 416
	EMsg_AISGetMarketingTreatments                                EMsg = 419 // removed
	EMsg_AISGetMarketingTreatmentsResponse                        EMsg = 420 // removed
	EMsg_AISRequestMarketingTreatmentUpdate                       EMsg = 421 // removed
	EMsg_AISTestAddPackage                                        EMsg = 422 // removed
	EMsg_AIGetAppGCFlags                                          EMsg = 423
	EMsg_AIGetAppGCFlagsResponse                                  EMsg = 424
	EMsg_AIGetAppList                                             EMsg = 425
	EMsg_AIGetAppListResponse                                     EMsg = 426
	EMsg_AIGetAppInfo                                             EMsg = 427 // removed
	EMsg_AIGetAppInfoResponse                                     EMsg = 428 // removed
	EMsg_AISGetCouponDefinition                                   EMsg = 429
	EMsg_AISGetCouponDefinitionResponse                           EMsg = 430
	EMsg_AISUpdateSlaveContentDescription                         EMsg = 431
	EMsg_AISUpdateSlaveContentDescriptionResponse                 EMsg = 432
	EMsg_AISTestEnableGC                                          EMsg = 433
	EMsg_BaseAM                                                   EMsg = 500
	EMsg_AMUpdateUserBanRequest                                   EMsg = 504
	EMsg_AMAddLicense                                             EMsg = 505
	EMsg_AMBeginProcessingLicenses                                EMsg = 507 // removed
	EMsg_AMSendSystemIMToUser                                     EMsg = 508
	EMsg_AMExtendLicense                                          EMsg = 509
	EMsg_AMAddMinutesToLicense                                    EMsg = 510
	EMsg_AMCancelLicense                                          EMsg = 511
	EMsg_AMInitPurchase                                           EMsg = 512
	EMsg_AMPurchaseResponse                                       EMsg = 513
	EMsg_AMGetFinalPrice                                          EMsg = 514
	EMsg_AMGetFinalPriceResponse                                  EMsg = 515
	EMsg_AMGetLegacyGameKey                                       EMsg = 516
	EMsg_AMGetLegacyGameKeyResponse                               EMsg = 517
	EMsg_AMFindHungTransactions                                   EMsg = 518
	EMsg_AMSetAccountTrustedRequest                               EMsg = 519
	EMsg_AMCompletePurchase                                       EMsg = 521
	EMsg_AMCancelPurchase                                         EMsg = 522
	EMsg_AMNewChallenge                                           EMsg = 523
	EMsg_AMLoadOEMTickets                                         EMsg = 524
	EMsg_AMFixPendingPurchase                                     EMsg = 525
	EMsg_AMFixPendingPurchaseResponse                             EMsg = 526
	EMsg_AMIsUserBanned                                           EMsg = 527
	EMsg_AMRegisterKey                                            EMsg = 528
	EMsg_AMLoadActivationCodes                                    EMsg = 529
	EMsg_AMLoadActivationCodesResponse                            EMsg = 530
	EMsg_AMLookupKeyResponse                                      EMsg = 531
	EMsg_AMLookupKey                                              EMsg = 532
	EMsg_AMChatCleanup                                            EMsg = 533
	EMsg_AMClanCleanup                                            EMsg = 534
	EMsg_AMFixPendingRefund                                       EMsg = 535
	EMsg_AMReverseChargeback                                      EMsg = 536
	EMsg_AMReverseChargebackResponse                              EMsg = 537
	EMsg_AMClanCleanupList                                        EMsg = 538
	EMsg_AMGetLicenses                                            EMsg = 539
	EMsg_AMGetLicensesResponse                                    EMsg = 540
	EMsg_AllowUserToPlayQuery                                     EMsg = 550
	EMsg_AllowUserToPlayResponse                                  EMsg = 551
	EMsg_AMVerfiyUser                                             EMsg = 552
	EMsg_AMClientNotPlaying                                       EMsg = 553
	EMsg_ClientRequestFriendship                                  EMsg = 554
	EMsg_AMRelayPublishStatus                                     EMsg = 555
	EMsg_AMResetCommunityContent                                  EMsg = 556 // removed
	EMsg_AMPrimePersonaStateCache                                 EMsg = 557 // removed
	EMsg_AMAllowUserContentQuery                                  EMsg = 558 // removed
	EMsg_AMAllowUserContentResponse                               EMsg = 559 // removed
	EMsg_AMInitPurchaseResponse                                   EMsg = 560
	EMsg_AMRevokePurchaseResponse                                 EMsg = 561
	EMsg_AMLockProfile                                            EMsg = 562 // removed
	EMsg_AMRefreshGuestPasses                                     EMsg = 563
	EMsg_AMInviteUserToClan                                       EMsg = 564
	EMsg_AMAcknowledgeClanInvite                                  EMsg = 565
	EMsg_AMGrantGuestPasses                                       EMsg = 566
	EMsg_AMClanDataUpdated                                        EMsg = 567
	EMsg_AMReloadAccount                                          EMsg = 568
	EMsg_AMClientChatMsgRelay                                     EMsg = 569
	EMsg_AMChatMulti                                              EMsg = 570
	EMsg_AMClientChatInviteRelay                                  EMsg = 571
	EMsg_AMChatInvite                                             EMsg = 572
	EMsg_AMClientJoinChatRelay                                    EMsg = 573
	EMsg_AMClientChatMemberInfoRelay                              EMsg = 574
	EMsg_AMPublishChatMemberInfo                                  EMsg = 575
	EMsg_AMClientAcceptFriendInvite                               EMsg = 576
	EMsg_AMChatEnter                                              EMsg = 577
	EMsg_AMClientPublishRemovalFromSource                         EMsg = 578
	EMsg_AMChatActionResult                                       EMsg = 579
	EMsg_AMFindAccounts                                           EMsg = 580
	EMsg_AMFindAccountsResponse                                   EMsg = 581
	EMsg_AMRequestAccountData                                     EMsg = 582
	EMsg_AMRequestAccountDataResponse                             EMsg = 583
	EMsg_AMSetAccountFlags                                        EMsg = 584
	EMsg_AMCreateClan                                             EMsg = 586
	EMsg_AMCreateClanResponse                                     EMsg = 587
	EMsg_AMGetClanDetails                                         EMsg = 588
	EMsg_AMGetClanDetailsResponse                                 EMsg = 589
	EMsg_AMSetPersonaName                                         EMsg = 590
	EMsg_AMSetAvatar                                              EMsg = 591
	EMsg_AMAuthenticateUser                                       EMsg = 592
	EMsg_AMAuthenticateUserResponse                               EMsg = 593
	EMsg_AMGetAccountFriendsCount                                 EMsg = 594 // removed
	EMsg_AMGetAccountFriendsCountResponse                         EMsg = 595 // removed
	EMsg_AMP2PIntroducerMessage                                   EMsg = 596
	EMsg_ClientChatAction                                         EMsg = 597
	EMsg_AMClientChatActionRelay                                  EMsg = 598
	EMsg_BaseVS                                                   EMsg = 600
	EMsg_ReqChallenge                                             EMsg = 600
	EMsg_VACResponse                                              EMsg = 601
	EMsg_ReqChallengeTest                                         EMsg = 602
	EMsg_VSMarkCheat                                              EMsg = 604
	EMsg_VSAddCheat                                               EMsg = 605
	EMsg_VSPurgeCodeModDB                                         EMsg = 606
	EMsg_VSGetChallengeResults                                    EMsg = 607
	EMsg_VSChallengeResultText                                    EMsg = 608
	EMsg_VSReportLingerer                                         EMsg = 609
	EMsg_VSRequestManagedChallenge                                EMsg = 610
	EMsg_VSLoadDBFinished                                         EMsg = 611
	EMsg_BaseDRMS                                                 EMsg = 625
	EMsg_DRMBuildBlobRequest                                      EMsg = 628
	EMsg_DRMBuildBlobResponse                                     EMsg = 629
	EMsg_DRMResolveGuidRequest                                    EMsg = 630
	EMsg_DRMResolveGuidResponse                                   EMsg = 631
	EMsg_DRMVariabilityReport                                     EMsg = 633
	EMsg_DRMVariabilityReportResponse                             EMsg = 634
	EMsg_DRMStabilityReport                                       EMsg = 635
	EMsg_DRMStabilityReportResponse                               EMsg = 636
	EMsg_DRMDetailsReportRequest                                  EMsg = 637
	EMsg_DRMDetailsReportResponse                                 EMsg = 638
	EMsg_DRMProcessFile                                           EMsg = 639
	EMsg_DRMAdminUpdate                                           EMsg = 640
	EMsg_DRMAdminUpdateResponse                                   EMsg = 641
	EMsg_DRMSync                                                  EMsg = 642
	EMsg_DRMSyncResponse                                          EMsg = 643
	EMsg_DRMProcessFileResponse                                   EMsg = 644
	EMsg_DRMEmptyGuidCache                                        EMsg = 645
	EMsg_DRMEmptyGuidCacheResponse                                EMsg = 646
	EMsg_BaseCS                                                   EMsg = 650
	EMsg_CSUserContentRequest                                     EMsg = 652 // removed
	EMsg_BaseClient                                               EMsg = 700
	EMsg_ClientLogOn_Deprecated                                   EMsg = 701 // removed
	EMsg_ClientAnonLogOn_Deprecated                               EMsg = 702 // removed
	EMsg_ClientHeartBeat                                          EMsg = 703
	EMsg_ClientVACResponse                                        EMsg = 704
	EMsg_ClientGamesPlayed_obsolete                               EMsg = 705 // removed
	EMsg_ClientLogOff                                             EMsg = 706
	EMsg_ClientNoUDPConnectivity                                  EMsg = 707
	EMsg_ClientInformOfCreateAccount                              EMsg = 708
	EMsg_ClientAckVACBan                                          EMsg = 709 // removed
	EMsg_ClientConnectionStats                                    EMsg = 710
	EMsg_ClientInitPurchase                                       EMsg = 711 // removed
	EMsg_ClientPingResponse                                       EMsg = 712
	EMsg_ClientRemoveFriend                                       EMsg = 714
	EMsg_ClientGamesPlayedNoDataBlob                              EMsg = 715
	EMsg_ClientChangeStatus                                       EMsg = 716
	EMsg_ClientVacStatusResponse                                  EMsg = 717
	EMsg_ClientFriendMsg                                          EMsg = 718
	EMsg_ClientGameConnect_obsolete                               EMsg = 719 // removed
	EMsg_ClientGamesPlayed2_obsolete                              EMsg = 720 // removed
	EMsg_ClientGameEnded_obsolete                                 EMsg = 721 // removed
	EMsg_ClientGetFinalPrice                                      EMsg = 722 // removed
	EMsg_ClientSystemIM                                           EMsg = 726
	EMsg_ClientSystemIMAck                                        EMsg = 727
	EMsg_ClientGetLicenses                                        EMsg = 728
	EMsg_ClientCancelLicense                                      EMsg = 729 // removed
	EMsg_ClientGetLegacyGameKey                                   EMsg = 730
	EMsg_ClientContentServerLogOn_Deprecated                      EMsg = 731 // removed
	EMsg_ClientAckVACBan2                                         EMsg = 732
	EMsg_ClientAckMessageByGID                                    EMsg = 735 // removed
	EMsg_ClientGetPurchaseReceipts                                EMsg = 736
	EMsg_ClientAckPurchaseReceipt                                 EMsg = 737 // removed
	EMsg_ClientGamesPlayed3_obsolete                              EMsg = 738 // removed
	EMsg_ClientSendGuestPass                                      EMsg = 739 // removed
	EMsg_ClientAckGuestPass                                       EMsg = 740
	EMsg_ClientRedeemGuestPass                                    EMsg = 741
	EMsg_ClientGamesPlayed                                        EMsg = 742
	EMsg_ClientRegisterKey                                        EMsg = 743
	EMsg_ClientInviteUserToClan                                   EMsg = 744
	EMsg_ClientAcknowledgeClanInvite                              EMsg = 745
	EMsg_ClientPurchaseWithMachineID                              EMsg = 746
	EMsg_ClientAppUsageEvent                                      EMsg = 747
	EMsg_ClientGetGiftTargetList                                  EMsg = 748 // removed
	EMsg_ClientGetGiftTargetListResponse                          EMsg = 749 // removed
	EMsg_ClientLogOnResponse                                      EMsg = 751
	EMsg_ClientVACChallenge                                       EMsg = 753 // removed
	EMsg_ClientSetHeartbeatRate                                   EMsg = 755
	EMsg_ClientNotLoggedOnDeprecated                              EMsg = 756 // removed
	EMsg_ClientLoggedOff                                          EMsg = 757
	EMsg_GSApprove                                                EMsg = 758
	EMsg_GSDeny                                                   EMsg = 759
	EMsg_GSKick                                                   EMsg = 760
	EMsg_ClientCreateAcctResponse                                 EMsg = 761
	EMsg_ClientPurchaseResponse                                   EMsg = 763
	EMsg_ClientPing                                               EMsg = 764
	EMsg_ClientNOP                                                EMsg = 765
	EMsg_ClientPersonaState                                       EMsg = 766
	EMsg_ClientFriendsList                                        EMsg = 767
	EMsg_ClientAccountInfo                                        EMsg = 768
	EMsg_ClientVacStatusQuery                                     EMsg = 770 // removed
	EMsg_ClientNewsUpdate                                         EMsg = 771
	EMsg_ClientGameConnectDeny                                    EMsg = 773
	EMsg_GSStatusReply                                            EMsg = 774
	EMsg_ClientGetFinalPriceResponse                              EMsg = 775 // removed
	EMsg_ClientGameConnectTokens                                  EMsg = 779
	EMsg_ClientLicenseList                                        EMsg = 780
	EMsg_ClientCancelLicenseResponse                              EMsg = 781 // removed
	EMsg_ClientVACBanStatus                                       EMsg = 782
	EMsg_ClientCMList                                             EMsg = 783
	EMsg_ClientEncryptPct                                         EMsg = 784
	EMsg_ClientGetLegacyGameKeyResponse                           EMsg = 785
	EMsg_ClientFavoritesList                                      EMsg = 786 // removed
	EMsg_CSUserContentApprove                                     EMsg = 787 // removed
	EMsg_CSUserContentDeny                                        EMsg = 788 // removed
	EMsg_ClientInitPurchaseResponse                               EMsg = 789 // removed
	EMsg_ClientAddFriend                                          EMsg = 791
	EMsg_ClientAddFriendResponse                                  EMsg = 792
	EMsg_ClientInviteFriend                                       EMsg = 793 // removed
	EMsg_ClientInviteFriendResponse                               EMsg = 794 // removed
	EMsg_ClientSendGuestPassResponse                              EMsg = 795 // removed
	EMsg_ClientAckGuestPassResponse                               EMsg = 796
	EMsg_ClientRedeemGuestPassResponse                            EMsg = 797
	EMsg_ClientUpdateGuestPassesList                              EMsg = 798
	EMsg_ClientChatMsg                                            EMsg = 799
	EMsg_ClientChatInvite                                         EMsg = 800
	EMsg_ClientJoinChat                                           EMsg = 801
	EMsg_ClientChatMemberInfo                                     EMsg = 802
	EMsg_ClientLogOnWithCredentials_Deprecated                    EMsg = 803 // removed
	EMsg_ClientPasswordChangeResponse                             EMsg = 805
	EMsg_ClientChatEnter                                          EMsg = 807
	EMsg_ClientFriendRemovedFromSource                            EMsg = 808
	EMsg_ClientCreateChat                                         EMsg = 809
	EMsg_ClientCreateChatResponse                                 EMsg = 810
	EMsg_ClientUpdateChatMetadata                                 EMsg = 811 // removed
	EMsg_ClientP2PIntroducerMessage                               EMsg = 813
	EMsg_ClientChatActionResult                                   EMsg = 814
	EMsg_ClientRequestFriendData                                  EMsg = 815
	EMsg_ClientGetUserStats                                       EMsg = 818
	EMsg_ClientGetUserStatsResponse                               EMsg = 819
	EMsg_ClientStoreUserStats                                     EMsg = 820
	EMsg_ClientStoreUserStatsResponse                             EMsg = 821
	EMsg_ClientClanState                                          EMsg = 822
	EMsg_ClientServiceModule                                      EMsg = 830
	EMsg_ClientServiceCall                                        EMsg = 831
	EMsg_ClientServiceCallResponse                                EMsg = 832
	EMsg_ClientPackageInfoRequest                                 EMsg = 833 // removed
	EMsg_ClientPackageInfoResponse                                EMsg = 834 // removed
	EMsg_ClientNatTraversalStatEvent                              EMsg = 839
	EMsg_ClientAppInfoRequest                                     EMsg = 840 // removed
	EMsg_ClientAppInfoResponse                                    EMsg = 841 // removed
	EMsg_ClientSteamUsageEvent                                    EMsg = 842
	EMsg_ClientCheckPassword                                      EMsg = 845
	EMsg_ClientResetPassword                                      EMsg = 846
	EMsg_ClientCheckPasswordResponse                              EMsg = 848
	EMsg_ClientResetPasswordResponse                              EMsg = 849
	EMsg_ClientSessionToken                                       EMsg = 850
	EMsg_ClientDRMProblemReport                                   EMsg = 851
	EMsg_ClientSetIgnoreFriend                                    EMsg = 855
	EMsg_ClientSetIgnoreFriendResponse                            EMsg = 856
	EMsg_ClientGetAppOwnershipTicket                              EMsg = 857
	EMsg_ClientGetAppOwnershipTicketResponse                      EMsg = 858
	EMsg_ClientGetLobbyListResponse                               EMsg = 860
	EMsg_ClientGetLobbyMetadata                                   EMsg = 861 // removed
	EMsg_ClientGetLobbyMetadataResponse                           EMsg = 862 // removed
	EMsg_ClientVTTCert                                            EMsg = 863 // removed
	EMsg_ClientAppInfoUpdate                                      EMsg = 866 // removed
	EMsg_ClientAppInfoChanges                                     EMsg = 867 // removed
	EMsg_ClientServerList                                         EMsg = 880
	EMsg_ClientEmailChangeResponse                                EMsg = 891
	EMsg_ClientSecretQAChangeResponse                             EMsg = 892
	EMsg_ClientDRMBlobRequest                                     EMsg = 896
	EMsg_ClientDRMBlobResponse                                    EMsg = 897
	EMsg_ClientLookupKey                                          EMsg = 898 // removed
	EMsg_ClientLookupKeyResponse                                  EMsg = 899 // removed
	EMsg_BaseGameServer                                           EMsg = 900
	EMsg_GSDisconnectNotice                                       EMsg = 901
	EMsg_GSStatus                                                 EMsg = 903
	EMsg_GSUserPlaying                                            EMsg = 905
	EMsg_GSStatus2                                                EMsg = 906
	EMsg_GSStatusUpdate_Unused                                    EMsg = 907
	EMsg_GSServerType                                             EMsg = 908
	EMsg_GSPlayerList                                             EMsg = 909
	EMsg_GSGetUserAchievementStatus                               EMsg = 910
	EMsg_GSGetUserAchievementStatusResponse                       EMsg = 911
	EMsg_GSGetPlayStats                                           EMsg = 918
	EMsg_GSGetPlayStatsResponse                                   EMsg = 919
	EMsg_GSGetUserGroupStatus                                     EMsg = 920
	EMsg_AMGetUserGroupStatus                                     EMsg = 921
	EMsg_AMGetUserGroupStatusResponse                             EMsg = 922
	EMsg_GSGetUserGroupStatusResponse                             EMsg = 923
	EMsg_GSGetReputation                                          EMsg = 936
	EMsg_GSGetReputationResponse                                  EMsg = 937
	EMsg_GSAssociateWithClan                                      EMsg = 938
	EMsg_GSAssociateWithClanResponse                              EMsg = 939
	EMsg_GSComputeNewPlayerCompatibility                          EMsg = 940
	EMsg_GSComputeNewPlayerCompatibilityResponse                  EMsg = 941
	EMsg_BaseAdmin                                                EMsg = 1000
	EMsg_AdminCmd                                                 EMsg = 1000
	EMsg_AdminCmdResponse                                         EMsg = 1004
	EMsg_AdminLogListenRequest                                    EMsg = 1005
	EMsg_AdminLogEvent                                            EMsg = 1006
	EMsg_LogSearchRequest                                         EMsg = 1007 // removed
	EMsg_LogSearchResponse                                        EMsg = 1008 // removed
	EMsg_LogSearchCancel                                          EMsg = 1009 // removed
	EMsg_UniverseData                                             EMsg = 1010
	EMsg_RequestStatHistory                                       EMsg = 1014 // removed
	EMsg_StatHistory                                              EMsg = 1015 // removed
	EMsg_AdminPwLogon                                             EMsg = 1017 // removed
	EMsg_AdminPwLogonResponse                                     EMsg = 1018 // removed
	EMsg_AdminSpew                                                EMsg = 1019
	EMsg_AdminConsoleTitle                                        EMsg = 1020
	EMsg_AdminGCSpew                                              EMsg = 1023
	EMsg_AdminGCCommand                                           EMsg = 1024
	EMsg_AdminGCGetCommandList                                    EMsg = 1025
	EMsg_AdminGCGetCommandListResponse                            EMsg = 1026
	EMsg_FBSConnectionData                                        EMsg = 1027
	EMsg_AdminMsgSpew                                             EMsg = 1028
	EMsg_BaseFBS                                                  EMsg = 1100
	EMsg_FBSReqVersion                                            EMsg = 1100
	EMsg_FBSVersionInfo                                           EMsg = 1101
	EMsg_FBSForceRefresh                                          EMsg = 1102
	EMsg_FBSForceBounce                                           EMsg = 1103
	EMsg_FBSDeployPackage                                         EMsg = 1104
	EMsg_FBSDeployResponse                                        EMsg = 1105
	EMsg_FBSUpdateBootstrapper                                    EMsg = 1106
	EMsg_FBSSetState                                              EMsg = 1107
	EMsg_FBSApplyOSUpdates                                        EMsg = 1108
	EMsg_FBSRunCMDScript                                          EMsg = 1109
	EMsg_FBSRebootBox                                             EMsg = 1110
	EMsg_FBSSetBigBrotherMode                                     EMsg = 1111
	EMsg_FBSMinidumpServer                                        EMsg = 1112
	EMsg_FBSSetShellCount_obsolete                                EMsg = 1113 // removed
	EMsg_FBSDeployHotFixPackage                                   EMsg = 1114
	EMsg_FBSDeployHotFixResponse                                  EMsg = 1115
	EMsg_FBSDownloadHotFix                                        EMsg = 1116
	EMsg_FBSDownloadHotFixResponse                                EMsg = 1117
	EMsg_FBSUpdateTargetConfigFile                                EMsg = 1118
	EMsg_FBSApplyAccountCred                                      EMsg = 1119
	EMsg_FBSApplyAccountCredResponse                              EMsg = 1120
	EMsg_FBSSetShellCount                                         EMsg = 1121
	EMsg_FBSTerminateShell                                        EMsg = 1122
	EMsg_FBSQueryGMForRequest                                     EMsg = 1123
	EMsg_FBSQueryGMResponse                                       EMsg = 1124
	EMsg_FBSTerminateZombies                                      EMsg = 1125
	EMsg_FBSInfoFromBootstrapper                                  EMsg = 1126
	EMsg_FBSRebootBoxResponse                                     EMsg = 1127
	EMsg_FBSBootstrapperPackageRequest                            EMsg = 1128
	EMsg_FBSBootstrapperPackageResponse                           EMsg = 1129
	EMsg_FBSBootstrapperGetPackageChunk                           EMsg = 1130
	EMsg_FBSBootstrapperGetPackageChunkResponse                   EMsg = 1131
	EMsg_FBSBootstrapperPackageTransferProgress                   EMsg = 1132
	EMsg_FBSRestartBootstrapper                                   EMsg = 1133
	EMsg_BaseFileXfer                                             EMsg = 1200
	EMsg_FileXferRequest                                          EMsg = 1200
	EMsg_FileXferResponse                                         EMsg = 1201
	EMsg_FileXferData                                             EMsg = 1202
	EMsg_FileXferEnd                                              EMsg = 1203
	EMsg_FileXferDataAck                                          EMsg = 1204
	EMsg_BaseChannelAuth                                          EMsg = 1300
	EMsg_ChannelAuthChallenge                                     EMsg = 1300
	EMsg_ChannelAuthResponse                                      EMsg = 1301
	EMsg_ChannelAuthResult                                        EMsg = 1302
	EMsg_ChannelEncryptRequest                                    EMsg = 1303
	EMsg_ChannelEncryptResponse                                   EMsg = 1304
	EMsg_ChannelEncryptResult                                     EMsg = 1305
	EMsg_BaseBS                                                   EMsg = 1400
	EMsg_BSPurchaseStart                                          EMsg = 1401
	EMsg_BSPurchaseResponse                                       EMsg = 1402 // removed
	EMsg_BSSettleNOVA                                             EMsg = 1404 // removed
	EMsg_BSSettleComplete                                         EMsg = 1406
	EMsg_BSBannedRequest                                          EMsg = 1407 // removed
	EMsg_BSInitPayPalTxn                                          EMsg = 1408
	EMsg_BSInitPayPalTxnResponse                                  EMsg = 1409
	EMsg_BSGetPayPalUserInfo                                      EMsg = 1410
	EMsg_BSGetPayPalUserInfoResponse                              EMsg = 1411
	EMsg_BSRefundTxn                                              EMsg = 1413 // removed
	EMsg_BSRefundTxnResponse                                      EMsg = 1414 // removed
	EMsg_BSGetEvents                                              EMsg = 1415 // removed
	EMsg_BSChaseRFRRequest                                        EMsg = 1416 // removed
	EMsg_BSPaymentInstrBan                                        EMsg = 1417
	EMsg_BSPaymentInstrBanResponse                                EMsg = 1418
	EMsg_BSProcessGCReports                                       EMsg = 1419 // removed
	EMsg_BSProcessPPReports                                       EMsg = 1420 // removed
	EMsg_BSInitGCBankXferTxn                                      EMsg = 1421
	EMsg_BSInitGCBankXferTxnResponse                              EMsg = 1422
	EMsg_BSQueryGCBankXferTxn                                     EMsg = 1423 // removed
	EMsg_BSQueryGCBankXferTxnResponse                             EMsg = 1424 // removed
	EMsg_BSCommitGCTxn                                            EMsg = 1425
	EMsg_BSQueryTransactionStatus                                 EMsg = 1426
	EMsg_BSQueryTransactionStatusResponse                         EMsg = 1427
	EMsg_BSQueryCBOrderStatus                                     EMsg = 1428 // removed
	EMsg_BSQueryCBOrderStatusResponse                             EMsg = 1429 // removed
	EMsg_BSRunRedFlagReport                                       EMsg = 1430 // removed
	EMsg_BSQueryPaymentInstUsage                                  EMsg = 1431
	EMsg_BSQueryPaymentInstResponse                               EMsg = 1432
	EMsg_BSQueryTxnExtendedInfo                                   EMsg = 1433
	EMsg_BSQueryTxnExtendedInfoResponse                           EMsg = 1434
	EMsg_BSUpdateConversionRates                                  EMsg = 1435
	EMsg_BSProcessUSBankReports                                   EMsg = 1436 // removed
	EMsg_BSPurchaseRunFraudChecks                                 EMsg = 1437
	EMsg_BSPurchaseRunFraudChecksResponse                         EMsg = 1438
	EMsg_BSStartShippingJobs                                      EMsg = 1439 // removed
	EMsg_BSQueryBankInformation                                   EMsg = 1440
	EMsg_BSQueryBankInformationResponse                           EMsg = 1441
	EMsg_BSValidateXsollaSignature                                EMsg = 1445
	EMsg_BSValidateXsollaSignatureResponse                        EMsg = 1446
	EMsg_BSQiwiWalletInvoice                                      EMsg = 1448
	EMsg_BSQiwiWalletInvoiceResponse                              EMsg = 1449
	EMsg_BSUpdateInventoryFromProPack                             EMsg = 1450
	EMsg_BSUpdateInventoryFromProPackResponse                     EMsg = 1451
	EMsg_BSSendShippingRequest                                    EMsg = 1452
	EMsg_BSSendShippingRequestResponse                            EMsg = 1453
	EMsg_BSGetProPackOrderStatus                                  EMsg = 1454
	EMsg_BSGetProPackOrderStatusResponse                          EMsg = 1455
	EMsg_BSCheckJobRunning                                        EMsg = 1456
	EMsg_BSCheckJobRunningResponse                                EMsg = 1457
	EMsg_BSResetPackagePurchaseRateLimit                          EMsg = 1458
	EMsg_BSResetPackagePurchaseRateLimitResponse                  EMsg = 1459
	EMsg_BSUpdatePaymentData                                      EMsg = 1460
	EMsg_BSUpdatePaymentDataResponse                              EMsg = 1461
	EMsg_BSGetBillingAddress                                      EMsg = 1462
	EMsg_BSGetBillingAddressResponse                              EMsg = 1463
	EMsg_BSGetCreditCardInfo                                      EMsg = 1464
	EMsg_BSGetCreditCardInfoResponse                              EMsg = 1465
	EMsg_BSRemoveExpiredPaymentData                               EMsg = 1468
	EMsg_BSRemoveExpiredPaymentDataResponse                       EMsg = 1469
	EMsg_BSConvertToCurrentKeys                                   EMsg = 1470
	EMsg_BSConvertToCurrentKeysResponse                           EMsg = 1471
	EMsg_BSInitPurchase                                           EMsg = 1472
	EMsg_BSInitPurchaseResponse                                   EMsg = 1473
	EMsg_BSCompletePurchase                                       EMsg = 1474
	EMsg_BSCompletePurchaseResponse                               EMsg = 1475
	EMsg_BSPruneCardUsageStats                                    EMsg = 1476
	EMsg_BSPruneCardUsageStatsResponse                            EMsg = 1477
	EMsg_BSStoreBankInformation                                   EMsg = 1478
	EMsg_BSStoreBankInformationResponse                           EMsg = 1479
	EMsg_BSVerifyPOSAKey                                          EMsg = 1480
	EMsg_BSVerifyPOSAKeyResponse                                  EMsg = 1481
	EMsg_BSReverseRedeemPOSAKey                                   EMsg = 1482
	EMsg_BSReverseRedeemPOSAKeyResponse                           EMsg = 1483
	EMsg_BSQueryFindCreditCard                                    EMsg = 1484
	EMsg_BSQueryFindCreditCardResponse                            EMsg = 1485
	EMsg_BSStatusInquiryPOSAKey                                   EMsg = 1486
	EMsg_BSStatusInquiryPOSAKeyResponse                           EMsg = 1487
	EMsg_BSValidateMoPaySignature                                 EMsg = 1488
	EMsg_BSValidateMoPaySignatureResponse                         EMsg = 1489
	EMsg_BSMoPayConfirmProductDelivery                            EMsg = 1490
	EMsg_BSMoPayConfirmProductDeliveryResponse                    EMsg = 1491
	EMsg_BSGenerateMoPayMD5                                       EMsg = 1492
	EMsg_BSGenerateMoPayMD5Response                               EMsg = 1493
	EMsg_BSBoaCompraConfirmProductDelivery                        EMsg = 1494
	EMsg_BSBoaCompraConfirmProductDeliveryResponse                EMsg = 1495
	EMsg_BSGenerateBoaCompraMD5                                   EMsg = 1496
	EMsg_BSGenerateBoaCompraMD5Response                           EMsg = 1497
	EMsg_BSCommitWPTxn                                            EMsg = 1498
	EMsg_BaseATS                                                  EMsg = 1500
	EMsg_ATSStartStressTest                                       EMsg = 1501
	EMsg_ATSStopStressTest                                        EMsg = 1502
	EMsg_ATSRunFailServerTest                                     EMsg = 1503
	EMsg_ATSUFSPerfTestTask                                       EMsg = 1504
	EMsg_ATSUFSPerfTestResponse                                   EMsg = 1505
	EMsg_ATSCycleTCM                                              EMsg = 1506
	EMsg_ATSInitDRMSStressTest                                    EMsg = 1507
	EMsg_ATSCallTest                                              EMsg = 1508
	EMsg_ATSCallTestReply                                         EMsg = 1509
	EMsg_ATSStartExternalStress                                   EMsg = 1510
	EMsg_ATSExternalStressJobStart                                EMsg = 1511
	EMsg_ATSExternalStressJobQueued                               EMsg = 1512
	EMsg_ATSExternalStressJobRunning                              EMsg = 1513
	EMsg_ATSExternalStressJobStopped                              EMsg = 1514
	EMsg_ATSExternalStressJobStopAll                              EMsg = 1515
	EMsg_ATSExternalStressActionResult                            EMsg = 1516
	EMsg_ATSStarted                                               EMsg = 1517
	EMsg_ATSCSPerfTestTask                                        EMsg = 1518
	EMsg_ATSCSPerfTestResponse                                    EMsg = 1519
	EMsg_BaseDP                                                   EMsg = 1600
	EMsg_DPSetPublishingState                                     EMsg = 1601
	EMsg_DPGamePlayedStats                                        EMsg = 1602 // removed
	EMsg_DPUniquePlayersStat                                      EMsg = 1603
	EMsg_DPStreamingUniquePlayersStat                             EMsg = 1604
	EMsg_DPVacInfractionStats                                     EMsg = 1605
	EMsg_DPVacBanStats                                            EMsg = 1606
	EMsg_DPBlockingStats                                          EMsg = 1607
	EMsg_DPNatTraversalStats                                      EMsg = 1608
	EMsg_DPSteamUsageEvent                                        EMsg = 1609 // removed
	EMsg_DPVacCertBanStats                                        EMsg = 1610
	EMsg_DPVacCafeBanStats                                        EMsg = 1611
	EMsg_DPCloudStats                                             EMsg = 1612
	EMsg_DPAchievementStats                                       EMsg = 1613
	EMsg_DPAccountCreationStats                                   EMsg = 1614
	EMsg_DPGetPlayerCount                                         EMsg = 1615
	EMsg_DPGetPlayerCountResponse                                 EMsg = 1616
	EMsg_DPGameServersPlayersStats                                EMsg = 1617
	EMsg_DPDownloadRateStatistics                                 EMsg = 1618 // removed
	EMsg_DPFacebookStatistics                                     EMsg = 1619
	EMsg_ClientDPCheckSpecialSurvey                               EMsg = 1620
	EMsg_ClientDPCheckSpecialSurveyResponse                       EMsg = 1621
	EMsg_ClientDPSendSpecialSurveyResponse                        EMsg = 1622
	EMsg_ClientDPSendSpecialSurveyResponseReply                   EMsg = 1623
	EMsg_DPStoreSaleStatistics                                    EMsg = 1624
	EMsg_ClientDPUpdateAppJobReport                               EMsg = 1625
	EMsg_ClientDPSteam2AppStarted                                 EMsg = 1627 // removed
	EMsg_DPUpdateContentEvent                                     EMsg = 1626
	EMsg_DPPartnerMicroTxns                                       EMsg = 1628
	EMsg_DPPartnerMicroTxnsResponse                               EMsg = 1629
	EMsg_ClientDPContentStatsReport                               EMsg = 1630
	EMsg_DPVRUniquePlayersStat                                    EMsg = 1631
	EMsg_BaseCM                                                   EMsg = 1700
	EMsg_CMSetAllowState                                          EMsg = 1701
	EMsg_CMSpewAllowState                                         EMsg = 1702
	EMsg_CMAppInfoResponseDeprecated                              EMsg = 1703 // removed
	EMsg_BaseDSS                                                  EMsg = 1800 // removed
	EMsg_DSSNewFile                                               EMsg = 1801 // removed
	EMsg_DSSCurrentFileList                                       EMsg = 1802 // removed
	EMsg_DSSSynchList                                             EMsg = 1803 // removed
	EMsg_DSSSynchListResponse                                     EMsg = 1804 // removed
	EMsg_DSSSynchSubscribe                                        EMsg = 1805 // removed
	EMsg_DSSSynchUnsubscribe                                      EMsg = 1806 // removed
	EMsg_BaseEPM                                                  EMsg = 1900 // removed
	EMsg_EPMStartProcess                                          EMsg = 1901 // removed
	EMsg_EPMStopProcess                                           EMsg = 1902 // removed
	EMsg_EPMRestartProcess                                        EMsg = 1903 // removed
	EMsg_BaseGC                                                   EMsg = 2200
	EMsg_GCSendClient                                             EMsg = 2200 // removed
	EMsg_AMRelayToGC                                              EMsg = 2201 // removed
	EMsg_GCUpdatePlayedState                                      EMsg = 2202 // removed
	EMsg_GCCmdRevive                                              EMsg = 2203
	EMsg_GCCmdBounce                                              EMsg = 2204 // removed
	EMsg_GCCmdForceBounce                                         EMsg = 2205 // removed
	EMsg_GCCmdDown                                                EMsg = 2206
	EMsg_GCCmdDeploy                                              EMsg = 2207
	EMsg_GCCmdDeployResponse                                      EMsg = 2208
	EMsg_GCCmdSwitch                                              EMsg = 2209
	EMsg_AMRefreshSessions                                        EMsg = 2210
	EMsg_GCUpdateGSState                                          EMsg = 2211 // removed
	EMsg_GCAchievementAwarded                                     EMsg = 2212
	EMsg_GCSystemMessage                                          EMsg = 2213
	EMsg_GCValidateSession                                        EMsg = 2214 // removed
	EMsg_GCValidateSessionResponse                                EMsg = 2215 // removed
	EMsg_GCCmdStatus                                              EMsg = 2216
	EMsg_GCRegisterWebInterfaces                                  EMsg = 2217 // removed
	EMsg_GCRegisterWebInterfaces_Deprecated                       EMsg = 2217 // removed
	EMsg_GCGetAccountDetails                                      EMsg = 2218 // removed
	EMsg_GCGetAccountDetails_DEPRECATED                           EMsg = 2218 // removed
	EMsg_GCInterAppMessage                                        EMsg = 2219
	EMsg_GCGetEmailTemplate                                       EMsg = 2220
	EMsg_GCGetEmailTemplateResponse                               EMsg = 2221
	EMsg_ISRelayToGCH                                             EMsg = 2222 // removed: renamed to GCHRelay
	EMsg_GCHRelay                                                 EMsg = 2222
	EMsg_GCHRelayClientToIS                                       EMsg = 2223 // removed: renamed to GCHRelayToClient
	EMsg_GCHRelayToClient                                         EMsg = 2223
	EMsg_GCHUpdateSession                                         EMsg = 2224
	EMsg_GCHRequestUpdateSession                                  EMsg = 2225
	EMsg_GCHRequestStatus                                         EMsg = 2226
	EMsg_GCHRequestStatusResponse                                 EMsg = 2227
	EMsg_GCHAccountVacStatusChange                                EMsg = 2228
	EMsg_GCHSpawnGC                                               EMsg = 2229
	EMsg_GCHSpawnGCResponse                                       EMsg = 2230
	EMsg_GCHKillGC                                                EMsg = 2231
	EMsg_GCHKillGCResponse                                        EMsg = 2232
	EMsg_GCHAccountTradeBanStatusChange                           EMsg = 2233
	EMsg_GCHAccountLockStatusChange                               EMsg = 2234
	EMsg_GCHVacVerificationChange                                 EMsg = 2235
	EMsg_GCHAccountPhoneNumberChange                              EMsg = 2236
	EMsg_GCHAccountTwoFactorChange                                EMsg = 2237
	EMsg_BaseP2P                                                  EMsg = 2500
	EMsg_P2PIntroducerMessage                                     EMsg = 2502
	EMsg_BaseSM                                                   EMsg = 2900
	EMsg_SMExpensiveReport                                        EMsg = 2902
	EMsg_SMHourlyReport                                           EMsg = 2903
	EMsg_SMFishingReport                                          EMsg = 2904
	EMsg_SMPartitionRenames                                       EMsg = 2905
	EMsg_SMMonitorSpace                                           EMsg = 2906
	EMsg_SMGetSchemaConversionResults                             EMsg = 2907 // removed
	EMsg_SMGetSchemaConversionResultsResponse                     EMsg = 2908 // removed
	EMsg_BaseTest                                                 EMsg = 3000
	EMsg_FailServer                                               EMsg = 3000
	EMsg_JobHeartbeatTest                                         EMsg = 3001
	EMsg_JobHeartbeatTestResponse                                 EMsg = 3002
	EMsg_BaseFTSRange                                             EMsg = 3100
	EMsg_FTSGetBrowseCounts                                       EMsg = 3101 // removed
	EMsg_FTSGetBrowseCountsResponse                               EMsg = 3102 // removed
	EMsg_FTSBrowseClans                                           EMsg = 3103 // removed
	EMsg_FTSBrowseClansResponse                                   EMsg = 3104 // removed
	EMsg_FTSSearchClansByLocation                                 EMsg = 3105 // removed
	EMsg_FTSSearchClansByLocationResponse                         EMsg = 3106 // removed
	EMsg_FTSSearchPlayersByLocation                               EMsg = 3107 // removed
	EMsg_FTSSearchPlayersByLocationResponse                       EMsg = 3108 // removed
	EMsg_FTSClanDeleted                                           EMsg = 3109 // removed
	EMsg_FTSSearch                                                EMsg = 3110 // removed
	EMsg_FTSSearchResponse                                        EMsg = 3111 // removed
	EMsg_FTSSearchStatus                                          EMsg = 3112 // removed
	EMsg_FTSSearchStatusResponse                                  EMsg = 3113 // removed
	EMsg_FTSGetGSPlayStats                                        EMsg = 3114 // removed
	EMsg_FTSGetGSPlayStatsResponse                                EMsg = 3115 // removed
	EMsg_FTSGetGSPlayStatsForServer                               EMsg = 3116 // removed
	EMsg_FTSGetGSPlayStatsForServerResponse                       EMsg = 3117 // removed
	EMsg_FTSReportIPUpdates                                       EMsg = 3118 // removed
	EMsg_BaseCCSRange                                             EMsg = 3150
	EMsg_CCSGetComments                                           EMsg = 3151 // removed
	EMsg_CCSGetCommentsResponse                                   EMsg = 3152 // removed
	EMsg_CCSAddComment                                            EMsg = 3153 // removed
	EMsg_CCSAddCommentResponse                                    EMsg = 3154 // removed
	EMsg_CCSDeleteComment                                         EMsg = 3155 // removed
	EMsg_CCSDeleteCommentResponse                                 EMsg = 3156 // removed
	EMsg_CCSPreloadComments                                       EMsg = 3157 // removed
	EMsg_CCSNotifyCommentCount                                    EMsg = 3158 // removed
	EMsg_CCSGetCommentsForNews                                    EMsg = 3159 // removed
	EMsg_CCSGetCommentsForNewsResponse                            EMsg = 3160 // removed
	EMsg_CCSDeleteAllCommentsByAuthor                             EMsg = 3161
	EMsg_CCSDeleteAllCommentsByAuthorResponse                     EMsg = 3162
	EMsg_BaseLBSRange                                             EMsg = 3200
	EMsg_LBSSetScore                                              EMsg = 3201
	EMsg_LBSSetScoreResponse                                      EMsg = 3202
	EMsg_LBSFindOrCreateLB                                        EMsg = 3203
	EMsg_LBSFindOrCreateLBResponse                                EMsg = 3204
	EMsg_LBSGetLBEntries                                          EMsg = 3205
	EMsg_LBSGetLBEntriesResponse                                  EMsg = 3206
	EMsg_LBSGetLBList                                             EMsg = 3207
	EMsg_LBSGetLBListResponse                                     EMsg = 3208
	EMsg_LBSSetLBDetails                                          EMsg = 3209
	EMsg_LBSDeleteLB                                              EMsg = 3210
	EMsg_LBSDeleteLBEntry                                         EMsg = 3211
	EMsg_LBSResetLB                                               EMsg = 3212
	EMsg_LBSResetLBResponse                                       EMsg = 3213
	EMsg_LBSDeleteLBResponse                                      EMsg = 3214
	EMsg_BaseOGS                                                  EMsg = 3400
	EMsg_OGSBeginSession                                          EMsg = 3401
	EMsg_OGSBeginSessionResponse                                  EMsg = 3402
	EMsg_OGSEndSession                                            EMsg = 3403
	EMsg_OGSEndSessionResponse                                    EMsg = 3404
	EMsg_OGSWriteAppSessionRow                                    EMsg = 3406
	EMsg_BaseBRP                                                  EMsg = 3600
	EMsg_BRPStartShippingJobs                                     EMsg = 3601
	EMsg_BRPProcessUSBankReports                                  EMsg = 3602
	EMsg_BRPProcessGCReports                                      EMsg = 3603
	EMsg_BRPProcessPPReports                                      EMsg = 3604
	EMsg_BRPSettleNOVA                                            EMsg = 3605 // removed
	EMsg_BRPSettleCB                                              EMsg = 3606 // removed
	EMsg_BRPCommitGC                                              EMsg = 3607
	EMsg_BRPCommitGCResponse                                      EMsg = 3608
	EMsg_BRPFindHungTransactions                                  EMsg = 3609
	EMsg_BRPCheckFinanceCloseOutDate                              EMsg = 3610
	EMsg_BRPProcessLicenses                                       EMsg = 3611
	EMsg_BRPProcessLicensesResponse                               EMsg = 3612
	EMsg_BRPRemoveExpiredPaymentData                              EMsg = 3613
	EMsg_BRPRemoveExpiredPaymentDataResponse                      EMsg = 3614
	EMsg_BRPConvertToCurrentKeys                                  EMsg = 3615
	EMsg_BRPConvertToCurrentKeysResponse                          EMsg = 3616
	EMsg_BRPPruneCardUsageStats                                   EMsg = 3617
	EMsg_BRPPruneCardUsageStatsResponse                           EMsg = 3618
	EMsg_BRPCheckActivationCodes                                  EMsg = 3619
	EMsg_BRPCheckActivationCodesResponse                          EMsg = 3620
	EMsg_BRPCommitWP                                              EMsg = 3621
	EMsg_BRPCommitWPResponse                                      EMsg = 3622
	EMsg_BRPProcessWPReports                                      EMsg = 3623
	EMsg_BRPProcessPaymentRules                                   EMsg = 3624
	EMsg_BRPProcessPartnerPayments                                EMsg = 3625
	EMsg_BRPCheckSettlementReports                                EMsg = 3626
	EMsg_BRPPostTaxToAvalara                                      EMsg = 3628
	EMsg_BRPPostTransactionTax                                    EMsg = 3629
	EMsg_BRPPostTransactionTaxResponse                            EMsg = 3630
	EMsg_BRPProcessIMReports                                      EMsg = 3631
	EMsg_BaseAMRange2                                             EMsg = 4000
	EMsg_AMCreateChat                                             EMsg = 4001
	EMsg_AMCreateChatResponse                                     EMsg = 4002
	EMsg_AMUpdateChatMetadata                                     EMsg = 4003 // removed
	EMsg_AMPublishChatMetadata                                    EMsg = 4004 // removed
	EMsg_AMSetProfileURL                                          EMsg = 4005
	EMsg_AMGetAccountEmailAddress                                 EMsg = 4006
	EMsg_AMGetAccountEmailAddressResponse                         EMsg = 4007
	EMsg_AMRequestFriendData                                      EMsg = 4008 // removed: renamed to AMRequestClanData
	EMsg_AMRequestClanData                                        EMsg = 4008
	EMsg_AMRouteToClients                                         EMsg = 4009
	EMsg_AMLeaveClan                                              EMsg = 4010
	EMsg_AMClanPermissions                                        EMsg = 4011
	EMsg_AMClanPermissionsResponse                                EMsg = 4012
	EMsg_AMCreateClanEvent                                        EMsg = 4013
	EMsg_AMCreateClanEventResponse                                EMsg = 4014
	EMsg_AMUpdateClanEvent                                        EMsg = 4015
	EMsg_AMUpdateClanEventResponse                                EMsg = 4016
	EMsg_AMGetClanEvents                                          EMsg = 4017
	EMsg_AMGetClanEventsResponse                                  EMsg = 4018
	EMsg_AMDeleteClanEvent                                        EMsg = 4019
	EMsg_AMDeleteClanEventResponse                                EMsg = 4020
	EMsg_AMSetClanPermissionSettings                              EMsg = 4021
	EMsg_AMSetClanPermissionSettingsResponse                      EMsg = 4022
	EMsg_AMGetClanPermissionSettings                              EMsg = 4023
	EMsg_AMGetClanPermissionSettingsResponse                      EMsg = 4024
	EMsg_AMPublishChatRoomInfo                                    EMsg = 4025
	EMsg_ClientChatRoomInfo                                       EMsg = 4026
	EMsg_AMCreateClanAnnouncement                                 EMsg = 4027 // removed
	EMsg_AMCreateClanAnnouncementResponse                         EMsg = 4028 // removed
	EMsg_AMUpdateClanAnnouncement                                 EMsg = 4029 // removed
	EMsg_AMUpdateClanAnnouncementResponse                         EMsg = 4030 // removed
	EMsg_AMGetClanAnnouncementsCount                              EMsg = 4031 // removed
	EMsg_AMGetClanAnnouncementsCountResponse                      EMsg = 4032 // removed
	EMsg_AMGetClanAnnouncements                                   EMsg = 4033 // removed
	EMsg_AMGetClanAnnouncementsResponse                           EMsg = 4034 // removed
	EMsg_AMDeleteClanAnnouncement                                 EMsg = 4035 // removed
	EMsg_AMDeleteClanAnnouncementResponse                         EMsg = 4036 // removed
	EMsg_AMGetSingleClanAnnouncement                              EMsg = 4037 // removed
	EMsg_AMGetSingleClanAnnouncementResponse                      EMsg = 4038 // removed
	EMsg_AMGetClanHistory                                         EMsg = 4039
	EMsg_AMGetClanHistoryResponse                                 EMsg = 4040
	EMsg_AMGetClanPermissionBits                                  EMsg = 4041
	EMsg_AMGetClanPermissionBitsResponse                          EMsg = 4042
	EMsg_AMSetClanPermissionBits                                  EMsg = 4043
	EMsg_AMSetClanPermissionBitsResponse                          EMsg = 4044
	EMsg_AMSessionInfoRequest                                     EMsg = 4045
	EMsg_AMSessionInfoResponse                                    EMsg = 4046
	EMsg_AMValidateWGToken                                        EMsg = 4047
	EMsg_AMGetSingleClanEvent                                     EMsg = 4048
	EMsg_AMGetSingleClanEventResponse                             EMsg = 4049
	EMsg_AMGetClanRank                                            EMsg = 4050
	EMsg_AMGetClanRankResponse                                    EMsg = 4051
	EMsg_AMSetClanRank                                            EMsg = 4052
	EMsg_AMSetClanRankResponse                                    EMsg = 4053
	EMsg_AMGetClanPOTW                                            EMsg = 4054
	EMsg_AMGetClanPOTWResponse                                    EMsg = 4055
	EMsg_AMSetClanPOTW                                            EMsg = 4056
	EMsg_AMSetClanPOTWResponse                                    EMsg = 4057
	EMsg_AMRequestChatMetadata                                    EMsg = 4058 // removed
	EMsg_AMDumpUser                                               EMsg = 4059
	EMsg_AMKickUserFromClan                                       EMsg = 4060
	EMsg_AMAddFounderToClan                                       EMsg = 4061
	EMsg_AMValidateWGTokenResponse                                EMsg = 4062
	EMsg_AMSetCommunityState                                      EMsg = 4063
	EMsg_AMSetAccountDetails                                      EMsg = 4064
	EMsg_AMGetChatBanList                                         EMsg = 4065
	EMsg_AMGetChatBanListResponse                                 EMsg = 4066
	EMsg_AMUnBanFromChat                                          EMsg = 4067
	EMsg_AMSetClanDetails                                         EMsg = 4068
	EMsg_AMGetAccountLinks                                        EMsg = 4069
	EMsg_AMGetAccountLinksResponse                                EMsg = 4070
	EMsg_AMSetAccountLinks                                        EMsg = 4071
	EMsg_AMSetAccountLinksResponse                                EMsg = 4072
	EMsg_AMGetUserGameStats                                       EMsg = 4073
	EMsg_AMGetUserGameStatsResponse                               EMsg = 4074
	EMsg_AMCheckClanMembership                                    EMsg = 4075
	EMsg_AMGetClanMembers                                         EMsg = 4076
	EMsg_AMGetClanMembersResponse                                 EMsg = 4077
	EMsg_AMJoinPublicClan                                         EMsg = 4078
	EMsg_AMNotifyChatOfClanChange                                 EMsg = 4079
	EMsg_AMResubmitPurchase                                       EMsg = 4080
	EMsg_AMAddFriend                                              EMsg = 4081
	EMsg_AMAddFriendResponse                                      EMsg = 4082
	EMsg_AMRemoveFriend                                           EMsg = 4083
	EMsg_AMDumpClan                                               EMsg = 4084
	EMsg_AMChangeClanOwner                                        EMsg = 4085
	EMsg_AMCancelEasyCollect                                      EMsg = 4086
	EMsg_AMCancelEasyCollectResponse                              EMsg = 4087
	EMsg_AMGetClanMembershipList                                  EMsg = 4088 // removed
	EMsg_AMGetClanMembershipListResponse                          EMsg = 4089 // removed
	EMsg_AMClansInCommon                                          EMsg = 4090
	EMsg_AMClansInCommonResponse                                  EMsg = 4091
	EMsg_AMIsValidAccountID                                       EMsg = 4092
	EMsg_AMConvertClan                                            EMsg = 4093
	EMsg_AMGetGiftTargetListRelay                                 EMsg = 4094 // removed
	EMsg_AMWipeFriendsList                                        EMsg = 4095
	EMsg_AMSetIgnored                                             EMsg = 4096
	EMsg_AMClansInCommonCountResponse                             EMsg = 4097
	EMsg_AMFriendsList                                            EMsg = 4098
	EMsg_AMFriendsListResponse                                    EMsg = 4099
	EMsg_AMFriendsInCommon                                        EMsg = 4100
	EMsg_AMFriendsInCommonResponse                                EMsg = 4101
	EMsg_AMFriendsInCommonCountResponse                           EMsg = 4102
	EMsg_AMClansInCommonCount                                     EMsg = 4103
	EMsg_AMChallengeVerdict                                       EMsg = 4104
	EMsg_AMChallengeNotification                                  EMsg = 4105
	EMsg_AMFindGSByIP                                             EMsg = 4106
	EMsg_AMFoundGSByIP                                            EMsg = 4107
	EMsg_AMGiftRevoked                                            EMsg = 4108
	EMsg_AMCreateAccountRecord                                    EMsg = 4109
	EMsg_AMUserClanList                                           EMsg = 4110
	EMsg_AMUserClanListResponse                                   EMsg = 4111
	EMsg_AMGetAccountDetails2                                     EMsg = 4112
	EMsg_AMGetAccountDetailsResponse2                             EMsg = 4113
	EMsg_AMSetCommunityProfileSettings                            EMsg = 4114
	EMsg_AMSetCommunityProfileSettingsResponse                    EMsg = 4115
	EMsg_AMGetCommunityPrivacyState                               EMsg = 4116
	EMsg_AMGetCommunityPrivacyStateResponse                       EMsg = 4117
	EMsg_AMCheckClanInviteRateLimiting                            EMsg = 4118
	EMsg_AMGetUserAchievementStatus                               EMsg = 4119
	EMsg_AMGetIgnored                                             EMsg = 4120
	EMsg_AMGetIgnoredResponse                                     EMsg = 4121
	EMsg_AMSetIgnoredResponse                                     EMsg = 4122
	EMsg_AMSetFriendRelationshipNone                              EMsg = 4123
	EMsg_AMGetFriendRelationship                                  EMsg = 4124
	EMsg_AMGetFriendRelationshipResponse                          EMsg = 4125
	EMsg_AMServiceModulesCache                                    EMsg = 4126
	EMsg_AMServiceModulesCall                                     EMsg = 4127
	EMsg_AMServiceModulesCallResponse                             EMsg = 4128
	EMsg_AMGetCaptchaDataForIP                                    EMsg = 4129
	EMsg_AMGetCaptchaDataForIPResponse                            EMsg = 4130
	EMsg_AMValidateCaptchaDataForIP                               EMsg = 4131
	EMsg_AMValidateCaptchaDataForIPResponse                       EMsg = 4132
	EMsg_AMTrackFailedAuthByIP                                    EMsg = 4133
	EMsg_AMGetCaptchaDataByGID                                    EMsg = 4134
	EMsg_AMGetCaptchaDataByGIDResponse                            EMsg = 4135
	EMsg_AMGetLobbyList                                           EMsg = 4136 // removed
	EMsg_AMGetLobbyListResponse                                   EMsg = 4137 // removed
	EMsg_AMGetLobbyMetadata                                       EMsg = 4138 // removed
	EMsg_AMGetLobbyMetadataResponse                               EMsg = 4139 // removed
	EMsg_CommunityAddFriendNews                                   EMsg = 4140
	EMsg_AMAddClanNews                                            EMsg = 4141 // removed
	EMsg_AMWriteNews                                              EMsg = 4142 // removed
	EMsg_AMFindClanUser                                           EMsg = 4143
	EMsg_AMFindClanUserResponse                                   EMsg = 4144
	EMsg_AMBanFromChat                                            EMsg = 4145
	EMsg_AMGetUserHistoryResponse                                 EMsg = 4146 // removed
	EMsg_AMGetUserNewsSubscriptions                               EMsg = 4147
	EMsg_AMGetUserNewsSubscriptionsResponse                       EMsg = 4148
	EMsg_AMSetUserNewsSubscriptions                               EMsg = 4149
	EMsg_AMGetUserNews                                            EMsg = 4150 // removed
	EMsg_AMGetUserNewsResponse                                    EMsg = 4151 // removed
	EMsg_AMSendQueuedEmails                                       EMsg = 4152
	EMsg_AMSetLicenseFlags                                        EMsg = 4153
	EMsg_AMGetUserHistory                                         EMsg = 4154 // removed
	EMsg_CommunityDeleteUserNews                                  EMsg = 4155
	EMsg_AMAllowUserFilesRequest                                  EMsg = 4156
	EMsg_AMAllowUserFilesResponse                                 EMsg = 4157
	EMsg_AMGetAccountStatus                                       EMsg = 4158
	EMsg_AMGetAccountStatusResponse                               EMsg = 4159
	EMsg_AMEditBanReason                                          EMsg = 4160
	EMsg_AMCheckClanMembershipResponse                            EMsg = 4161
	EMsg_AMProbeClanMembershipList                                EMsg = 4162
	EMsg_AMProbeClanMembershipListResponse                        EMsg = 4163
	EMsg_AMGetFriendsLobbies                                      EMsg = 4165
	EMsg_AMGetFriendsLobbiesResponse                              EMsg = 4166
	EMsg_AMGetUserFriendNewsResponse                              EMsg = 4172
	EMsg_CommunityGetUserFriendNews                               EMsg = 4173
	EMsg_AMGetUserClansNewsResponse                               EMsg = 4174
	EMsg_AMGetUserClansNews                                       EMsg = 4175
	EMsg_AMStoreInitPurchase                                      EMsg = 4176 // removed
	EMsg_AMStoreInitPurchaseResponse                              EMsg = 4177 // removed
	EMsg_AMStoreGetFinalPrice                                     EMsg = 4178 // removed
	EMsg_AMStoreGetFinalPriceResponse                             EMsg = 4179 // removed
	EMsg_AMStoreCompletePurchase                                  EMsg = 4180 // removed
	EMsg_AMStoreCancelPurchase                                    EMsg = 4181 // removed
	EMsg_AMStorePurchaseResponse                                  EMsg = 4182 // removed
	EMsg_AMCreateAccountRecordInSteam3                            EMsg = 4183 // removed
	EMsg_AMGetPreviousCBAccount                                   EMsg = 4184
	EMsg_AMGetPreviousCBAccountResponse                           EMsg = 4185
	EMsg_AMUpdateBillingAddress                                   EMsg = 4186 // removed
	EMsg_AMUpdateBillingAddressResponse                           EMsg = 4187 // removed
	EMsg_AMGetBillingAddress                                      EMsg = 4188 // removed
	EMsg_AMGetBillingAddressResponse                              EMsg = 4189 // removed
	EMsg_AMGetUserLicenseHistory                                  EMsg = 4190
	EMsg_AMGetUserLicenseHistoryResponse                          EMsg = 4191
	EMsg_AMSupportChangePassword                                  EMsg = 4194
	EMsg_AMSupportChangeEmail                                     EMsg = 4195
	EMsg_AMSupportChangeSecretQA                                  EMsg = 4196 // removed
	EMsg_AMResetUserVerificationGSByIP                            EMsg = 4197
	EMsg_AMUpdateGSPlayStats                                      EMsg = 4198
	EMsg_AMSupportEnableOrDisable                                 EMsg = 4199
	EMsg_AMGetComments                                            EMsg = 4200 // removed
	EMsg_AMGetCommentsResponse                                    EMsg = 4201 // removed
	EMsg_AMAddComment                                             EMsg = 4202 // removed
	EMsg_AMAddCommentResponse                                     EMsg = 4203 // removed
	EMsg_AMDeleteComment                                          EMsg = 4204 // removed
	EMsg_AMDeleteCommentResponse                                  EMsg = 4205 // removed
	EMsg_AMGetPurchaseStatus                                      EMsg = 4206
	EMsg_AMSupportIsAccountEnabled                                EMsg = 4209
	EMsg_AMSupportIsAccountEnabledResponse                        EMsg = 4210
	EMsg_AMGetUserStats                                           EMsg = 4211
	EMsg_AMSupportKickSession                                     EMsg = 4212
	EMsg_AMGSSearch                                               EMsg = 4213
	EMsg_MarketingMessageUpdate                                   EMsg = 4216
	EMsg_AMRouteFriendMsg                                         EMsg = 4219
	EMsg_AMTicketAuthRequestOrResponse                            EMsg = 4220
	EMsg_AMVerifyDepotManagementRights                            EMsg = 4222
	EMsg_AMVerifyDepotManagementRightsResponse                    EMsg = 4223
	EMsg_AMAddFreeLicense                                         EMsg = 4224
	EMsg_AMGetUserFriendsMinutesPlayed                            EMsg = 4225 // removed
	EMsg_AMGetUserFriendsMinutesPlayedResponse                    EMsg = 4226 // removed
	EMsg_AMGetUserMinutesPlayed                                   EMsg = 4227 // removed
	EMsg_AMGetUserMinutesPlayedResponse                           EMsg = 4228 // removed
	EMsg_AMValidateEmailLink                                      EMsg = 4231
	EMsg_AMValidateEmailLinkResponse                              EMsg = 4232
	EMsg_AMAddUsersToMarketingTreatment                           EMsg = 4234 // removed
	EMsg_AMStoreUserStats                                         EMsg = 4236
	EMsg_AMGetUserGameplayInfo                                    EMsg = 4237 // removed
	EMsg_AMGetUserGameplayInfoResponse                            EMsg = 4238 // removed
	EMsg_AMGetCardList                                            EMsg = 4239 // removed
	EMsg_AMGetCardListResponse                                    EMsg = 4240 // removed
	EMsg_AMDeleteStoredCard                                       EMsg = 4241
	EMsg_AMRevokeLegacyGameKeys                                   EMsg = 4242
	EMsg_AMGetWalletDetails                                       EMsg = 4244
	EMsg_AMGetWalletDetailsResponse                               EMsg = 4245
	EMsg_AMDeleteStoredPaymentInfo                                EMsg = 4246
	EMsg_AMGetStoredPaymentSummary                                EMsg = 4247
	EMsg_AMGetStoredPaymentSummaryResponse                        EMsg = 4248
	EMsg_AMGetWalletConversionRate                                EMsg = 4249
	EMsg_AMGetWalletConversionRateResponse                        EMsg = 4250
	EMsg_AMConvertWallet                                          EMsg = 4251
	EMsg_AMConvertWalletResponse                                  EMsg = 4252
	EMsg_AMRelayGetFriendsWhoPlayGame                             EMsg = 4253 // removed
	EMsg_AMRelayGetFriendsWhoPlayGameResponse                     EMsg = 4254 // removed
	EMsg_AMSetPreApproval                                         EMsg = 4255
	EMsg_AMSetPreApprovalResponse                                 EMsg = 4256
	EMsg_AMMarketingTreatmentUpdate                               EMsg = 4257 // removed
	EMsg_AMCreateRefund                                           EMsg = 4258
	EMsg_AMCreateRefundResponse                                   EMsg = 4259
	EMsg_AMCreateChargeback                                       EMsg = 4260
	EMsg_AMCreateChargebackResponse                               EMsg = 4261
	EMsg_AMCreateDispute                                          EMsg = 4262
	EMsg_AMCreateDisputeResponse                                  EMsg = 4263
	EMsg_AMClearDispute                                           EMsg = 4264
	EMsg_AMClearDisputeResponse                                   EMsg = 4265
	EMsg_AMPlayerNicknameList                                     EMsg = 4266
	EMsg_AMPlayerNicknameListResponse                             EMsg = 4267
	EMsg_AMSetDRMTestConfig                                       EMsg = 4268
	EMsg_AMGetUserCurrentGameInfo                                 EMsg = 4269
	EMsg_AMGetUserCurrentGameInfoResponse                         EMsg = 4270
	EMsg_AMGetGSPlayerList                                        EMsg = 4271
	EMsg_AMGetGSPlayerListResponse                                EMsg = 4272
	EMsg_AMUpdatePersonaStateCache                                EMsg = 4275 // removed
	EMsg_AMGetGameMembers                                         EMsg = 4276
	EMsg_AMGetGameMembersResponse                                 EMsg = 4277
	EMsg_AMGetSteamIDForMicroTxn                                  EMsg = 4278
	EMsg_AMGetSteamIDForMicroTxnResponse                          EMsg = 4279
	EMsg_AMAddPublisherUser                                       EMsg = 4280
	EMsg_AMRemovePublisherUser                                    EMsg = 4281
	EMsg_AMGetUserLicenseList                                     EMsg = 4282
	EMsg_AMGetUserLicenseListResponse                             EMsg = 4283
	EMsg_AMReloadGameGroupPolicy                                  EMsg = 4284
	EMsg_AMAddFreeLicenseResponse                                 EMsg = 4285
	EMsg_AMVACStatusUpdate                                        EMsg = 4286
	EMsg_AMGetAccountDetails                                      EMsg = 4287
	EMsg_AMGetAccountDetailsResponse                              EMsg = 4288
	EMsg_AMGetPlayerLinkDetails                                   EMsg = 4289
	EMsg_AMGetPlayerLinkDetailsResponse                           EMsg = 4290
	EMsg_AMSubscribeToPersonaFeed                                 EMsg = 4291 // removed
	EMsg_AMGetUserVacBanList                                      EMsg = 4292 // removed
	EMsg_AMGetUserVacBanListResponse                              EMsg = 4293 // removed
	EMsg_AMGetAccountFlagsForWGSpoofing                           EMsg = 4294
	EMsg_AMGetAccountFlagsForWGSpoofingResponse                   EMsg = 4295
	EMsg_AMGetFriendsWishlistInfo                                 EMsg = 4296 // removed
	EMsg_AMGetFriendsWishlistInfoResponse                         EMsg = 4297 // removed
	EMsg_AMGetClanOfficers                                        EMsg = 4298
	EMsg_AMGetClanOfficersResponse                                EMsg = 4299
	EMsg_AMNameChange                                             EMsg = 4300
	EMsg_AMGetNameHistory                                         EMsg = 4301
	EMsg_AMGetNameHistoryResponse                                 EMsg = 4302
	EMsg_AMUpdateProviderStatus                                   EMsg = 4305
	EMsg_AMClearPersonaMetadataBlob                               EMsg = 4306 // removed
	EMsg_AMSupportRemoveAccountSecurity                           EMsg = 4307
	EMsg_AMIsAccountInCaptchaGracePeriod                          EMsg = 4308
	EMsg_AMIsAccountInCaptchaGracePeriodResponse                  EMsg = 4309
	EMsg_AMAccountPS3Unlink                                       EMsg = 4310
	EMsg_AMAccountPS3UnlinkResponse                               EMsg = 4311
	EMsg_AMStoreUserStatsResponse                                 EMsg = 4312
	EMsg_AMGetAccountPSNInfo                                      EMsg = 4313
	EMsg_AMGetAccountPSNInfoResponse                              EMsg = 4314
	EMsg_AMAuthenticatedPlayerList                                EMsg = 4315
	EMsg_AMGetUserGifts                                           EMsg = 4316
	EMsg_AMGetUserGiftsResponse                                   EMsg = 4317
	EMsg_AMTransferLockedGifts                                    EMsg = 4320
	EMsg_AMTransferLockedGiftsResponse                            EMsg = 4321
	EMsg_AMPlayerHostedOnGameServer                               EMsg = 4322
	EMsg_AMGetAccountBanInfo                                      EMsg = 4323
	EMsg_AMGetAccountBanInfoResponse                              EMsg = 4324
	EMsg_AMRecordBanEnforcement                                   EMsg = 4325
	EMsg_AMRollbackGiftTransfer                                   EMsg = 4326
	EMsg_AMRollbackGiftTransferResponse                           EMsg = 4327
	EMsg_AMHandlePendingTransaction                               EMsg = 4328
	EMsg_AMRequestClanDetails                                     EMsg = 4329
	EMsg_AMDeleteStoredPaypalAgreement                            EMsg = 4330
	EMsg_AMGameServerUpdate                                       EMsg = 4331
	EMsg_AMGameServerRemove                                       EMsg = 4332
	EMsg_AMGetPaypalAgreements                                    EMsg = 4333
	EMsg_AMGetPaypalAgreementsResponse                            EMsg = 4334
	EMsg_AMGameServerPlayerCompatibilityCheck                     EMsg = 4335
	EMsg_AMGameServerPlayerCompatibilityCheckResponse             EMsg = 4336
	EMsg_AMRenewLicense                                           EMsg = 4337
	EMsg_AMGetAccountCommunityBanInfo                             EMsg = 4338
	EMsg_AMGetAccountCommunityBanInfoResponse                     EMsg = 4339
	EMsg_AMGameServerAccountChangePassword                        EMsg = 4340
	EMsg_AMGameServerAccountDeleteAccount                         EMsg = 4341
	EMsg_AMRenewAgreement                                         EMsg = 4342
	EMsg_AMSendEmail                                              EMsg = 4343 // removed
	EMsg_AMXsollaPayment                                          EMsg = 4344
	EMsg_AMXsollaPaymentResponse                                  EMsg = 4345
	EMsg_AMAcctAllowedToPurchase                                  EMsg = 4346
	EMsg_AMAcctAllowedToPurchaseResponse                          EMsg = 4347
	EMsg_AMSwapKioskDeposit                                       EMsg = 4348
	EMsg_AMSwapKioskDepositResponse                               EMsg = 4349
	EMsg_AMSetUserGiftUnowned                                     EMsg = 4350
	EMsg_AMSetUserGiftUnownedResponse                             EMsg = 4351
	EMsg_AMClaimUnownedUserGift                                   EMsg = 4352
	EMsg_AMClaimUnownedUserGiftResponse                           EMsg = 4353
	EMsg_AMSetClanName                                            EMsg = 4354
	EMsg_AMSetClanNameResponse                                    EMsg = 4355
	EMsg_AMGrantCoupon                                            EMsg = 4356
	EMsg_AMGrantCouponResponse                                    EMsg = 4357
	EMsg_AMIsPackageRestrictedInUserCountry                       EMsg = 4358
	EMsg_AMIsPackageRestrictedInUserCountryResponse               EMsg = 4359
	EMsg_AMHandlePendingTransactionResponse                       EMsg = 4360
	EMsg_AMGrantGuestPasses2                                      EMsg = 4361
	EMsg_AMGrantGuestPasses2Response                              EMsg = 4362
	EMsg_AMSessionQuery                                           EMsg = 4363
	EMsg_AMSessionQueryResponse                                   EMsg = 4364
	EMsg_AMGetPlayerBanDetails                                    EMsg = 4365
	EMsg_AMGetPlayerBanDetailsResponse                            EMsg = 4366
	EMsg_AMFinalizePurchase                                       EMsg = 4367
	EMsg_AMFinalizePurchaseResponse                               EMsg = 4368
	EMsg_AMPersonaChangeResponse                                  EMsg = 4372
	EMsg_AMGetClanDetailsForForumCreation                         EMsg = 4373
	EMsg_AMGetClanDetailsForForumCreationResponse                 EMsg = 4374
	EMsg_AMGetPendingNotificationCount                            EMsg = 4375
	EMsg_AMGetPendingNotificationCountResponse                    EMsg = 4376
	EMsg_AMPasswordHashUpgrade                                    EMsg = 4377
	EMsg_AMMoPayPayment                                           EMsg = 4378
	EMsg_AMMoPayPaymentResponse                                   EMsg = 4379
	EMsg_AMBoaCompraPayment                                       EMsg = 4380
	EMsg_AMBoaCompraPaymentResponse                               EMsg = 4381
	EMsg_AMExpireCaptchaByGID                                     EMsg = 4382
	EMsg_AMCompleteExternalPurchase                               EMsg = 4383
	EMsg_AMCompleteExternalPurchaseResponse                       EMsg = 4384
	EMsg_AMResolveNegativeWalletCredits                           EMsg = 4385
	EMsg_AMResolveNegativeWalletCreditsResponse                   EMsg = 4386
	EMsg_AMPayelpPayment                                          EMsg = 4387
	EMsg_AMPayelpPaymentResponse                                  EMsg = 4388
	EMsg_AMPlayerGetClanBasicDetails                              EMsg = 4389
	EMsg_AMPlayerGetClanBasicDetailsResponse                      EMsg = 4390
	EMsg_AMMOLPayment                                             EMsg = 4391
	EMsg_AMMOLPaymentResponse                                     EMsg = 4392
	EMsg_GetUserIPCountry                                         EMsg = 4393
	EMsg_GetUserIPCountryResponse                                 EMsg = 4394
	EMsg_NotificationOfSuspiciousActivity                         EMsg = 4395
	EMsg_AMDegicaPayment                                          EMsg = 4396
	EMsg_AMDegicaPaymentResponse                                  EMsg = 4397
	EMsg_AMEClubPayment                                           EMsg = 4398
	EMsg_AMEClubPaymentResponse                                   EMsg = 4399
	EMsg_AMPayPalPaymentsHubPayment                               EMsg = 4400
	EMsg_AMPayPalPaymentsHubPaymentResponse                       EMsg = 4401
	EMsg_AMTwoFactorRecoverAuthenticatorRequest                   EMsg = 4402
	EMsg_AMTwoFactorRecoverAuthenticatorResponse                  EMsg = 4403
	EMsg_AMSmart2PayPayment                                       EMsg = 4404
	EMsg_AMSmart2PayPaymentResponse                               EMsg = 4405
	EMsg_AMValidatePasswordResetCodeAndSendSmsRequest             EMsg = 4406
	EMsg_AMValidatePasswordResetCodeAndSendSmsResponse            EMsg = 4407
	EMsg_AMGetAccountResetDetailsRequest                          EMsg = 4408
	EMsg_AMGetAccountResetDetailsResponse                         EMsg = 4409
	EMsg_AMBitPayPayment                                          EMsg = 4410
	EMsg_AMBitPayPaymentResponse                                  EMsg = 4411
	EMsg_AMSendAccountInfoUpdate                                  EMsg = 4412
	EMsg_BasePSRange                                              EMsg = 5000
	EMsg_PSCreateShoppingCart                                     EMsg = 5001
	EMsg_PSCreateShoppingCartResponse                             EMsg = 5002
	EMsg_PSIsValidShoppingCart                                    EMsg = 5003
	EMsg_PSIsValidShoppingCartResponse                            EMsg = 5004
	EMsg_PSAddPackageToShoppingCart                               EMsg = 5005
	EMsg_PSAddPackageToShoppingCartResponse                       EMsg = 5006
	EMsg_PSRemoveLineItemFromShoppingCart                         EMsg = 5007
	EMsg_PSRemoveLineItemFromShoppingCartResponse                 EMsg = 5008
	EMsg_PSGetShoppingCartContents                                EMsg = 5009
	EMsg_PSGetShoppingCartContentsResponse                        EMsg = 5010
	EMsg_PSAddWalletCreditToShoppingCart                          EMsg = 5011
	EMsg_PSAddWalletCreditToShoppingCartResponse                  EMsg = 5012
	EMsg_BaseUFSRange                                             EMsg = 5200
	EMsg_ClientUFSUploadFileRequest                               EMsg = 5202
	EMsg_ClientUFSUploadFileResponse                              EMsg = 5203
	EMsg_ClientUFSUploadFileChunk                                 EMsg = 5204
	EMsg_ClientUFSUploadFileFinished                              EMsg = 5205
	EMsg_ClientUFSGetFileListForApp                               EMsg = 5206
	EMsg_ClientUFSGetFileListForAppResponse                       EMsg = 5207
	EMsg_ClientUFSDownloadRequest                                 EMsg = 5210
	EMsg_ClientUFSDownloadResponse                                EMsg = 5211
	EMsg_ClientUFSDownloadChunk                                   EMsg = 5212
	EMsg_ClientUFSLoginRequest                                    EMsg = 5213
	EMsg_ClientUFSLoginResponse                                   EMsg = 5214
	EMsg_UFSReloadPartitionInfo                                   EMsg = 5215
	EMsg_ClientUFSTransferHeartbeat                               EMsg = 5216
	EMsg_UFSSynchronizeFile                                       EMsg = 5217
	EMsg_UFSSynchronizeFileResponse                               EMsg = 5218
	EMsg_ClientUFSDeleteFileRequest                               EMsg = 5219
	EMsg_ClientUFSDeleteFileResponse                              EMsg = 5220
	EMsg_UFSDownloadRequest                                       EMsg = 5221 // removed
	EMsg_UFSDownloadResponse                                      EMsg = 5222 // removed
	EMsg_UFSDownloadChunk                                         EMsg = 5223 // removed
	EMsg_ClientUFSGetUGCDetails                                   EMsg = 5226
	EMsg_ClientUFSGetUGCDetailsResponse                           EMsg = 5227
	EMsg_UFSUpdateFileFlags                                       EMsg = 5228
	EMsg_UFSUpdateFileFlagsResponse                               EMsg = 5229
	EMsg_ClientUFSGetSingleFileInfo                               EMsg = 5230
	EMsg_ClientUFSGetSingleFileInfoResponse                       EMsg = 5231
	EMsg_ClientUFSShareFile                                       EMsg = 5232
	EMsg_ClientUFSShareFileResponse                               EMsg = 5233
	EMsg_UFSReloadAccount                                         EMsg = 5234
	EMsg_UFSReloadAccountResponse                                 EMsg = 5235
	EMsg_UFSUpdateRecordBatched                                   EMsg = 5236
	EMsg_UFSUpdateRecordBatchedResponse                           EMsg = 5237
	EMsg_UFSMigrateFile                                           EMsg = 5238
	EMsg_UFSMigrateFileResponse                                   EMsg = 5239
	EMsg_UFSGetUGCURLs                                            EMsg = 5240
	EMsg_UFSGetUGCURLsResponse                                    EMsg = 5241
	EMsg_UFSHttpUploadFileFinishRequest                           EMsg = 5242
	EMsg_UFSHttpUploadFileFinishResponse                          EMsg = 5243
	EMsg_UFSDownloadStartRequest                                  EMsg = 5244
	EMsg_UFSDownloadStartResponse                                 EMsg = 5245
	EMsg_UFSDownloadChunkRequest                                  EMsg = 5246
	EMsg_UFSDownloadChunkResponse                                 EMsg = 5247
	EMsg_UFSDownloadFinishRequest                                 EMsg = 5248
	EMsg_UFSDownloadFinishResponse                                EMsg = 5249
	EMsg_UFSFlushURLCache                                         EMsg = 5250
	EMsg_UFSUploadCommit                                          EMsg = 5251
	EMsg_UFSUploadCommitResponse                                  EMsg = 5252
	EMsg_UFSMigrateFileAppID                                      EMsg = 5253
	EMsg_UFSMigrateFileAppIDResponse                              EMsg = 5254
	EMsg_BaseClient2                                              EMsg = 5400
	EMsg_ClientRequestForgottenPasswordEmail                      EMsg = 5401
	EMsg_ClientRequestForgottenPasswordEmailResponse              EMsg = 5402
	EMsg_ClientCreateAccountResponse                              EMsg = 5403
	EMsg_ClientResetForgottenPassword                             EMsg = 5404
	EMsg_ClientResetForgottenPasswordResponse                     EMsg = 5405
	EMsg_ClientCreateAccount2                                     EMsg = 5406
	EMsg_ClientInformOfResetForgottenPassword                     EMsg = 5407
	EMsg_ClientInformOfResetForgottenPasswordResponse             EMsg = 5408
	EMsg_ClientAnonUserLogOn_Deprecated                           EMsg = 5409 // removed
	EMsg_ClientGamesPlayedWithDataBlob                            EMsg = 5410
	EMsg_ClientUpdateUserGameInfo                                 EMsg = 5411
	EMsg_ClientFileToDownload                                     EMsg = 5412
	EMsg_ClientFileToDownloadResponse                             EMsg = 5413
	EMsg_ClientLBSSetScore                                        EMsg = 5414
	EMsg_ClientLBSSetScoreResponse                                EMsg = 5415
	EMsg_ClientLBSFindOrCreateLB                                  EMsg = 5416
	EMsg_ClientLBSFindOrCreateLBResponse                          EMsg = 5417
	EMsg_ClientLBSGetLBEntries                                    EMsg = 5418
	EMsg_ClientLBSGetLBEntriesResponse                            EMsg = 5419
	EMsg_ClientMarketingMessageUpdate                             EMsg = 5420 // removed
	EMsg_ClientChatDeclined                                       EMsg = 5426
	EMsg_ClientFriendMsgIncoming                                  EMsg = 5427
	EMsg_ClientAuthList_Deprecated                                EMsg = 5428 // removed
	EMsg_ClientTicketAuthComplete                                 EMsg = 5429
	EMsg_ClientIsLimitedAccount                                   EMsg = 5430
	EMsg_ClientRequestAuthList                                    EMsg = 5431
	EMsg_ClientAuthList                                           EMsg = 5432
	EMsg_ClientStat                                               EMsg = 5433
	EMsg_ClientP2PConnectionInfo                                  EMsg = 5434
	EMsg_ClientP2PConnectionFailInfo                              EMsg = 5435
	EMsg_ClientGetNumberOfCurrentPlayers                          EMsg = 5436 // removed
	EMsg_ClientGetNumberOfCurrentPlayersResponse                  EMsg = 5437 // removed
	EMsg_ClientGetDepotDecryptionKey                              EMsg = 5438
	EMsg_ClientGetDepotDecryptionKeyResponse                      EMsg = 5439
	EMsg_GSPerformHardwareSurvey                                  EMsg = 5440
	EMsg_ClientGetAppBetaPasswords                                EMsg = 5441 // removed
	EMsg_ClientGetAppBetaPasswordsResponse                        EMsg = 5442 // removed
	EMsg_ClientEnableTestLicense                                  EMsg = 5443
	EMsg_ClientEnableTestLicenseResponse                          EMsg = 5444
	EMsg_ClientDisableTestLicense                                 EMsg = 5445
	EMsg_ClientDisableTestLicenseResponse                         EMsg = 5446
	EMsg_ClientRequestValidationMail                              EMsg = 5448
	EMsg_ClientRequestValidationMailResponse                      EMsg = 5449
	EMsg_ClientCheckAppBetaPassword                               EMsg = 5450
	EMsg_ClientCheckAppBetaPasswordResponse                       EMsg = 5451
	EMsg_ClientToGC                                               EMsg = 5452
	EMsg_ClientFromGC                                             EMsg = 5453
	EMsg_ClientRequestChangeMail                                  EMsg = 5454
	EMsg_ClientRequestChangeMailResponse                          EMsg = 5455
	EMsg_ClientEmailAddrInfo                                      EMsg = 5456
	EMsg_ClientPasswordChange3                                    EMsg = 5457
	EMsg_ClientEmailChange3                                       EMsg = 5458
	EMsg_ClientPersonalQAChange3                                  EMsg = 5459
	EMsg_ClientResetForgottenPassword3                            EMsg = 5460
	EMsg_ClientRequestForgottenPasswordEmail3                     EMsg = 5461
	EMsg_ClientCreateAccount3                                     EMsg = 5462 // removed
	EMsg_ClientNewLoginKey                                        EMsg = 5463
	EMsg_ClientNewLoginKeyAccepted                                EMsg = 5464
	EMsg_ClientLogOnWithHash_Deprecated                           EMsg = 5465 // removed
	EMsg_ClientStoreUserStats2                                    EMsg = 5466
	EMsg_ClientStatsUpdated                                       EMsg = 5467
	EMsg_ClientActivateOEMLicense                                 EMsg = 5468
	EMsg_ClientRegisterOEMMachine                                 EMsg = 5469
	EMsg_ClientRegisterOEMMachineResponse                         EMsg = 5470
	EMsg_ClientRequestedClientStats                               EMsg = 5480
	EMsg_ClientStat2Int32                                         EMsg = 5481
	EMsg_ClientStat2                                              EMsg = 5482
	EMsg_ClientVerifyPassword                                     EMsg = 5483
	EMsg_ClientVerifyPasswordResponse                             EMsg = 5484
	EMsg_ClientDRMDownloadRequest                                 EMsg = 5485
	EMsg_ClientDRMDownloadResponse                                EMsg = 5486
	EMsg_ClientDRMFinalResult                                     EMsg = 5487
	EMsg_ClientGetFriendsWhoPlayGame                              EMsg = 5488
	EMsg_ClientGetFriendsWhoPlayGameResponse                      EMsg = 5489
	EMsg_ClientOGSBeginSession                                    EMsg = 5490
	EMsg_ClientOGSBeginSessionResponse                            EMsg = 5491
	EMsg_ClientOGSEndSession                                      EMsg = 5492
	EMsg_ClientOGSEndSessionResponse                              EMsg = 5493
	EMsg_ClientOGSWriteRow                                        EMsg = 5494
	EMsg_ClientDRMTest                                            EMsg = 5495
	EMsg_ClientDRMTestResult                                      EMsg = 5496
	EMsg_ClientServerUnavailable                                  EMsg = 5500
	EMsg_ClientServersAvailable                                   EMsg = 5501
	EMsg_ClientRegisterAuthTicketWithCM                           EMsg = 5502
	EMsg_ClientGCMsgFailed                                        EMsg = 5503
	EMsg_ClientMicroTxnAuthRequest                                EMsg = 5504
	EMsg_ClientMicroTxnAuthorize                                  EMsg = 5505
	EMsg_ClientMicroTxnAuthorizeResponse                          EMsg = 5506
	EMsg_ClientAppMinutesPlayedData                               EMsg = 5507
	EMsg_ClientGetMicroTxnInfo                                    EMsg = 5508
	EMsg_ClientGetMicroTxnInfoResponse                            EMsg = 5509
	EMsg_ClientMarketingMessageUpdate2                            EMsg = 5510
	EMsg_ClientDeregisterWithServer                               EMsg = 5511
	EMsg_ClientSubscribeToPersonaFeed                             EMsg = 5512
	EMsg_ClientLogon                                              EMsg = 5514
	EMsg_ClientGetClientDetails                                   EMsg = 5515
	EMsg_ClientGetClientDetailsResponse                           EMsg = 5516
	EMsg_ClientReportOverlayDetourFailure                         EMsg = 5517
	EMsg_ClientGetClientAppList                                   EMsg = 5518
	EMsg_ClientGetClientAppListResponse                           EMsg = 5519
	EMsg_ClientInstallClientApp                                   EMsg = 5520
	EMsg_ClientInstallClientAppResponse                           EMsg = 5521
	EMsg_ClientUninstallClientApp                                 EMsg = 5522
	EMsg_ClientUninstallClientAppResponse                         EMsg = 5523
	EMsg_ClientSetClientAppUpdateState                            EMsg = 5524
	EMsg_ClientSetClientAppUpdateStateResponse                    EMsg = 5525
	EMsg_ClientRequestEncryptedAppTicket                          EMsg = 5526
	EMsg_ClientRequestEncryptedAppTicketResponse                  EMsg = 5527
	EMsg_ClientWalletInfoUpdate                                   EMsg = 5528
	EMsg_ClientLBSSetUGC                                          EMsg = 5529
	EMsg_ClientLBSSetUGCResponse                                  EMsg = 5530
	EMsg_ClientAMGetClanOfficers                                  EMsg = 5531
	EMsg_ClientAMGetClanOfficersResponse                          EMsg = 5532
	EMsg_ClientCheckFileSignature                                 EMsg = 5533 // removed
	EMsg_ClientCheckFileSignatureResponse                         EMsg = 5534 // removed
	EMsg_ClientFriendProfileInfo                                  EMsg = 5535
	EMsg_ClientFriendProfileInfoResponse                          EMsg = 5536
	EMsg_ClientUpdateMachineAuth                                  EMsg = 5537
	EMsg_ClientUpdateMachineAuthResponse                          EMsg = 5538
	EMsg_ClientReadMachineAuth                                    EMsg = 5539
	EMsg_ClientReadMachineAuthResponse                            EMsg = 5540
	EMsg_ClientRequestMachineAuth                                 EMsg = 5541
	EMsg_ClientRequestMachineAuthResponse                         EMsg = 5542
	EMsg_ClientScreenshotsChanged                                 EMsg = 5543
	EMsg_ClientEmailChange4                                       EMsg = 5544
	EMsg_ClientEmailChangeResponse4                               EMsg = 5545
	EMsg_ClientGetCDNAuthToken                                    EMsg = 5546
	EMsg_ClientGetCDNAuthTokenResponse                            EMsg = 5547
	EMsg_ClientDownloadRateStatistics                             EMsg = 5548
	EMsg_ClientRequestAccountData                                 EMsg = 5549
	EMsg_ClientRequestAccountDataResponse                         EMsg = 5550
	EMsg_ClientResetForgottenPassword4                            EMsg = 5551
	EMsg_ClientHideFriend                                         EMsg = 5552
	EMsg_ClientFriendsGroupsList                                  EMsg = 5553
	EMsg_ClientGetClanActivityCounts                              EMsg = 5554
	EMsg_ClientGetClanActivityCountsResponse                      EMsg = 5555
	EMsg_ClientOGSReportString                                    EMsg = 5556
	EMsg_ClientOGSReportBug                                       EMsg = 5557
	EMsg_ClientSentLogs                                           EMsg = 5558
	EMsg_ClientLogonGameServer                                    EMsg = 5559
	EMsg_AMClientCreateFriendsGroup                               EMsg = 5560
	EMsg_AMClientCreateFriendsGroupResponse                       EMsg = 5561
	EMsg_AMClientDeleteFriendsGroup                               EMsg = 5562
	EMsg_AMClientDeleteFriendsGroupResponse                       EMsg = 5563
	EMsg_AMClientRenameFriendsGroup                               EMsg = 5564
	EMsg_AMClientRenameFriendsGroupResponse                       EMsg = 5565
	EMsg_AMClientAddFriendToGroup                                 EMsg = 5566
	EMsg_AMClientAddFriendToGroupResponse                         EMsg = 5567
	EMsg_AMClientRemoveFriendFromGroup                            EMsg = 5568
	EMsg_AMClientRemoveFriendFromGroupResponse                    EMsg = 5569
	EMsg_ClientAMGetPersonaNameHistory                            EMsg = 5570
	EMsg_ClientAMGetPersonaNameHistoryResponse                    EMsg = 5571
	EMsg_ClientRequestFreeLicense                                 EMsg = 5572
	EMsg_ClientRequestFreeLicenseResponse                         EMsg = 5573
	EMsg_ClientDRMDownloadRequestWithCrashData                    EMsg = 5574
	EMsg_ClientAuthListAck                                        EMsg = 5575
	EMsg_ClientItemAnnouncements                                  EMsg = 5576
	EMsg_ClientRequestItemAnnouncements                           EMsg = 5577
	EMsg_ClientFriendMsgEchoToSender                              EMsg = 5578
	EMsg_ClientChangeSteamGuardOptions                            EMsg = 5579 // removed
	EMsg_ClientChangeSteamGuardOptionsResponse                    EMsg = 5580 // removed
	EMsg_ClientOGSGameServerPingSample                            EMsg = 5581
	EMsg_ClientCommentNotifications                               EMsg = 5582
	EMsg_ClientRequestCommentNotifications                        EMsg = 5583
	EMsg_ClientPersonaChangeResponse                              EMsg = 5584
	EMsg_ClientRequestWebAPIAuthenticateUserNonce                 EMsg = 5585
	EMsg_ClientRequestWebAPIAuthenticateUserNonceResponse         EMsg = 5586
	EMsg_ClientPlayerNicknameList                                 EMsg = 5587
	EMsg_AMClientSetPlayerNickname                                EMsg = 5588
	EMsg_AMClientSetPlayerNicknameResponse                        EMsg = 5589
	EMsg_ClientRequestOAuthTokenForApp                            EMsg = 5590 // removed
	EMsg_ClientRequestOAuthTokenForAppResponse                    EMsg = 5591 // removed
	EMsg_ClientCreateAccountProto                                 EMsg = 5590
	EMsg_ClientCreateAccountProtoResponse                         EMsg = 5591
	EMsg_ClientGetNumberOfCurrentPlayersDP                        EMsg = 5592
	EMsg_ClientGetNumberOfCurrentPlayersDPResponse                EMsg = 5593
	EMsg_ClientServiceMethod                                      EMsg = 5594
	EMsg_ClientServiceMethodResponse                              EMsg = 5595
	EMsg_ClientFriendUserStatusPublished                          EMsg = 5596
	EMsg_ClientCurrentUIMode                                      EMsg = 5597
	EMsg_ClientVanityURLChangedNotification                       EMsg = 5598
	EMsg_ClientUserNotifications                                  EMsg = 5599
	EMsg_BaseDFS                                                  EMsg = 5600
	EMsg_DFSGetFile                                               EMsg = 5601
	EMsg_DFSInstallLocalFile                                      EMsg = 5602
	EMsg_DFSConnection                                            EMsg = 5603
	EMsg_DFSConnectionReply                                       EMsg = 5604
	EMsg_ClientDFSAuthenticateRequest                             EMsg = 5605
	EMsg_ClientDFSAuthenticateResponse                            EMsg = 5606
	EMsg_ClientDFSEndSession                                      EMsg = 5607
	EMsg_DFSPurgeFile                                             EMsg = 5608
	EMsg_DFSRouteFile                                             EMsg = 5609
	EMsg_DFSGetFileFromServer                                     EMsg = 5610
	EMsg_DFSAcceptedResponse                                      EMsg = 5611
	EMsg_DFSRequestPingback                                       EMsg = 5612
	EMsg_DFSRecvTransmitFile                                      EMsg = 5613
	EMsg_DFSSendTransmitFile                                      EMsg = 5614
	EMsg_DFSRequestPingback2                                      EMsg = 5615
	EMsg_DFSResponsePingback2                                     EMsg = 5616
	EMsg_ClientDFSDownloadStatus                                  EMsg = 5617
	EMsg_DFSStartTransfer                                         EMsg = 5618
	EMsg_DFSTransferComplete                                      EMsg = 5619
	EMsg_DFSRouteFileResponse                                     EMsg = 5620
	EMsg_ClientNetworkingCertRequest                              EMsg = 5621
	EMsg_ClientNetworkingCertRequestResponse                      EMsg = 5622
	EMsg_BaseMDS                                                  EMsg = 5800
	EMsg_ClientMDSLoginRequest                                    EMsg = 5801 // removed
	EMsg_ClientMDSLoginResponse                                   EMsg = 5802 // removed
	EMsg_ClientMDSUploadManifestRequest                           EMsg = 5803 // removed
	EMsg_ClientMDSUploadManifestResponse                          EMsg = 5804 // removed
	EMsg_ClientMDSTransmitManifestDataChunk                       EMsg = 5805 // removed
	EMsg_ClientMDSHeartbeat                                       EMsg = 5806 // removed
	EMsg_ClientMDSUploadDepotChunks                               EMsg = 5807 // removed
	EMsg_ClientMDSUploadDepotChunksResponse                       EMsg = 5808 // removed
	EMsg_ClientMDSInitDepotBuildRequest                           EMsg = 5809 // removed
	EMsg_ClientMDSInitDepotBuildResponse                          EMsg = 5810 // removed
	EMsg_AMToMDSGetDepotDecryptionKey                             EMsg = 5812
	EMsg_MDSToAMGetDepotDecryptionKeyResponse                     EMsg = 5813
	EMsg_MDSGetVersionsForDepot                                   EMsg = 5814 // removed
	EMsg_MDSGetVersionsForDepotResponse                           EMsg = 5815 // removed
	EMsg_MDSSetPublicVersionForDepot                              EMsg = 5816 // removed
	EMsg_MDSSetPublicVersionForDepotResponse                      EMsg = 5817 // removed
	EMsg_ClientMDSInitWorkshopBuildRequest                        EMsg = 5816 // removed
	EMsg_ClientMDSInitWorkshopBuildResponse                       EMsg = 5817 // removed
	EMsg_ClientMDSGetDepotManifest                                EMsg = 5818 // removed
	EMsg_ClientMDSGetDepotManifestResponse                        EMsg = 5819 // removed
	EMsg_ClientMDSGetDepotManifestChunk                           EMsg = 5820 // removed
	EMsg_ClientMDSUploadRateTest                                  EMsg = 5823 // removed
	EMsg_ClientMDSUploadRateTestResponse                          EMsg = 5824 // removed
	EMsg_MDSDownloadDepotChunksAck                                EMsg = 5825 // removed
	EMsg_MDSContentServerStatsBroadcast                           EMsg = 5826 // removed
	EMsg_MDSContentServerConfigRequest                            EMsg = 5827
	EMsg_MDSContentServerConfig                                   EMsg = 5828
	EMsg_MDSGetDepotManifest                                      EMsg = 5829
	EMsg_MDSGetDepotManifestResponse                              EMsg = 5830
	EMsg_MDSGetDepotManifestChunk                                 EMsg = 5831
	EMsg_MDSGetDepotChunk                                         EMsg = 5832
	EMsg_MDSGetDepotChunkResponse                                 EMsg = 5833
	EMsg_MDSGetDepotChunkChunk                                    EMsg = 5834
	EMsg_MDSUpdateContentServerConfig                             EMsg = 5835 // removed
	EMsg_MDSGetServerListForUser                                  EMsg = 5836
	EMsg_MDSGetServerListForUserResponse                          EMsg = 5837
	EMsg_ClientMDSRegisterAppBuild                                EMsg = 5838 // removed
	EMsg_ClientMDSRegisterAppBuildResponse                        EMsg = 5839 // removed
	EMsg_ClientMDSSetAppBuildLive                                 EMsg = 5840 // removed
	EMsg_ClientMDSSetAppBuildLiveResponse                         EMsg = 5841 // removed
	EMsg_ClientMDSGetPrevDepotBuild                               EMsg = 5842 // removed
	EMsg_ClientMDSGetPrevDepotBuildResponse                       EMsg = 5843 // removed
	EMsg_MDSToCSFlushChunk                                        EMsg = 5844
	EMsg_ClientMDSSignInstallScript                               EMsg = 5845 // removed
	EMsg_ClientMDSSignInstallScriptResponse                       EMsg = 5846 // removed
	EMsg_MDSMigrateChunk                                          EMsg = 5847
	EMsg_MDSMigrateChunkResponse                                  EMsg = 5848
	EMsg_CSBase                                                   EMsg = 6200
	EMsg_CSPing                                                   EMsg = 6201
	EMsg_CSPingResponse                                           EMsg = 6202
	EMsg_GMSBase                                                  EMsg = 6400
	EMsg_GMSGameServerReplicate                                   EMsg = 6401
	EMsg_ClientGMSServerQuery                                     EMsg = 6403
	EMsg_GMSClientServerQueryResponse                             EMsg = 6404
	EMsg_AMGMSGameServerUpdate                                    EMsg = 6405
	EMsg_AMGMSGameServerRemove                                    EMsg = 6406
	EMsg_GameServerOutOfDate                                      EMsg = 6407
	EMsg_DeviceAuthorizationBase                                  EMsg = 6500
	EMsg_ClientAuthorizeLocalDeviceRequest                        EMsg = 6501
	EMsg_ClientAuthorizeLocalDevice                               EMsg = 6502 // removed
	EMsg_ClientAuthorizeLocalDeviceResponse                       EMsg = 6502
	EMsg_ClientDeauthorizeDeviceRequest                           EMsg = 6503
	EMsg_ClientDeauthorizeDevice                                  EMsg = 6504
	EMsg_ClientUseLocalDeviceAuthorizations                       EMsg = 6505
	EMsg_ClientGetAuthorizedDevices                               EMsg = 6506
	EMsg_ClientGetAuthorizedDevicesResponse                       EMsg = 6507
	EMsg_AMNotifySessionDeviceAuthorized                          EMsg = 6508
	EMsg_ClientAuthorizeLocalDeviceNotification                   EMsg = 6509
	EMsg_MMSBase                                                  EMsg = 6600
	EMsg_ClientMMSCreateLobby                                     EMsg = 6601
	EMsg_ClientMMSCreateLobbyResponse                             EMsg = 6602
	EMsg_ClientMMSJoinLobby                                       EMsg = 6603
	EMsg_ClientMMSJoinLobbyResponse                               EMsg = 6604
	EMsg_ClientMMSLeaveLobby                                      EMsg = 6605
	EMsg_ClientMMSLeaveLobbyResponse                              EMsg = 6606
	EMsg_ClientMMSGetLobbyList                                    EMsg = 6607
	EMsg_ClientMMSGetLobbyListResponse                            EMsg = 6608
	EMsg_ClientMMSSetLobbyData                                    EMsg = 6609
	EMsg_ClientMMSSetLobbyDataResponse                            EMsg = 6610
	EMsg_ClientMMSGetLobbyData                                    EMsg = 6611
	EMsg_ClientMMSLobbyData                                       EMsg = 6612
	EMsg_ClientMMSSendLobbyChatMsg                                EMsg = 6613
	EMsg_ClientMMSLobbyChatMsg                                    EMsg = 6614
	EMsg_ClientMMSSetLobbyOwner                                   EMsg = 6615
	EMsg_ClientMMSSetLobbyOwnerResponse                           EMsg = 6616
	EMsg_ClientMMSSetLobbyGameServer                              EMsg = 6617
	EMsg_ClientMMSLobbyGameServerSet                              EMsg = 6618
	EMsg_ClientMMSUserJoinedLobby                                 EMsg = 6619
	EMsg_ClientMMSUserLeftLobby                                   EMsg = 6620
	EMsg_ClientMMSInviteToLobby                                   EMsg = 6621
	EMsg_ClientMMSFlushFrenemyListCache                           EMsg = 6622
	EMsg_ClientMMSFlushFrenemyListCacheResponse                   EMsg = 6623
	EMsg_ClientMMSSetLobbyLinked                                  EMsg = 6624
	EMsg_NonStdMsgBase                                            EMsg = 6800
	EMsg_NonStdMsgMemcached                                       EMsg = 6801
	EMsg_NonStdMsgHTTPServer                                      EMsg = 6802
	EMsg_NonStdMsgHTTPClient                                      EMsg = 6803
	EMsg_NonStdMsgWGResponse                                      EMsg = 6804
	EMsg_NonStdMsgPHPSimulator                                    EMsg = 6805
	EMsg_NonStdMsgChase                                           EMsg = 6806
	EMsg_NonStdMsgDFSTransfer                                     EMsg = 6807
	EMsg_NonStdMsgTests                                           EMsg = 6808
	EMsg_NonStdMsgUMQpipeAAPL                                     EMsg = 6809
	EMsg_NonStdMsgSyslog                                          EMsg = 6810 // removed
	EMsg_NonStdMsgLogsink                                         EMsg = 6811
	EMsg_NonStdMsgSteam2Emulator                                  EMsg = 6812
	EMsg_NonStdMsgRTMPServer                                      EMsg = 6813
	EMsg_UDSBase                                                  EMsg = 7000
	EMsg_ClientUDSP2PSessionStarted                               EMsg = 7001
	EMsg_ClientUDSP2PSessionEnded                                 EMsg = 7002
	EMsg_UDSRenderUserAuth                                        EMsg = 7003
	EMsg_UDSRenderUserAuthResponse                                EMsg = 7004
	EMsg_ClientUDSInviteToGame                                    EMsg = 7005
	EMsg_UDSFindSession                                           EMsg = 7006 // removed: renamed to UDSHasSession
	EMsg_UDSHasSession                                            EMsg = 7006
	EMsg_UDSFindSessionResponse                                   EMsg = 7007 // removed: renamed to UDSHasSessionResponse
	EMsg_UDSHasSessionResponse                                    EMsg = 7007
	EMsg_MPASBase                                                 EMsg = 7100
	EMsg_MPASVacBanReset                                          EMsg = 7101
	EMsg_KGSBase                                                  EMsg = 7200
	EMsg_KGSAllocateKeyRange                                      EMsg = 7201 // removed
	EMsg_KGSAllocateKeyRangeResponse                              EMsg = 7202 // removed
	EMsg_KGSGenerateKeys                                          EMsg = 7203 // removed
	EMsg_KGSGenerateKeysResponse                                  EMsg = 7204 // removed
	EMsg_KGSRemapKeys                                             EMsg = 7205 // removed
	EMsg_KGSRemapKeysResponse                                     EMsg = 7206 // removed
	EMsg_KGSGenerateGameStopWCKeys                                EMsg = 7207 // removed
	EMsg_KGSGenerateGameStopWCKeysResponse                        EMsg = 7208 // removed
	EMsg_UCMBase                                                  EMsg = 7300
	EMsg_ClientUCMAddScreenshot                                   EMsg = 7301
	EMsg_ClientUCMAddScreenshotResponse                           EMsg = 7302
	EMsg_UCMValidateObjectExists                                  EMsg = 7303 // removed
	EMsg_UCMValidateObjectExistsResponse                          EMsg = 7304 // removed
	EMsg_UCMResetCommunityContent                                 EMsg = 7307
	EMsg_UCMResetCommunityContentResponse                         EMsg = 7308
	EMsg_ClientUCMDeleteScreenshot                                EMsg = 7309
	EMsg_ClientUCMDeleteScreenshotResponse                        EMsg = 7310
	EMsg_ClientUCMPublishFile                                     EMsg = 7311
	EMsg_ClientUCMPublishFileResponse                             EMsg = 7312
	EMsg_ClientUCMGetPublishedFileDetails                         EMsg = 7313 // removed
	EMsg_ClientUCMGetPublishedFileDetailsResponse                 EMsg = 7314 // removed
	EMsg_ClientUCMDeletePublishedFile                             EMsg = 7315
	EMsg_ClientUCMDeletePublishedFileResponse                     EMsg = 7316
	EMsg_ClientUCMEnumerateUserPublishedFiles                     EMsg = 7317
	EMsg_ClientUCMEnumerateUserPublishedFilesResponse             EMsg = 7318
	EMsg_ClientUCMSubscribePublishedFile                          EMsg = 7319 // removed
	EMsg_ClientUCMSubscribePublishedFileResponse                  EMsg = 7320 // removed
	EMsg_ClientUCMEnumerateUserSubscribedFiles                    EMsg = 7321
	EMsg_ClientUCMEnumerateUserSubscribedFilesResponse            EMsg = 7322
	EMsg_ClientUCMUnsubscribePublishedFile                        EMsg = 7323 // removed
	EMsg_ClientUCMUnsubscribePublishedFileResponse                EMsg = 7324 // removed
	EMsg_ClientUCMUpdatePublishedFile                             EMsg = 7325
	EMsg_ClientUCMUpdatePublishedFileResponse                     EMsg = 7326
	EMsg_UCMUpdatePublishedFile                                   EMsg = 7327
	EMsg_UCMUpdatePublishedFileResponse                           EMsg = 7328
	EMsg_UCMDeletePublishedFile                                   EMsg = 7329
	EMsg_UCMDeletePublishedFileResponse                           EMsg = 7330
	EMsg_UCMUpdatePublishedFileStat                               EMsg = 7331
	EMsg_UCMUpdatePublishedFileBan                                EMsg = 7332
	EMsg_UCMUpdatePublishedFileBanResponse                        EMsg = 7333
	EMsg_UCMUpdateTaggedScreenshot                                EMsg = 7334 // removed
	EMsg_UCMAddTaggedScreenshot                                   EMsg = 7335 // removed
	EMsg_UCMRemoveTaggedScreenshot                                EMsg = 7336 // removed
	EMsg_UCMReloadPublishedFile                                   EMsg = 7337
	EMsg_UCMReloadUserFileListCaches                              EMsg = 7338
	EMsg_UCMPublishedFileReported                                 EMsg = 7339
	EMsg_UCMUpdatePublishedFileIncompatibleStatus                 EMsg = 7340
	EMsg_UCMPublishedFilePreviewAdd                               EMsg = 7341
	EMsg_UCMPublishedFilePreviewAddResponse                       EMsg = 7342
	EMsg_UCMPublishedFilePreviewRemove                            EMsg = 7343
	EMsg_UCMPublishedFilePreviewRemoveResponse                    EMsg = 7344
	EMsg_UCMPublishedFilePreviewChangeSortOrder                   EMsg = 7345 // removed
	EMsg_UCMPublishedFilePreviewChangeSortOrderResponse           EMsg = 7346 // removed
	EMsg_ClientUCMPublishedFileSubscribed                         EMsg = 7347
	EMsg_ClientUCMPublishedFileUnsubscribed                       EMsg = 7348
	EMsg_UCMPublishedFileSubscribed                               EMsg = 7349
	EMsg_UCMPublishedFileUnsubscribed                             EMsg = 7350
	EMsg_UCMPublishFile                                           EMsg = 7351
	EMsg_UCMPublishFileResponse                                   EMsg = 7352
	EMsg_UCMPublishedFileChildAdd                                 EMsg = 7353
	EMsg_UCMPublishedFileChildAddResponse                         EMsg = 7354
	EMsg_UCMPublishedFileChildRemove                              EMsg = 7355
	EMsg_UCMPublishedFileChildRemoveResponse                      EMsg = 7356
	EMsg_UCMPublishedFileChildChangeSortOrder                     EMsg = 7357 // removed
	EMsg_UCMPublishedFileChildChangeSortOrderResponse             EMsg = 7358 // removed
	EMsg_UCMPublishedFileParentChanged                            EMsg = 7359
	EMsg_ClientUCMGetPublishedFilesForUser                        EMsg = 7360
	EMsg_ClientUCMGetPublishedFilesForUserResponse                EMsg = 7361
	EMsg_UCMGetPublishedFilesForUser                              EMsg = 7362 // removed
	EMsg_UCMGetPublishedFilesForUserResponse                      EMsg = 7363 // removed
	EMsg_ClientUCMSetUserPublishedFileAction                      EMsg = 7364
	EMsg_ClientUCMSetUserPublishedFileActionResponse              EMsg = 7365
	EMsg_ClientUCMEnumeratePublishedFilesByUserAction             EMsg = 7366
	EMsg_ClientUCMEnumeratePublishedFilesByUserActionResponse     EMsg = 7367
	EMsg_ClientUCMPublishedFileDeleted                            EMsg = 7368
	EMsg_UCMGetUserSubscribedFiles                                EMsg = 7369
	EMsg_UCMGetUserSubscribedFilesResponse                        EMsg = 7370
	EMsg_UCMFixStatsPublishedFile                                 EMsg = 7371
	EMsg_UCMDeleteOldScreenshot                                   EMsg = 7372 // removed
	EMsg_UCMDeleteOldScreenshotResponse                           EMsg = 7373 // removed
	EMsg_UCMDeleteOldVideo                                        EMsg = 7374 // removed
	EMsg_UCMDeleteOldVideoResponse                                EMsg = 7375 // removed
	EMsg_UCMUpdateOldScreenshotPrivacy                            EMsg = 7376 // removed
	EMsg_UCMUpdateOldScreenshotPrivacyResponse                    EMsg = 7377 // removed
	EMsg_ClientUCMEnumerateUserSubscribedFilesWithUpdates         EMsg = 7378
	EMsg_ClientUCMEnumerateUserSubscribedFilesWithUpdatesResponse EMsg = 7379
	EMsg_UCMPublishedFileContentUpdated                           EMsg = 7380
	EMsg_UCMPublishedFileUpdated                                  EMsg = 7381
	EMsg_ClientWorkshopItemChangesRequest                         EMsg = 7382
	EMsg_ClientWorkshopItemChangesResponse                        EMsg = 7383
	EMsg_ClientWorkshopItemInfoRequest                            EMsg = 7384
	EMsg_ClientWorkshopItemInfoResponse                           EMsg = 7385
	EMsg_FSBase                                                   EMsg = 7500
	EMsg_ClientRichPresenceUpload                                 EMsg = 7501
	EMsg_ClientRichPresenceRequest                                EMsg = 7502
	EMsg_ClientRichPresenceInfo                                   EMsg = 7503
	EMsg_FSRichPresenceRequest                                    EMsg = 7504
	EMsg_FSRichPresenceResponse                                   EMsg = 7505
	EMsg_FSComputeFrenematrix                                     EMsg = 7506
	EMsg_FSComputeFrenematrixResponse                             EMsg = 7507
	EMsg_FSPlayStatusNotification                                 EMsg = 7508
	EMsg_FSPublishPersonaStatus                                   EMsg = 7509
	EMsg_FSAddOrRemoveFollower                                    EMsg = 7510
	EMsg_FSAddOrRemoveFollowerResponse                            EMsg = 7511
	EMsg_FSUpdateFollowingList                                    EMsg = 7512
	EMsg_FSCommentNotification                                    EMsg = 7513
	EMsg_FSCommentNotificationViewed                              EMsg = 7514
	EMsg_ClientFSGetFollowerCount                                 EMsg = 7515
	EMsg_ClientFSGetFollowerCountResponse                         EMsg = 7516
	EMsg_ClientFSGetIsFollowing                                   EMsg = 7517
	EMsg_ClientFSGetIsFollowingResponse                           EMsg = 7518
	EMsg_ClientFSEnumerateFollowingList                           EMsg = 7519
	EMsg_ClientFSEnumerateFollowingListResponse                   EMsg = 7520
	EMsg_FSGetPendingNotificationCount                            EMsg = 7521
	EMsg_FSGetPendingNotificationCountResponse                    EMsg = 7522
	EMsg_ClientFSOfflineMessageNotification                       EMsg = 7523 // obsolete: Renamed to ClientChatOfflineMessageNotification
	EMsg_ClientFSRequestOfflineMessageCount                       EMsg = 7524 // obsolete: Renamed to ClientChatRequestOfflineMessageCount
	EMsg_ClientFSGetFriendMessageHistory                          EMsg = 7525 // obsolete: Renamed to ClientChatGetFriendMessageHistory
	EMsg_ClientFSGetFriendMessageHistoryResponse                  EMsg = 7526 // obsolete: Renamed to ClientChatGetFriendMessageHistoryResponse
	EMsg_ClientFSGetFriendMessageHistoryForOfflineMessages        EMsg = 7527 // obsolete: Renamed to ClientChatGetFriendMessageHistoryForOfflineMessages
	EMsg_ClientChatOfflineMessageNotification                     EMsg = 7523
	EMsg_ClientChatRequestOfflineMessageCount                     EMsg = 7524
	EMsg_ClientChatGetFriendMessageHistory                        EMsg = 7525
	EMsg_ClientChatGetFriendMessageHistoryResponse                EMsg = 7526
	EMsg_ClientChatGetFriendMessageHistoryForOfflineMessages      EMsg = 7527
	EMsg_ClientFSGetFriendsSteamLevels                            EMsg = 7528
	EMsg_ClientFSGetFriendsSteamLevelsResponse                    EMsg = 7529
	EMsg_FSRequestFriendData                                      EMsg = 7530
	EMsg_DRMRange2                                                EMsg = 7600
	EMsg_CEGVersionSetEnableDisableRequest                        EMsg = 7600
	EMsg_CEGVersionSetEnableDisableResponse                       EMsg = 7601
	EMsg_CEGPropStatusDRMSRequest                                 EMsg = 7602
	EMsg_CEGPropStatusDRMSResponse                                EMsg = 7603
	EMsg_CEGWhackFailureReportRequest                             EMsg = 7604
	EMsg_CEGWhackFailureReportResponse                            EMsg = 7605
	EMsg_DRMSFetchVersionSet                                      EMsg = 7606
	EMsg_DRMSFetchVersionSetResponse                              EMsg = 7607
	EMsg_EconBase                                                 EMsg = 7700
	EMsg_EconTrading_InitiateTradeRequest                         EMsg = 7701
	EMsg_EconTrading_InitiateTradeProposed                        EMsg = 7702
	EMsg_EconTrading_InitiateTradeResponse                        EMsg = 7703
	EMsg_EconTrading_InitiateTradeResult                          EMsg = 7704
	EMsg_EconTrading_StartSession                                 EMsg = 7705
	EMsg_EconTrading_CancelTradeRequest                           EMsg = 7706
	EMsg_EconFlushInventoryCache                                  EMsg = 7707
	EMsg_EconFlushInventoryCacheResponse                          EMsg = 7708
	EMsg_EconCDKeyProcessTransaction                              EMsg = 7711
	EMsg_EconCDKeyProcessTransactionResponse                      EMsg = 7712
	EMsg_EconGetErrorLogs                                         EMsg = 7713
	EMsg_EconGetErrorLogsResponse                                 EMsg = 7714
	EMsg_RMRange                                                  EMsg = 7800
	EMsg_RMTestVerisignOTP                                        EMsg = 7800
	EMsg_RMTestVerisignOTPResponse                                EMsg = 7801
	EMsg_RMDeleteMemcachedKeys                                    EMsg = 7803
	EMsg_RMRemoteInvoke                                           EMsg = 7804
	EMsg_BadLoginIPList                                           EMsg = 7805
	EMsg_RMMsgTraceAddTrigger                                     EMsg = 7806
	EMsg_RMMsgTraceRemoveTrigger                                  EMsg = 7807
	EMsg_RMMsgTraceEvent                                          EMsg = 7808
	EMsg_UGSBase                                                  EMsg = 7900
	EMsg_UGSUpdateGlobalStats                                     EMsg = 7900
	EMsg_ClientUGSGetGlobalStats                                  EMsg = 7901
	EMsg_ClientUGSGetGlobalStatsResponse                          EMsg = 7902
	EMsg_StoreBase                                                EMsg = 8000
	EMsg_StoreUpdateRecommendationCount                           EMsg = 8000 // removed
	EMsg_UMQBase                                                  EMsg = 8100
	EMsg_UMQLogonRequest                                          EMsg = 8100
	EMsg_UMQLogonResponse                                         EMsg = 8101
	EMsg_UMQLogoffRequest                                         EMsg = 8102
	EMsg_UMQLogoffResponse                                        EMsg = 8103
	EMsg_UMQSendChatMessage                                       EMsg = 8104
	EMsg_UMQIncomingChatMessage                                   EMsg = 8105
	EMsg_UMQPoll                                                  EMsg = 8106
	EMsg_UMQPollResults                                           EMsg = 8107
	EMsg_UMQ2AM_ClientMsgBatch                                    EMsg = 8108
	EMsg_UMQEnqueueMobileSalePromotions                           EMsg = 8109 // removed
	EMsg_UMQEnqueueMobileAnnouncements                            EMsg = 8110 // removed
	EMsg_WorkshopBase                                             EMsg = 8200
	EMsg_WorkshopAcceptTOSRequest                                 EMsg = 8200 // removed
	EMsg_WorkshopAcceptTOSResponse                                EMsg = 8201 // removed
	EMsg_WebAPIBase                                               EMsg = 8300
	EMsg_WebAPIValidateOAuth2Token                                EMsg = 8300
	EMsg_WebAPIValidateOAuth2TokenResponse                        EMsg = 8301
	EMsg_WebAPIInvalidateTokensForAccount                         EMsg = 8302 // removed
	EMsg_WebAPIRegisterGCInterfaces                               EMsg = 8303
	EMsg_WebAPIInvalidateOAuthClientCache                         EMsg = 8304
	EMsg_WebAPIInvalidateOAuthTokenCache                          EMsg = 8305
	EMsg_WebAPISetSecrets                                         EMsg = 8306
	EMsg_BackpackBase                                             EMsg = 8400
	EMsg_BackpackAddToCurrency                                    EMsg = 8401
	EMsg_BackpackAddToCurrencyResponse                            EMsg = 8402
	EMsg_CREBase                                                  EMsg = 8500
	EMsg_CRERankByTrend                                           EMsg = 8501 // removed
	EMsg_CRERankByTrendResponse                                   EMsg = 8502 // removed
	EMsg_CREItemVoteSummary                                       EMsg = 8503
	EMsg_CREItemVoteSummaryResponse                               EMsg = 8504
	EMsg_CRERankByVote                                            EMsg = 8505 // removed
	EMsg_CRERankByVoteResponse                                    EMsg = 8506 // removed
	EMsg_CREUpdateUserPublishedItemVote                           EMsg = 8507
	EMsg_CREUpdateUserPublishedItemVoteResponse                   EMsg = 8508
	EMsg_CREGetUserPublishedItemVoteDetails                       EMsg = 8509
	EMsg_CREGetUserPublishedItemVoteDetailsResponse               EMsg = 8510
	EMsg_CREEnumeratePublishedFiles                               EMsg = 8511
	EMsg_CREEnumeratePublishedFilesResponse                       EMsg = 8512
	EMsg_CREPublishedFileVoteAdded                                EMsg = 8513
	EMsg_SecretsBase                                              EMsg = 8600
	EMsg_SecretsRequestCredentialPair                             EMsg = 8600
	EMsg_SecretsCredentialPairResponse                            EMsg = 8601
	EMsg_SecretsRequestServerIdentity                             EMsg = 8602 // removed
	EMsg_SecretsServerIdentityResponse                            EMsg = 8603 // removed
	EMsg_SecretsUpdateServerIdentities                            EMsg = 8604 // removed
	EMsg_BoxMonitorBase                                           EMsg = 8700
	EMsg_BoxMonitorReportRequest                                  EMsg = 8700
	EMsg_BoxMonitorReportResponse                                 EMsg = 8701
	EMsg_LogsinkBase                                              EMsg = 8800
	EMsg_LogsinkWriteReport                                       EMsg = 8800
	EMsg_PICSBase                                                 EMsg = 8900
	EMsg_ClientPICSChangesSinceRequest                            EMsg = 8901
	EMsg_ClientPICSChangesSinceResponse                           EMsg = 8902
	EMsg_ClientPICSProductInfoRequest                             EMsg = 8903
	EMsg_ClientPICSProductInfoResponse                            EMsg = 8904
	EMsg_ClientPICSAccessTokenRequest                             EMsg = 8905
	EMsg_ClientPICSAccessTokenResponse                            EMsg = 8906
	EMsg_WorkerProcess                                            EMsg = 9000
	EMsg_WorkerProcessPingRequest                                 EMsg = 9000
	EMsg_WorkerProcessPingResponse                                EMsg = 9001
	EMsg_WorkerProcessShutdown                                    EMsg = 9002
	EMsg_DRMWorkerProcess                                         EMsg = 9100
	EMsg_DRMWorkerProcessDRMAndSign                               EMsg = 9100
	EMsg_DRMWorkerProcessDRMAndSignResponse                       EMsg = 9101
	EMsg_DRMWorkerProcessSteamworksInfoRequest                    EMsg = 9102
	EMsg_DRMWorkerProcessSteamworksInfoResponse                   EMsg = 9103
	EMsg_DRMWorkerProcessInstallDRMDLLRequest                     EMsg = 9104
	EMsg_DRMWorkerProcessInstallDRMDLLResponse                    EMsg = 9105
	EMsg_DRMWorkerProcessSecretIdStringRequest                    EMsg = 9106
	EMsg_DRMWorkerProcessSecretIdStringResponse                   EMsg = 9107
	EMsg_DRMWorkerProcessGetDRMGuidsFromFileRequest               EMsg = 9108 // removed
	EMsg_DRMWorkerProcessGetDRMGuidsFromFileResponse              EMsg = 9109 // removed
	EMsg_DRMWorkerProcessInstallProcessedFilesRequest             EMsg = 9110
	EMsg_DRMWorkerProcessInstallProcessedFilesResponse            EMsg = 9111
	EMsg_DRMWorkerProcessExamineBlobRequest                       EMsg = 9112
	EMsg_DRMWorkerProcessExamineBlobResponse                      EMsg = 9113
	EMsg_DRMWorkerProcessDescribeSecretRequest                    EMsg = 9114
	EMsg_DRMWorkerProcessDescribeSecretResponse                   EMsg = 9115
	EMsg_DRMWorkerProcessBackfillOriginalRequest                  EMsg = 9116
	EMsg_DRMWorkerProcessBackfillOriginalResponse                 EMsg = 9117
	EMsg_DRMWorkerProcessValidateDRMDLLRequest                    EMsg = 9118
	EMsg_DRMWorkerProcessValidateDRMDLLResponse                   EMsg = 9119
	EMsg_DRMWorkerProcessValidateFileRequest                      EMsg = 9120
	EMsg_DRMWorkerProcessValidateFileResponse                     EMsg = 9121
	EMsg_DRMWorkerProcessSplitAndInstallRequest                   EMsg = 9122
	EMsg_DRMWorkerProcessSplitAndInstallResponse                  EMsg = 9123
	EMsg_DRMWorkerProcessGetBlobRequest                           EMsg = 9124
	EMsg_DRMWorkerProcessGetBlobResponse                          EMsg = 9125
	EMsg_DRMWorkerProcessEvaluateCrashRequest                     EMsg = 9126
	EMsg_DRMWorkerProcessEvaluateCrashResponse                    EMsg = 9127
	EMsg_DRMWorkerProcessAnalyzeFileRequest                       EMsg = 9128
	EMsg_DRMWorkerProcessAnalyzeFileResponse                      EMsg = 9129
	EMsg_DRMWorkerProcessUnpackBlobRequest                        EMsg = 9130
	EMsg_DRMWorkerProcessUnpackBlobResponse                       EMsg = 9131
	EMsg_DRMWorkerProcessInstallAllRequest                        EMsg = 9132
	EMsg_DRMWorkerProcessInstallAllResponse                       EMsg = 9133
	EMsg_TestWorkerProcess                                        EMsg = 9200
	EMsg_TestWorkerProcessLoadUnloadModuleRequest                 EMsg = 9200
	EMsg_TestWorkerProcessLoadUnloadModuleResponse                EMsg = 9201
	EMsg_TestWorkerProcessServiceModuleCallRequest                EMsg = 9202
	EMsg_TestWorkerProcessServiceModuleCallResponse               EMsg = 9203
	EMsg_QuestServerBase                                          EMsg = 9300
	EMsg_ClientGetEmoticonList                                    EMsg = 9330
	EMsg_ClientEmoticonList                                       EMsg = 9331
	EMsg_ClientSharedLibraryBase                                  EMsg = 9400 // removed: renamed to SLCBase
	EMsg_SLCBase                                                  EMsg = 9400
	EMsg_SLCUserSessionStatus                                     EMsg = 9400
	EMsg_SLCRequestUserSessionStatus                              EMsg = 9401
	EMsg_SLCSharedLicensesLockStatus                              EMsg = 9402
	EMsg_ClientSharedLicensesLockStatus                           EMsg = 9403 // removed
	EMsg_ClientSharedLicensesStopPlaying                          EMsg = 9404 // removed
	EMsg_ClientSharedLibraryLockStatus                            EMsg = 9405
	EMsg_ClientSharedLibraryStopPlaying                           EMsg = 9406
	EMsg_SLCOwnerLibraryChanged                                   EMsg = 9407
	EMsg_SLCSharedLibraryChanged                                  EMsg = 9408
	EMsg_RemoteClientBase                                         EMsg = 9500
	EMsg_RemoteClientAuth                                         EMsg = 9500
	EMsg_RemoteClientAuthResponse                                 EMsg = 9501
	EMsg_RemoteClientAppStatus                                    EMsg = 9502
	EMsg_RemoteClientStartStream                                  EMsg = 9503
	EMsg_RemoteClientStartStreamResponse                          EMsg = 9504
	EMsg_RemoteClientPing                                         EMsg = 9505
	EMsg_RemoteClientPingResponse                                 EMsg = 9506
	EMsg_ClientUnlockStreaming                                    EMsg = 9507
	EMsg_ClientUnlockStreamingResponse                            EMsg = 9508
	EMsg_RemoteClientAcceptEULA                                   EMsg = 9509
	EMsg_RemoteClientGetControllerConfig                          EMsg = 9510
	EMsg_RemoteClientGetControllerConfigResposne                  EMsg = 9511
	EMsg_RemoteClientStreamingEnabled                             EMsg = 9512
	EMsg_ClientUnlockHEVC                                         EMsg = 9513
	EMsg_ClientUnlockHEVCResponse                                 EMsg = 9514
	EMsg_ClientConcurrentSessionsBase                             EMsg = 9600
	EMsg_ClientPlayingSessionState                                EMsg = 9600
	EMsg_ClientKickPlayingSession                                 EMsg = 9601
	EMsg_ClientBroadcastBase                                      EMsg = 9700
	EMsg_ClientBroadcastInit                                      EMsg = 9700
	EMsg_ClientBroadcastFrames                                    EMsg = 9701
	EMsg_ClientBroadcastDisconnect                                EMsg = 9702
	EMsg_ClientBroadcastScreenshot                                EMsg = 9703
	EMsg_ClientBroadcastUploadConfig                              EMsg = 9704
	EMsg_BaseClient3                                              EMsg = 9800
	EMsg_ClientVoiceCallPreAuthorize                              EMsg = 9800
	EMsg_ClientVoiceCallPreAuthorizeResponse                      EMsg = 9801
	EMsg_ClientServerTimestampRequest                             EMsg = 9802
)

var EMsg_name = map[EMsg]string{
	EMsg_Invalid:                                                  "EMsg_Invalid",
	EMsg_Multi:                                                    "EMsg_Multi",
	EMsg_ProtobufWrapped:                                          "EMsg_ProtobufWrapped",
	EMsg_BaseGeneral:                                              "EMsg_BaseGeneral",
	EMsg_DestJobFailed:                                            "EMsg_DestJobFailed",
	EMsg_Alert:                                                    "EMsg_Alert",
	EMsg_SCIDRequest:                                              "EMsg_SCIDRequest",
	EMsg_SCIDResponse:                                             "EMsg_SCIDResponse",
	EMsg_JobHeartbeat:                                             "EMsg_JobHeartbeat",
	EMsg_HubConnect:                                               "EMsg_HubConnect",
	EMsg_Subscribe:                                                "EMsg_Subscribe",
	EMsg_RouteMessage:                                             "EMsg_RouteMessage",
	EMsg_RemoteSysID:                                              "EMsg_RemoteSysID",
	EMsg_AMCreateAccountResponse:                                  "EMsg_AMCreateAccountResponse",
	EMsg_WGRequest:                                                "EMsg_WGRequest",
	EMsg_WGResponse:                                               "EMsg_WGResponse",
	EMsg_KeepAlive:                                                "EMsg_KeepAlive",
	EMsg_WebAPIJobRequest:                                         "EMsg_WebAPIJobRequest",
	EMsg_WebAPIJobResponse:                                        "EMsg_WebAPIJobResponse",
	EMsg_ClientSessionStart:                                       "EMsg_ClientSessionStart",
	EMsg_ClientSessionEnd:                                         "EMsg_ClientSessionEnd",
	EMsg_ClientSessionUpdate:                                      "EMsg_ClientSessionUpdate",
	EMsg_StatsDeprecated:                                          "EMsg_StatsDeprecated",
	EMsg_Ping:                                                     "EMsg_Ping",
	EMsg_PingResponse:                                             "EMsg_PingResponse",
	EMsg_Stats:                                                    "EMsg_Stats",
	EMsg_RequestFullStatsBlock:                                    "EMsg_RequestFullStatsBlock",
	EMsg_LoadDBOCacheItem:                                         "EMsg_LoadDBOCacheItem",
	EMsg_LoadDBOCacheItemResponse:                                 "EMsg_LoadDBOCacheItemResponse",
	EMsg_InvalidateDBOCacheItems:                                  "EMsg_InvalidateDBOCacheItems",
	EMsg_ServiceMethod:                                            "EMsg_ServiceMethod",
	EMsg_ServiceMethodResponse:                                    "EMsg_ServiceMethodResponse",
	EMsg_ClientPackageVersions:                                    "EMsg_ClientPackageVersions",
	EMsg_TimestampRequest:                                         "EMsg_TimestampRequest",
	EMsg_TimestampResponse:                                        "EMsg_TimestampResponse",
	EMsg_ServiceMethodCallFromClient:                              "EMsg_ServiceMethodCallFromClient",
	EMsg_ServiceMethodSendToClient:                                "EMsg_ServiceMethodSendToClient",
	EMsg_BaseShell:                                                "EMsg_BaseShell",
	EMsg_Exit:                                                     "EMsg_Exit",
	EMsg_DirRequest:                                               "EMsg_DirRequest",
	EMsg_DirResponse:                                              "EMsg_DirResponse",
	EMsg_ZipRequest:                                               "EMsg_ZipRequest",
	EMsg_ZipResponse:                                              "EMsg_ZipResponse",
	EMsg_UpdateRecordResponse:                                     "EMsg_UpdateRecordResponse",
	EMsg_UpdateCreditCardRequest:                                  "EMsg_UpdateCreditCardRequest",
	EMsg_UpdateUserBanResponse:                                    "EMsg_UpdateUserBanResponse",
	EMsg_PrepareToExit:                                            "EMsg_PrepareToExit",
	EMsg_ContentDescriptionUpdate:                                 "EMsg_ContentDescriptionUpdate",
	EMsg_TestResetServer:                                          "EMsg_TestResetServer",
	EMsg_UniverseChanged:                                          "EMsg_UniverseChanged",
	EMsg_ShellConfigInfoUpdate:                                    "EMsg_ShellConfigInfoUpdate",
	EMsg_RequestWindowsEventLogEntries:                            "EMsg_RequestWindowsEventLogEntries",
	EMsg_ProvideWindowsEventLogEntries:                            "EMsg_ProvideWindowsEventLogEntries",
	EMsg_ShellSearchLogs:                                          "EMsg_ShellSearchLogs",
	EMsg_ShellSearchLogsResponse:                                  "EMsg_ShellSearchLogsResponse",
	EMsg_ShellCheckWindowsUpdates:                                 "EMsg_ShellCheckWindowsUpdates",
	EMsg_ShellCheckWindowsUpdatesResponse:                         "EMsg_ShellCheckWindowsUpdatesResponse",
	EMsg_ShellFlushUserLicenseCache:                               "EMsg_ShellFlushUserLicenseCache",
	EMsg_BaseGM:                                                   "EMsg_BaseGM",
	EMsg_ShellFailed:                                              "EMsg_ShellFailed",
	EMsg_ExitShells:                                               "EMsg_ExitShells",
	EMsg_ExitShell:                                                "EMsg_ExitShell",
	EMsg_GracefulExitShell:                                        "EMsg_GracefulExitShell",
	EMsg_NotifyWatchdog:                                           "EMsg_NotifyWatchdog",
	EMsg_LicenseProcessingComplete:                                "EMsg_LicenseProcessingComplete",
	EMsg_SetTestFlag:                                              "EMsg_SetTestFlag",
	EMsg_QueuedEmailsComplete:                                     "EMsg_QueuedEmailsComplete",
	EMsg_GMReportPHPError:                                         "EMsg_GMReportPHPError",
	EMsg_GMDRMSync:                                                "EMsg_GMDRMSync",
	EMsg_PhysicalBoxInventory:                                     "EMsg_PhysicalBoxInventory",
	EMsg_UpdateConfigFile:                                         "EMsg_UpdateConfigFile",
	EMsg_TestInitDB:                                               "EMsg_TestInitDB",
	EMsg_GMWriteConfigToSQL:                                       "EMsg_GMWriteConfigToSQL",
	EMsg_GMLoadActivationCodes:                                    "EMsg_GMLoadActivationCodes",
	EMsg_GMQueueForFBS:                                            "EMsg_GMQueueForFBS",
	EMsg_GMSchemaConversionResults:                                "EMsg_GMSchemaConversionResults",
	EMsg_GMSchemaConversionResultsResponse:                        "EMsg_GMSchemaConversionResultsResponse",
	EMsg_GMWriteShellFailureToSQL:                                 "EMsg_GMWriteShellFailureToSQL",
	EMsg_GMWriteStatsToSOS:                                        "EMsg_GMWriteStatsToSOS",
	EMsg_GMGetServiceMethodRouting:                                "EMsg_GMGetServiceMethodRouting",
	EMsg_GMGetServiceMethodRoutingResponse:                        "EMsg_GMGetServiceMethodRoutingResponse",
	EMsg_GMConvertUserWallets:                                     "EMsg_GMConvertUserWallets",
	EMsg_BaseAIS:                                                  "EMsg_BaseAIS",
	EMsg_AISRefreshContentDescription:                             "EMsg_AISRefreshContentDescription",
	EMsg_AISRequestContentDescription:                             "EMsg_AISRequestContentDescription",
	EMsg_AISUpdateAppInfo:                                         "EMsg_AISUpdateAppInfo",
	EMsg_AISUpdatePackageCosts:                                    "EMsg_AISUpdatePackageCosts",
	EMsg_AISGetPackageChangeNumber:                                "EMsg_AISGetPackageChangeNumber",
	EMsg_AISGetPackageChangeNumberResponse:                        "EMsg_AISGetPackageChangeNumberResponse",
	EMsg_AISAppInfoTableChanged:                                   "EMsg_AISAppInfoTableChanged",
	EMsg_AISUpdatePackageCostsResponse:                            "EMsg_AISUpdatePackageCostsResponse",
	EMsg_AISCreateMarketingMessage:                                "EMsg_AISCreateMarketingMessage",
	EMsg_AISCreateMarketingMessageResponse:                        "EMsg_AISCreateMarketingMessageResponse",
	EMsg_AISGetMarketingMessage:                                   "EMsg_AISGetMarketingMessage",
	EMsg_AISGetMarketingMessageResponse:                           "EMsg_AISGetMarketingMessageResponse",
	EMsg_AISUpdateMarketingMessage:                                "EMsg_AISUpdateMarketingMessage",
	EMsg_AISUpdateMarketingMessageResponse:                        "EMsg_AISUpdateMarketingMessageResponse",
	EMsg_AISRequestMarketingMessageUpdate:                         "EMsg_AISRequestMarketingMessageUpdate",
	EMsg_AISDeleteMarketingMessage:                                "EMsg_AISDeleteMarketingMessage",
	EMsg_AISGetMarketingTreatments:                                "EMsg_AISGetMarketingTreatments",
	EMsg_AISGetMarketingTreatmentsResponse:                        "EMsg_AISGetMarketingTreatmentsResponse",
	EMsg_AISRequestMarketingTreatmentUpdate:                       "EMsg_AISRequestMarketingTreatmentUpdate",
	EMsg_AISTestAddPackage:                                        "EMsg_AISTestAddPackage",
	EMsg_AIGetAppGCFlags:                                          "EMsg_AIGetAppGCFlags",
	EMsg_AIGetAppGCFlagsResponse:                                  "EMsg_AIGetAppGCFlagsResponse",
	EMsg_AIGetAppList:                                             "EMsg_AIGetAppList",
	EMsg_AIGetAppListResponse:                                     "EMsg_AIGetAppListResponse",
	EMsg_AIGetAppInfo:                                             "EMsg_AIGetAppInfo",
	EMsg_AIGetAppInfoResponse:                                     "EMsg_AIGetAppInfoResponse",
	EMsg_AISGetCouponDefinition:                                   "EMsg_AISGetCouponDefinition",
	EMsg_AISGetCouponDefinitionResponse:                           "EMsg_AISGetCouponDefinitionResponse",
	EMsg_AISUpdateSlaveContentDescription:                         "EMsg_AISUpdateSlaveContentDescription",
	EMsg_AISUpdateSlaveContentDescriptionResponse:                 "EMsg_AISUpdateSlaveContentDescriptionResponse",
	EMsg_AISTestEnableGC:                                          "EMsg_AISTestEnableGC",
	EMsg_BaseAM:                                                   "EMsg_BaseAM",
	EMsg_AMUpdateUserBanRequest:                                   "EMsg_AMUpdateUserBanRequest",
	EMsg_AMAddLicense:                                             "EMsg_AMAddLicense",
	EMsg_AMBeginProcessingLicenses:                                "EMsg_AMBeginProcessingLicenses",
	EMsg_AMSendSystemIMToUser:                                     "EMsg_AMSendSystemIMToUser",
	EMsg_AMExtendLicense:                                          "EMsg_AMExtendLicense",
	EMsg_AMAddMinutesToLicense:                                    "EMsg_AMAddMinutesToLicense",
	EMsg_AMCancelLicense:                                          "EMsg_AMCancelLicense",
	EMsg_AMInitPurchase:                                           "EMsg_AMInitPurchase",
	EMsg_AMPurchaseResponse:                                       "EMsg_AMPurchaseResponse",
	EMsg_AMGetFinalPrice:                                          "EMsg_AMGetFinalPrice",
	EMsg_AMGetFinalPriceResponse:                                  "EMsg_AMGetFinalPriceResponse",
	EMsg_AMGetLegacyGameKey:                                       "EMsg_AMGetLegacyGameKey",
	EMsg_AMGetLegacyGameKeyResponse:                               "EMsg_AMGetLegacyGameKeyResponse",
	EMsg_AMFindHungTransactions:                                   "EMsg_AMFindHungTransactions",
	EMsg_AMSetAccountTrustedRequest:                               "EMsg_AMSetAccountTrustedRequest",
	EMsg_AMCompletePurchase:                                       "EMsg_AMCompletePurchase",
	EMsg_AMCancelPurchase:                                         "EMsg_AMCancelPurchase",
	EMsg_AMNewChallenge:                                           "EMsg_AMNewChallenge",
	EMsg_AMLoadOEMTickets:                                         "EMsg_AMLoadOEMTickets",
	EMsg_AMFixPendingPurchase:                                     "EMsg_AMFixPendingPurchase",
	EMsg_AMFixPendingPurchaseResponse:                             "EMsg_AMFixPendingPurchaseResponse",
	EMsg_AMIsUserBanned:                                           "EMsg_AMIsUserBanned",
	EMsg_AMRegisterKey:                                            "EMsg_AMRegisterKey",
	EMsg_AMLoadActivationCodes:                                    "EMsg_AMLoadActivationCodes",
	EMsg_AMLoadActivationCodesResponse:                            "EMsg_AMLoadActivationCodesResponse",
	EMsg_AMLookupKeyResponse:                                      "EMsg_AMLookupKeyResponse",
	EMsg_AMLookupKey:                                              "EMsg_AMLookupKey",
	EMsg_AMChatCleanup:                                            "EMsg_AMChatCleanup",
	EMsg_AMClanCleanup:                                            "EMsg_AMClanCleanup",
	EMsg_AMFixPendingRefund:                                       "EMsg_AMFixPendingRefund",
	EMsg_AMReverseChargeback:                                      "EMsg_AMReverseChargeback",
	EMsg_AMReverseChargebackResponse:                              "EMsg_AMReverseChargebackResponse",
	EMsg_AMClanCleanupList:                                        "EMsg_AMClanCleanupList",
	EMsg_AMGetLicenses:                                            "EMsg_AMGetLicenses",
	EMsg_AMGetLicensesResponse:                                    "EMsg_AMGetLicensesResponse",
	EMsg_AllowUserToPlayQuery:                                     "EMsg_AllowUserToPlayQuery",
	EMsg_AllowUserToPlayResponse:                                  "EMsg_AllowUserToPlayResponse",
	EMsg_AMVerfiyUser:                                             "EMsg_AMVerfiyUser",
	EMsg_AMClientNotPlaying:                                       "EMsg_AMClientNotPlaying",
	EMsg_ClientRequestFriendship:                                  "EMsg_ClientRequestFriendship",
	EMsg_AMRelayPublishStatus:                                     "EMsg_AMRelayPublishStatus",
	EMsg_AMResetCommunityContent:                                  "EMsg_AMResetCommunityContent",
	EMsg_AMPrimePersonaStateCache:                                 "EMsg_AMPrimePersonaStateCache",
	EMsg_AMAllowUserContentQuery:                                  "EMsg_AMAllowUserContentQuery",
	EMsg_AMAllowUserContentResponse:                               "EMsg_AMAllowUserContentResponse",
	EMsg_AMInitPurchaseResponse:                                   "EMsg_AMInitPurchaseResponse",
	EMsg_AMRevokePurchaseResponse:                                 "EMsg_AMRevokePurchaseResponse",
	EMsg_AMLockProfile:                                            "EMsg_AMLockProfile",
	EMsg_AMRefreshGuestPasses:                                     "EMsg_AMRefreshGuestPasses",
	EMsg_AMInviteUserToClan:                                       "EMsg_AMInviteUserToClan",
	EMsg_AMAcknowledgeClanInvite:                                  "EMsg_AMAcknowledgeClanInvite",
	EMsg_AMGrantGuestPasses:                                       "EMsg_AMGrantGuestPasses",
	EMsg_AMClanDataUpdated:                                        "EMsg_AMClanDataUpdated",
	EMsg_AMReloadAccount:                                          "EMsg_AMReloadAccount",
	EMsg_AMClientChatMsgRelay:                                     "EMsg_AMClientChatMsgRelay",
	EMsg_AMChatMulti:                                              "EMsg_AMChatMulti",
	EMsg_AMClientChatInviteRelay:                                  "EMsg_AMClientChatInviteRelay",
	EMsg_AMChatInvite:                                             "EMsg_AMChatInvite",
	EMsg_AMClientJoinChatRelay:                                    "EMsg_AMClientJoinChatRelay",
	EMsg_AMClientChatMemberInfoRelay:                              "EMsg_AMClientChatMemberInfoRelay",
	EMsg_AMPublishChatMemberInfo:                                  "EMsg_AMPublishChatMemberInfo",
	EMsg_AMClientAcceptFriendInvite:                               "EMsg_AMClientAcceptFriendInvite",
	EMsg_AMChatEnter:                                              "EMsg_AMChatEnter",
	EMsg_AMClientPublishRemovalFromSource:                         "EMsg_AMClientPublishRemovalFromSource",
	EMsg_AMChatActionResult:                                       "EMsg_AMChatActionResult",
	EMsg_AMFindAccounts:                                           "EMsg_AMFindAccounts",
	EMsg_AMFindAccountsResponse:                                   "EMsg_AMFindAccountsResponse",
	EMsg_AMRequestAccountData:                                     "EMsg_AMRequestAccountData",
	EMsg_AMRequestAccountDataResponse:                             "EMsg_AMRequestAccountDataResponse",
	EMsg_AMSetAccountFlags:                                        "EMsg_AMSetAccountFlags",
	EMsg_AMCreateClan:                                             "EMsg_AMCreateClan",
	EMsg_AMCreateClanResponse:                                     "EMsg_AMCreateClanResponse",
	EMsg_AMGetClanDetails:                                         "EMsg_AMGetClanDetails",
	EMsg_AMGetClanDetailsResponse:                                 "EMsg_AMGetClanDetailsResponse",
	EMsg_AMSetPersonaName:                                         "EMsg_AMSetPersonaName",
	EMsg_AMSetAvatar:                                              "EMsg_AMSetAvatar",
	EMsg_AMAuthenticateUser:                                       "EMsg_AMAuthenticateUser",
	EMsg_AMAuthenticateUserResponse:                               "EMsg_AMAuthenticateUserResponse",
	EMsg_AMGetAccountFriendsCount:                                 "EMsg_AMGetAccountFriendsCount",
	EMsg_AMGetAccountFriendsCountResponse:                         "EMsg_AMGetAccountFriendsCountResponse",
	EMsg_AMP2PIntroducerMessage:                                   "EMsg_AMP2PIntroducerMessage",
	EMsg_ClientChatAction:                                         "EMsg_ClientChatAction",
	EMsg_AMClientChatActionRelay:                                  "EMsg_AMClientChatActionRelay",
	EMsg_BaseVS:                                                   "EMsg_BaseVS",
	EMsg_VACResponse:                                              "EMsg_VACResponse",
	EMsg_ReqChallengeTest:                                         "EMsg_ReqChallengeTest",
	EMsg_VSMarkCheat:                                              "EMsg_VSMarkCheat",
	EMsg_VSAddCheat:                                               "EMsg_VSAddCheat",
	EMsg_VSPurgeCodeModDB:                                         "EMsg_VSPurgeCodeModDB",
	EMsg_VSGetChallengeResults:                                    "EMsg_VSGetChallengeResults",
	EMsg_VSChallengeResultText:                                    "EMsg_VSChallengeResultText",
	EMsg_VSReportLingerer:                                         "EMsg_VSReportLingerer",
	EMsg_VSRequestManagedChallenge:                                "EMsg_VSRequestManagedChallenge",
	EMsg_VSLoadDBFinished:                                         "EMsg_VSLoadDBFinished",
	EMsg_BaseDRMS:                                                 "EMsg_BaseDRMS",
	EMsg_DRMBuildBlobRequest:                                      "EMsg_DRMBuildBlobRequest",
	EMsg_DRMBuildBlobResponse:                                     "EMsg_DRMBuildBlobResponse",
	EMsg_DRMResolveGuidRequest:                                    "EMsg_DRMResolveGuidRequest",
	EMsg_DRMResolveGuidResponse:                                   "EMsg_DRMResolveGuidResponse",
	EMsg_DRMVariabilityReport:                                     "EMsg_DRMVariabilityReport",
	EMsg_DRMVariabilityReportResponse:                             "EMsg_DRMVariabilityReportResponse",
	EMsg_DRMStabilityReport:                                       "EMsg_DRMStabilityReport",
	EMsg_DRMStabilityReportResponse:                               "EMsg_DRMStabilityReportResponse",
	EMsg_DRMDetailsReportRequest:                                  "EMsg_DRMDetailsReportRequest",
	EMsg_DRMDetailsReportResponse:                                 "EMsg_DRMDetailsReportResponse",
	EMsg_DRMProcessFile:                                           "EMsg_DRMProcessFile",
	EMsg_DRMAdminUpdate:                                           "EMsg_DRMAdminUpdate",
	EMsg_DRMAdminUpdateResponse:                                   "EMsg_DRMAdminUpdateResponse",
	EMsg_DRMSync:                                                  "EMsg_DRMSync",
	EMsg_DRMSyncResponse:                                          "EMsg_DRMSyncResponse",
	EMsg_DRMProcessFileResponse:                                   "EMsg_DRMProcessFileResponse",
	EMsg_DRMEmptyGuidCache:                                        "EMsg_DRMEmptyGuidCache",
	EMsg_DRMEmptyGuidCacheResponse:                                "EMsg_DRMEmptyGuidCacheResponse",
	EMsg_BaseCS:                                                   "EMsg_BaseCS",
	EMsg_CSUserContentRequest:                                     "EMsg_CSUserContentRequest",
	EMsg_BaseClient:                                               "EMsg_BaseClient",
	EMsg_ClientLogOn_Deprecated:                                   "EMsg_ClientLogOn_Deprecated",
	EMsg_ClientAnonLogOn_Deprecated:                               "EMsg_ClientAnonLogOn_Deprecated",
	EMsg_ClientHeartBeat:                                          "EMsg_ClientHeartBeat",
	EMsg_ClientVACResponse:                                        "EMsg_ClientVACResponse",
	EMsg_ClientGamesPlayed_obsolete:                               "EMsg_ClientGamesPlayed_obsolete",
	EMsg_ClientLogOff:                                             "EMsg_ClientLogOff",
	EMsg_ClientNoUDPConnectivity:                                  "EMsg_ClientNoUDPConnectivity",
	EMsg_ClientInformOfCreateAccount:                              "EMsg_ClientInformOfCreateAccount",
	EMsg_ClientAckVACBan:                                          "EMsg_ClientAckVACBan",
	EMsg_ClientConnectionStats:                                    "EMsg_ClientConnectionStats",
	EMsg_ClientInitPurchase:                                       "EMsg_ClientInitPurchase",
	EMsg_ClientPingResponse:                                       "EMsg_ClientPingResponse",
	EMsg_ClientRemoveFriend:                                       "EMsg_ClientRemoveFriend",
	EMsg_ClientGamesPlayedNoDataBlob:                              "EMsg_ClientGamesPlayedNoDataBlob",
	EMsg_ClientChangeStatus:                                       "EMsg_ClientChangeStatus",
	EMsg_ClientVacStatusResponse:                                  "EMsg_ClientVacStatusResponse",
	EMsg_ClientFriendMsg:                                          "EMsg_ClientFriendMsg",
	EMsg_ClientGameConnect_obsolete:                               "EMsg_ClientGameConnect_obsolete",
	EMsg_ClientGamesPlayed2_obsolete:                              "EMsg_ClientGamesPlayed2_obsolete",
	EMsg_ClientGameEnded_obsolete:                                 "EMsg_ClientGameEnded_obsolete",
	EMsg_ClientGetFinalPrice:                                      "EMsg_ClientGetFinalPrice",
	EMsg_ClientSystemIM:                                           "EMsg_ClientSystemIM",
	EMsg_ClientSystemIMAck:                                        "EMsg_ClientSystemIMAck",
	EMsg_ClientGetLicenses:                                        "EMsg_ClientGetLicenses",
	EMsg_ClientCancelLicense:                                      "EMsg_ClientCancelLicense",
	EMsg_ClientGetLegacyGameKey:                                   "EMsg_ClientGetLegacyGameKey",
	EMsg_ClientContentServerLogOn_Deprecated:                      "EMsg_ClientContentServerLogOn_Deprecated",
	EMsg_ClientAckVACBan2:                                         "EMsg_ClientAckVACBan2",
	EMsg_ClientAckMessageByGID:                                    "EMsg_ClientAckMessageByGID",
	EMsg_ClientGetPurchaseReceipts:                                "EMsg_ClientGetPurchaseReceipts",
	EMsg_ClientAckPurchaseReceipt:                                 "EMsg_ClientAckPurchaseReceipt",
	EMsg_ClientGamesPlayed3_obsolete:                              "EMsg_ClientGamesPlayed3_obsolete",
	EMsg_ClientSendGuestPass:                                      "EMsg_ClientSendGuestPass",
	EMsg_ClientAckGuestPass:                                       "EMsg_ClientAckGuestPass",
	EMsg_ClientRedeemGuestPass:                                    "EMsg_ClientRedeemGuestPass",
	EMsg_ClientGamesPlayed:                                        "EMsg_ClientGamesPlayed",
	EMsg_ClientRegisterKey:                                        "EMsg_ClientRegisterKey",
	EMsg_ClientInviteUserToClan:                                   "EMsg_ClientInviteUserToClan",
	EMsg_ClientAcknowledgeClanInvite:                              "EMsg_ClientAcknowledgeClanInvite",
	EMsg_ClientPurchaseWithMachineID:                              "EMsg_ClientPurchaseWithMachineID",
	EMsg_ClientAppUsageEvent:                                      "EMsg_ClientAppUsageEvent",
	EMsg_ClientGetGiftTargetList:                                  "EMsg_ClientGetGiftTargetList",
	EMsg_ClientGetGiftTargetListResponse:                          "EMsg_ClientGetGiftTargetListResponse",
	EMsg_ClientLogOnResponse:                                      "EMsg_ClientLogOnResponse",
	EMsg_ClientVACChallenge:                                       "EMsg_ClientVACChallenge",
	EMsg_ClientSetHeartbeatRate:                                   "EMsg_ClientSetHeartbeatRate",
	EMsg_ClientNotLoggedOnDeprecated:                              "EMsg_ClientNotLoggedOnDeprecated",
	EMsg_ClientLoggedOff:                                          "EMsg_ClientLoggedOff",
	EMsg_GSApprove:                                                "EMsg_GSApprove",
	EMsg_GSDeny:                                                   "EMsg_GSDeny",
	EMsg_GSKick:                                                   "EMsg_GSKick",
	EMsg_ClientCreateAcctResponse:                                 "EMsg_ClientCreateAcctResponse",
	EMsg_ClientPurchaseResponse:                                   "EMsg_ClientPurchaseResponse",
	EMsg_ClientPing:                                               "EMsg_ClientPing",
	EMsg_ClientNOP:                                                "EMsg_ClientNOP",
	EMsg_ClientPersonaState:                                       "EMsg_ClientPersonaState",
	EMsg_ClientFriendsList:                                        "EMsg_ClientFriendsList",
	EMsg_ClientAccountInfo:                                        "EMsg_ClientAccountInfo",
	EMsg_ClientVacStatusQuery:                                     "EMsg_ClientVacStatusQuery",
	EMsg_ClientNewsUpdate:                                         "EMsg_ClientNewsUpdate",
	EMsg_ClientGameConnectDeny:                                    "EMsg_ClientGameConnectDeny",
	EMsg_GSStatusReply:                                            "EMsg_GSStatusReply",
	EMsg_ClientGetFinalPriceResponse:                              "EMsg_ClientGetFinalPriceResponse",
	EMsg_ClientGameConnectTokens:                                  "EMsg_ClientGameConnectTokens",
	EMsg_ClientLicenseList:                                        "EMsg_ClientLicenseList",
	EMsg_ClientCancelLicenseResponse:                              "EMsg_ClientCancelLicenseResponse",
	EMsg_ClientVACBanStatus:                                       "EMsg_ClientVACBanStatus",
	EMsg_ClientCMList:                                             "EMsg_ClientCMList",
	EMsg_ClientEncryptPct:                                         "EMsg_ClientEncryptPct",
	EMsg_ClientGetLegacyGameKeyResponse:                           "EMsg_ClientGetLegacyGameKeyResponse",
	EMsg_ClientFavoritesList:                                      "EMsg_ClientFavoritesList",
	EMsg_CSUserContentApprove:                                     "EMsg_CSUserContentApprove",
	EMsg_CSUserContentDeny:                                        "EMsg_CSUserContentDeny",
	EMsg_ClientInitPurchaseResponse:                               "EMsg_ClientInitPurchaseResponse",
	EMsg_ClientAddFriend:                                          "EMsg_ClientAddFriend",
	EMsg_ClientAddFriendResponse:                                  "EMsg_ClientAddFriendResponse",
	EMsg_ClientInviteFriend:                                       "EMsg_ClientInviteFriend",
	EMsg_ClientInviteFriendResponse:                               "EMsg_ClientInviteFriendResponse",
	EMsg_ClientSendGuestPassResponse:                              "EMsg_ClientSendGuestPassResponse",
	EMsg_ClientAckGuestPassResponse:                               "EMsg_ClientAckGuestPassResponse",
	EMsg_ClientRedeemGuestPassResponse:                            "EMsg_ClientRedeemGuestPassResponse",
	EMsg_ClientUpdateGuestPassesList:                              "EMsg_ClientUpdateGuestPassesList",
	EMsg_ClientChatMsg:                                            "EMsg_ClientChatMsg",
	EMsg_ClientChatInvite:                                         "EMsg_ClientChatInvite",
	EMsg_ClientJoinChat:                                           "EMsg_ClientJoinChat",
	EMsg_ClientChatMemberInfo:                                     "EMsg_ClientChatMemberInfo",
	EMsg_ClientLogOnWithCredentials_Deprecated:                    "EMsg_ClientLogOnWithCredentials_Deprecated",
	EMsg_ClientPasswordChangeResponse:                             "EMsg_ClientPasswordChangeResponse",
	EMsg_ClientChatEnter:                                          "EMsg_ClientChatEnter",
	EMsg_ClientFriendRemovedFromSource:                            "EMsg_ClientFriendRemovedFromSource",
	EMsg_ClientCreateChat:                                         "EMsg_ClientCreateChat",
	EMsg_ClientCreateChatResponse:                                 "EMsg_ClientCreateChatResponse",
	EMsg_ClientUpdateChatMetadata:                                 "EMsg_ClientUpdateChatMetadata",
	EMsg_ClientP2PIntroducerMessage:                               "EMsg_ClientP2PIntroducerMessage",
	EMsg_ClientChatActionResult:                                   "EMsg_ClientChatActionResult",
	EMsg_ClientRequestFriendData:                                  "EMsg_ClientRequestFriendData",
	EMsg_ClientGetUserStats:                                       "EMsg_ClientGetUserStats",
	EMsg_ClientGetUserStatsResponse:                               "EMsg_ClientGetUserStatsResponse",
	EMsg_ClientStoreUserStats:                                     "EMsg_ClientStoreUserStats",
	EMsg_ClientStoreUserStatsResponse:                             "EMsg_ClientStoreUserStatsResponse",
	EMsg_ClientClanState:                                          "EMsg_ClientClanState",
	EMsg_ClientServiceModule:                                      "EMsg_ClientServiceModule",
	EMsg_ClientServiceCall:                                        "EMsg_ClientServiceCall",
	EMsg_ClientServiceCallResponse:                                "EMsg_ClientServiceCallResponse",
	EMsg_ClientPackageInfoRequest:                                 "EMsg_ClientPackageInfoRequest",
	EMsg_ClientPackageInfoResponse:                                "EMsg_ClientPackageInfoResponse",
	EMsg_ClientNatTraversalStatEvent:                              "EMsg_ClientNatTraversalStatEvent",
	EMsg_ClientAppInfoRequest:                                     "EMsg_ClientAppInfoRequest",
	EMsg_ClientAppInfoResponse:                                    "EMsg_ClientAppInfoResponse",
	EMsg_ClientSteamUsageEvent:                                    "EMsg_ClientSteamUsageEvent",
	EMsg_ClientCheckPassword:                                      "EMsg_ClientCheckPassword",
	EMsg_ClientResetPassword:                                      "EMsg_ClientResetPassword",
	EMsg_ClientCheckPasswordResponse:                              "EMsg_ClientCheckPasswordResponse",
	EMsg_ClientResetPasswordResponse:                              "EMsg_ClientResetPasswordResponse",
	EMsg_ClientSessionToken:                                       "EMsg_ClientSessionToken",
	EMsg_ClientDRMProblemReport:                                   "EMsg_ClientDRMProblemReport",
	EMsg_ClientSetIgnoreFriend:                                    "EMsg_ClientSetIgnoreFriend",
	EMsg_ClientSetIgnoreFriendResponse:                            "EMsg_ClientSetIgnoreFriendResponse",
	EMsg_ClientGetAppOwnershipTicket:                              "EMsg_ClientGetAppOwnershipTicket",
	EMsg_ClientGetAppOwnershipTicketResponse:                      "EMsg_ClientGetAppOwnershipTicketResponse",
	EMsg_ClientGetLobbyListResponse:                               "EMsg_ClientGetLobbyListResponse",
	EMsg_ClientGetLobbyMetadata:                                   "EMsg_ClientGetLobbyMetadata",
	EMsg_ClientGetLobbyMetadataResponse:                           "EMsg_ClientGetLobbyMetadataResponse",
	EMsg_ClientVTTCert:                                            "EMsg_ClientVTTCert",
	EMsg_ClientAppInfoUpdate:                                      "EMsg_ClientAppInfoUpdate",
	EMsg_ClientAppInfoChanges:                                     "EMsg_ClientAppInfoChanges",
	EMsg_ClientServerList:                                         "EMsg_ClientServerList",
	EMsg_ClientEmailChangeResponse:                                "EMsg_ClientEmailChangeResponse",
	EMsg_ClientSecretQAChangeResponse:                             "EMsg_ClientSecretQAChangeResponse",
	EMsg_ClientDRMBlobRequest:                                     "EMsg_ClientDRMBlobRequest",
	EMsg_ClientDRMBlobResponse:                                    "EMsg_ClientDRMBlobResponse",
	EMsg_ClientLookupKey:                                          "EMsg_ClientLookupKey",
	EMsg_ClientLookupKeyResponse:                                  "EMsg_ClientLookupKeyResponse",
	EMsg_BaseGameServer:                                           "EMsg_BaseGameServer",
	EMsg_GSDisconnectNotice:                                       "EMsg_GSDisconnectNotice",
	EMsg_GSStatus:                                                 "EMsg_GSStatus",
	EMsg_GSUserPlaying:                                            "EMsg_GSUserPlaying",
	EMsg_GSStatus2:                                                "EMsg_GSStatus2",
	EMsg_GSStatusUpdate_Unused:                                    "EMsg_GSStatusUpdate_Unused",
	EMsg_GSServerType:                                             "EMsg_GSServerType",
	EMsg_GSPlayerList:                                             "EMsg_GSPlayerList",
	EMsg_GSGetUserAchievementStatus:                               "EMsg_GSGetUserAchievementStatus",
	EMsg_GSGetUserAchievementStatusResponse:                       "EMsg_GSGetUserAchievementStatusResponse",
	EMsg_GSGetPlayStats:                                           "EMsg_GSGetPlayStats",
	EMsg_GSGetPlayStatsResponse:                                   "EMsg_GSGetPlayStatsResponse",
	EMsg_GSGetUserGroupStatus:                                     "EMsg_GSGetUserGroupStatus",
	EMsg_AMGetUserGroupStatus:                                     "EMsg_AMGetUserGroupStatus",
	EMsg_AMGetUserGroupStatusResponse:                             "EMsg_AMGetUserGroupStatusResponse",
	EMsg_GSGetUserGroupStatusResponse:                             "EMsg_GSGetUserGroupStatusResponse",
	EMsg_GSGetReputation:                                          "EMsg_GSGetReputation",
	EMsg_GSGetReputationResponse:                                  "EMsg_GSGetReputationResponse",
	EMsg_GSAssociateWithClan:                                      "EMsg_GSAssociateWithClan",
	EMsg_GSAssociateWithClanResponse:                              "EMsg_GSAssociateWithClanResponse",
	EMsg_GSComputeNewPlayerCompatibility:                          "EMsg_GSComputeNewPlayerCompatibility",
	EMsg_GSComputeNewPlayerCompatibilityResponse:                  "EMsg_GSComputeNewPlayerCompatibilityResponse",
	EMsg_BaseAdmin:                                                "EMsg_BaseAdmin",
	EMsg_AdminCmdResponse:                                         "EMsg_AdminCmdResponse",
	EMsg_AdminLogListenRequest:                                    "EMsg_AdminLogListenRequest",
	EMsg_AdminLogEvent:                                            "EMsg_AdminLogEvent",
	EMsg_LogSearchRequest:                                         "EMsg_LogSearchRequest",
	EMsg_LogSearchResponse:                                        "EMsg_LogSearchResponse",
	EMsg_LogSearchCancel:                                          "EMsg_LogSearchCancel",
	EMsg_UniverseData:                                             "EMsg_UniverseData",
	EMsg_RequestStatHistory:                                       "EMsg_RequestStatHistory",
	EMsg_StatHistory:                                              "EMsg_StatHistory",
	EMsg_AdminPwLogon:                                             "EMsg_AdminPwLogon",
	EMsg_AdminPwLogonResponse:                                     "EMsg_AdminPwLogonResponse",
	EMsg_AdminSpew:                                                "EMsg_AdminSpew",
	EMsg_AdminConsoleTitle:                                        "EMsg_AdminConsoleTitle",
	EMsg_AdminGCSpew:                                              "EMsg_AdminGCSpew",
	EMsg_AdminGCCommand:                                           "EMsg_AdminGCCommand",
	EMsg_AdminGCGetCommandList:                                    "EMsg_AdminGCGetCommandList",
	EMsg_AdminGCGetCommandListResponse:                            "EMsg_AdminGCGetCommandListResponse",
	EMsg_FBSConnectionData:                                        "EMsg_FBSConnectionData",
	EMsg_AdminMsgSpew:                                             "EMsg_AdminMsgSpew",
	EMsg_BaseFBS:                                                  "EMsg_BaseFBS",
	EMsg_FBSVersionInfo:                                           "EMsg_FBSVersionInfo",
	EMsg_FBSForceRefresh:                                          "EMsg_FBSForceRefresh",
	EMsg_FBSForceBounce:                                           "EMsg_FBSForceBounce",
	EMsg_FBSDeployPackage:                                         "EMsg_FBSDeployPackage",
	EMsg_FBSDeployResponse:                                        "EMsg_FBSDeployResponse",
	EMsg_FBSUpdateBootstrapper:                                    "EMsg_FBSUpdateBootstrapper",
	EMsg_FBSSetState:                                              "EMsg_FBSSetState",
	EMsg_FBSApplyOSUpdates:                                        "EMsg_FBSApplyOSUpdates",
	EMsg_FBSRunCMDScript:                                          "EMsg_FBSRunCMDScript",
	EMsg_FBSRebootBox:                                             "EMsg_FBSRebootBox",
	EMsg_FBSSetBigBrotherMode:                                     "EMsg_FBSSetBigBrotherMode",
	EMsg_FBSMinidumpServer:                                        "EMsg_FBSMinidumpServer",
	EMsg_FBSSetShellCount_obsolete:                                "EMsg_FBSSetShellCount_obsolete",
	EMsg_FBSDeployHotFixPackage:                                   "EMsg_FBSDeployHotFixPackage",
	EMsg_FBSDeployHotFixResponse:                                  "EMsg_FBSDeployHotFixResponse",
	EMsg_FBSDownloadHotFix:                                        "EMsg_FBSDownloadHotFix",
	EMsg_FBSDownloadHotFixResponse:                                "EMsg_FBSDownloadHotFixResponse",
	EMsg_FBSUpdateTargetConfigFile:                                "EMsg_FBSUpdateTargetConfigFile",
	EMsg_FBSApplyAccountCred:                                      "EMsg_FBSApplyAccountCred",
	EMsg_FBSApplyAccountCredResponse:                              "EMsg_FBSApplyAccountCredResponse",
	EMsg_FBSSetShellCount:                                         "EMsg_FBSSetShellCount",
	EMsg_FBSTerminateShell:                                        "EMsg_FBSTerminateShell",
	EMsg_FBSQueryGMForRequest:                                     "EMsg_FBSQueryGMForRequest",
	EMsg_FBSQueryGMResponse:                                       "EMsg_FBSQueryGMResponse",
	EMsg_FBSTerminateZombies:                                      "EMsg_FBSTerminateZombies",
	EMsg_FBSInfoFromBootstrapper:                                  "EMsg_FBSInfoFromBootstrapper",
	EMsg_FBSRebootBoxResponse:                                     "EMsg_FBSRebootBoxResponse",
	EMsg_FBSBootstrapperPackageRequest:                            "EMsg_FBSBootstrapperPackageRequest",
	EMsg_FBSBootstrapperPackageResponse:                           "EMsg_FBSBootstrapperPackageResponse",
	EMsg_FBSBootstrapperGetPackageChunk:                           "EMsg_FBSBootstrapperGetPackageChunk",
	EMsg_FBSBootstrapperGetPackageChunkResponse:                   "EMsg_FBSBootstrapperGetPackageChunkResponse",
	EMsg_FBSBootstrapperPackageTransferProgress:                   "EMsg_FBSBootstrapperPackageTransferProgress",
	EMsg_FBSRestartBootstrapper:                                   "EMsg_FBSRestartBootstrapper",
	EMsg_BaseFileXfer:                                             "EMsg_BaseFileXfer",
	EMsg_FileXferResponse:                                         "EMsg_FileXferResponse",
	EMsg_FileXferData:                                             "EMsg_FileXferData",
	EMsg_FileXferEnd:                                              "EMsg_FileXferEnd",
	EMsg_FileXferDataAck:                                          "EMsg_FileXferDataAck",
	EMsg_BaseChannelAuth:                                          "EMsg_BaseChannelAuth",
	EMsg_ChannelAuthResponse:                                      "EMsg_ChannelAuthResponse",
	EMsg_ChannelAuthResult:                                        "EMsg_ChannelAuthResult",
	EMsg_ChannelEncryptRequest:                                    "EMsg_ChannelEncryptRequest",
	EMsg_ChannelEncryptResponse:                                   "EMsg_ChannelEncryptResponse",
	EMsg_ChannelEncryptResult:                                     "EMsg_ChannelEncryptResult",
	EMsg_BaseBS:                                                   "EMsg_BaseBS",
	EMsg_BSPurchaseStart:                                          "EMsg_BSPurchaseStart",
	EMsg_BSPurchaseResponse:                                       "EMsg_BSPurchaseResponse",
	EMsg_BSSettleNOVA:                                             "EMsg_BSSettleNOVA",
	EMsg_BSSettleComplete:                                         "EMsg_BSSettleComplete",
	EMsg_BSBannedRequest:                                          "EMsg_BSBannedRequest",
	EMsg_BSInitPayPalTxn:                                          "EMsg_BSInitPayPalTxn",
	EMsg_BSInitPayPalTxnResponse:                                  "EMsg_BSInitPayPalTxnResponse",
	EMsg_BSGetPayPalUserInfo:                                      "EMsg_BSGetPayPalUserInfo",
	EMsg_BSGetPayPalUserInfoResponse:                              "EMsg_BSGetPayPalUserInfoResponse",
	EMsg_BSRefundTxn:                                              "EMsg_BSRefundTxn",
	EMsg_BSRefundTxnResponse:                                      "EMsg_BSRefundTxnResponse",
	EMsg_BSGetEvents:                                              "EMsg_BSGetEvents",
	EMsg_BSChaseRFRRequest:                                        "EMsg_BSChaseRFRRequest",
	EMsg_BSPaymentInstrBan:                                        "EMsg_BSPaymentInstrBan",
	EMsg_BSPaymentInstrBanResponse:                                "EMsg_BSPaymentInstrBanResponse",
	EMsg_BSProcessGCReports:                                       "EMsg_BSProcessGCReports",
	EMsg_BSProcessPPReports:                                       "EMsg_BSProcessPPReports",
	EMsg_BSInitGCBankXferTxn:                                      "EMsg_BSInitGCBankXferTxn",
	EMsg_BSInitGCBankXferTxnResponse:                              "EMsg_BSInitGCBankXferTxnResponse",
	EMsg_BSQueryGCBankXferTxn:                                     "EMsg_BSQueryGCBankXferTxn",
	EMsg_BSQueryGCBankXferTxnResponse:                             "EMsg_BSQueryGCBankXferTxnResponse",
	EMsg_BSCommitGCTxn:                                            "EMsg_BSCommitGCTxn",
	EMsg_BSQueryTransactionStatus:                                 "EMsg_BSQueryTransactionStatus",
	EMsg_BSQueryTransactionStatusResponse:                         "EMsg_BSQueryTransactionStatusResponse",
	EMsg_BSQueryCBOrderStatus:                                     "EMsg_BSQueryCBOrderStatus",
	EMsg_BSQueryCBOrderStatusResponse:                             "EMsg_BSQueryCBOrderStatusResponse",
	EMsg_BSRunRedFlagReport:                                       "EMsg_BSRunRedFlagReport",
	EMsg_BSQueryPaymentInstUsage:                                  "EMsg_BSQueryPaymentInstUsage",
	EMsg_BSQueryPaymentInstResponse:                               "EMsg_BSQueryPaymentInstResponse",
	EMsg_BSQueryTxnExtendedInfo:                                   "EMsg_BSQueryTxnExtendedInfo",
	EMsg_BSQueryTxnExtendedInfoResponse:                           "EMsg_BSQueryTxnExtendedInfoResponse",
	EMsg_BSUpdateConversionRates:                                  "EMsg_BSUpdateConversionRates",
	EMsg_BSProcessUSBankReports:                                   "EMsg_BSProcessUSBankReports",
	EMsg_BSPurchaseRunFraudChecks:                                 "EMsg_BSPurchaseRunFraudChecks",
	EMsg_BSPurchaseRunFraudChecksResponse:                         "EMsg_BSPurchaseRunFraudChecksResponse",
	EMsg_BSStartShippingJobs:                                      "EMsg_BSStartShippingJobs",
	EMsg_BSQueryBankInformation:                                   "EMsg_BSQueryBankInformation",
	EMsg_BSQueryBankInformationResponse:                           "EMsg_BSQueryBankInformationResponse",
	EMsg_BSValidateXsollaSignature:                                "EMsg_BSValidateXsollaSignature",
	EMsg_BSValidateXsollaSignatureResponse:                        "EMsg_BSValidateXsollaSignatureResponse",
	EMsg_BSQiwiWalletInvoice:                                      "EMsg_BSQiwiWalletInvoice",
	EMsg_BSQiwiWalletInvoiceResponse:                              "EMsg_BSQiwiWalletInvoiceResponse",
	EMsg_BSUpdateInventoryFromProPack:                             "EMsg_BSUpdateInventoryFromProPack",
	EMsg_BSUpdateInventoryFromProPackResponse:                     "EMsg_BSUpdateInventoryFromProPackResponse",
	EMsg_BSSendShippingRequest:                                    "EMsg_BSSendShippingRequest",
	EMsg_BSSendShippingRequestResponse:                            "EMsg_BSSendShippingRequestResponse",
	EMsg_BSGetProPackOrderStatus:                                  "EMsg_BSGetProPackOrderStatus",
	EMsg_BSGetProPackOrderStatusResponse:                          "EMsg_BSGetProPackOrderStatusResponse",
	EMsg_BSCheckJobRunning:                                        "EMsg_BSCheckJobRunning",
	EMsg_BSCheckJobRunningResponse:                                "EMsg_BSCheckJobRunningResponse",
	EMsg_BSResetPackagePurchaseRateLimit:                          "EMsg_BSResetPackagePurchaseRateLimit",
	EMsg_BSResetPackagePurchaseRateLimitResponse:                  "EMsg_BSResetPackagePurchaseRateLimitResponse",
	EMsg_BSUpdatePaymentData:                                      "EMsg_BSUpdatePaymentData",
	EMsg_BSUpdatePaymentDataResponse:                              "EMsg_BSUpdatePaymentDataResponse",
	EMsg_BSGetBillingAddress:                                      "EMsg_BSGetBillingAddress",
	EMsg_BSGetBillingAddressResponse:                              "EMsg_BSGetBillingAddressResponse",
	EMsg_BSGetCreditCardInfo:                                      "EMsg_BSGetCreditCardInfo",
	EMsg_BSGetCreditCardInfoResponse:                              "EMsg_BSGetCreditCardInfoResponse",
	EMsg_BSRemoveExpiredPaymentData:                               "EMsg_BSRemoveExpiredPaymentData",
	EMsg_BSRemoveExpiredPaymentDataResponse:                       "EMsg_BSRemoveExpiredPaymentDataResponse",
	EMsg_BSConvertToCurrentKeys:                                   "EMsg_BSConvertToCurrentKeys",
	EMsg_BSConvertToCurrentKeysResponse:                           "EMsg_BSConvertToCurrentKeysResponse",
	EMsg_BSInitPurchase:                                           "EMsg_BSInitPurchase",
	EMsg_BSInitPurchaseResponse:                                   "EMsg_BSInitPurchaseResponse",
	EMsg_BSCompletePurchase:                                       "EMsg_BSCompletePurchase",
	EMsg_BSCompletePurchaseResponse:                               "EMsg_BSCompletePurchaseResponse",
	EMsg_BSPruneCardUsageStats:                                    "EMsg_BSPruneCardUsageStats",
	EMsg_BSPruneCardUsageStatsResponse:                            "EMsg_BSPruneCardUsageStatsResponse",
	EMsg_BSStoreBankInformation:                                   "EMsg_BSStoreBankInformation",
	EMsg_BSStoreBankInformationResponse:                           "EMsg_BSStoreBankInformationResponse",
	EMsg_BSVerifyPOSAKey:                                          "EMsg_BSVerifyPOSAKey",
	EMsg_BSVerifyPOSAKeyResponse:                                  "EMsg_BSVerifyPOSAKeyResponse",
	EMsg_BSReverseRedeemPOSAKey:                                   "EMsg_BSReverseRedeemPOSAKey",
	EMsg_BSReverseRedeemPOSAKeyResponse:                           "EMsg_BSReverseRedeemPOSAKeyResponse",
	EMsg_BSQueryFindCreditCard:                                    "EMsg_BSQueryFindCreditCard",
	EMsg_BSQueryFindCreditCardResponse:                            "EMsg_BSQueryFindCreditCardResponse",
	EMsg_BSStatusInquiryPOSAKey:                                   "EMsg_BSStatusInquiryPOSAKey",
	EMsg_BSStatusInquiryPOSAKeyResponse:                           "EMsg_BSStatusInquiryPOSAKeyResponse",
	EMsg_BSValidateMoPaySignature:                                 "EMsg_BSValidateMoPaySignature",
	EMsg_BSValidateMoPaySignatureResponse:                         "EMsg_BSValidateMoPaySignatureResponse",
	EMsg_BSMoPayConfirmProductDelivery:                            "EMsg_BSMoPayConfirmProductDelivery",
	EMsg_BSMoPayConfirmProductDeliveryResponse:                    "EMsg_BSMoPayConfirmProductDeliveryResponse",
	EMsg_BSGenerateMoPayMD5:                                       "EMsg_BSGenerateMoPayMD5",
	EMsg_BSGenerateMoPayMD5Response:                               "EMsg_BSGenerateMoPayMD5Response",
	EMsg_BSBoaCompraConfirmProductDelivery:                        "EMsg_BSBoaCompraConfirmProductDelivery",
	EMsg_BSBoaCompraConfirmProductDeliveryResponse:                "EMsg_BSBoaCompraConfirmProductDeliveryResponse",
	EMsg_BSGenerateBoaCompraMD5:                                   "EMsg_BSGenerateBoaCompraMD5",
	EMsg_BSGenerateBoaCompraMD5Response:                           "EMsg_BSGenerateBoaCompraMD5Response",
	EMsg_BSCommitWPTxn:                                            "EMsg_BSCommitWPTxn",
	EMsg_BaseATS:                                                  "EMsg_BaseATS",
	EMsg_ATSStartStressTest:                                       "EMsg_ATSStartStressTest",
	EMsg_ATSStopStressTest:                                        "EMsg_ATSStopStressTest",
	EMsg_ATSRunFailServerTest:                                     "EMsg_ATSRunFailServerTest",
	EMsg_ATSUFSPerfTestTask:                                       "EMsg_ATSUFSPerfTestTask",
	EMsg_ATSUFSPerfTestResponse:                                   "EMsg_ATSUFSPerfTestResponse",
	EMsg_ATSCycleTCM:                                              "EMsg_ATSCycleTCM",
	EMsg_ATSInitDRMSStressTest:                                    "EMsg_ATSInitDRMSStressTest",
	EMsg_ATSCallTest:                                              "EMsg_ATSCallTest",
	EMsg_ATSCallTestReply:                                         "EMsg_ATSCallTestReply",
	EMsg_ATSStartExternalStress:                                   "EMsg_ATSStartExternalStress",
	EMsg_ATSExternalStressJobStart:                                "EMsg_ATSExternalStressJobStart",
	EMsg_ATSExternalStressJobQueued:                               "EMsg_ATSExternalStressJobQueued",
	EMsg_ATSExternalStressJobRunning:                              "EMsg_ATSExternalStressJobRunning",
	EMsg_ATSExternalStressJobStopped:                              "EMsg_ATSExternalStressJobStopped",
	EMsg_ATSExternalStressJobStopAll:                              "EMsg_ATSExternalStressJobStopAll",
	EMsg_ATSExternalStressActionResult:                            "EMsg_ATSExternalStressActionResult",
	EMsg_ATSStarted:                                               "EMsg_ATSStarted",
	EMsg_ATSCSPerfTestTask:                                        "EMsg_ATSCSPerfTestTask",
	EMsg_ATSCSPerfTestResponse:                                    "EMsg_ATSCSPerfTestResponse",
	EMsg_BaseDP:                                                   "EMsg_BaseDP",
	EMsg_DPSetPublishingState:                                     "EMsg_DPSetPublishingState",
	EMsg_DPGamePlayedStats:                                        "EMsg_DPGamePlayedStats",
	EMsg_DPUniquePlayersStat:                                      "EMsg_DPUniquePlayersStat",
	EMsg_DPStreamingUniquePlayersStat:                             "EMsg_DPStreamingUniquePlayersStat",
	EMsg_DPVacInfractionStats:                                     "EMsg_DPVacInfractionStats",
	EMsg_DPVacBanStats:                                            "EMsg_DPVacBanStats",
	EMsg_DPBlockingStats:                                          "EMsg_DPBlockingStats",
	EMsg_DPNatTraversalStats:                                      "EMsg_DPNatTraversalStats",
	EMsg_DPSteamUsageEvent:                                        "EMsg_DPSteamUsageEvent",
	EMsg_DPVacCertBanStats:                                        "EMsg_DPVacCertBanStats",
	EMsg_DPVacCafeBanStats:                                        "EMsg_DPVacCafeBanStats",
	EMsg_DPCloudStats:                                             "EMsg_DPCloudStats",
	EMsg_DPAchievementStats:                                       "EMsg_DPAchievementStats",
	EMsg_DPAccountCreationStats:                                   "EMsg_DPAccountCreationStats",
	EMsg_DPGetPlayerCount:                                         "EMsg_DPGetPlayerCount",
	EMsg_DPGetPlayerCountResponse:                                 "EMsg_DPGetPlayerCountResponse",
	EMsg_DPGameServersPlayersStats:                                "EMsg_DPGameServersPlayersStats",
	EMsg_DPDownloadRateStatistics:                                 "EMsg_DPDownloadRateStatistics",
	EMsg_DPFacebookStatistics:                                     "EMsg_DPFacebookStatistics",
	EMsg_ClientDPCheckSpecialSurvey:                               "EMsg_ClientDPCheckSpecialSurvey",
	EMsg_ClientDPCheckSpecialSurveyResponse:                       "EMsg_ClientDPCheckSpecialSurveyResponse",
	EMsg_ClientDPSendSpecialSurveyResponse:                        "EMsg_ClientDPSendSpecialSurveyResponse",
	EMsg_ClientDPSendSpecialSurveyResponseReply:                   "EMsg_ClientDPSendSpecialSurveyResponseReply",
	EMsg_DPStoreSaleStatistics:                                    "EMsg_DPStoreSaleStatistics",
	EMsg_ClientDPUpdateAppJobReport:                               "EMsg_ClientDPUpdateAppJobReport",
	EMsg_ClientDPSteam2AppStarted:                                 "EMsg_ClientDPSteam2AppStarted",
	EMsg_DPUpdateContentEvent:                                     "EMsg_DPUpdateContentEvent",
	EMsg_DPPartnerMicroTxns:                                       "EMsg_DPPartnerMicroTxns",
	EMsg_DPPartnerMicroTxnsResponse:                               "EMsg_DPPartnerMicroTxnsResponse",
	EMsg_ClientDPContentStatsReport:                               "EMsg_ClientDPContentStatsReport",
	EMsg_DPVRUniquePlayersStat:                                    "EMsg_DPVRUniquePlayersStat",
	EMsg_BaseCM:                                                   "EMsg_BaseCM",
	EMsg_CMSetAllowState:                                          "EMsg_CMSetAllowState",
	EMsg_CMSpewAllowState:                                         "EMsg_CMSpewAllowState",
	EMsg_CMAppInfoResponseDeprecated:                              "EMsg_CMAppInfoResponseDeprecated",
	EMsg_BaseDSS:                                                  "EMsg_BaseDSS",
	EMsg_DSSNewFile:                                               "EMsg_DSSNewFile",
	EMsg_DSSCurrentFileList:                                       "EMsg_DSSCurrentFileList",
	EMsg_DSSSynchList:                                             "EMsg_DSSSynchList",
	EMsg_DSSSynchListResponse:                                     "EMsg_DSSSynchListResponse",
	EMsg_DSSSynchSubscribe:                                        "EMsg_DSSSynchSubscribe",
	EMsg_DSSSynchUnsubscribe:                                      "EMsg_DSSSynchUnsubscribe",
	EMsg_BaseEPM:                                                  "EMsg_BaseEPM",
	EMsg_EPMStartProcess:                                          "EMsg_EPMStartProcess",
	EMsg_EPMStopProcess:                                           "EMsg_EPMStopProcess",
	EMsg_EPMRestartProcess:                                        "EMsg_EPMRestartProcess",
	EMsg_BaseGC:                                                   "EMsg_BaseGC",
	EMsg_AMRelayToGC:                                              "EMsg_AMRelayToGC",
	EMsg_GCUpdatePlayedState:                                      "EMsg_GCUpdatePlayedState",
	EMsg_GCCmdRevive:                                              "EMsg_GCCmdRevive",
	EMsg_GCCmdBounce:                                              "EMsg_GCCmdBounce",
	EMsg_GCCmdForceBounce:                                         "EMsg_GCCmdForceBounce",
	EMsg_GCCmdDown:                                                "EMsg_GCCmdDown",
	EMsg_GCCmdDeploy:                                              "EMsg_GCCmdDeploy",
	EMsg_GCCmdDeployResponse:                                      "EMsg_GCCmdDeployResponse",
	EMsg_GCCmdSwitch:                                              "EMsg_GCCmdSwitch",
	EMsg_AMRefreshSessions:                                        "EMsg_AMRefreshSessions",
	EMsg_GCUpdateGSState:                                          "EMsg_GCUpdateGSState",
	EMsg_GCAchievementAwarded:                                     "EMsg_GCAchievementAwarded",
	EMsg_GCSystemMessage:                                          "EMsg_GCSystemMessage",
	EMsg_GCValidateSession:                                        "EMsg_GCValidateSession",
	EMsg_GCValidateSessionResponse:                                "EMsg_GCValidateSessionResponse",
	EMsg_GCCmdStatus:                                              "EMsg_GCCmdStatus",
	EMsg_GCRegisterWebInterfaces:                                  "EMsg_GCRegisterWebInterfaces",
	EMsg_GCGetAccountDetails:                                      "EMsg_GCGetAccountDetails",
	EMsg_GCInterAppMessage:                                        "EMsg_GCInterAppMessage",
	EMsg_GCGetEmailTemplate:                                       "EMsg_GCGetEmailTemplate",
	EMsg_GCGetEmailTemplateResponse:                               "EMsg_GCGetEmailTemplateResponse",
	EMsg_GCHRelay:                                                 "EMsg_GCHRelay",
	EMsg_GCHRelayToClient:                                         "EMsg_GCHRelayToClient",
	EMsg_GCHUpdateSession:                                         "EMsg_GCHUpdateSession",
	EMsg_GCHRequestUpdateSession:                                  "EMsg_GCHRequestUpdateSession",
	EMsg_GCHRequestStatus:                                         "EMsg_GCHRequestStatus",
	EMsg_GCHRequestStatusResponse:                                 "EMsg_GCHRequestStatusResponse",
	EMsg_GCHAccountVacStatusChange:                                "EMsg_GCHAccountVacStatusChange",
	EMsg_GCHSpawnGC:                                               "EMsg_GCHSpawnGC",
	EMsg_GCHSpawnGCResponse:                                       "EMsg_GCHSpawnGCResponse",
	EMsg_GCHKillGC:                                                "EMsg_GCHKillGC",
	EMsg_GCHKillGCResponse:                                        "EMsg_GCHKillGCResponse",
	EMsg_GCHAccountTradeBanStatusChange:                           "EMsg_GCHAccountTradeBanStatusChange",
	EMsg_GCHAccountLockStatusChange:                               "EMsg_GCHAccountLockStatusChange",
	EMsg_GCHVacVerificationChange:                                 "EMsg_GCHVacVerificationChange",
	EMsg_GCHAccountPhoneNumberChange:                              "EMsg_GCHAccountPhoneNumberChange",
	EMsg_GCHAccountTwoFactorChange:                                "EMsg_GCHAccountTwoFactorChange",
	EMsg_BaseP2P:                                                  "EMsg_BaseP2P",
	EMsg_P2PIntroducerMessage:                                     "EMsg_P2PIntroducerMessage",
	EMsg_BaseSM:                                                   "EMsg_BaseSM",
	EMsg_SMExpensiveReport:                                        "EMsg_SMExpensiveReport",
	EMsg_SMHourlyReport:                                           "EMsg_SMHourlyReport",
	EMsg_SMFishingReport:                                          "EMsg_SMFishingReport",
	EMsg_SMPartitionRenames:                                       "EMsg_SMPartitionRenames",
	EMsg_SMMonitorSpace:                                           "EMsg_SMMonitorSpace",
	EMsg_SMGetSchemaConversionResults:                             "EMsg_SMGetSchemaConversionResults",
	EMsg_SMGetSchemaConversionResultsResponse:                     "EMsg_SMGetSchemaConversionResultsResponse",
	EMsg_BaseTest:                                                 "EMsg_BaseTest",
	EMsg_JobHeartbeatTest:                                         "EMsg_JobHeartbeatTest",
	EMsg_JobHeartbeatTestResponse:                                 "EMsg_JobHeartbeatTestResponse",
	EMsg_BaseFTSRange:                                             "EMsg_BaseFTSRange",
	EMsg_FTSGetBrowseCounts:                                       "EMsg_FTSGetBrowseCounts",
	EMsg_FTSGetBrowseCountsResponse:                               "EMsg_FTSGetBrowseCountsResponse",
	EMsg_FTSBrowseClans:                                           "EMsg_FTSBrowseClans",
	EMsg_FTSBrowseClansResponse:                                   "EMsg_FTSBrowseClansResponse",
	EMsg_FTSSearchClansByLocation:                                 "EMsg_FTSSearchClansByLocation",
	EMsg_FTSSearchClansByLocationResponse:                         "EMsg_FTSSearchClansByLocationResponse",
	EMsg_FTSSearchPlayersByLocation:                               "EMsg_FTSSearchPlayersByLocation",
	EMsg_FTSSearchPlayersByLocationResponse:                       "EMsg_FTSSearchPlayersByLocationResponse",
	EMsg_FTSClanDeleted:                                           "EMsg_FTSClanDeleted",
	EMsg_FTSSearch:                                                "EMsg_FTSSearch",
	EMsg_FTSSearchResponse:                                        "EMsg_FTSSearchResponse",
	EMsg_FTSSearchStatus:                                          "EMsg_FTSSearchStatus",
	EMsg_FTSSearchStatusResponse:                                  "EMsg_FTSSearchStatusResponse",
	EMsg_FTSGetGSPlayStats:                                        "EMsg_FTSGetGSPlayStats",
	EMsg_FTSGetGSPlayStatsResponse:                                "EMsg_FTSGetGSPlayStatsResponse",
	EMsg_FTSGetGSPlayStatsForServer:                               "EMsg_FTSGetGSPlayStatsForServer",
	EMsg_FTSGetGSPlayStatsForServerResponse:                       "EMsg_FTSGetGSPlayStatsForServerResponse",
	EMsg_FTSReportIPUpdates:                                       "EMsg_FTSReportIPUpdates",
	EMsg_BaseCCSRange:                                             "EMsg_BaseCCSRange",
	EMsg_CCSGetComments:                                           "EMsg_CCSGetComments",
	EMsg_CCSGetCommentsResponse:                                   "EMsg_CCSGetCommentsResponse",
	EMsg_CCSAddComment:                                            "EMsg_CCSAddComment",
	EMsg_CCSAddCommentResponse:                                    "EMsg_CCSAddCommentResponse",
	EMsg_CCSDeleteComment:                                         "EMsg_CCSDeleteComment",
	EMsg_CCSDeleteCommentResponse:                                 "EMsg_CCSDeleteCommentResponse",
	EMsg_CCSPreloadComments:                                       "EMsg_CCSPreloadComments",
	EMsg_CCSNotifyCommentCount:                                    "EMsg_CCSNotifyCommentCount",
	EMsg_CCSGetCommentsForNews:                                    "EMsg_CCSGetCommentsForNews",
	EMsg_CCSGetCommentsForNewsResponse:                            "EMsg_CCSGetCommentsForNewsResponse",
	EMsg_CCSDeleteAllCommentsByAuthor:                             "EMsg_CCSDeleteAllCommentsByAuthor",
	EMsg_CCSDeleteAllCommentsByAuthorResponse:                     "EMsg_CCSDeleteAllCommentsByAuthorResponse",
	EMsg_BaseLBSRange:                                             "EMsg_BaseLBSRange",
	EMsg_LBSSetScore:                                              "EMsg_LBSSetScore",
	EMsg_LBSSetScoreResponse:                                      "EMsg_LBSSetScoreResponse",
	EMsg_LBSFindOrCreateLB:                                        "EMsg_LBSFindOrCreateLB",
	EMsg_LBSFindOrCreateLBResponse:                                "EMsg_LBSFindOrCreateLBResponse",
	EMsg_LBSGetLBEntries:                                          "EMsg_LBSGetLBEntries",
	EMsg_LBSGetLBEntriesResponse:                                  "EMsg_LBSGetLBEntriesResponse",
	EMsg_LBSGetLBList:                                             "EMsg_LBSGetLBList",
	EMsg_LBSGetLBListResponse:                                     "EMsg_LBSGetLBListResponse",
	EMsg_LBSSetLBDetails:                                          "EMsg_LBSSetLBDetails",
	EMsg_LBSDeleteLB:                                              "EMsg_LBSDeleteLB",
	EMsg_LBSDeleteLBEntry:                                         "EMsg_LBSDeleteLBEntry",
	EMsg_LBSResetLB:                                               "EMsg_LBSResetLB",
	EMsg_LBSResetLBResponse:                                       "EMsg_LBSResetLBResponse",
	EMsg_LBSDeleteLBResponse:                                      "EMsg_LBSDeleteLBResponse",
	EMsg_BaseOGS:                                                  "EMsg_BaseOGS",
	EMsg_OGSBeginSession:                                          "EMsg_OGSBeginSession",
	EMsg_OGSBeginSessionResponse:                                  "EMsg_OGSBeginSessionResponse",
	EMsg_OGSEndSession:                                            "EMsg_OGSEndSession",
	EMsg_OGSEndSessionResponse:                                    "EMsg_OGSEndSessionResponse",
	EMsg_OGSWriteAppSessionRow:                                    "EMsg_OGSWriteAppSessionRow",
	EMsg_BaseBRP:                                                  "EMsg_BaseBRP",
	EMsg_BRPStartShippingJobs:                                     "EMsg_BRPStartShippingJobs",
	EMsg_BRPProcessUSBankReports:                                  "EMsg_BRPProcessUSBankReports",
	EMsg_BRPProcessGCReports:                                      "EMsg_BRPProcessGCReports",
	EMsg_BRPProcessPPReports:                                      "EMsg_BRPProcessPPReports",
	EMsg_BRPSettleNOVA:                                            "EMsg_BRPSettleNOVA",
	EMsg_BRPSettleCB:                                              "EMsg_BRPSettleCB",
	EMsg_BRPCommitGC:                                              "EMsg_BRPCommitGC",
	EMsg_BRPCommitGCResponse:                                      "EMsg_BRPCommitGCResponse",
	EMsg_BRPFindHungTransactions:                                  "EMsg_BRPFindHungTransactions",
	EMsg_BRPCheckFinanceCloseOutDate:                              "EMsg_BRPCheckFinanceCloseOutDate",
	EMsg_BRPProcessLicenses:                                       "EMsg_BRPProcessLicenses",
	EMsg_BRPProcessLicensesResponse:                               "EMsg_BRPProcessLicensesResponse",
	EMsg_BRPRemoveExpiredPaymentData:                              "EMsg_BRPRemoveExpiredPaymentData",
	EMsg_BRPRemoveExpiredPaymentDataResponse:                      "EMsg_BRPRemoveExpiredPaymentDataResponse",
	EMsg_BRPConvertToCurrentKeys:                                  "EMsg_BRPConvertToCurrentKeys",
	EMsg_BRPConvertToCurrentKeysResponse:                          "EMsg_BRPConvertToCurrentKeysResponse",
	EMsg_BRPPruneCardUsageStats:                                   "EMsg_BRPPruneCardUsageStats",
	EMsg_BRPPruneCardUsageStatsResponse:                           "EMsg_BRPPruneCardUsageStatsResponse",
	EMsg_BRPCheckActivationCodes:                                  "EMsg_BRPCheckActivationCodes",
	EMsg_BRPCheckActivationCodesResponse:                          "EMsg_BRPCheckActivationCodesResponse",
	EMsg_BRPCommitWP:                                              "EMsg_BRPCommitWP",
	EMsg_BRPCommitWPResponse:                                      "EMsg_BRPCommitWPResponse",
	EMsg_BRPProcessWPReports:                                      "EMsg_BRPProcessWPReports",
	EMsg_BRPProcessPaymentRules:                                   "EMsg_BRPProcessPaymentRules",
	EMsg_BRPProcessPartnerPayments:                                "EMsg_BRPProcessPartnerPayments",
	EMsg_BRPCheckSettlementReports:                                "EMsg_BRPCheckSettlementReports",
	EMsg_BRPPostTaxToAvalara:                                      "EMsg_BRPPostTaxToAvalara",
	EMsg_BRPPostTransactionTax:                                    "EMsg_BRPPostTransactionTax",
	EMsg_BRPPostTransactionTaxResponse:                            "EMsg_BRPPostTransactionTaxResponse",
	EMsg_BRPProcessIMReports:                                      "EMsg_BRPProcessIMReports",
	EMsg_BaseAMRange2:                                             "EMsg_BaseAMRange2",
	EMsg_AMCreateChat:                                             "EMsg_AMCreateChat",
	EMsg_AMCreateChatResponse:                                     "EMsg_AMCreateChatResponse",
	EMsg_AMUpdateChatMetadata:                                     "EMsg_AMUpdateChatMetadata",
	EMsg_AMPublishChatMetadata:                                    "EMsg_AMPublishChatMetadata",
	EMsg_AMSetProfileURL:                                          "EMsg_AMSetProfileURL",
	EMsg_AMGetAccountEmailAddress:                                 "EMsg_AMGetAccountEmailAddress",
	EMsg_AMGetAccountEmailAddressResponse:                         "EMsg_AMGetAccountEmailAddressResponse",
	EMsg_AMRequestClanData:                                        "EMsg_AMRequestClanData",
	EMsg_AMRouteToClients:                                         "EMsg_AMRouteToClients",
	EMsg_AMLeaveClan:                                              "EMsg_AMLeaveClan",
	EMsg_AMClanPermissions:                                        "EMsg_AMClanPermissions",
	EMsg_AMClanPermissionsResponse:                                "EMsg_AMClanPermissionsResponse",
	EMsg_AMCreateClanEvent:                                        "EMsg_AMCreateClanEvent",
	EMsg_AMCreateClanEventResponse:                                "EMsg_AMCreateClanEventResponse",
	EMsg_AMUpdateClanEvent:                                        "EMsg_AMUpdateClanEvent",
	EMsg_AMUpdateClanEventResponse:                                "EMsg_AMUpdateClanEventResponse",
	EMsg_AMGetClanEvents:                                          "EMsg_AMGetClanEvents",
	EMsg_AMGetClanEventsResponse:                                  "EMsg_AMGetClanEventsResponse",
	EMsg_AMDeleteClanEvent:                                        "EMsg_AMDeleteClanEvent",
	EMsg_AMDeleteClanEventResponse:                                "EMsg_AMDeleteClanEventResponse",
	EMsg_AMSetClanPermissionSettings:                              "EMsg_AMSetClanPermissionSettings",
	EMsg_AMSetClanPermissionSettingsResponse:                      "EMsg_AMSetClanPermissionSettingsResponse",
	EMsg_AMGetClanPermissionSettings:                              "EMsg_AMGetClanPermissionSettings",
	EMsg_AMGetClanPermissionSettingsResponse:                      "EMsg_AMGetClanPermissionSettingsResponse",
	EMsg_AMPublishChatRoomInfo:                                    "EMsg_AMPublishChatRoomInfo",
	EMsg_ClientChatRoomInfo:                                       "EMsg_ClientChatRoomInfo",
	EMsg_AMCreateClanAnnouncement:                                 "EMsg_AMCreateClanAnnouncement",
	EMsg_AMCreateClanAnnouncementResponse:                         "EMsg_AMCreateClanAnnouncementResponse",
	EMsg_AMUpdateClanAnnouncement:                                 "EMsg_AMUpdateClanAnnouncement",
	EMsg_AMUpdateClanAnnouncementResponse:                         "EMsg_AMUpdateClanAnnouncementResponse",
	EMsg_AMGetClanAnnouncementsCount:                              "EMsg_AMGetClanAnnouncementsCount",
	EMsg_AMGetClanAnnouncementsCountResponse:                      "EMsg_AMGetClanAnnouncementsCountResponse",
	EMsg_AMGetClanAnnouncements:                                   "EMsg_AMGetClanAnnouncements",
	EMsg_AMGetClanAnnouncementsResponse:                           "EMsg_AMGetClanAnnouncementsResponse",
	EMsg_AMDeleteClanAnnouncement:                                 "EMsg_AMDeleteClanAnnouncement",
	EMsg_AMDeleteClanAnnouncementResponse:                         "EMsg_AMDeleteClanAnnouncementResponse",
	EMsg_AMGetSingleClanAnnouncement:                              "EMsg_AMGetSingleClanAnnouncement",
	EMsg_AMGetSingleClanAnnouncementResponse:                      "EMsg_AMGetSingleClanAnnouncementResponse",
	EMsg_AMGetClanHistory:                                         "EMsg_AMGetClanHistory",
	EMsg_AMGetClanHistoryResponse:                                 "EMsg_AMGetClanHistoryResponse",
	EMsg_AMGetClanPermissionBits:                                  "EMsg_AMGetClanPermissionBits",
	EMsg_AMGetClanPermissionBitsResponse:                          "EMsg_AMGetClanPermissionBitsResponse",
	EMsg_AMSetClanPermissionBits:                                  "EMsg_AMSetClanPermissionBits",
	EMsg_AMSetClanPermissionBitsResponse:                          "EMsg_AMSetClanPermissionBitsResponse",
	EMsg_AMSessionInfoRequest:                                     "EMsg_AMSessionInfoRequest",
	EMsg_AMSessionInfoResponse:                                    "EMsg_AMSessionInfoResponse",
	EMsg_AMValidateWGToken:                                        "EMsg_AMValidateWGToken",
	EMsg_AMGetSingleClanEvent:                                     "EMsg_AMGetSingleClanEvent",
	EMsg_AMGetSingleClanEventResponse:                             "EMsg_AMGetSingleClanEventResponse",
	EMsg_AMGetClanRank:                                            "EMsg_AMGetClanRank",
	EMsg_AMGetClanRankResponse:                                    "EMsg_AMGetClanRankResponse",
	EMsg_AMSetClanRank:                                            "EMsg_AMSetClanRank",
	EMsg_AMSetClanRankResponse:                                    "EMsg_AMSetClanRankResponse",
	EMsg_AMGetClanPOTW:                                            "EMsg_AMGetClanPOTW",
	EMsg_AMGetClanPOTWResponse:                                    "EMsg_AMGetClanPOTWResponse",
	EMsg_AMSetClanPOTW:                                            "EMsg_AMSetClanPOTW",
	EMsg_AMSetClanPOTWResponse:                                    "EMsg_AMSetClanPOTWResponse",
	EMsg_AMRequestChatMetadata:                                    "EMsg_AMRequestChatMetadata",
	EMsg_AMDumpUser:                                               "EMsg_AMDumpUser",
	EMsg_AMKickUserFromClan:                                       "EMsg_AMKickUserFromClan",
	EMsg_AMAddFounderToClan:                                       "EMsg_AMAddFounderToClan",
	EMsg_AMValidateWGTokenResponse:                                "EMsg_AMValidateWGTokenResponse",
	EMsg_AMSetCommunityState:                                      "EMsg_AMSetCommunityState",
	EMsg_AMSetAccountDetails:                                      "EMsg_AMSetAccountDetails",
	EMsg_AMGetChatBanList:                                         "EMsg_AMGetChatBanList",
	EMsg_AMGetChatBanListResponse:                                 "EMsg_AMGetChatBanListResponse",
	EMsg_AMUnBanFromChat:                                          "EMsg_AMUnBanFromChat",
	EMsg_AMSetClanDetails:                                         "EMsg_AMSetClanDetails",
	EMsg_AMGetAccountLinks:                                        "EMsg_AMGetAccountLinks",
	EMsg_AMGetAccountLinksResponse:                                "EMsg_AMGetAccountLinksResponse",
	EMsg_AMSetAccountLinks:                                        "EMsg_AMSetAccountLinks",
	EMsg_AMSetAccountLinksResponse:                                "EMsg_AMSetAccountLinksResponse",
	EMsg_AMGetUserGameStats:                                       "EMsg_AMGetUserGameStats",
	EMsg_AMGetUserGameStatsResponse:                               "EMsg_AMGetUserGameStatsResponse",
	EMsg_AMCheckClanMembership:                                    "EMsg_AMCheckClanMembership",
	EMsg_AMGetClanMembers:                                         "EMsg_AMGetClanMembers",
	EMsg_AMGetClanMembersResponse:                                 "EMsg_AMGetClanMembersResponse",
	EMsg_AMJoinPublicClan:                                         "EMsg_AMJoinPublicClan",
	EMsg_AMNotifyChatOfClanChange:                                 "EMsg_AMNotifyChatOfClanChange",
	EMsg_AMResubmitPurchase:                                       "EMsg_AMResubmitPurchase",
	EMsg_AMAddFriend:                                              "EMsg_AMAddFriend",
	EMsg_AMAddFriendResponse:                                      "EMsg_AMAddFriendResponse",
	EMsg_AMRemoveFriend:                                           "EMsg_AMRemoveFriend",
	EMsg_AMDumpClan:                                               "EMsg_AMDumpClan",
	EMsg_AMChangeClanOwner:                                        "EMsg_AMChangeClanOwner",
	EMsg_AMCancelEasyCollect:                                      "EMsg_AMCancelEasyCollect",
	EMsg_AMCancelEasyCollectResponse:                              "EMsg_AMCancelEasyCollectResponse",
	EMsg_AMGetClanMembershipList:                                  "EMsg_AMGetClanMembershipList",
	EMsg_AMGetClanMembershipListResponse:                          "EMsg_AMGetClanMembershipListResponse",
	EMsg_AMClansInCommon:                                          "EMsg_AMClansInCommon",
	EMsg_AMClansInCommonResponse:                                  "EMsg_AMClansInCommonResponse",
	EMsg_AMIsValidAccountID:                                       "EMsg_AMIsValidAccountID",
	EMsg_AMConvertClan:                                            "EMsg_AMConvertClan",
	EMsg_AMGetGiftTargetListRelay:                                 "EMsg_AMGetGiftTargetListRelay",
	EMsg_AMWipeFriendsList:                                        "EMsg_AMWipeFriendsList",
	EMsg_AMSetIgnored:                                             "EMsg_AMSetIgnored",
	EMsg_AMClansInCommonCountResponse:                             "EMsg_AMClansInCommonCountResponse",
	EMsg_AMFriendsList:                                            "EMsg_AMFriendsList",
	EMsg_AMFriendsListResponse:                                    "EMsg_AMFriendsListResponse",
	EMsg_AMFriendsInCommon:                                        "EMsg_AMFriendsInCommon",
	EMsg_AMFriendsInCommonResponse:                                "EMsg_AMFriendsInCommonResponse",
	EMsg_AMFriendsInCommonCountResponse:                           "EMsg_AMFriendsInCommonCountResponse",
	EMsg_AMClansInCommonCount:                                     "EMsg_AMClansInCommonCount",
	EMsg_AMChallengeVerdict:                                       "EMsg_AMChallengeVerdict",
	EMsg_AMChallengeNotification:                                  "EMsg_AMChallengeNotification",
	EMsg_AMFindGSByIP:                                             "EMsg_AMFindGSByIP",
	EMsg_AMFoundGSByIP:                                            "EMsg_AMFoundGSByIP",
	EMsg_AMGiftRevoked:                                            "EMsg_AMGiftRevoked",
	EMsg_AMCreateAccountRecord:                                    "EMsg_AMCreateAccountRecord",
	EMsg_AMUserClanList:                                           "EMsg_AMUserClanList",
	EMsg_AMUserClanListResponse:                                   "EMsg_AMUserClanListResponse",
	EMsg_AMGetAccountDetails2:                                     "EMsg_AMGetAccountDetails2",
	EMsg_AMGetAccountDetailsResponse2:                             "EMsg_AMGetAccountDetailsResponse2",
	EMsg_AMSetCommunityProfileSettings:                            "EMsg_AMSetCommunityProfileSettings",
	EMsg_AMSetCommunityProfileSettingsResponse:                    "EMsg_AMSetCommunityProfileSettingsResponse",
	EMsg_AMGetCommunityPrivacyState:                               "EMsg_AMGetCommunityPrivacyState",
	EMsg_AMGetCommunityPrivacyStateResponse:                       "EMsg_AMGetCommunityPrivacyStateResponse",
	EMsg_AMCheckClanInviteRateLimiting:                            "EMsg_AMCheckClanInviteRateLimiting",
	EMsg_AMGetUserAchievementStatus:                               "EMsg_AMGetUserAchievementStatus",
	EMsg_AMGetIgnored:                                             "EMsg_AMGetIgnored",
	EMsg_AMGetIgnoredResponse:                                     "EMsg_AMGetIgnoredResponse",
	EMsg_AMSetIgnoredResponse:                                     "EMsg_AMSetIgnoredResponse",
	EMsg_AMSetFriendRelationshipNone:                              "EMsg_AMSetFriendRelationshipNone",
	EMsg_AMGetFriendRelationship:                                  "EMsg_AMGetFriendRelationship",
	EMsg_AMGetFriendRelationshipResponse:                          "EMsg_AMGetFriendRelationshipResponse",
	EMsg_AMServiceModulesCache:                                    "EMsg_AMServiceModulesCache",
	EMsg_AMServiceModulesCall:                                     "EMsg_AMServiceModulesCall",
	EMsg_AMServiceModulesCallResponse:                             "EMsg_AMServiceModulesCallResponse",
	EMsg_AMGetCaptchaDataForIP:                                    "EMsg_AMGetCaptchaDataForIP",
	EMsg_AMGetCaptchaDataForIPResponse:                            "EMsg_AMGetCaptchaDataForIPResponse",
	EMsg_AMValidateCaptchaDataForIP:                               "EMsg_AMValidateCaptchaDataForIP",
	EMsg_AMValidateCaptchaDataForIPResponse:                       "EMsg_AMValidateCaptchaDataForIPResponse",
	EMsg_AMTrackFailedAuthByIP:                                    "EMsg_AMTrackFailedAuthByIP",
	EMsg_AMGetCaptchaDataByGID:                                    "EMsg_AMGetCaptchaDataByGID",
	EMsg_AMGetCaptchaDataByGIDResponse:                            "EMsg_AMGetCaptchaDataByGIDResponse",
	EMsg_AMGetLobbyList:                                           "EMsg_AMGetLobbyList",
	EMsg_AMGetLobbyListResponse:                                   "EMsg_AMGetLobbyListResponse",
	EMsg_AMGetLobbyMetadata:                                       "EMsg_AMGetLobbyMetadata",
	EMsg_AMGetLobbyMetadataResponse:                               "EMsg_AMGetLobbyMetadataResponse",
	EMsg_CommunityAddFriendNews:                                   "EMsg_CommunityAddFriendNews",
	EMsg_AMAddClanNews:                                            "EMsg_AMAddClanNews",
	EMsg_AMWriteNews:                                              "EMsg_AMWriteNews",
	EMsg_AMFindClanUser:                                           "EMsg_AMFindClanUser",
	EMsg_AMFindClanUserResponse:                                   "EMsg_AMFindClanUserResponse",
	EMsg_AMBanFromChat:                                            "EMsg_AMBanFromChat",
	EMsg_AMGetUserHistoryResponse:                                 "EMsg_AMGetUserHistoryResponse",
	EMsg_AMGetUserNewsSubscriptions:                               "EMsg_AMGetUserNewsSubscriptions",
	EMsg_AMGetUserNewsSubscriptionsResponse:                       "EMsg_AMGetUserNewsSubscriptionsResponse",
	EMsg_AMSetUserNewsSubscriptions:                               "EMsg_AMSetUserNewsSubscriptions",
	EMsg_AMGetUserNews:                                            "EMsg_AMGetUserNews",
	EMsg_AMGetUserNewsResponse:                                    "EMsg_AMGetUserNewsResponse",
	EMsg_AMSendQueuedEmails:                                       "EMsg_AMSendQueuedEmails",
	EMsg_AMSetLicenseFlags:                                        "EMsg_AMSetLicenseFlags",
	EMsg_AMGetUserHistory:                                         "EMsg_AMGetUserHistory",
	EMsg_CommunityDeleteUserNews:                                  "EMsg_CommunityDeleteUserNews",
	EMsg_AMAllowUserFilesRequest:                                  "EMsg_AMAllowUserFilesRequest",
	EMsg_AMAllowUserFilesResponse:                                 "EMsg_AMAllowUserFilesResponse",
	EMsg_AMGetAccountStatus:                                       "EMsg_AMGetAccountStatus",
	EMsg_AMGetAccountStatusResponse:                               "EMsg_AMGetAccountStatusResponse",
	EMsg_AMEditBanReason:                                          "EMsg_AMEditBanReason",
	EMsg_AMCheckClanMembershipResponse:                            "EMsg_AMCheckClanMembershipResponse",
	EMsg_AMProbeClanMembershipList:                                "EMsg_AMProbeClanMembershipList",
	EMsg_AMProbeClanMembershipListResponse:                        "EMsg_AMProbeClanMembershipListResponse",
	EMsg_AMGetFriendsLobbies:                                      "EMsg_AMGetFriendsLobbies",
	EMsg_AMGetFriendsLobbiesResponse:                              "EMsg_AMGetFriendsLobbiesResponse",
	EMsg_AMGetUserFriendNewsResponse:                              "EMsg_AMGetUserFriendNewsResponse",
	EMsg_CommunityGetUserFriendNews:                               "EMsg_CommunityGetUserFriendNews",
	EMsg_AMGetUserClansNewsResponse:                               "EMsg_AMGetUserClansNewsResponse",
	EMsg_AMGetUserClansNews:                                       "EMsg_AMGetUserClansNews",
	EMsg_AMStoreInitPurchase:                                      "EMsg_AMStoreInitPurchase",
	EMsg_AMStoreInitPurchaseResponse:                              "EMsg_AMStoreInitPurchaseResponse",
	EMsg_AMStoreGetFinalPrice:                                     "EMsg_AMStoreGetFinalPrice",
	EMsg_AMStoreGetFinalPriceResponse:                             "EMsg_AMStoreGetFinalPriceResponse",
	EMsg_AMStoreCompletePurchase:                                  "EMsg_AMStoreCompletePurchase",
	EMsg_AMStoreCancelPurchase:                                    "EMsg_AMStoreCancelPurchase",
	EMsg_AMStorePurchaseResponse:                                  "EMsg_AMStorePurchaseResponse",
	EMsg_AMCreateAccountRecordInSteam3:                            "EMsg_AMCreateAccountRecordInSteam3",
	EMsg_AMGetPreviousCBAccount:                                   "EMsg_AMGetPreviousCBAccount",
	EMsg_AMGetPreviousCBAccountResponse:                           "EMsg_AMGetPreviousCBAccountResponse",
	EMsg_AMUpdateBillingAddress:                                   "EMsg_AMUpdateBillingAddress",
	EMsg_AMUpdateBillingAddressResponse:                           "EMsg_AMUpdateBillingAddressResponse",
	EMsg_AMGetBillingAddress:                                      "EMsg_AMGetBillingAddress",
	EMsg_AMGetBillingAddressResponse:                              "EMsg_AMGetBillingAddressResponse",
	EMsg_AMGetUserLicenseHistory:                                  "EMsg_AMGetUserLicenseHistory",
	EMsg_AMGetUserLicenseHistoryResponse:                          "EMsg_AMGetUserLicenseHistoryResponse",
	EMsg_AMSupportChangePassword:                                  "EMsg_AMSupportChangePassword",
	EMsg_AMSupportChangeEmail:                                     "EMsg_AMSupportChangeEmail",
	EMsg_AMSupportChangeSecretQA:                                  "EMsg_AMSupportChangeSecretQA",
	EMsg_AMResetUserVerificationGSByIP:                            "EMsg_AMResetUserVerificationGSByIP",
	EMsg_AMUpdateGSPlayStats:                                      "EMsg_AMUpdateGSPlayStats",
	EMsg_AMSupportEnableOrDisable:                                 "EMsg_AMSupportEnableOrDisable",
	EMsg_AMGetComments:                                            "EMsg_AMGetComments",
	EMsg_AMGetCommentsResponse:                                    "EMsg_AMGetCommentsResponse",
	EMsg_AMAddComment:                                             "EMsg_AMAddComment",
	EMsg_AMAddCommentResponse:                                     "EMsg_AMAddCommentResponse",
	EMsg_AMDeleteComment:                                          "EMsg_AMDeleteComment",
	EMsg_AMDeleteCommentResponse:                                  "EMsg_AMDeleteCommentResponse",
	EMsg_AMGetPurchaseStatus:                                      "EMsg_AMGetPurchaseStatus",
	EMsg_AMSupportIsAccountEnabled:                                "EMsg_AMSupportIsAccountEnabled",
	EMsg_AMSupportIsAccountEnabledResponse:                        "EMsg_AMSupportIsAccountEnabledResponse",
	EMsg_AMGetUserStats:                                           "EMsg_AMGetUserStats",
	EMsg_AMSupportKickSession:                                     "EMsg_AMSupportKickSession",
	EMsg_AMGSSearch:                                               "EMsg_AMGSSearch",
	EMsg_MarketingMessageUpdate:                                   "EMsg_MarketingMessageUpdate",
	EMsg_AMRouteFriendMsg:                                         "EMsg_AMRouteFriendMsg",
	EMsg_AMTicketAuthRequestOrResponse:                            "EMsg_AMTicketAuthRequestOrResponse",
	EMsg_AMVerifyDepotManagementRights:                            "EMsg_AMVerifyDepotManagementRights",
	EMsg_AMVerifyDepotManagementRightsResponse:                    "EMsg_AMVerifyDepotManagementRightsResponse",
	EMsg_AMAddFreeLicense:                                         "EMsg_AMAddFreeLicense",
	EMsg_AMGetUserFriendsMinutesPlayed:                            "EMsg_AMGetUserFriendsMinutesPlayed",
	EMsg_AMGetUserFriendsMinutesPlayedResponse:                    "EMsg_AMGetUserFriendsMinutesPlayedResponse",
	EMsg_AMGetUserMinutesPlayed:                                   "EMsg_AMGetUserMinutesPlayed",
	EMsg_AMGetUserMinutesPlayedResponse:                           "EMsg_AMGetUserMinutesPlayedResponse",
	EMsg_AMValidateEmailLink:                                      "EMsg_AMValidateEmailLink",
	EMsg_AMValidateEmailLinkResponse:                              "EMsg_AMValidateEmailLinkResponse",
	EMsg_AMAddUsersToMarketingTreatment:                           "EMsg_AMAddUsersToMarketingTreatment",
	EMsg_AMStoreUserStats:                                         "EMsg_AMStoreUserStats",
	EMsg_AMGetUserGameplayInfo:                                    "EMsg_AMGetUserGameplayInfo",
	EMsg_AMGetUserGameplayInfoResponse:                            "EMsg_AMGetUserGameplayInfoResponse",
	EMsg_AMGetCardList:                                            "EMsg_AMGetCardList",
	EMsg_AMGetCardListResponse:                                    "EMsg_AMGetCardListResponse",
	EMsg_AMDeleteStoredCard:                                       "EMsg_AMDeleteStoredCard",
	EMsg_AMRevokeLegacyGameKeys:                                   "EMsg_AMRevokeLegacyGameKeys",
	EMsg_AMGetWalletDetails:                                       "EMsg_AMGetWalletDetails",
	EMsg_AMGetWalletDetailsResponse:                               "EMsg_AMGetWalletDetailsResponse",
	EMsg_AMDeleteStoredPaymentInfo:                                "EMsg_AMDeleteStoredPaymentInfo",
	EMsg_AMGetStoredPaymentSummary:                                "EMsg_AMGetStoredPaymentSummary",
	EMsg_AMGetStoredPaymentSummaryResponse:                        "EMsg_AMGetStoredPaymentSummaryResponse",
	EMsg_AMGetWalletConversionRate:                                "EMsg_AMGetWalletConversionRate",
	EMsg_AMGetWalletConversionRateResponse:                        "EMsg_AMGetWalletConversionRateResponse",
	EMsg_AMConvertWallet:                                          "EMsg_AMConvertWallet",
	EMsg_AMConvertWalletResponse:                                  "EMsg_AMConvertWalletResponse",
	EMsg_AMRelayGetFriendsWhoPlayGame:                             "EMsg_AMRelayGetFriendsWhoPlayGame",
	EMsg_AMRelayGetFriendsWhoPlayGameResponse:                     "EMsg_AMRelayGetFriendsWhoPlayGameResponse",
	EMsg_AMSetPreApproval:                                         "EMsg_AMSetPreApproval",
	EMsg_AMSetPreApprovalResponse:                                 "EMsg_AMSetPreApprovalResponse",
	EMsg_AMMarketingTreatmentUpdate:                               "EMsg_AMMarketingTreatmentUpdate",
	EMsg_AMCreateRefund:                                           "EMsg_AMCreateRefund",
	EMsg_AMCreateRefundResponse:                                   "EMsg_AMCreateRefundResponse",
	EMsg_AMCreateChargeback:                                       "EMsg_AMCreateChargeback",
	EMsg_AMCreateChargebackResponse:                               "EMsg_AMCreateChargebackResponse",
	EMsg_AMCreateDispute:                                          "EMsg_AMCreateDispute",
	EMsg_AMCreateDisputeResponse:                                  "EMsg_AMCreateDisputeResponse",
	EMsg_AMClearDispute:                                           "EMsg_AMClearDispute",
	EMsg_AMClearDisputeResponse:                                   "EMsg_AMClearDisputeResponse",
	EMsg_AMPlayerNicknameList:                                     "EMsg_AMPlayerNicknameList",
	EMsg_AMPlayerNicknameListResponse:                             "EMsg_AMPlayerNicknameListResponse",
	EMsg_AMSetDRMTestConfig:                                       "EMsg_AMSetDRMTestConfig",
	EMsg_AMGetUserCurrentGameInfo:                                 "EMsg_AMGetUserCurrentGameInfo",
	EMsg_AMGetUserCurrentGameInfoResponse:                         "EMsg_AMGetUserCurrentGameInfoResponse",
	EMsg_AMGetGSPlayerList:                                        "EMsg_AMGetGSPlayerList",
	EMsg_AMGetGSPlayerListResponse:                                "EMsg_AMGetGSPlayerListResponse",
	EMsg_AMUpdatePersonaStateCache:                                "EMsg_AMUpdatePersonaStateCache",
	EMsg_AMGetGameMembers:                                         "EMsg_AMGetGameMembers",
	EMsg_AMGetGameMembersResponse:                                 "EMsg_AMGetGameMembersResponse",
	EMsg_AMGetSteamIDForMicroTxn:                                  "EMsg_AMGetSteamIDForMicroTxn",
	EMsg_AMGetSteamIDForMicroTxnResponse:                          "EMsg_AMGetSteamIDForMicroTxnResponse",
	EMsg_AMAddPublisherUser:                                       "EMsg_AMAddPublisherUser",
	EMsg_AMRemovePublisherUser:                                    "EMsg_AMRemovePublisherUser",
	EMsg_AMGetUserLicenseList:                                     "EMsg_AMGetUserLicenseList",
	EMsg_AMGetUserLicenseListResponse:                             "EMsg_AMGetUserLicenseListResponse",
	EMsg_AMReloadGameGroupPolicy:                                  "EMsg_AMReloadGameGroupPolicy",
	EMsg_AMAddFreeLicenseResponse:                                 "EMsg_AMAddFreeLicenseResponse",
	EMsg_AMVACStatusUpdate:                                        "EMsg_AMVACStatusUpdate",
	EMsg_AMGetAccountDetails:                                      "EMsg_AMGetAccountDetails",
	EMsg_AMGetAccountDetailsResponse:                              "EMsg_AMGetAccountDetailsResponse",
	EMsg_AMGetPlayerLinkDetails:                                   "EMsg_AMGetPlayerLinkDetails",
	EMsg_AMGetPlayerLinkDetailsResponse:                           "EMsg_AMGetPlayerLinkDetailsResponse",
	EMsg_AMSubscribeToPersonaFeed:                                 "EMsg_AMSubscribeToPersonaFeed",
	EMsg_AMGetUserVacBanList:                                      "EMsg_AMGetUserVacBanList",
	EMsg_AMGetUserVacBanListResponse:                              "EMsg_AMGetUserVacBanListResponse",
	EMsg_AMGetAccountFlagsForWGSpoofing:                           "EMsg_AMGetAccountFlagsForWGSpoofing",
	EMsg_AMGetAccountFlagsForWGSpoofingResponse:                   "EMsg_AMGetAccountFlagsForWGSpoofingResponse",
	EMsg_AMGetFriendsWishlistInfo:                                 "EMsg_AMGetFriendsWishlistInfo",
	EMsg_AMGetFriendsWishlistInfoResponse:                         "EMsg_AMGetFriendsWishlistInfoResponse",
	EMsg_AMGetClanOfficers:                                        "EMsg_AMGetClanOfficers",
	EMsg_AMGetClanOfficersResponse:                                "EMsg_AMGetClanOfficersResponse",
	EMsg_AMNameChange:                                             "EMsg_AMNameChange",
	EMsg_AMGetNameHistory:                                         "EMsg_AMGetNameHistory",
	EMsg_AMGetNameHistoryResponse:                                 "EMsg_AMGetNameHistoryResponse",
	EMsg_AMUpdateProviderStatus:                                   "EMsg_AMUpdateProviderStatus",
	EMsg_AMClearPersonaMetadataBlob:                               "EMsg_AMClearPersonaMetadataBlob",
	EMsg_AMSupportRemoveAccountSecurity:                           "EMsg_AMSupportRemoveAccountSecurity",
	EMsg_AMIsAccountInCaptchaGracePeriod:                          "EMsg_AMIsAccountInCaptchaGracePeriod",
	EMsg_AMIsAccountInCaptchaGracePeriodResponse:                  "EMsg_AMIsAccountInCaptchaGracePeriodResponse",
	EMsg_AMAccountPS3Unlink:                                       "EMsg_AMAccountPS3Unlink",
	EMsg_AMAccountPS3UnlinkResponse:                               "EMsg_AMAccountPS3UnlinkResponse",
	EMsg_AMStoreUserStatsResponse:                                 "EMsg_AMStoreUserStatsResponse",
	EMsg_AMGetAccountPSNInfo:                                      "EMsg_AMGetAccountPSNInfo",
	EMsg_AMGetAccountPSNInfoResponse:                              "EMsg_AMGetAccountPSNInfoResponse",
	EMsg_AMAuthenticatedPlayerList:                                "EMsg_AMAuthenticatedPlayerList",
	EMsg_AMGetUserGifts:                                           "EMsg_AMGetUserGifts",
	EMsg_AMGetUserGiftsResponse:                                   "EMsg_AMGetUserGiftsResponse",
	EMsg_AMTransferLockedGifts:                                    "EMsg_AMTransferLockedGifts",
	EMsg_AMTransferLockedGiftsResponse:                            "EMsg_AMTransferLockedGiftsResponse",
	EMsg_AMPlayerHostedOnGameServer:                               "EMsg_AMPlayerHostedOnGameServer",
	EMsg_AMGetAccountBanInfo:                                      "EMsg_AMGetAccountBanInfo",
	EMsg_AMGetAccountBanInfoResponse:                              "EMsg_AMGetAccountBanInfoResponse",
	EMsg_AMRecordBanEnforcement:                                   "EMsg_AMRecordBanEnforcement",
	EMsg_AMRollbackGiftTransfer:                                   "EMsg_AMRollbackGiftTransfer",
	EMsg_AMRollbackGiftTransferResponse:                           "EMsg_AMRollbackGiftTransferResponse",
	EMsg_AMHandlePendingTransaction:                               "EMsg_AMHandlePendingTransaction",
	EMsg_AMRequestClanDetails:                                     "EMsg_AMRequestClanDetails",
	EMsg_AMDeleteStoredPaypalAgreement:                            "EMsg_AMDeleteStoredPaypalAgreement",
	EMsg_AMGameServerUpdate:                                       "EMsg_AMGameServerUpdate",
	EMsg_AMGameServerRemove:                                       "EMsg_AMGameServerRemove",
	EMsg_AMGetPaypalAgreements:                                    "EMsg_AMGetPaypalAgreements",
	EMsg_AMGetPaypalAgreementsResponse:                            "EMsg_AMGetPaypalAgreementsResponse",
	EMsg_AMGameServerPlayerCompatibilityCheck:                     "EMsg_AMGameServerPlayerCompatibilityCheck",
	EMsg_AMGameServerPlayerCompatibilityCheckResponse:             "EMsg_AMGameServerPlayerCompatibilityCheckResponse",
	EMsg_AMRenewLicense:                                           "EMsg_AMRenewLicense",
	EMsg_AMGetAccountCommunityBanInfo:                             "EMsg_AMGetAccountCommunityBanInfo",
	EMsg_AMGetAccountCommunityBanInfoResponse:                     "EMsg_AMGetAccountCommunityBanInfoResponse",
	EMsg_AMGameServerAccountChangePassword:                        "EMsg_AMGameServerAccountChangePassword",
	EMsg_AMGameServerAccountDeleteAccount:                         "EMsg_AMGameServerAccountDeleteAccount",
	EMsg_AMRenewAgreement:                                         "EMsg_AMRenewAgreement",
	EMsg_AMSendEmail:                                              "EMsg_AMSendEmail",
	EMsg_AMXsollaPayment:                                          "EMsg_AMXsollaPayment",
	EMsg_AMXsollaPaymentResponse:                                  "EMsg_AMXsollaPaymentResponse",
	EMsg_AMAcctAllowedToPurchase:                                  "EMsg_AMAcctAllowedToPurchase",
	EMsg_AMAcctAllowedToPurchaseResponse:                          "EMsg_AMAcctAllowedToPurchaseResponse",
	EMsg_AMSwapKioskDeposit:                                       "EMsg_AMSwapKioskDeposit",
	EMsg_AMSwapKioskDepositResponse:                               "EMsg_AMSwapKioskDepositResponse",
	EMsg_AMSetUserGiftUnowned:                                     "EMsg_AMSetUserGiftUnowned",
	EMsg_AMSetUserGiftUnownedResponse:                             "EMsg_AMSetUserGiftUnownedResponse",
	EMsg_AMClaimUnownedUserGift:                                   "EMsg_AMClaimUnownedUserGift",
	EMsg_AMClaimUnownedUserGiftResponse:                           "EMsg_AMClaimUnownedUserGiftResponse",
	EMsg_AMSetClanName:                                            "EMsg_AMSetClanName",
	EMsg_AMSetClanNameResponse:                                    "EMsg_AMSetClanNameResponse",
	EMsg_AMGrantCoupon:                                            "EMsg_AMGrantCoupon",
	EMsg_AMGrantCouponResponse:                                    "EMsg_AMGrantCouponResponse",
	EMsg_AMIsPackageRestrictedInUserCountry:                       "EMsg_AMIsPackageRestrictedInUserCountry",
	EMsg_AMIsPackageRestrictedInUserCountryResponse:               "EMsg_AMIsPackageRestrictedInUserCountryResponse",
	EMsg_AMHandlePendingTransactionResponse:                       "EMsg_AMHandlePendingTransactionResponse",
	EMsg_AMGrantGuestPasses2:                                      "EMsg_AMGrantGuestPasses2",
	EMsg_AMGrantGuestPasses2Response:                              "EMsg_AMGrantGuestPasses2Response",
	EMsg_AMSessionQuery:                                           "EMsg_AMSessionQuery",
	EMsg_AMSessionQueryResponse:                                   "EMsg_AMSessionQueryResponse",
	EMsg_AMGetPlayerBanDetails:                                    "EMsg_AMGetPlayerBanDetails",
	EMsg_AMGetPlayerBanDetailsResponse:                            "EMsg_AMGetPlayerBanDetailsResponse",
	EMsg_AMFinalizePurchase:                                       "EMsg_AMFinalizePurchase",
	EMsg_AMFinalizePurchaseResponse:                               "EMsg_AMFinalizePurchaseResponse",
	EMsg_AMPersonaChangeResponse:                                  "EMsg_AMPersonaChangeResponse",
	EMsg_AMGetClanDetailsForForumCreation:                         "EMsg_AMGetClanDetailsForForumCreation",
	EMsg_AMGetClanDetailsForForumCreationResponse:                 "EMsg_AMGetClanDetailsForForumCreationResponse",
	EMsg_AMGetPendingNotificationCount:                            "EMsg_AMGetPendingNotificationCount",
	EMsg_AMGetPendingNotificationCountResponse:                    "EMsg_AMGetPendingNotificationCountResponse",
	EMsg_AMPasswordHashUpgrade:                                    "EMsg_AMPasswordHashUpgrade",
	EMsg_AMMoPayPayment:                                           "EMsg_AMMoPayPayment",
	EMsg_AMMoPayPaymentResponse:                                   "EMsg_AMMoPayPaymentResponse",
	EMsg_AMBoaCompraPayment:                                       "EMsg_AMBoaCompraPayment",
	EMsg_AMBoaCompraPaymentResponse:                               "EMsg_AMBoaCompraPaymentResponse",
	EMsg_AMExpireCaptchaByGID:                                     "EMsg_AMExpireCaptchaByGID",
	EMsg_AMCompleteExternalPurchase:                               "EMsg_AMCompleteExternalPurchase",
	EMsg_AMCompleteExternalPurchaseResponse:                       "EMsg_AMCompleteExternalPurchaseResponse",
	EMsg_AMResolveNegativeWalletCredits:                           "EMsg_AMResolveNegativeWalletCredits",
	EMsg_AMResolveNegativeWalletCreditsResponse:                   "EMsg_AMResolveNegativeWalletCreditsResponse",
	EMsg_AMPayelpPayment:                                          "EMsg_AMPayelpPayment",
	EMsg_AMPayelpPaymentResponse:                                  "EMsg_AMPayelpPaymentResponse",
	EMsg_AMPlayerGetClanBasicDetails:                              "EMsg_AMPlayerGetClanBasicDetails",
	EMsg_AMPlayerGetClanBasicDetailsResponse:                      "EMsg_AMPlayerGetClanBasicDetailsResponse",
	EMsg_AMMOLPayment:                                             "EMsg_AMMOLPayment",
	EMsg_AMMOLPaymentResponse:                                     "EMsg_AMMOLPaymentResponse",
	EMsg_GetUserIPCountry:                                         "EMsg_GetUserIPCountry",
	EMsg_GetUserIPCountryResponse:                                 "EMsg_GetUserIPCountryResponse",
	EMsg_NotificationOfSuspiciousActivity:                         "EMsg_NotificationOfSuspiciousActivity",
	EMsg_AMDegicaPayment:                                          "EMsg_AMDegicaPayment",
	EMsg_AMDegicaPaymentResponse:                                  "EMsg_AMDegicaPaymentResponse",
	EMsg_AMEClubPayment:                                           "EMsg_AMEClubPayment",
	EMsg_AMEClubPaymentResponse:                                   "EMsg_AMEClubPaymentResponse",
	EMsg_AMPayPalPaymentsHubPayment:                               "EMsg_AMPayPalPaymentsHubPayment",
	EMsg_AMPayPalPaymentsHubPaymentResponse:                       "EMsg_AMPayPalPaymentsHubPaymentResponse",
	EMsg_AMTwoFactorRecoverAuthenticatorRequest:                   "EMsg_AMTwoFactorRecoverAuthenticatorRequest",
	EMsg_AMTwoFactorRecoverAuthenticatorResponse:                  "EMsg_AMTwoFactorRecoverAuthenticatorResponse",
	EMsg_AMSmart2PayPayment:                                       "EMsg_AMSmart2PayPayment",
	EMsg_AMSmart2PayPaymentResponse:                               "EMsg_AMSmart2PayPaymentResponse",
	EMsg_AMValidatePasswordResetCodeAndSendSmsRequest:             "EMsg_AMValidatePasswordResetCodeAndSendSmsRequest",
	EMsg_AMValidatePasswordResetCodeAndSendSmsResponse:            "EMsg_AMValidatePasswordResetCodeAndSendSmsResponse",
	EMsg_AMGetAccountResetDetailsRequest:                          "EMsg_AMGetAccountResetDetailsRequest",
	EMsg_AMGetAccountResetDetailsResponse:                         "EMsg_AMGetAccountResetDetailsResponse",
	EMsg_AMBitPayPayment:                                          "EMsg_AMBitPayPayment",
	EMsg_AMBitPayPaymentResponse:                                  "EMsg_AMBitPayPaymentResponse",
	EMsg_AMSendAccountInfoUpdate:                                  "EMsg_AMSendAccountInfoUpdate",
	EMsg_BasePSRange:                                              "EMsg_BasePSRange",
	EMsg_PSCreateShoppingCart:                                     "EMsg_PSCreateShoppingCart",
	EMsg_PSCreateShoppingCartResponse:                             "EMsg_PSCreateShoppingCartResponse",
	EMsg_PSIsValidShoppingCart:                                    "EMsg_PSIsValidShoppingCart",
	EMsg_PSIsValidShoppingCartResponse:                            "EMsg_PSIsValidShoppingCartResponse",
	EMsg_PSAddPackageToShoppingCart:                               "EMsg_PSAddPackageToShoppingCart",
	EMsg_PSAddPackageToShoppingCartResponse:                       "EMsg_PSAddPackageToShoppingCartResponse",
	EMsg_PSRemoveLineItemFromShoppingCart:                         "EMsg_PSRemoveLineItemFromShoppingCart",
	EMsg_PSRemoveLineItemFromShoppingCartResponse:                 "EMsg_PSRemoveLineItemFromShoppingCartResponse",
	EMsg_PSGetShoppingCartContents:                                "EMsg_PSGetShoppingCartContents",
	EMsg_PSGetShoppingCartContentsResponse:                        "EMsg_PSGetShoppingCartContentsResponse",
	EMsg_PSAddWalletCreditToShoppingCart:                          "EMsg_PSAddWalletCreditToShoppingCart",
	EMsg_PSAddWalletCreditToShoppingCartResponse:                  "EMsg_PSAddWalletCreditToShoppingCartResponse",
	EMsg_BaseUFSRange:                                             "EMsg_BaseUFSRange",
	EMsg_ClientUFSUploadFileRequest:                               "EMsg_ClientUFSUploadFileRequest",
	EMsg_ClientUFSUploadFileResponse:                              "EMsg_ClientUFSUploadFileResponse",
	EMsg_ClientUFSUploadFileChunk:                                 "EMsg_ClientUFSUploadFileChunk",
	EMsg_ClientUFSUploadFileFinished:                              "EMsg_ClientUFSUploadFileFinished",
	EMsg_ClientUFSGetFileListForApp:                               "EMsg_ClientUFSGetFileListForApp",
	EMsg_ClientUFSGetFileListForAppResponse:                       "EMsg_ClientUFSGetFileListForAppResponse",
	EMsg_ClientUFSDownloadRequest:                                 "EMsg_ClientUFSDownloadRequest",
	EMsg_ClientUFSDownloadResponse:                                "EMsg_ClientUFSDownloadResponse",
	EMsg_ClientUFSDownloadChunk:                                   "EMsg_ClientUFSDownloadChunk",
	EMsg_ClientUFSLoginRequest:                                    "EMsg_ClientUFSLoginRequest",
	EMsg_ClientUFSLoginResponse:                                   "EMsg_ClientUFSLoginResponse",
	EMsg_UFSReloadPartitionInfo:                                   "EMsg_UFSReloadPartitionInfo",
	EMsg_ClientUFSTransferHeartbeat:                               "EMsg_ClientUFSTransferHeartbeat",
	EMsg_UFSSynchronizeFile:                                       "EMsg_UFSSynchronizeFile",
	EMsg_UFSSynchronizeFileResponse:                               "EMsg_UFSSynchronizeFileResponse",
	EMsg_ClientUFSDeleteFileRequest:                               "EMsg_ClientUFSDeleteFileRequest",
	EMsg_ClientUFSDeleteFileResponse:                              "EMsg_ClientUFSDeleteFileResponse",
	EMsg_UFSDownloadRequest:                                       "EMsg_UFSDownloadRequest",
	EMsg_UFSDownloadResponse:                                      "EMsg_UFSDownloadResponse",
	EMsg_UFSDownloadChunk:                                         "EMsg_UFSDownloadChunk",
	EMsg_ClientUFSGetUGCDetails:                                   "EMsg_ClientUFSGetUGCDetails",
	EMsg_ClientUFSGetUGCDetailsResponse:                           "EMsg_ClientUFSGetUGCDetailsResponse",
	EMsg_UFSUpdateFileFlags:                                       "EMsg_UFSUpdateFileFlags",
	EMsg_UFSUpdateFileFlagsResponse:                               "EMsg_UFSUpdateFileFlagsResponse",
	EMsg_ClientUFSGetSingleFileInfo:                               "EMsg_ClientUFSGetSingleFileInfo",
	EMsg_ClientUFSGetSingleFileInfoResponse:                       "EMsg_ClientUFSGetSingleFileInfoResponse",
	EMsg_ClientUFSShareFile:                                       "EMsg_ClientUFSShareFile",
	EMsg_ClientUFSShareFileResponse:                               "EMsg_ClientUFSShareFileResponse",
	EMsg_UFSReloadAccount:                                         "EMsg_UFSReloadAccount",
	EMsg_UFSReloadAccountResponse:                                 "EMsg_UFSReloadAccountResponse",
	EMsg_UFSUpdateRecordBatched:                                   "EMsg_UFSUpdateRecordBatched",
	EMsg_UFSUpdateRecordBatchedResponse:                           "EMsg_UFSUpdateRecordBatchedResponse",
	EMsg_UFSMigrateFile:                                           "EMsg_UFSMigrateFile",
	EMsg_UFSMigrateFileResponse:                                   "EMsg_UFSMigrateFileResponse",
	EMsg_UFSGetUGCURLs:                                            "EMsg_UFSGetUGCURLs",
	EMsg_UFSGetUGCURLsResponse:                                    "EMsg_UFSGetUGCURLsResponse",
	EMsg_UFSHttpUploadFileFinishRequest:                           "EMsg_UFSHttpUploadFileFinishRequest",
	EMsg_UFSHttpUploadFileFinishResponse:                          "EMsg_UFSHttpUploadFileFinishResponse",
	EMsg_UFSDownloadStartRequest:                                  "EMsg_UFSDownloadStartRequest",
	EMsg_UFSDownloadStartResponse:                                 "EMsg_UFSDownloadStartResponse",
	EMsg_UFSDownloadChunkRequest:                                  "EMsg_UFSDownloadChunkRequest",
	EMsg_UFSDownloadChunkResponse:                                 "EMsg_UFSDownloadChunkResponse",
	EMsg_UFSDownloadFinishRequest:                                 "EMsg_UFSDownloadFinishRequest",
	EMsg_UFSDownloadFinishResponse:                                "EMsg_UFSDownloadFinishResponse",
	EMsg_UFSFlushURLCache:                                         "EMsg_UFSFlushURLCache",
	EMsg_UFSUploadCommit:                                          "EMsg_UFSUploadCommit",
	EMsg_UFSUploadCommitResponse:                                  "EMsg_UFSUploadCommitResponse",
	EMsg_UFSMigrateFileAppID:                                      "EMsg_UFSMigrateFileAppID",
	EMsg_UFSMigrateFileAppIDResponse:                              "EMsg_UFSMigrateFileAppIDResponse",
	EMsg_BaseClient2:                                              "EMsg_BaseClient2",
	EMsg_ClientRequestForgottenPasswordEmail:                      "EMsg_ClientRequestForgottenPasswordEmail",
	EMsg_ClientRequestForgottenPasswordEmailResponse:              "EMsg_ClientRequestForgottenPasswordEmailResponse",
	EMsg_ClientCreateAccountResponse:                              "EMsg_ClientCreateAccountResponse",
	EMsg_ClientResetForgottenPassword:                             "EMsg_ClientResetForgottenPassword",
	EMsg_ClientResetForgottenPasswordResponse:                     "EMsg_ClientResetForgottenPasswordResponse",
	EMsg_ClientCreateAccount2:                                     "EMsg_ClientCreateAccount2",
	EMsg_ClientInformOfResetForgottenPassword:                     "EMsg_ClientInformOfResetForgottenPassword",
	EMsg_ClientInformOfResetForgottenPasswordResponse:             "EMsg_ClientInformOfResetForgottenPasswordResponse",
	EMsg_ClientAnonUserLogOn_Deprecated:                           "EMsg_ClientAnonUserLogOn_Deprecated",
	EMsg_ClientGamesPlayedWithDataBlob:                            "EMsg_ClientGamesPlayedWithDataBlob",
	EMsg_ClientUpdateUserGameInfo:                                 "EMsg_ClientUpdateUserGameInfo",
	EMsg_ClientFileToDownload:                                     "EMsg_ClientFileToDownload",
	EMsg_ClientFileToDownloadResponse:                             "EMsg_ClientFileToDownloadResponse",
	EMsg_ClientLBSSetScore:                                        "EMsg_ClientLBSSetScore",
	EMsg_ClientLBSSetScoreResponse:                                "EMsg_ClientLBSSetScoreResponse",
	EMsg_ClientLBSFindOrCreateLB:                                  "EMsg_ClientLBSFindOrCreateLB",
	EMsg_ClientLBSFindOrCreateLBResponse:                          "EMsg_ClientLBSFindOrCreateLBResponse",
	EMsg_ClientLBSGetLBEntries:                                    "EMsg_ClientLBSGetLBEntries",
	EMsg_ClientLBSGetLBEntriesResponse:                            "EMsg_ClientLBSGetLBEntriesResponse",
	EMsg_ClientMarketingMessageUpdate:                             "EMsg_ClientMarketingMessageUpdate",
	EMsg_ClientChatDeclined:                                       "EMsg_ClientChatDeclined",
	EMsg_ClientFriendMsgIncoming:                                  "EMsg_ClientFriendMsgIncoming",
	EMsg_ClientAuthList_Deprecated:                                "EMsg_ClientAuthList_Deprecated",
	EMsg_ClientTicketAuthComplete:                                 "EMsg_ClientTicketAuthComplete",
	EMsg_ClientIsLimitedAccount:                                   "EMsg_ClientIsLimitedAccount",
	EMsg_ClientRequestAuthList:                                    "EMsg_ClientRequestAuthList",
	EMsg_ClientAuthList:                                           "EMsg_ClientAuthList",
	EMsg_ClientStat:                                               "EMsg_ClientStat",
	EMsg_ClientP2PConnectionInfo:                                  "EMsg_ClientP2PConnectionInfo",
	EMsg_ClientP2PConnectionFailInfo:                              "EMsg_ClientP2PConnectionFailInfo",
	EMsg_ClientGetNumberOfCurrentPlayers:                          "EMsg_ClientGetNumberOfCurrentPlayers",
	EMsg_ClientGetNumberOfCurrentPlayersResponse:                  "EMsg_ClientGetNumberOfCurrentPlayersResponse",
	EMsg_ClientGetDepotDecryptionKey:                              "EMsg_ClientGetDepotDecryptionKey",
	EMsg_ClientGetDepotDecryptionKeyResponse:                      "EMsg_ClientGetDepotDecryptionKeyResponse",
	EMsg_GSPerformHardwareSurvey:                                  "EMsg_GSPerformHardwareSurvey",
	EMsg_ClientGetAppBetaPasswords:                                "EMsg_ClientGetAppBetaPasswords",
	EMsg_ClientGetAppBetaPasswordsResponse:                        "EMsg_ClientGetAppBetaPasswordsResponse",
	EMsg_ClientEnableTestLicense:                                  "EMsg_ClientEnableTestLicense",
	EMsg_ClientEnableTestLicenseResponse:                          "EMsg_ClientEnableTestLicenseResponse",
	EMsg_ClientDisableTestLicense:                                 "EMsg_ClientDisableTestLicense",
	EMsg_ClientDisableTestLicenseResponse:                         "EMsg_ClientDisableTestLicenseResponse",
	EMsg_ClientRequestValidationMail:                              "EMsg_ClientRequestValidationMail",
	EMsg_ClientRequestValidationMailResponse:                      "EMsg_ClientRequestValidationMailResponse",
	EMsg_ClientCheckAppBetaPassword:                               "EMsg_ClientCheckAppBetaPassword",
	EMsg_ClientCheckAppBetaPasswordResponse:                       "EMsg_ClientCheckAppBetaPasswordResponse",
	EMsg_ClientToGC:                                               "EMsg_ClientToGC",
	EMsg_ClientFromGC:                                             "EMsg_ClientFromGC",
	EMsg_ClientRequestChangeMail:                                  "EMsg_ClientRequestChangeMail",
	EMsg_ClientRequestChangeMailResponse:                          "EMsg_ClientRequestChangeMailResponse",
	EMsg_ClientEmailAddrInfo:                                      "EMsg_ClientEmailAddrInfo",
	EMsg_ClientPasswordChange3:                                    "EMsg_ClientPasswordChange3",
	EMsg_ClientEmailChange3:                                       "EMsg_ClientEmailChange3",
	EMsg_ClientPersonalQAChange3:                                  "EMsg_ClientPersonalQAChange3",
	EMsg_ClientResetForgottenPassword3:                            "EMsg_ClientResetForgottenPassword3",
	EMsg_ClientRequestForgottenPasswordEmail3:                     "EMsg_ClientRequestForgottenPasswordEmail3",
	EMsg_ClientCreateAccount3:                                     "EMsg_ClientCreateAccount3",
	EMsg_ClientNewLoginKey:                                        "EMsg_ClientNewLoginKey",
	EMsg_ClientNewLoginKeyAccepted:                                "EMsg_ClientNewLoginKeyAccepted",
	EMsg_ClientLogOnWithHash_Deprecated:                           "EMsg_ClientLogOnWithHash_Deprecated",
	EMsg_ClientStoreUserStats2:                                    "EMsg_ClientStoreUserStats2",
	EMsg_ClientStatsUpdated:                                       "EMsg_ClientStatsUpdated",
	EMsg_ClientActivateOEMLicense:                                 "EMsg_ClientActivateOEMLicense",
	EMsg_ClientRegisterOEMMachine:                                 "EMsg_ClientRegisterOEMMachine",
	EMsg_ClientRegisterOEMMachineResponse:                         "EMsg_ClientRegisterOEMMachineResponse",
	EMsg_ClientRequestedClientStats:                               "EMsg_ClientRequestedClientStats",
	EMsg_ClientStat2Int32:                                         "EMsg_ClientStat2Int32",
	EMsg_ClientStat2:                                              "EMsg_ClientStat2",
	EMsg_ClientVerifyPassword:                                     "EMsg_ClientVerifyPassword",
	EMsg_ClientVerifyPasswordResponse:                             "EMsg_ClientVerifyPasswordResponse",
	EMsg_ClientDRMDownloadRequest:                                 "EMsg_ClientDRMDownloadRequest",
	EMsg_ClientDRMDownloadResponse:                                "EMsg_ClientDRMDownloadResponse",
	EMsg_ClientDRMFinalResult:                                     "EMsg_ClientDRMFinalResult",
	EMsg_ClientGetFriendsWhoPlayGame:                              "EMsg_ClientGetFriendsWhoPlayGame",
	EMsg_ClientGetFriendsWhoPlayGameResponse:                      "EMsg_ClientGetFriendsWhoPlayGameResponse",
	EMsg_ClientOGSBeginSession:                                    "EMsg_ClientOGSBeginSession",
	EMsg_ClientOGSBeginSessionResponse:                            "EMsg_ClientOGSBeginSessionResponse",
	EMsg_ClientOGSEndSession:                                      "EMsg_ClientOGSEndSession",
	EMsg_ClientOGSEndSessionResponse:                              "EMsg_ClientOGSEndSessionResponse",
	EMsg_ClientOGSWriteRow:                                        "EMsg_ClientOGSWriteRow",
	EMsg_ClientDRMTest:                                            "EMsg_ClientDRMTest",
	EMsg_ClientDRMTestResult:                                      "EMsg_ClientDRMTestResult",
	EMsg_ClientServerUnavailable:                                  "EMsg_ClientServerUnavailable",
	EMsg_ClientServersAvailable:                                   "EMsg_ClientServersAvailable",
	EMsg_ClientRegisterAuthTicketWithCM:                           "EMsg_ClientRegisterAuthTicketWithCM",
	EMsg_ClientGCMsgFailed:                                        "EMsg_ClientGCMsgFailed",
	EMsg_ClientMicroTxnAuthRequest:                                "EMsg_ClientMicroTxnAuthRequest",
	EMsg_ClientMicroTxnAuthorize:                                  "EMsg_ClientMicroTxnAuthorize",
	EMsg_ClientMicroTxnAuthorizeResponse:                          "EMsg_ClientMicroTxnAuthorizeResponse",
	EMsg_ClientAppMinutesPlayedData:                               "EMsg_ClientAppMinutesPlayedData",
	EMsg_ClientGetMicroTxnInfo:                                    "EMsg_ClientGetMicroTxnInfo",
	EMsg_ClientGetMicroTxnInfoResponse:                            "EMsg_ClientGetMicroTxnInfoResponse",
	EMsg_ClientMarketingMessageUpdate2:                            "EMsg_ClientMarketingMessageUpdate2",
	EMsg_ClientDeregisterWithServer:                               "EMsg_ClientDeregisterWithServer",
	EMsg_ClientSubscribeToPersonaFeed:                             "EMsg_ClientSubscribeToPersonaFeed",
	EMsg_ClientLogon:                                              "EMsg_ClientLogon",
	EMsg_ClientGetClientDetails:                                   "EMsg_ClientGetClientDetails",
	EMsg_ClientGetClientDetailsResponse:                           "EMsg_ClientGetClientDetailsResponse",
	EMsg_ClientReportOverlayDetourFailure:                         "EMsg_ClientReportOverlayDetourFailure",
	EMsg_ClientGetClientAppList:                                   "EMsg_ClientGetClientAppList",
	EMsg_ClientGetClientAppListResponse:                           "EMsg_ClientGetClientAppListResponse",
	EMsg_ClientInstallClientApp:                                   "EMsg_ClientInstallClientApp",
	EMsg_ClientInstallClientAppResponse:                           "EMsg_ClientInstallClientAppResponse",
	EMsg_ClientUninstallClientApp:                                 "EMsg_ClientUninstallClientApp",
	EMsg_ClientUninstallClientAppResponse:                         "EMsg_ClientUninstallClientAppResponse",
	EMsg_ClientSetClientAppUpdateState:                            "EMsg_ClientSetClientAppUpdateState",
	EMsg_ClientSetClientAppUpdateStateResponse:                    "EMsg_ClientSetClientAppUpdateStateResponse",
	EMsg_ClientRequestEncryptedAppTicket:                          "EMsg_ClientRequestEncryptedAppTicket",
	EMsg_ClientRequestEncryptedAppTicketResponse:                  "EMsg_ClientRequestEncryptedAppTicketResponse",
	EMsg_ClientWalletInfoUpdate:                                   "EMsg_ClientWalletInfoUpdate",
	EMsg_ClientLBSSetUGC:                                          "EMsg_ClientLBSSetUGC",
	EMsg_ClientLBSSetUGCResponse:                                  "EMsg_ClientLBSSetUGCResponse",
	EMsg_ClientAMGetClanOfficers:                                  "EMsg_ClientAMGetClanOfficers",
	EMsg_ClientAMGetClanOfficersResponse:                          "EMsg_ClientAMGetClanOfficersResponse",
	EMsg_ClientCheckFileSignature:                                 "EMsg_ClientCheckFileSignature",
	EMsg_ClientCheckFileSignatureResponse:                         "EMsg_ClientCheckFileSignatureResponse",
	EMsg_ClientFriendProfileInfo:                                  "EMsg_ClientFriendProfileInfo",
	EMsg_ClientFriendProfileInfoResponse:                          "EMsg_ClientFriendProfileInfoResponse",
	EMsg_ClientUpdateMachineAuth:                                  "EMsg_ClientUpdateMachineAuth",
	EMsg_ClientUpdateMachineAuthResponse:                          "EMsg_ClientUpdateMachineAuthResponse",
	EMsg_ClientReadMachineAuth:                                    "EMsg_ClientReadMachineAuth",
	EMsg_ClientReadMachineAuthResponse:                            "EMsg_ClientReadMachineAuthResponse",
	EMsg_ClientRequestMachineAuth:                                 "EMsg_ClientRequestMachineAuth",
	EMsg_ClientRequestMachineAuthResponse:                         "EMsg_ClientRequestMachineAuthResponse",
	EMsg_ClientScreenshotsChanged:                                 "EMsg_ClientScreenshotsChanged",
	EMsg_ClientEmailChange4:                                       "EMsg_ClientEmailChange4",
	EMsg_ClientEmailChangeResponse4:                               "EMsg_ClientEmailChangeResponse4",
	EMsg_ClientGetCDNAuthToken:                                    "EMsg_ClientGetCDNAuthToken",
	EMsg_ClientGetCDNAuthTokenResponse:                            "EMsg_ClientGetCDNAuthTokenResponse",
	EMsg_ClientDownloadRateStatistics:                             "EMsg_ClientDownloadRateStatistics",
	EMsg_ClientRequestAccountData:                                 "EMsg_ClientRequestAccountData",
	EMsg_ClientRequestAccountDataResponse:                         "EMsg_ClientRequestAccountDataResponse",
	EMsg_ClientResetForgottenPassword4:                            "EMsg_ClientResetForgottenPassword4",
	EMsg_ClientHideFriend:                                         "EMsg_ClientHideFriend",
	EMsg_ClientFriendsGroupsList:                                  "EMsg_ClientFriendsGroupsList",
	EMsg_ClientGetClanActivityCounts:                              "EMsg_ClientGetClanActivityCounts",
	EMsg_ClientGetClanActivityCountsResponse:                      "EMsg_ClientGetClanActivityCountsResponse",
	EMsg_ClientOGSReportString:                                    "EMsg_ClientOGSReportString",
	EMsg_ClientOGSReportBug:                                       "EMsg_ClientOGSReportBug",
	EMsg_ClientSentLogs:                                           "EMsg_ClientSentLogs",
	EMsg_ClientLogonGameServer:                                    "EMsg_ClientLogonGameServer",
	EMsg_AMClientCreateFriendsGroup:                               "EMsg_AMClientCreateFriendsGroup",
	EMsg_AMClientCreateFriendsGroupResponse:                       "EMsg_AMClientCreateFriendsGroupResponse",
	EMsg_AMClientDeleteFriendsGroup:                               "EMsg_AMClientDeleteFriendsGroup",
	EMsg_AMClientDeleteFriendsGroupResponse:                       "EMsg_AMClientDeleteFriendsGroupResponse",
	EMsg_AMClientRenameFriendsGroup:                               "EMsg_AMClientRenameFriendsGroup",
	EMsg_AMClientRenameFriendsGroupResponse:                       "EMsg_AMClientRenameFriendsGroupResponse",
	EMsg_AMClientAddFriendToGroup:                                 "EMsg_AMClientAddFriendToGroup",
	EMsg_AMClientAddFriendToGroupResponse:                         "EMsg_AMClientAddFriendToGroupResponse",
	EMsg_AMClientRemoveFriendFromGroup:                            "EMsg_AMClientRemoveFriendFromGroup",
	EMsg_AMClientRemoveFriendFromGroupResponse:                    "EMsg_AMClientRemoveFriendFromGroupResponse",
	EMsg_ClientAMGetPersonaNameHistory:                            "EMsg_ClientAMGetPersonaNameHistory",
	EMsg_ClientAMGetPersonaNameHistoryResponse:                    "EMsg_ClientAMGetPersonaNameHistoryResponse",
	EMsg_ClientRequestFreeLicense:                                 "EMsg_ClientRequestFreeLicense",
	EMsg_ClientRequestFreeLicenseResponse:                         "EMsg_ClientRequestFreeLicenseResponse",
	EMsg_ClientDRMDownloadRequestWithCrashData:                    "EMsg_ClientDRMDownloadRequestWithCrashData",
	EMsg_ClientAuthListAck:                                        "EMsg_ClientAuthListAck",
	EMsg_ClientItemAnnouncements:                                  "EMsg_ClientItemAnnouncements",
	EMsg_ClientRequestItemAnnouncements:                           "EMsg_ClientRequestItemAnnouncements",
	EMsg_ClientFriendMsgEchoToSender:                              "EMsg_ClientFriendMsgEchoToSender",
	EMsg_ClientChangeSteamGuardOptions:                            "EMsg_ClientChangeSteamGuardOptions",
	EMsg_ClientChangeSteamGuardOptionsResponse:                    "EMsg_ClientChangeSteamGuardOptionsResponse",
	EMsg_ClientOGSGameServerPingSample:                            "EMsg_ClientOGSGameServerPingSample",
	EMsg_ClientCommentNotifications:                               "EMsg_ClientCommentNotifications",
	EMsg_ClientRequestCommentNotifications:                        "EMsg_ClientRequestCommentNotifications",
	EMsg_ClientPersonaChangeResponse:                              "EMsg_ClientPersonaChangeResponse",
	EMsg_ClientRequestWebAPIAuthenticateUserNonce:                 "EMsg_ClientRequestWebAPIAuthenticateUserNonce",
	EMsg_ClientRequestWebAPIAuthenticateUserNonceResponse:         "EMsg_ClientRequestWebAPIAuthenticateUserNonceResponse",
	EMsg_ClientPlayerNicknameList:                                 "EMsg_ClientPlayerNicknameList",
	EMsg_AMClientSetPlayerNickname:                                "EMsg_AMClientSetPlayerNickname",
	EMsg_AMClientSetPlayerNicknameResponse:                        "EMsg_AMClientSetPlayerNicknameResponse",
	EMsg_ClientCreateAccountProto:                                 "EMsg_ClientCreateAccountProto",
	EMsg_ClientCreateAccountProtoResponse:                         "EMsg_ClientCreateAccountProtoResponse",
	EMsg_ClientGetNumberOfCurrentPlayersDP:                        "EMsg_ClientGetNumberOfCurrentPlayersDP",
	EMsg_ClientGetNumberOfCurrentPlayersDPResponse:                "EMsg_ClientGetNumberOfCurrentPlayersDPResponse",
	EMsg_ClientServiceMethod:                                      "EMsg_ClientServiceMethod",
	EMsg_ClientServiceMethodResponse:                              "EMsg_ClientServiceMethodResponse",
	EMsg_ClientFriendUserStatusPublished:                          "EMsg_ClientFriendUserStatusPublished",
	EMsg_ClientCurrentUIMode:                                      "EMsg_ClientCurrentUIMode",
	EMsg_ClientVanityURLChangedNotification:                       "EMsg_ClientVanityURLChangedNotification",
	EMsg_ClientUserNotifications:                                  "EMsg_ClientUserNotifications",
	EMsg_BaseDFS:                                                  "EMsg_BaseDFS",
	EMsg_DFSGetFile:                                               "EMsg_DFSGetFile",
	EMsg_DFSInstallLocalFile:                                      "EMsg_DFSInstallLocalFile",
	EMsg_DFSConnection:                                            "EMsg_DFSConnection",
	EMsg_DFSConnectionReply:                                       "EMsg_DFSConnectionReply",
	EMsg_ClientDFSAuthenticateRequest:                             "EMsg_ClientDFSAuthenticateRequest",
	EMsg_ClientDFSAuthenticateResponse:                            "EMsg_ClientDFSAuthenticateResponse",
	EMsg_ClientDFSEndSession:                                      "EMsg_ClientDFSEndSession",
	EMsg_DFSPurgeFile:                                             "EMsg_DFSPurgeFile",
	EMsg_DFSRouteFile:                                             "EMsg_DFSRouteFile",
	EMsg_DFSGetFileFromServer:                                     "EMsg_DFSGetFileFromServer",
	EMsg_DFSAcceptedResponse:                                      "EMsg_DFSAcceptedResponse",
	EMsg_DFSRequestPingback:                                       "EMsg_DFSRequestPingback",
	EMsg_DFSRecvTransmitFile:                                      "EMsg_DFSRecvTransmitFile",
	EMsg_DFSSendTransmitFile:                                      "EMsg_DFSSendTransmitFile",
	EMsg_DFSRequestPingback2:                                      "EMsg_DFSRequestPingback2",
	EMsg_DFSResponsePingback2:                                     "EMsg_DFSResponsePingback2",
	EMsg_ClientDFSDownloadStatus:                                  "EMsg_ClientDFSDownloadStatus",
	EMsg_DFSStartTransfer:                                         "EMsg_DFSStartTransfer",
	EMsg_DFSTransferComplete:                                      "EMsg_DFSTransferComplete",
	EMsg_DFSRouteFileResponse:                                     "EMsg_DFSRouteFileResponse",
	EMsg_ClientNetworkingCertRequest:                              "EMsg_ClientNetworkingCertRequest",
	EMsg_ClientNetworkingCertRequestResponse:                      "EMsg_ClientNetworkingCertRequestResponse",
	EMsg_BaseMDS:                                                  "EMsg_BaseMDS",
	EMsg_ClientMDSLoginRequest:                                    "EMsg_ClientMDSLoginRequest",
	EMsg_ClientMDSLoginResponse:                                   "EMsg_ClientMDSLoginResponse",
	EMsg_ClientMDSUploadManifestRequest:                           "EMsg_ClientMDSUploadManifestRequest",
	EMsg_ClientMDSUploadManifestResponse:                          "EMsg_ClientMDSUploadManifestResponse",
	EMsg_ClientMDSTransmitManifestDataChunk:                       "EMsg_ClientMDSTransmitManifestDataChunk",
	EMsg_ClientMDSHeartbeat:                                       "EMsg_ClientMDSHeartbeat",
	EMsg_ClientMDSUploadDepotChunks:                               "EMsg_ClientMDSUploadDepotChunks",
	EMsg_ClientMDSUploadDepotChunksResponse:                       "EMsg_ClientMDSUploadDepotChunksResponse",
	EMsg_ClientMDSInitDepotBuildRequest:                           "EMsg_ClientMDSInitDepotBuildRequest",
	EMsg_ClientMDSInitDepotBuildResponse:                          "EMsg_ClientMDSInitDepotBuildResponse",
	EMsg_AMToMDSGetDepotDecryptionKey:                             "EMsg_AMToMDSGetDepotDecryptionKey",
	EMsg_MDSToAMGetDepotDecryptionKeyResponse:                     "EMsg_MDSToAMGetDepotDecryptionKeyResponse",
	EMsg_MDSGetVersionsForDepot:                                   "EMsg_MDSGetVersionsForDepot",
	EMsg_MDSGetVersionsForDepotResponse:                           "EMsg_MDSGetVersionsForDepotResponse",
	EMsg_MDSSetPublicVersionForDepot:                              "EMsg_MDSSetPublicVersionForDepot",
	EMsg_MDSSetPublicVersionForDepotResponse:                      "EMsg_MDSSetPublicVersionForDepotResponse",
	EMsg_ClientMDSGetDepotManifest:                                "EMsg_ClientMDSGetDepotManifest",
	EMsg_ClientMDSGetDepotManifestResponse:                        "EMsg_ClientMDSGetDepotManifestResponse",
	EMsg_ClientMDSGetDepotManifestChunk:                           "EMsg_ClientMDSGetDepotManifestChunk",
	EMsg_ClientMDSUploadRateTest:                                  "EMsg_ClientMDSUploadRateTest",
	EMsg_ClientMDSUploadRateTestResponse:                          "EMsg_ClientMDSUploadRateTestResponse",
	EMsg_MDSDownloadDepotChunksAck:                                "EMsg_MDSDownloadDepotChunksAck",
	EMsg_MDSContentServerStatsBroadcast:                           "EMsg_MDSContentServerStatsBroadcast",
	EMsg_MDSContentServerConfigRequest:                            "EMsg_MDSContentServerConfigRequest",
	EMsg_MDSContentServerConfig:                                   "EMsg_MDSContentServerConfig",
	EMsg_MDSGetDepotManifest:                                      "EMsg_MDSGetDepotManifest",
	EMsg_MDSGetDepotManifestResponse:                              "EMsg_MDSGetDepotManifestResponse",
	EMsg_MDSGetDepotManifestChunk:                                 "EMsg_MDSGetDepotManifestChunk",
	EMsg_MDSGetDepotChunk:                                         "EMsg_MDSGetDepotChunk",
	EMsg_MDSGetDepotChunkResponse:                                 "EMsg_MDSGetDepotChunkResponse",
	EMsg_MDSGetDepotChunkChunk:                                    "EMsg_MDSGetDepotChunkChunk",
	EMsg_MDSUpdateContentServerConfig:                             "EMsg_MDSUpdateContentServerConfig",
	EMsg_MDSGetServerListForUser:                                  "EMsg_MDSGetServerListForUser",
	EMsg_MDSGetServerListForUserResponse:                          "EMsg_MDSGetServerListForUserResponse",
	EMsg_ClientMDSRegisterAppBuild:                                "EMsg_ClientMDSRegisterAppBuild",
	EMsg_ClientMDSRegisterAppBuildResponse:                        "EMsg_ClientMDSRegisterAppBuildResponse",
	EMsg_ClientMDSSetAppBuildLive:                                 "EMsg_ClientMDSSetAppBuildLive",
	EMsg_ClientMDSSetAppBuildLiveResponse:                         "EMsg_ClientMDSSetAppBuildLiveResponse",
	EMsg_ClientMDSGetPrevDepotBuild:                               "EMsg_ClientMDSGetPrevDepotBuild",
	EMsg_ClientMDSGetPrevDepotBuildResponse:                       "EMsg_ClientMDSGetPrevDepotBuildResponse",
	EMsg_MDSToCSFlushChunk:                                        "EMsg_MDSToCSFlushChunk",
	EMsg_ClientMDSSignInstallScript:                               "EMsg_ClientMDSSignInstallScript",
	EMsg_ClientMDSSignInstallScriptResponse:                       "EMsg_ClientMDSSignInstallScriptResponse",
	EMsg_MDSMigrateChunk:                                          "EMsg_MDSMigrateChunk",
	EMsg_MDSMigrateChunkResponse:                                  "EMsg_MDSMigrateChunkResponse",
	EMsg_CSBase:                                                   "EMsg_CSBase",
	EMsg_CSPing:                                                   "EMsg_CSPing",
	EMsg_CSPingResponse:                                           "EMsg_CSPingResponse",
	EMsg_GMSBase:                                                  "EMsg_GMSBase",
	EMsg_GMSGameServerReplicate:                                   "EMsg_GMSGameServerReplicate",
	EMsg_ClientGMSServerQuery:                                     "EMsg_ClientGMSServerQuery",
	EMsg_GMSClientServerQueryResponse:                             "EMsg_GMSClientServerQueryResponse",
	EMsg_AMGMSGameServerUpdate:                                    "EMsg_AMGMSGameServerUpdate",
	EMsg_AMGMSGameServerRemove:                                    "EMsg_AMGMSGameServerRemove",
	EMsg_GameServerOutOfDate:                                      "EMsg_GameServerOutOfDate",
	EMsg_DeviceAuthorizationBase:                                  "EMsg_DeviceAuthorizationBase",
	EMsg_ClientAuthorizeLocalDeviceRequest:                        "EMsg_ClientAuthorizeLocalDeviceRequest",
	EMsg_ClientAuthorizeLocalDeviceResponse:                       "EMsg_ClientAuthorizeLocalDeviceResponse",
	EMsg_ClientDeauthorizeDeviceRequest:                           "EMsg_ClientDeauthorizeDeviceRequest",
	EMsg_ClientDeauthorizeDevice:                                  "EMsg_ClientDeauthorizeDevice",
	EMsg_ClientUseLocalDeviceAuthorizations:                       "EMsg_ClientUseLocalDeviceAuthorizations",
	EMsg_ClientGetAuthorizedDevices:                               "EMsg_ClientGetAuthorizedDevices",
	EMsg_ClientGetAuthorizedDevicesResponse:                       "EMsg_ClientGetAuthorizedDevicesResponse",
	EMsg_AMNotifySessionDeviceAuthorized:                          "EMsg_AMNotifySessionDeviceAuthorized",
	EMsg_ClientAuthorizeLocalDeviceNotification:                   "EMsg_ClientAuthorizeLocalDeviceNotification",
	EMsg_MMSBase:                                                  "EMsg_MMSBase",
	EMsg_ClientMMSCreateLobby:                                     "EMsg_ClientMMSCreateLobby",
	EMsg_ClientMMSCreateLobbyResponse:                             "EMsg_ClientMMSCreateLobbyResponse",
	EMsg_ClientMMSJoinLobby:                                       "EMsg_ClientMMSJoinLobby",
	EMsg_ClientMMSJoinLobbyResponse:                               "EMsg_ClientMMSJoinLobbyResponse",
	EMsg_ClientMMSLeaveLobby:                                      "EMsg_ClientMMSLeaveLobby",
	EMsg_ClientMMSLeaveLobbyResponse:                              "EMsg_ClientMMSLeaveLobbyResponse",
	EMsg_ClientMMSGetLobbyList:                                    "EMsg_ClientMMSGetLobbyList",
	EMsg_ClientMMSGetLobbyListResponse:                            "EMsg_ClientMMSGetLobbyListResponse",
	EMsg_ClientMMSSetLobbyData:                                    "EMsg_ClientMMSSetLobbyData",
	EMsg_ClientMMSSetLobbyDataResponse:                            "EMsg_ClientMMSSetLobbyDataResponse",
	EMsg_ClientMMSGetLobbyData:                                    "EMsg_ClientMMSGetLobbyData",
	EMsg_ClientMMSLobbyData:                                       "EMsg_ClientMMSLobbyData",
	EMsg_ClientMMSSendLobbyChatMsg:                                "EMsg_ClientMMSSendLobbyChatMsg",
	EMsg_ClientMMSLobbyChatMsg:                                    "EMsg_ClientMMSLobbyChatMsg",
	EMsg_ClientMMSSetLobbyOwner:                                   "EMsg_ClientMMSSetLobbyOwner",
	EMsg_ClientMMSSetLobbyOwnerResponse:                           "EMsg_ClientMMSSetLobbyOwnerResponse",
	EMsg_ClientMMSSetLobbyGameServer:                              "EMsg_ClientMMSSetLobbyGameServer",
	EMsg_ClientMMSLobbyGameServerSet:                              "EMsg_ClientMMSLobbyGameServerSet",
	EMsg_ClientMMSUserJoinedLobby:                                 "EMsg_ClientMMSUserJoinedLobby",
	EMsg_ClientMMSUserLeftLobby:                                   "EMsg_ClientMMSUserLeftLobby",
	EMsg_ClientMMSInviteToLobby:                                   "EMsg_ClientMMSInviteToLobby",
	EMsg_ClientMMSFlushFrenemyListCache:                           "EMsg_ClientMMSFlushFrenemyListCache",
	EMsg_ClientMMSFlushFrenemyListCacheResponse:                   "EMsg_ClientMMSFlushFrenemyListCacheResponse",
	EMsg_ClientMMSSetLobbyLinked:                                  "EMsg_ClientMMSSetLobbyLinked",
	EMsg_NonStdMsgBase:                                            "EMsg_NonStdMsgBase",
	EMsg_NonStdMsgMemcached:                                       "EMsg_NonStdMsgMemcached",
	EMsg_NonStdMsgHTTPServer:                                      "EMsg_NonStdMsgHTTPServer",
	EMsg_NonStdMsgHTTPClient:                                      "EMsg_NonStdMsgHTTPClient",
	EMsg_NonStdMsgWGResponse:                                      "EMsg_NonStdMsgWGResponse",
	EMsg_NonStdMsgPHPSimulator:                                    "EMsg_NonStdMsgPHPSimulator",
	EMsg_NonStdMsgChase:                                           "EMsg_NonStdMsgChase",
	EMsg_NonStdMsgDFSTransfer:                                     "EMsg_NonStdMsgDFSTransfer",
	EMsg_NonStdMsgTests:                                           "EMsg_NonStdMsgTests",
	EMsg_NonStdMsgUMQpipeAAPL:                                     "EMsg_NonStdMsgUMQpipeAAPL",
	EMsg_NonStdMsgSyslog:                                          "EMsg_NonStdMsgSyslog",
	EMsg_NonStdMsgLogsink:                                         "EMsg_NonStdMsgLogsink",
	EMsg_NonStdMsgSteam2Emulator:                                  "EMsg_NonStdMsgSteam2Emulator",
	EMsg_NonStdMsgRTMPServer:                                      "EMsg_NonStdMsgRTMPServer",
	EMsg_UDSBase:                                                  "EMsg_UDSBase",
	EMsg_ClientUDSP2PSessionStarted:                               "EMsg_ClientUDSP2PSessionStarted",
	EMsg_ClientUDSP2PSessionEnded:                                 "EMsg_ClientUDSP2PSessionEnded",
	EMsg_UDSRenderUserAuth:                                        "EMsg_UDSRenderUserAuth",
	EMsg_UDSRenderUserAuthResponse:                                "EMsg_UDSRenderUserAuthResponse",
	EMsg_ClientUDSInviteToGame:                                    "EMsg_ClientUDSInviteToGame",
	EMsg_UDSHasSession:                                            "EMsg_UDSHasSession",
	EMsg_UDSHasSessionResponse:                                    "EMsg_UDSHasSessionResponse",
	EMsg_MPASBase:                                                 "EMsg_MPASBase",
	EMsg_MPASVacBanReset:                                          "EMsg_MPASVacBanReset",
	EMsg_KGSBase:                                                  "EMsg_KGSBase",
	EMsg_KGSAllocateKeyRange:                                      "EMsg_KGSAllocateKeyRange",
	EMsg_KGSAllocateKeyRangeResponse:                              "EMsg_KGSAllocateKeyRangeResponse",
	EMsg_KGSGenerateKeys:                                          "EMsg_KGSGenerateKeys",
	EMsg_KGSGenerateKeysResponse:                                  "EMsg_KGSGenerateKeysResponse",
	EMsg_KGSRemapKeys:                                             "EMsg_KGSRemapKeys",
	EMsg_KGSRemapKeysResponse:                                     "EMsg_KGSRemapKeysResponse",
	EMsg_KGSGenerateGameStopWCKeys:                                "EMsg_KGSGenerateGameStopWCKeys",
	EMsg_KGSGenerateGameStopWCKeysResponse:                        "EMsg_KGSGenerateGameStopWCKeysResponse",
	EMsg_UCMBase:                                                  "EMsg_UCMBase",
	EMsg_ClientUCMAddScreenshot:                                   "EMsg_ClientUCMAddScreenshot",
	EMsg_ClientUCMAddScreenshotResponse:                           "EMsg_ClientUCMAddScreenshotResponse",
	EMsg_UCMValidateObjectExists:                                  "EMsg_UCMValidateObjectExists",
	EMsg_UCMValidateObjectExistsResponse:                          "EMsg_UCMValidateObjectExistsResponse",
	EMsg_UCMResetCommunityContent:                                 "EMsg_UCMResetCommunityContent",
	EMsg_UCMResetCommunityContentResponse:                         "EMsg_UCMResetCommunityContentResponse",
	EMsg_ClientUCMDeleteScreenshot:                                "EMsg_ClientUCMDeleteScreenshot",
	EMsg_ClientUCMDeleteScreenshotResponse:                        "EMsg_ClientUCMDeleteScreenshotResponse",
	EMsg_ClientUCMPublishFile:                                     "EMsg_ClientUCMPublishFile",
	EMsg_ClientUCMPublishFileResponse:                             "EMsg_ClientUCMPublishFileResponse",
	EMsg_ClientUCMGetPublishedFileDetails:                         "EMsg_ClientUCMGetPublishedFileDetails",
	EMsg_ClientUCMGetPublishedFileDetailsResponse:                 "EMsg_ClientUCMGetPublishedFileDetailsResponse",
	EMsg_ClientUCMDeletePublishedFile:                             "EMsg_ClientUCMDeletePublishedFile",
	EMsg_ClientUCMDeletePublishedFileResponse:                     "EMsg_ClientUCMDeletePublishedFileResponse",
	EMsg_ClientUCMEnumerateUserPublishedFiles:                     "EMsg_ClientUCMEnumerateUserPublishedFiles",
	EMsg_ClientUCMEnumerateUserPublishedFilesResponse:             "EMsg_ClientUCMEnumerateUserPublishedFilesResponse",
	EMsg_ClientUCMSubscribePublishedFile:                          "EMsg_ClientUCMSubscribePublishedFile",
	EMsg_ClientUCMSubscribePublishedFileResponse:                  "EMsg_ClientUCMSubscribePublishedFileResponse",
	EMsg_ClientUCMEnumerateUserSubscribedFiles:                    "EMsg_ClientUCMEnumerateUserSubscribedFiles",
	EMsg_ClientUCMEnumerateUserSubscribedFilesResponse:            "EMsg_ClientUCMEnumerateUserSubscribedFilesResponse",
	EMsg_ClientUCMUnsubscribePublishedFile:                        "EMsg_ClientUCMUnsubscribePublishedFile",
	EMsg_ClientUCMUnsubscribePublishedFileResponse:                "EMsg_ClientUCMUnsubscribePublishedFileResponse",
	EMsg_ClientUCMUpdatePublishedFile:                             "EMsg_ClientUCMUpdatePublishedFile",
	EMsg_ClientUCMUpdatePublishedFileResponse:                     "EMsg_ClientUCMUpdatePublishedFileResponse",
	EMsg_UCMUpdatePublishedFile:                                   "EMsg_UCMUpdatePublishedFile",
	EMsg_UCMUpdatePublishedFileResponse:                           "EMsg_UCMUpdatePublishedFileResponse",
	EMsg_UCMDeletePublishedFile:                                   "EMsg_UCMDeletePublishedFile",
	EMsg_UCMDeletePublishedFileResponse:                           "EMsg_UCMDeletePublishedFileResponse",
	EMsg_UCMUpdatePublishedFileStat:                               "EMsg_UCMUpdatePublishedFileStat",
	EMsg_UCMUpdatePublishedFileBan:                                "EMsg_UCMUpdatePublishedFileBan",
	EMsg_UCMUpdatePublishedFileBanResponse:                        "EMsg_UCMUpdatePublishedFileBanResponse",
	EMsg_UCMUpdateTaggedScreenshot:                                "EMsg_UCMUpdateTaggedScreenshot",
	EMsg_UCMAddTaggedScreenshot:                                   "EMsg_UCMAddTaggedScreenshot",
	EMsg_UCMRemoveTaggedScreenshot:                                "EMsg_UCMRemoveTaggedScreenshot",
	EMsg_UCMReloadPublishedFile:                                   "EMsg_UCMReloadPublishedFile",
	EMsg_UCMReloadUserFileListCaches:                              "EMsg_UCMReloadUserFileListCaches",
	EMsg_UCMPublishedFileReported:                                 "EMsg_UCMPublishedFileReported",
	EMsg_UCMUpdatePublishedFileIncompatibleStatus:                 "EMsg_UCMUpdatePublishedFileIncompatibleStatus",
	EMsg_UCMPublishedFilePreviewAdd:                               "EMsg_UCMPublishedFilePreviewAdd",
	EMsg_UCMPublishedFilePreviewAddResponse:                       "EMsg_UCMPublishedFilePreviewAddResponse",
	EMsg_UCMPublishedFilePreviewRemove:                            "EMsg_UCMPublishedFilePreviewRemove",
	EMsg_UCMPublishedFilePreviewRemoveResponse:                    "EMsg_UCMPublishedFilePreviewRemoveResponse",
	EMsg_UCMPublishedFilePreviewChangeSortOrder:                   "EMsg_UCMPublishedFilePreviewChangeSortOrder",
	EMsg_UCMPublishedFilePreviewChangeSortOrderResponse:           "EMsg_UCMPublishedFilePreviewChangeSortOrderResponse",
	EMsg_ClientUCMPublishedFileSubscribed:                         "EMsg_ClientUCMPublishedFileSubscribed",
	EMsg_ClientUCMPublishedFileUnsubscribed:                       "EMsg_ClientUCMPublishedFileUnsubscribed",
	EMsg_UCMPublishedFileSubscribed:                               "EMsg_UCMPublishedFileSubscribed",
	EMsg_UCMPublishedFileUnsubscribed:                             "EMsg_UCMPublishedFileUnsubscribed",
	EMsg_UCMPublishFile:                                           "EMsg_UCMPublishFile",
	EMsg_UCMPublishFileResponse:                                   "EMsg_UCMPublishFileResponse",
	EMsg_UCMPublishedFileChildAdd:                                 "EMsg_UCMPublishedFileChildAdd",
	EMsg_UCMPublishedFileChildAddResponse:                         "EMsg_UCMPublishedFileChildAddResponse",
	EMsg_UCMPublishedFileChildRemove:                              "EMsg_UCMPublishedFileChildRemove",
	EMsg_UCMPublishedFileChildRemoveResponse:                      "EMsg_UCMPublishedFileChildRemoveResponse",
	EMsg_UCMPublishedFileChildChangeSortOrder:                     "EMsg_UCMPublishedFileChildChangeSortOrder",
	EMsg_UCMPublishedFileChildChangeSortOrderResponse:             "EMsg_UCMPublishedFileChildChangeSortOrderResponse",
	EMsg_UCMPublishedFileParentChanged:                            "EMsg_UCMPublishedFileParentChanged",
	EMsg_ClientUCMGetPublishedFilesForUser:                        "EMsg_ClientUCMGetPublishedFilesForUser",
	EMsg_ClientUCMGetPublishedFilesForUserResponse:                "EMsg_ClientUCMGetPublishedFilesForUserResponse",
	EMsg_UCMGetPublishedFilesForUser:                              "EMsg_UCMGetPublishedFilesForUser",
	EMsg_UCMGetPublishedFilesForUserResponse:                      "EMsg_UCMGetPublishedFilesForUserResponse",
	EMsg_ClientUCMSetUserPublishedFileAction:                      "EMsg_ClientUCMSetUserPublishedFileAction",
	EMsg_ClientUCMSetUserPublishedFileActionResponse:              "EMsg_ClientUCMSetUserPublishedFileActionResponse",
	EMsg_ClientUCMEnumeratePublishedFilesByUserAction:             "EMsg_ClientUCMEnumeratePublishedFilesByUserAction",
	EMsg_ClientUCMEnumeratePublishedFilesByUserActionResponse:     "EMsg_ClientUCMEnumeratePublishedFilesByUserActionResponse",
	EMsg_ClientUCMPublishedFileDeleted:                            "EMsg_ClientUCMPublishedFileDeleted",
	EMsg_UCMGetUserSubscribedFiles:                                "EMsg_UCMGetUserSubscribedFiles",
	EMsg_UCMGetUserSubscribedFilesResponse:                        "EMsg_UCMGetUserSubscribedFilesResponse",
	EMsg_UCMFixStatsPublishedFile:                                 "EMsg_UCMFixStatsPublishedFile",
	EMsg_UCMDeleteOldScreenshot:                                   "EMsg_UCMDeleteOldScreenshot",
	EMsg_UCMDeleteOldScreenshotResponse:                           "EMsg_UCMDeleteOldScreenshotResponse",
	EMsg_UCMDeleteOldVideo:                                        "EMsg_UCMDeleteOldVideo",
	EMsg_UCMDeleteOldVideoResponse:                                "EMsg_UCMDeleteOldVideoResponse",
	EMsg_UCMUpdateOldScreenshotPrivacy:                            "EMsg_UCMUpdateOldScreenshotPrivacy",
	EMsg_UCMUpdateOldScreenshotPrivacyResponse:                    "EMsg_UCMUpdateOldScreenshotPrivacyResponse",
	EMsg_ClientUCMEnumerateUserSubscribedFilesWithUpdates:         "EMsg_ClientUCMEnumerateUserSubscribedFilesWithUpdates",
	EMsg_ClientUCMEnumerateUserSubscribedFilesWithUpdatesResponse: "EMsg_ClientUCMEnumerateUserSubscribedFilesWithUpdatesResponse",
	EMsg_UCMPublishedFileContentUpdated:                           "EMsg_UCMPublishedFileContentUpdated",
	EMsg_UCMPublishedFileUpdated:                                  "EMsg_UCMPublishedFileUpdated",
	EMsg_ClientWorkshopItemChangesRequest:                         "EMsg_ClientWorkshopItemChangesRequest",
	EMsg_ClientWorkshopItemChangesResponse:                        "EMsg_ClientWorkshopItemChangesResponse",
	EMsg_ClientWorkshopItemInfoRequest:                            "EMsg_ClientWorkshopItemInfoRequest",
	EMsg_ClientWorkshopItemInfoResponse:                           "EMsg_ClientWorkshopItemInfoResponse",
	EMsg_FSBase:                                                   "EMsg_FSBase",
	EMsg_ClientRichPresenceUpload:                                 "EMsg_ClientRichPresenceUpload",
	EMsg_ClientRichPresenceRequest:                                "EMsg_ClientRichPresenceRequest",
	EMsg_ClientRichPresenceInfo:                                   "EMsg_ClientRichPresenceInfo",
	EMsg_FSRichPresenceRequest:                                    "EMsg_FSRichPresenceRequest",
	EMsg_FSRichPresenceResponse:                                   "EMsg_FSRichPresenceResponse",
	EMsg_FSComputeFrenematrix:                                     "EMsg_FSComputeFrenematrix",
	EMsg_FSComputeFrenematrixResponse:                             "EMsg_FSComputeFrenematrixResponse",
	EMsg_FSPlayStatusNotification:                                 "EMsg_FSPlayStatusNotification",
	EMsg_FSPublishPersonaStatus:                                   "EMsg_FSPublishPersonaStatus",
	EMsg_FSAddOrRemoveFollower:                                    "EMsg_FSAddOrRemoveFollower",
	EMsg_FSAddOrRemoveFollowerResponse:                            "EMsg_FSAddOrRemoveFollowerResponse",
	EMsg_FSUpdateFollowingList:                                    "EMsg_FSUpdateFollowingList",
	EMsg_FSCommentNotification:                                    "EMsg_FSCommentNotification",
	EMsg_FSCommentNotificationViewed:                              "EMsg_FSCommentNotificationViewed",
	EMsg_ClientFSGetFollowerCount:                                 "EMsg_ClientFSGetFollowerCount",
	EMsg_ClientFSGetFollowerCountResponse:                         "EMsg_ClientFSGetFollowerCountResponse",
	EMsg_ClientFSGetIsFollowing:                                   "EMsg_ClientFSGetIsFollowing",
	EMsg_ClientFSGetIsFollowingResponse:                           "EMsg_ClientFSGetIsFollowingResponse",
	EMsg_ClientFSEnumerateFollowingList:                           "EMsg_ClientFSEnumerateFollowingList",
	EMsg_ClientFSEnumerateFollowingListResponse:                   "EMsg_ClientFSEnumerateFollowingListResponse",
	EMsg_FSGetPendingNotificationCount:                            "EMsg_FSGetPendingNotificationCount",
	EMsg_FSGetPendingNotificationCountResponse:                    "EMsg_FSGetPendingNotificationCountResponse",
	EMsg_ClientChatOfflineMessageNotification:                     "EMsg_ClientChatOfflineMessageNotification",
	EMsg_ClientChatRequestOfflineMessageCount:                     "EMsg_ClientChatRequestOfflineMessageCount",
	EMsg_ClientChatGetFriendMessageHistory:                        "EMsg_ClientChatGetFriendMessageHistory",
	EMsg_ClientChatGetFriendMessageHistoryResponse:                "EMsg_ClientChatGetFriendMessageHistoryResponse",
	EMsg_ClientChatGetFriendMessageHistoryForOfflineMessages:      "EMsg_ClientChatGetFriendMessageHistoryForOfflineMessages",
	EMsg_ClientFSGetFriendsSteamLevels:                            "EMsg_ClientFSGetFriendsSteamLevels",
	EMsg_ClientFSGetFriendsSteamLevelsResponse:                    "EMsg_ClientFSGetFriendsSteamLevelsResponse",
	EMsg_FSRequestFriendData:                                      "EMsg_FSRequestFriendData",
	EMsg_DRMRange2:                                                "EMsg_DRMRange2",
	EMsg_CEGVersionSetEnableDisableResponse:                       "EMsg_CEGVersionSetEnableDisableResponse",
	EMsg_CEGPropStatusDRMSRequest:                                 "EMsg_CEGPropStatusDRMSRequest",
	EMsg_CEGPropStatusDRMSResponse:                                "EMsg_CEGPropStatusDRMSResponse",
	EMsg_CEGWhackFailureReportRequest:                             "EMsg_CEGWhackFailureReportRequest",
	EMsg_CEGWhackFailureReportResponse:                            "EMsg_CEGWhackFailureReportResponse",
	EMsg_DRMSFetchVersionSet:                                      "EMsg_DRMSFetchVersionSet",
	EMsg_DRMSFetchVersionSetResponse:                              "EMsg_DRMSFetchVersionSetResponse",
	EMsg_EconBase:                                                 "EMsg_EconBase",
	EMsg_EconTrading_InitiateTradeRequest:                         "EMsg_EconTrading_InitiateTradeRequest",
	EMsg_EconTrading_InitiateTradeProposed:                        "EMsg_EconTrading_InitiateTradeProposed",
	EMsg_EconTrading_InitiateTradeResponse:                        "EMsg_EconTrading_InitiateTradeResponse",
	EMsg_EconTrading_InitiateTradeResult:                          "EMsg_EconTrading_InitiateTradeResult",
	EMsg_EconTrading_StartSession:                                 "EMsg_EconTrading_StartSession",
	EMsg_EconTrading_CancelTradeRequest:                           "EMsg_EconTrading_CancelTradeRequest",
	EMsg_EconFlushInventoryCache:                                  "EMsg_EconFlushInventoryCache",
	EMsg_EconFlushInventoryCacheResponse:                          "EMsg_EconFlushInventoryCacheResponse",
	EMsg_EconCDKeyProcessTransaction:                              "EMsg_EconCDKeyProcessTransaction",
	EMsg_EconCDKeyProcessTransactionResponse:                      "EMsg_EconCDKeyProcessTransactionResponse",
	EMsg_EconGetErrorLogs:                                         "EMsg_EconGetErrorLogs",
	EMsg_EconGetErrorLogsResponse:                                 "EMsg_EconGetErrorLogsResponse",
	EMsg_RMRange:                                                  "EMsg_RMRange",
	EMsg_RMTestVerisignOTPResponse:                                "EMsg_RMTestVerisignOTPResponse",
	EMsg_RMDeleteMemcachedKeys:                                    "EMsg_RMDeleteMemcachedKeys",
	EMsg_RMRemoteInvoke:                                           "EMsg_RMRemoteInvoke",
	EMsg_BadLoginIPList:                                           "EMsg_BadLoginIPList",
	EMsg_RMMsgTraceAddTrigger:                                     "EMsg_RMMsgTraceAddTrigger",
	EMsg_RMMsgTraceRemoveTrigger:                                  "EMsg_RMMsgTraceRemoveTrigger",
	EMsg_RMMsgTraceEvent:                                          "EMsg_RMMsgTraceEvent",
	EMsg_UGSBase:                                                  "EMsg_UGSBase",
	EMsg_ClientUGSGetGlobalStats:                                  "EMsg_ClientUGSGetGlobalStats",
	EMsg_ClientUGSGetGlobalStatsResponse:                          "EMsg_ClientUGSGetGlobalStatsResponse",
	EMsg_StoreBase:                                                "EMsg_StoreBase",
	EMsg_UMQBase:                                                  "EMsg_UMQBase",
	EMsg_UMQLogonResponse:                                         "EMsg_UMQLogonResponse",
	EMsg_UMQLogoffRequest:                                         "EMsg_UMQLogoffRequest",
	EMsg_UMQLogoffResponse:                                        "EMsg_UMQLogoffResponse",
	EMsg_UMQSendChatMessage:                                       "EMsg_UMQSendChatMessage",
	EMsg_UMQIncomingChatMessage:                                   "EMsg_UMQIncomingChatMessage",
	EMsg_UMQPoll:                                                  "EMsg_UMQPoll",
	EMsg_UMQPollResults:                                           "EMsg_UMQPollResults",
	EMsg_UMQ2AM_ClientMsgBatch:                                    "EMsg_UMQ2AM_ClientMsgBatch",
	EMsg_UMQEnqueueMobileSalePromotions:                           "EMsg_UMQEnqueueMobileSalePromotions",
	EMsg_UMQEnqueueMobileAnnouncements:                            "EMsg_UMQEnqueueMobileAnnouncements",
	EMsg_WorkshopBase:                                             "EMsg_WorkshopBase",
	EMsg_WorkshopAcceptTOSResponse:                                "EMsg_WorkshopAcceptTOSResponse",
	EMsg_WebAPIBase:                                               "EMsg_WebAPIBase",
	EMsg_WebAPIValidateOAuth2TokenResponse:                        "EMsg_WebAPIValidateOAuth2TokenResponse",
	EMsg_WebAPIInvalidateTokensForAccount:                         "EMsg_WebAPIInvalidateTokensForAccount",
	EMsg_WebAPIRegisterGCInterfaces:                               "EMsg_WebAPIRegisterGCInterfaces",
	EMsg_WebAPIInvalidateOAuthClientCache:                         "EMsg_WebAPIInvalidateOAuthClientCache",
	EMsg_WebAPIInvalidateOAuthTokenCache:                          "EMsg_WebAPIInvalidateOAuthTokenCache",
	EMsg_WebAPISetSecrets:                                         "EMsg_WebAPISetSecrets",
	EMsg_BackpackBase:                                             "EMsg_BackpackBase",
	EMsg_BackpackAddToCurrency:                                    "EMsg_BackpackAddToCurrency",
	EMsg_BackpackAddToCurrencyResponse:                            "EMsg_BackpackAddToCurrencyResponse",
	EMsg_CREBase:                                                  "EMsg_CREBase",
	EMsg_CRERankByTrend:                                           "EMsg_CRERankByTrend",
	EMsg_CRERankByTrendResponse:                                   "EMsg_CRERankByTrendResponse",
	EMsg_CREItemVoteSummary:                                       "EMsg_CREItemVoteSummary",
	EMsg_CREItemVoteSummaryResponse:                               "EMsg_CREItemVoteSummaryResponse",
	EMsg_CRERankByVote:                                            "EMsg_CRERankByVote",
	EMsg_CRERankByVoteResponse:                                    "EMsg_CRERankByVoteResponse",
	EMsg_CREUpdateUserPublishedItemVote:                           "EMsg_CREUpdateUserPublishedItemVote",
	EMsg_CREUpdateUserPublishedItemVoteResponse:                   "EMsg_CREUpdateUserPublishedItemVoteResponse",
	EMsg_CREGetUserPublishedItemVoteDetails:                       "EMsg_CREGetUserPublishedItemVoteDetails",
	EMsg_CREGetUserPublishedItemVoteDetailsResponse:               "EMsg_CREGetUserPublishedItemVoteDetailsResponse",
	EMsg_CREEnumeratePublishedFiles:                               "EMsg_CREEnumeratePublishedFiles",
	EMsg_CREEnumeratePublishedFilesResponse:                       "EMsg_CREEnumeratePublishedFilesResponse",
	EMsg_CREPublishedFileVoteAdded:                                "EMsg_CREPublishedFileVoteAdded",
	EMsg_SecretsBase:                                              "EMsg_SecretsBase",
	EMsg_SecretsCredentialPairResponse:                            "EMsg_SecretsCredentialPairResponse",
	EMsg_SecretsRequestServerIdentity:                             "EMsg_SecretsRequestServerIdentity",
	EMsg_SecretsServerIdentityResponse:                            "EMsg_SecretsServerIdentityResponse",
	EMsg_SecretsUpdateServerIdentities:                            "EMsg_SecretsUpdateServerIdentities",
	EMsg_BoxMonitorBase:                                           "EMsg_BoxMonitorBase",
	EMsg_BoxMonitorReportResponse:                                 "EMsg_BoxMonitorReportResponse",
	EMsg_LogsinkBase:                                              "EMsg_LogsinkBase",
	EMsg_PICSBase:                                                 "EMsg_PICSBase",
	EMsg_ClientPICSChangesSinceRequest:                            "EMsg_ClientPICSChangesSinceRequest",
	EMsg_ClientPICSChangesSinceResponse:                           "EMsg_ClientPICSChangesSinceResponse",
	EMsg_ClientPICSProductInfoRequest:                             "EMsg_ClientPICSProductInfoRequest",
	EMsg_ClientPICSProductInfoResponse:                            "EMsg_ClientPICSProductInfoResponse",
	EMsg_ClientPICSAccessTokenRequest:                             "EMsg_ClientPICSAccessTokenRequest",
	EMsg_ClientPICSAccessTokenResponse:                            "EMsg_ClientPICSAccessTokenResponse",
	EMsg_WorkerProcess:                                            "EMsg_WorkerProcess",
	EMsg_WorkerProcessPingResponse:                                "EMsg_WorkerProcessPingResponse",
	EMsg_WorkerProcessShutdown:                                    "EMsg_WorkerProcessShutdown",
	EMsg_DRMWorkerProcess:                                         "EMsg_DRMWorkerProcess",
	EMsg_DRMWorkerProcessDRMAndSignResponse:                       "EMsg_DRMWorkerProcessDRMAndSignResponse",
	EMsg_DRMWorkerProcessSteamworksInfoRequest:                    "EMsg_DRMWorkerProcessSteamworksInfoRequest",
	EMsg_DRMWorkerProcessSteamworksInfoResponse:                   "EMsg_DRMWorkerProcessSteamworksInfoResponse",
	EMsg_DRMWorkerProcessInstallDRMDLLRequest:                     "EMsg_DRMWorkerProcessInstallDRMDLLRequest",
	EMsg_DRMWorkerProcessInstallDRMDLLResponse:                    "EMsg_DRMWorkerProcessInstallDRMDLLResponse",
	EMsg_DRMWorkerProcessSecretIdStringRequest:                    "EMsg_DRMWorkerProcessSecretIdStringRequest",
	EMsg_DRMWorkerProcessSecretIdStringResponse:                   "EMsg_DRMWorkerProcessSecretIdStringResponse",
	EMsg_DRMWorkerProcessGetDRMGuidsFromFileRequest:               "EMsg_DRMWorkerProcessGetDRMGuidsFromFileRequest",
	EMsg_DRMWorkerProcessGetDRMGuidsFromFileResponse:              "EMsg_DRMWorkerProcessGetDRMGuidsFromFileResponse",
	EMsg_DRMWorkerProcessInstallProcessedFilesRequest:             "EMsg_DRMWorkerProcessInstallProcessedFilesRequest",
	EMsg_DRMWorkerProcessInstallProcessedFilesResponse:            "EMsg_DRMWorkerProcessInstallProcessedFilesResponse",
	EMsg_DRMWorkerProcessExamineBlobRequest:                       "EMsg_DRMWorkerProcessExamineBlobRequest",
	EMsg_DRMWorkerProcessExamineBlobResponse:                      "EMsg_DRMWorkerProcessExamineBlobResponse",
	EMsg_DRMWorkerProcessDescribeSecretRequest:                    "EMsg_DRMWorkerProcessDescribeSecretRequest",
	EMsg_DRMWorkerProcessDescribeSecretResponse:                   "EMsg_DRMWorkerProcessDescribeSecretResponse",
	EMsg_DRMWorkerProcessBackfillOriginalRequest:                  "EMsg_DRMWorkerProcessBackfillOriginalRequest",
	EMsg_DRMWorkerProcessBackfillOriginalResponse:                 "EMsg_DRMWorkerProcessBackfillOriginalResponse",
	EMsg_DRMWorkerProcessValidateDRMDLLRequest:                    "EMsg_DRMWorkerProcessValidateDRMDLLRequest",
	EMsg_DRMWorkerProcessValidateDRMDLLResponse:                   "EMsg_DRMWorkerProcessValidateDRMDLLResponse",
	EMsg_DRMWorkerProcessValidateFileRequest:                      "EMsg_DRMWorkerProcessValidateFileRequest",
	EMsg_DRMWorkerProcessValidateFileResponse:                     "EMsg_DRMWorkerProcessValidateFileResponse",
	EMsg_DRMWorkerProcessSplitAndInstallRequest:                   "EMsg_DRMWorkerProcessSplitAndInstallRequest",
	EMsg_DRMWorkerProcessSplitAndInstallResponse:                  "EMsg_DRMWorkerProcessSplitAndInstallResponse",
	EMsg_DRMWorkerProcessGetBlobRequest:                           "EMsg_DRMWorkerProcessGetBlobRequest",
	EMsg_DRMWorkerProcessGetBlobResponse:                          "EMsg_DRMWorkerProcessGetBlobResponse",
	EMsg_DRMWorkerProcessEvaluateCrashRequest:                     "EMsg_DRMWorkerProcessEvaluateCrashRequest",
	EMsg_DRMWorkerProcessEvaluateCrashResponse:                    "EMsg_DRMWorkerProcessEvaluateCrashResponse",
	EMsg_DRMWorkerProcessAnalyzeFileRequest:                       "EMsg_DRMWorkerProcessAnalyzeFileRequest",
	EMsg_DRMWorkerProcessAnalyzeFileResponse:                      "EMsg_DRMWorkerProcessAnalyzeFileResponse",
	EMsg_DRMWorkerProcessUnpackBlobRequest:                        "EMsg_DRMWorkerProcessUnpackBlobRequest",
	EMsg_DRMWorkerProcessUnpackBlobResponse:                       "EMsg_DRMWorkerProcessUnpackBlobResponse",
	EMsg_DRMWorkerProcessInstallAllRequest:                        "EMsg_DRMWorkerProcessInstallAllRequest",
	EMsg_DRMWorkerProcessInstallAllResponse:                       "EMsg_DRMWorkerProcessInstallAllResponse",
	EMsg_TestWorkerProcess:                                        "EMsg_TestWorkerProcess",
	EMsg_TestWorkerProcessLoadUnloadModuleResponse:                "EMsg_TestWorkerProcessLoadUnloadModuleResponse",
	EMsg_TestWorkerProcessServiceModuleCallRequest:                "EMsg_TestWorkerProcessServiceModuleCallRequest",
	EMsg_TestWorkerProcessServiceModuleCallResponse:               "EMsg_TestWorkerProcessServiceModuleCallResponse",
	EMsg_QuestServerBase:                                          "EMsg_QuestServerBase",
	EMsg_ClientGetEmoticonList:                                    "EMsg_ClientGetEmoticonList",
	EMsg_ClientEmoticonList:                                       "EMsg_ClientEmoticonList",
	EMsg_SLCBase:                                                  "EMsg_SLCBase",
	EMsg_SLCRequestUserSessionStatus:                              "EMsg_SLCRequestUserSessionStatus",
	EMsg_SLCSharedLicensesLockStatus:                              "EMsg_SLCSharedLicensesLockStatus",
	EMsg_ClientSharedLicensesLockStatus:                           "EMsg_ClientSharedLicensesLockStatus",
	EMsg_ClientSharedLicensesStopPlaying:                          "EMsg_ClientSharedLicensesStopPlaying",
	EMsg_ClientSharedLibraryLockStatus:                            "EMsg_ClientSharedLibraryLockStatus",
	EMsg_ClientSharedLibraryStopPlaying:                           "EMsg_ClientSharedLibraryStopPlaying",
	EMsg_SLCOwnerLibraryChanged:                                   "EMsg_SLCOwnerLibraryChanged",
	EMsg_SLCSharedLibraryChanged:                                  "EMsg_SLCSharedLibraryChanged",
	EMsg_RemoteClientBase:                                         "EMsg_RemoteClientBase",
	EMsg_RemoteClientAuthResponse:                                 "EMsg_RemoteClientAuthResponse",
	EMsg_RemoteClientAppStatus:                                    "EMsg_RemoteClientAppStatus",
	EMsg_RemoteClientStartStream:                                  "EMsg_RemoteClientStartStream",
	EMsg_RemoteClientStartStreamResponse:                          "EMsg_RemoteClientStartStreamResponse",
	EMsg_RemoteClientPing:                                         "EMsg_RemoteClientPing",
	EMsg_RemoteClientPingResponse:                                 "EMsg_RemoteClientPingResponse",
	EMsg_ClientUnlockStreaming:                                    "EMsg_ClientUnlockStreaming",
	EMsg_ClientUnlockStreamingResponse:                            "EMsg_ClientUnlockStreamingResponse",
	EMsg_RemoteClientAcceptEULA:                                   "EMsg_RemoteClientAcceptEULA",
	EMsg_RemoteClientGetControllerConfig:                          "EMsg_RemoteClientGetControllerConfig",
	EMsg_RemoteClientGetControllerConfigResposne:                  "EMsg_RemoteClientGetControllerConfigResposne",
	EMsg_RemoteClientStreamingEnabled:                             "EMsg_RemoteClientStreamingEnabled",
	EMsg_ClientUnlockHEVC:                                         "EMsg_ClientUnlockHEVC",
	EMsg_ClientUnlockHEVCResponse:                                 "EMsg_ClientUnlockHEVCResponse",
	EMsg_ClientConcurrentSessionsBase:                             "EMsg_ClientConcurrentSessionsBase",
	EMsg_ClientKickPlayingSession:                                 "EMsg_ClientKickPlayingSession",
	EMsg_ClientBroadcastBase:                                      "EMsg_ClientBroadcastBase",
	EMsg_ClientBroadcastFrames:                                    "EMsg_ClientBroadcastFrames",
	EMsg_ClientBroadcastDisconnect:                                "EMsg_ClientBroadcastDisconnect",
	EMsg_ClientBroadcastScreenshot:                                "EMsg_ClientBroadcastScreenshot",
	EMsg_ClientBroadcastUploadConfig:                              "EMsg_ClientBroadcastUploadConfig",
	EMsg_BaseClient3:                                              "EMsg_BaseClient3",
	EMsg_ClientVoiceCallPreAuthorizeResponse:                      "EMsg_ClientVoiceCallPreAuthorizeResponse",
	EMsg_ClientServerTimestampRequest:                             "EMsg_ClientServerTimestampRequest",
}

func (e EMsg) String() string {
	if s, ok := EMsg_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EMsg_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EResult int32

const (
	EResult_Invalid                                 EResult = 0
	EResult_OK                                      EResult = 1
	EResult_Fail                                    EResult = 2
	EResult_NoConnection                            EResult = 3
	EResult_InvalidPassword                         EResult = 5
	EResult_LoggedInElsewhere                       EResult = 6
	EResult_InvalidProtocolVer                      EResult = 7
	EResult_InvalidParam                            EResult = 8
	EResult_FileNotFound                            EResult = 9
	EResult_Busy                                    EResult = 10
	EResult_InvalidState                            EResult = 11
	EResult_InvalidName                             EResult = 12
	EResult_InvalidEmail                            EResult = 13
	EResult_DuplicateName                           EResult = 14
	EResult_AccessDenied                            EResult = 15
	EResult_Timeout                                 EResult = 16
	EResult_Banned                                  EResult = 17
	EResult_AccountNotFound                         EResult = 18
	EResult_InvalidSteamID                          EResult = 19
	EResult_ServiceUnavailable                      EResult = 20
	EResult_NotLoggedOn                             EResult = 21
	EResult_Pending                                 EResult = 22
	EResult_EncryptionFailure                       EResult = 23
	EResult_InsufficientPrivilege                   EResult = 24
	EResult_LimitExceeded                           EResult = 25
	EResult_Revoked                                 EResult = 26
	EResult_Expired                                 EResult = 27
	EResult_AlreadyRedeemed                         EResult = 28
	EResult_DuplicateRequest                        EResult = 29
	EResult_AlreadyOwned                            EResult = 30
	EResult_IPNotFound                              EResult = 31
	EResult_PersistFailed                           EResult = 32
	EResult_LockingFailed                           EResult = 33
	EResult_LogonSessionReplaced                    EResult = 34
	EResult_ConnectFailed                           EResult = 35
	EResult_HandshakeFailed                         EResult = 36
	EResult_IOFailure                               EResult = 37
	EResult_RemoteDisconnect                        EResult = 38
	EResult_ShoppingCartNotFound                    EResult = 39
	EResult_Blocked                                 EResult = 40
	EResult_Ignored                                 EResult = 41
	EResult_NoMatch                                 EResult = 42
	EResult_AccountDisabled                         EResult = 43
	EResult_ServiceReadOnly                         EResult = 44
	EResult_AccountNotFeatured                      EResult = 45
	EResult_AdministratorOK                         EResult = 46
	EResult_ContentVersion                          EResult = 47
	EResult_TryAnotherCM                            EResult = 48
	EResult_PasswordRequiredToKickSession           EResult = 49
	EResult_AlreadyLoggedInElsewhere                EResult = 50
	EResult_Suspended                               EResult = 51
	EResult_Cancelled                               EResult = 52
	EResult_DataCorruption                          EResult = 53
	EResult_DiskFull                                EResult = 54
	EResult_RemoteCallFailed                        EResult = 55
	EResult_PasswordNotSet                          EResult = 56 // removed: renamed to PasswordUnset
	EResult_PasswordUnset                           EResult = 56
	EResult_ExternalAccountUnlinked                 EResult = 57
	EResult_PSNTicketInvalid                        EResult = 58
	EResult_ExternalAccountAlreadyLinked            EResult = 59
	EResult_RemoteFileConflict                      EResult = 60
	EResult_IllegalPassword                         EResult = 61
	EResult_SameAsPreviousValue                     EResult = 62
	EResult_AccountLogonDenied                      EResult = 63
	EResult_CannotUseOldPassword                    EResult = 64
	EResult_InvalidLoginAuthCode                    EResult = 65
	EResult_AccountLogonDeniedNoMailSent            EResult = 66 // removed: renamed to AccountLogonDeniedNoMail
	EResult_AccountLogonDeniedNoMail                EResult = 66
	EResult_HardwareNotCapableOfIPT                 EResult = 67
	EResult_IPTInitError                            EResult = 68
	EResult_ParentalControlRestricted               EResult = 69
	EResult_FacebookQueryError                      EResult = 70
	EResult_ExpiredLoginAuthCode                    EResult = 71
	EResult_IPLoginRestrictionFailed                EResult = 72
	EResult_AccountLocked                           EResult = 73 // removed: renamed to AccountLockedDown
	EResult_AccountLockedDown                       EResult = 73
	EResult_AccountLogonDeniedVerifiedEmailRequired EResult = 74
	EResult_NoMatchingURL                           EResult = 75
	EResult_BadResponse                             EResult = 76
	EResult_RequirePasswordReEntry                  EResult = 77
	EResult_ValueOutOfRange                         EResult = 78
	EResult_UnexpectedError                         EResult = 79
	EResult_Disabled                                EResult = 80
	EResult_InvalidCEGSubmission                    EResult = 81
	EResult_RestrictedDevice                        EResult = 82
	EResult_RegionLocked                            EResult = 83
	EResult_RateLimitExceeded                       EResult = 84
	EResult_AccountLogonDeniedNeedTwoFactorCode     EResult = 85 // removed: renamed to AccountLoginDeniedNeedTwoFactor
	EResult_AccountLoginDeniedNeedTwoFactor         EResult = 85
	EResult_ItemOrEntryHasBeenDeleted               EResult = 86 // removed: renamed to ItemDeleted
	EResult_ItemDeleted                             EResult = 86
	EResult_AccountLoginDeniedThrottle              EResult = 87
	EResult_TwoFactorCodeMismatch                   EResult = 88
	EResult_TwoFactorActivationCodeMismatch         EResult = 89
	EResult_AccountAssociatedToMultiplePlayers      EResult = 90 // removed: renamed to AccountAssociatedToMultiplePartners
	EResult_AccountAssociatedToMultiplePartners     EResult = 90
	EResult_NotModified                             EResult = 91
	EResult_NoMobileDeviceAvailable                 EResult = 92 // removed: renamed to NoMobileDevice
	EResult_NoMobileDevice                          EResult = 92
	EResult_TimeIsOutOfSync                         EResult = 93 // removed: renamed to TimeNotSynced
	EResult_TimeNotSynced                           EResult = 93
	EResult_SMSCodeFailed                           EResult = 94
	EResult_TooManyAccountsAccessThisResource       EResult = 95 // removed: renamed to AccountLimitExceeded
	EResult_AccountLimitExceeded                    EResult = 95
	EResult_AccountActivityLimitExceeded            EResult = 96
	EResult_PhoneActivityLimitExceeded              EResult = 97
	EResult_RefundToWallet                          EResult = 98
	EResult_EmailSendFailure                        EResult = 99
	EResult_NotSettled                              EResult = 100
	EResult_NeedCaptcha                             EResult = 101
	EResult_GSLTDenied                              EResult = 102
	EResult_GSOwnerDenied                           EResult = 103
	EResult_InvalidItemType                         EResult = 104
	EResult_IPBanned                                EResult = 105
	EResult_GSLTExpired                             EResult = 106
	EResult_InsufficientFunds                       EResult = 107
	EResult_TooManyPending                          EResult = 108
	EResult_NoSiteLicensesFound                     EResult = 109
	EResult_WGNetworkSendExceeded                   EResult = 110
	EResult_AccountNotFriends                       EResult = 111
	EResult_LimitedUserAccount                      EResult = 112
)

var EResult_name = map[EResult]string{
	EResult_Invalid:                                 "EResult_Invalid",
	EResult_OK:                                      "EResult_OK",
	EResult_Fail:                                    "EResult_Fail",
	EResult_NoConnection:                            "EResult_NoConnection",
	EResult_InvalidPassword:                         "EResult_InvalidPassword",
	EResult_LoggedInElsewhere:                       "EResult_LoggedInElsewhere",
	EResult_InvalidProtocolVer:                      "EResult_InvalidProtocolVer",
	EResult_InvalidParam:                            "EResult_InvalidParam",
	EResult_FileNotFound:                            "EResult_FileNotFound",
	EResult_Busy:                                    "EResult_Busy",
	EResult_InvalidState:                            "EResult_InvalidState",
	EResult_InvalidName:                             "EResult_InvalidName",
	EResult_InvalidEmail:                            "EResult_InvalidEmail",
	EResult_DuplicateName:                           "EResult_DuplicateName",
	EResult_AccessDenied:                            "EResult_AccessDenied",
	EResult_Timeout:                                 "EResult_Timeout",
	EResult_Banned:                                  "EResult_Banned",
	EResult_AccountNotFound:                         "EResult_AccountNotFound",
	EResult_InvalidSteamID:                          "EResult_InvalidSteamID",
	EResult_ServiceUnavailable:                      "EResult_ServiceUnavailable",
	EResult_NotLoggedOn:                             "EResult_NotLoggedOn",
	EResult_Pending:                                 "EResult_Pending",
	EResult_EncryptionFailure:                       "EResult_EncryptionFailure",
	EResult_InsufficientPrivilege:                   "EResult_InsufficientPrivilege",
	EResult_LimitExceeded:                           "EResult_LimitExceeded",
	EResult_Revoked:                                 "EResult_Revoked",
	EResult_Expired:                                 "EResult_Expired",
	EResult_AlreadyRedeemed:                         "EResult_AlreadyRedeemed",
	EResult_DuplicateRequest:                        "EResult_DuplicateRequest",
	EResult_AlreadyOwned:                            "EResult_AlreadyOwned",
	EResult_IPNotFound:                              "EResult_IPNotFound",
	EResult_PersistFailed:                           "EResult_PersistFailed",
	EResult_LockingFailed:                           "EResult_LockingFailed",
	EResult_LogonSessionReplaced:                    "EResult_LogonSessionReplaced",
	EResult_ConnectFailed:                           "EResult_ConnectFailed",
	EResult_HandshakeFailed:                         "EResult_HandshakeFailed",
	EResult_IOFailure:                               "EResult_IOFailure",
	EResult_RemoteDisconnect:                        "EResult_RemoteDisconnect",
	EResult_ShoppingCartNotFound:                    "EResult_ShoppingCartNotFound",
	EResult_Blocked:                                 "EResult_Blocked",
	EResult_Ignored:                                 "EResult_Ignored",
	EResult_NoMatch:                                 "EResult_NoMatch",
	EResult_AccountDisabled:                         "EResult_AccountDisabled",
	EResult_ServiceReadOnly:                         "EResult_ServiceReadOnly",
	EResult_AccountNotFeatured:                      "EResult_AccountNotFeatured",
	EResult_AdministratorOK:                         "EResult_AdministratorOK",
	EResult_ContentVersion:                          "EResult_ContentVersion",
	EResult_TryAnotherCM:                            "EResult_TryAnotherCM",
	EResult_PasswordRequiredToKickSession:           "EResult_PasswordRequiredToKickSession",
	EResult_AlreadyLoggedInElsewhere:                "EResult_AlreadyLoggedInElsewhere",
	EResult_Suspended:                               "EResult_Suspended",
	EResult_Cancelled:                               "EResult_Cancelled",
	EResult_DataCorruption:                          "EResult_DataCorruption",
	EResult_DiskFull:                                "EResult_DiskFull",
	EResult_RemoteCallFailed:                        "EResult_RemoteCallFailed",
	EResult_PasswordUnset:                           "EResult_PasswordUnset",
	EResult_ExternalAccountUnlinked:                 "EResult_ExternalAccountUnlinked",
	EResult_PSNTicketInvalid:                        "EResult_PSNTicketInvalid",
	EResult_ExternalAccountAlreadyLinked:            "EResult_ExternalAccountAlreadyLinked",
	EResult_RemoteFileConflict:                      "EResult_RemoteFileConflict",
	EResult_IllegalPassword:                         "EResult_IllegalPassword",
	EResult_SameAsPreviousValue:                     "EResult_SameAsPreviousValue",
	EResult_AccountLogonDenied:                      "EResult_AccountLogonDenied",
	EResult_CannotUseOldPassword:                    "EResult_CannotUseOldPassword",
	EResult_InvalidLoginAuthCode:                    "EResult_InvalidLoginAuthCode",
	EResult_AccountLogonDeniedNoMail:                "EResult_AccountLogonDeniedNoMail",
	EResult_HardwareNotCapableOfIPT:                 "EResult_HardwareNotCapableOfIPT",
	EResult_IPTInitError:                            "EResult_IPTInitError",
	EResult_ParentalControlRestricted:               "EResult_ParentalControlRestricted",
	EResult_FacebookQueryError:                      "EResult_FacebookQueryError",
	EResult_ExpiredLoginAuthCode:                    "EResult_ExpiredLoginAuthCode",
	EResult_IPLoginRestrictionFailed:                "EResult_IPLoginRestrictionFailed",
	EResult_AccountLockedDown:                       "EResult_AccountLockedDown",
	EResult_AccountLogonDeniedVerifiedEmailRequired: "EResult_AccountLogonDeniedVerifiedEmailRequired",
	EResult_NoMatchingURL:                           "EResult_NoMatchingURL",
	EResult_BadResponse:                             "EResult_BadResponse",
	EResult_RequirePasswordReEntry:                  "EResult_RequirePasswordReEntry",
	EResult_ValueOutOfRange:                         "EResult_ValueOutOfRange",
	EResult_UnexpectedError:                         "EResult_UnexpectedError",
	EResult_Disabled:                                "EResult_Disabled",
	EResult_InvalidCEGSubmission:                    "EResult_InvalidCEGSubmission",
	EResult_RestrictedDevice:                        "EResult_RestrictedDevice",
	EResult_RegionLocked:                            "EResult_RegionLocked",
	EResult_RateLimitExceeded:                       "EResult_RateLimitExceeded",
	EResult_AccountLoginDeniedNeedTwoFactor:         "EResult_AccountLoginDeniedNeedTwoFactor",
	EResult_ItemDeleted:                             "EResult_ItemDeleted",
	EResult_AccountLoginDeniedThrottle:              "EResult_AccountLoginDeniedThrottle",
	EResult_TwoFactorCodeMismatch:                   "EResult_TwoFactorCodeMismatch",
	EResult_TwoFactorActivationCodeMismatch:         "EResult_TwoFactorActivationCodeMismatch",
	EResult_AccountAssociatedToMultiplePartners:     "EResult_AccountAssociatedToMultiplePartners",
	EResult_NotModified:                             "EResult_NotModified",
	EResult_NoMobileDevice:                          "EResult_NoMobileDevice",
	EResult_TimeNotSynced:                           "EResult_TimeNotSynced",
	EResult_SMSCodeFailed:                           "EResult_SMSCodeFailed",
	EResult_AccountLimitExceeded:                    "EResult_AccountLimitExceeded",
	EResult_AccountActivityLimitExceeded:            "EResult_AccountActivityLimitExceeded",
	EResult_PhoneActivityLimitExceeded:              "EResult_PhoneActivityLimitExceeded",
	EResult_RefundToWallet:                          "EResult_RefundToWallet",
	EResult_EmailSendFailure:                        "EResult_EmailSendFailure",
	EResult_NotSettled:                              "EResult_NotSettled",
	EResult_NeedCaptcha:                             "EResult_NeedCaptcha",
	EResult_GSLTDenied:                              "EResult_GSLTDenied",
	EResult_GSOwnerDenied:                           "EResult_GSOwnerDenied",
	EResult_InvalidItemType:                         "EResult_InvalidItemType",
	EResult_IPBanned:                                "EResult_IPBanned",
	EResult_GSLTExpired:                             "EResult_GSLTExpired",
	EResult_InsufficientFunds:                       "EResult_InsufficientFunds",
	EResult_TooManyPending:                          "EResult_TooManyPending",
	EResult_NoSiteLicensesFound:                     "EResult_NoSiteLicensesFound",
	EResult_WGNetworkSendExceeded:                   "EResult_WGNetworkSendExceeded",
	EResult_AccountNotFriends:                       "EResult_AccountNotFriends",
	EResult_LimitedUserAccount:                      "EResult_LimitedUserAccount",
}

func (e EResult) String() string {
	if s, ok := EResult_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EResult_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EUniverse int32

const (
	EUniverse_Invalid  EUniverse = 0
	EUniverse_Public   EUniverse = 1
	EUniverse_Beta     EUniverse = 2
	EUniverse_Internal EUniverse = 3
	EUniverse_Dev      EUniverse = 4
	EUniverse_Max      EUniverse = 5
)

var EUniverse_name = map[EUniverse]string{
	EUniverse_Invalid:  "EUniverse_Invalid",
	EUniverse_Public:   "EUniverse_Public",
	EUniverse_Beta:     "EUniverse_Beta",
	EUniverse_Internal: "EUniverse_Internal",
	EUniverse_Dev:      "EUniverse_Dev",
	EUniverse_Max:      "EUniverse_Max",
}

func (e EUniverse) String() string {
	if s, ok := EUniverse_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EUniverse_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EChatEntryType int32

const (
	EChatEntryType_Invalid          EChatEntryType = 0
	EChatEntryType_ChatMsg          EChatEntryType = 1
	EChatEntryType_Typing           EChatEntryType = 2
	EChatEntryType_InviteGame       EChatEntryType = 3
	EChatEntryType_Emote            EChatEntryType = 4 // removed: No longer supported by clients
	EChatEntryType_LobbyGameStart   EChatEntryType = 5 // removed: Listen for LobbyGameCreated_t callback instead
	EChatEntryType_LeftConversation EChatEntryType = 6
	EChatEntryType_Entered          EChatEntryType = 7
	EChatEntryType_WasKicked        EChatEntryType = 8
	EChatEntryType_WasBanned        EChatEntryType = 9
	EChatEntryType_Disconnected     EChatEntryType = 10
	EChatEntryType_HistoricalChat   EChatEntryType = 11
	EChatEntryType_Reserved1        EChatEntryType = 12
	EChatEntryType_Reserved2        EChatEntryType = 13
	EChatEntryType_LinkBlocked      EChatEntryType = 14
)

var EChatEntryType_name = map[EChatEntryType]string{
	EChatEntryType_Invalid:          "EChatEntryType_Invalid",
	EChatEntryType_ChatMsg:          "EChatEntryType_ChatMsg",
	EChatEntryType_Typing:           "EChatEntryType_Typing",
	EChatEntryType_InviteGame:       "EChatEntryType_InviteGame",
	EChatEntryType_Emote:            "EChatEntryType_Emote",
	EChatEntryType_LobbyGameStart:   "EChatEntryType_LobbyGameStart",
	EChatEntryType_LeftConversation: "EChatEntryType_LeftConversation",
	EChatEntryType_Entered:          "EChatEntryType_Entered",
	EChatEntryType_WasKicked:        "EChatEntryType_WasKicked",
	EChatEntryType_WasBanned:        "EChatEntryType_WasBanned",
	EChatEntryType_Disconnected:     "EChatEntryType_Disconnected",
	EChatEntryType_HistoricalChat:   "EChatEntryType_HistoricalChat",
	EChatEntryType_Reserved1:        "EChatEntryType_Reserved1",
	EChatEntryType_Reserved2:        "EChatEntryType_Reserved2",
	EChatEntryType_LinkBlocked:      "EChatEntryType_LinkBlocked",
}

func (e EChatEntryType) String() string {
	if s, ok := EChatEntryType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EChatEntryType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EPersonaState int32

const (
	EPersonaState_Offline        EPersonaState = 0
	EPersonaState_Online         EPersonaState = 1
	EPersonaState_Busy           EPersonaState = 2
	EPersonaState_Away           EPersonaState = 3
	EPersonaState_Snooze         EPersonaState = 4
	EPersonaState_LookingToTrade EPersonaState = 5
	EPersonaState_LookingToPlay  EPersonaState = 6
	EPersonaState_Invisible      EPersonaState = 7
	EPersonaState_Max            EPersonaState = 8
)

var EPersonaState_name = map[EPersonaState]string{
	EPersonaState_Offline:        "EPersonaState_Offline",
	EPersonaState_Online:         "EPersonaState_Online",
	EPersonaState_Busy:           "EPersonaState_Busy",
	EPersonaState_Away:           "EPersonaState_Away",
	EPersonaState_Snooze:         "EPersonaState_Snooze",
	EPersonaState_LookingToTrade: "EPersonaState_LookingToTrade",
	EPersonaState_LookingToPlay:  "EPersonaState_LookingToPlay",
	EPersonaState_Invisible:      "EPersonaState_Invisible",
	EPersonaState_Max:            "EPersonaState_Max",
}

func (e EPersonaState) String() string {
	if s, ok := EPersonaState_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EPersonaState_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EAccountType int32

const (
	EAccountType_Invalid        EAccountType = 0
	EAccountType_Individual     EAccountType = 1
	EAccountType_Multiseat      EAccountType = 2
	EAccountType_GameServer     EAccountType = 3
	EAccountType_AnonGameServer EAccountType = 4
	EAccountType_Pending        EAccountType = 5
	EAccountType_ContentServer  EAccountType = 6
	EAccountType_Clan           EAccountType = 7
	EAccountType_Chat           EAccountType = 8
	EAccountType_ConsoleUser    EAccountType = 9
	EAccountType_AnonUser       EAccountType = 10
	EAccountType_Max            EAccountType = 11
)

var EAccountType_name = map[EAccountType]string{
	EAccountType_Invalid:        "EAccountType_Invalid",
	EAccountType_Individual:     "EAccountType_Individual",
	EAccountType_Multiseat:      "EAccountType_Multiseat",
	EAccountType_GameServer:     "EAccountType_GameServer",
	EAccountType_AnonGameServer: "EAccountType_AnonGameServer",
	EAccountType_Pending:        "EAccountType_Pending",
	EAccountType_ContentServer:  "EAccountType_ContentServer",
	EAccountType_Clan:           "EAccountType_Clan",
	EAccountType_Chat:           "EAccountType_Chat",
	EAccountType_ConsoleUser:    "EAccountType_ConsoleUser",
	EAccountType_AnonUser:       "EAccountType_AnonUser",
	EAccountType_Max:            "EAccountType_Max",
}

func (e EAccountType) String() string {
	if s, ok := EAccountType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EAccountType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EFriendRelationship int32

const (
	EFriendRelationship_None             EFriendRelationship = 0
	EFriendRelationship_Blocked          EFriendRelationship = 1
	EFriendRelationship_RequestRecipient EFriendRelationship = 2
	EFriendRelationship_Friend           EFriendRelationship = 3
	EFriendRelationship_RequestInitiator EFriendRelationship = 4
	EFriendRelationship_Ignored          EFriendRelationship = 5
	EFriendRelationship_IgnoredFriend    EFriendRelationship = 6
	EFriendRelationship_SuggestedFriend  EFriendRelationship = 7
	EFriendRelationship_Max              EFriendRelationship = 8
)

var EFriendRelationship_name = map[EFriendRelationship]string{
	EFriendRelationship_None:             "EFriendRelationship_None",
	EFriendRelationship_Blocked:          "EFriendRelationship_Blocked",
	EFriendRelationship_RequestRecipient: "EFriendRelationship_RequestRecipient",
	EFriendRelationship_Friend:           "EFriendRelationship_Friend",
	EFriendRelationship_RequestInitiator: "EFriendRelationship_RequestInitiator",
	EFriendRelationship_Ignored:          "EFriendRelationship_Ignored",
	EFriendRelationship_IgnoredFriend:    "EFriendRelationship_IgnoredFriend",
	EFriendRelationship_SuggestedFriend:  "EFriendRelationship_SuggestedFriend",
	EFriendRelationship_Max:              "EFriendRelationship_Max",
}

func (e EFriendRelationship) String() string {
	if s, ok := EFriendRelationship_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EFriendRelationship_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EAccountFlags int32

const (
	EAccountFlags_NormalUser                 EAccountFlags = 0
	EAccountFlags_PersonaNameSet             EAccountFlags = 1
	EAccountFlags_Unbannable                 EAccountFlags = 2
	EAccountFlags_PasswordSet                EAccountFlags = 4
	EAccountFlags_Support                    EAccountFlags = 8
	EAccountFlags_Admin                      EAccountFlags = 16
	EAccountFlags_Supervisor                 EAccountFlags = 32
	EAccountFlags_AppEditor                  EAccountFlags = 64
	EAccountFlags_HWIDSet                    EAccountFlags = 128
	EAccountFlags_PersonalQASet              EAccountFlags = 256
	EAccountFlags_VacBeta                    EAccountFlags = 512
	EAccountFlags_Debug                      EAccountFlags = 1024
	EAccountFlags_Disabled                   EAccountFlags = 2048
	EAccountFlags_LimitedUser                EAccountFlags = 4096
	EAccountFlags_LimitedUserForce           EAccountFlags = 8192
	EAccountFlags_EmailValidated             EAccountFlags = 16384
	EAccountFlags_MarketingTreatment         EAccountFlags = 32768
	EAccountFlags_OGGInviteOptOut            EAccountFlags = 65536
	EAccountFlags_ForcePasswordChange        EAccountFlags = 131072
	EAccountFlags_ForceEmailVerification     EAccountFlags = 262144
	EAccountFlags_LogonExtraSecurity         EAccountFlags = 524288
	EAccountFlags_LogonExtraSecurityDisabled EAccountFlags = 1048576
	EAccountFlags_Steam2MigrationComplete    EAccountFlags = 2097152
	EAccountFlags_NeedLogs                   EAccountFlags = 4194304
	EAccountFlags_Lockdown                   EAccountFlags = 8388608
	EAccountFlags_MasterAppEditor            EAccountFlags = 16777216
	EAccountFlags_BannedFromWebAPI           EAccountFlags = 33554432
	EAccountFlags_ClansOnlyFromFriends       EAccountFlags = 67108864
	EAccountFlags_GlobalModerator            EAccountFlags = 134217728
	EAccountFlags_ParentalSettings           EAccountFlags = 268435456
	EAccountFlags_ThirdPartySupport          EAccountFlags = 536870912
	EAccountFlags_NeedsSSANextSteamLogon     EAccountFlags = 1073741824
)

var EAccountFlags_name = map[EAccountFlags]string{
	EAccountFlags_NormalUser:                 "EAccountFlags_NormalUser",
	EAccountFlags_PersonaNameSet:             "EAccountFlags_PersonaNameSet",
	EAccountFlags_Unbannable:                 "EAccountFlags_Unbannable",
	EAccountFlags_PasswordSet:                "EAccountFlags_PasswordSet",
	EAccountFlags_Support:                    "EAccountFlags_Support",
	EAccountFlags_Admin:                      "EAccountFlags_Admin",
	EAccountFlags_Supervisor:                 "EAccountFlags_Supervisor",
	EAccountFlags_AppEditor:                  "EAccountFlags_AppEditor",
	EAccountFlags_HWIDSet:                    "EAccountFlags_HWIDSet",
	EAccountFlags_PersonalQASet:              "EAccountFlags_PersonalQASet",
	EAccountFlags_VacBeta:                    "EAccountFlags_VacBeta",
	EAccountFlags_Debug:                      "EAccountFlags_Debug",
	EAccountFlags_Disabled:                   "EAccountFlags_Disabled",
	EAccountFlags_LimitedUser:                "EAccountFlags_LimitedUser",
	EAccountFlags_LimitedUserForce:           "EAccountFlags_LimitedUserForce",
	EAccountFlags_EmailValidated:             "EAccountFlags_EmailValidated",
	EAccountFlags_MarketingTreatment:         "EAccountFlags_MarketingTreatment",
	EAccountFlags_OGGInviteOptOut:            "EAccountFlags_OGGInviteOptOut",
	EAccountFlags_ForcePasswordChange:        "EAccountFlags_ForcePasswordChange",
	EAccountFlags_ForceEmailVerification:     "EAccountFlags_ForceEmailVerification",
	EAccountFlags_LogonExtraSecurity:         "EAccountFlags_LogonExtraSecurity",
	EAccountFlags_LogonExtraSecurityDisabled: "EAccountFlags_LogonExtraSecurityDisabled",
	EAccountFlags_Steam2MigrationComplete:    "EAccountFlags_Steam2MigrationComplete",
	EAccountFlags_NeedLogs:                   "EAccountFlags_NeedLogs",
	EAccountFlags_Lockdown:                   "EAccountFlags_Lockdown",
	EAccountFlags_MasterAppEditor:            "EAccountFlags_MasterAppEditor",
	EAccountFlags_BannedFromWebAPI:           "EAccountFlags_BannedFromWebAPI",
	EAccountFlags_ClansOnlyFromFriends:       "EAccountFlags_ClansOnlyFromFriends",
	EAccountFlags_GlobalModerator:            "EAccountFlags_GlobalModerator",
	EAccountFlags_ParentalSettings:           "EAccountFlags_ParentalSettings",
	EAccountFlags_ThirdPartySupport:          "EAccountFlags_ThirdPartySupport",
	EAccountFlags_NeedsSSANextSteamLogon:     "EAccountFlags_NeedsSSANextSteamLogon",
}

func (e EAccountFlags) String() string {
	if s, ok := EAccountFlags_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EAccountFlags_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EClanPermission int32

const (
	EClanPermission_Nobody                EClanPermission = 0
	EClanPermission_Owner                 EClanPermission = 1
	EClanPermission_Officer               EClanPermission = 2
	EClanPermission_OwnerAndOfficer       EClanPermission = 3
	EClanPermission_Member                EClanPermission = 4
	EClanPermission_Moderator             EClanPermission = 8
	EClanPermission_OwnerOfficerModerator EClanPermission = EClanPermission_Owner | EClanPermission_Officer | EClanPermission_Moderator
	EClanPermission_AllMembers            EClanPermission = EClanPermission_Owner | EClanPermission_Officer | EClanPermission_Moderator | EClanPermission_Member
	EClanPermission_OGGGameOwner          EClanPermission = 16
	EClanPermission_NonMember             EClanPermission = 128
	EClanPermission_MemberAllowed         EClanPermission = EClanPermission_NonMember | EClanPermission_Member
	EClanPermission_ModeratorAllowed      EClanPermission = EClanPermission_NonMember | EClanPermission_Member | EClanPermission_Moderator
	EClanPermission_OfficerAllowed        EClanPermission = EClanPermission_NonMember | EClanPermission_Member | EClanPermission_Moderator | EClanPermission_Officer
	EClanPermission_OwnerAllowed          EClanPermission = EClanPermission_NonMember | EClanPermission_Member | EClanPermission_Moderator | EClanPermission_Officer | EClanPermission_Owner
	EClanPermission_Anybody               EClanPermission = EClanPermission_NonMember | EClanPermission_Member | EClanPermission_Moderator | EClanPermission_Officer | EClanPermission_Owner
)

var EClanPermission_name = map[EClanPermission]string{
	EClanPermission_Nobody:                "EClanPermission_Nobody",
	EClanPermission_Owner:                 "EClanPermission_Owner",
	EClanPermission_Officer:               "EClanPermission_Officer",
	EClanPermission_OwnerAndOfficer:       "EClanPermission_OwnerAndOfficer",
	EClanPermission_Member:                "EClanPermission_Member",
	EClanPermission_Moderator:             "EClanPermission_Moderator",
	EClanPermission_OwnerOfficerModerator: "EClanPermission_OwnerOfficerModerator",
	EClanPermission_AllMembers:            "EClanPermission_AllMembers",
	EClanPermission_OGGGameOwner:          "EClanPermission_OGGGameOwner",
	EClanPermission_NonMember:             "EClanPermission_NonMember",
	EClanPermission_MemberAllowed:         "EClanPermission_MemberAllowed",
	EClanPermission_ModeratorAllowed:      "EClanPermission_ModeratorAllowed",
	EClanPermission_OfficerAllowed:        "EClanPermission_OfficerAllowed",
	EClanPermission_OwnerAllowed:          "EClanPermission_OwnerAllowed",
}

func (e EClanPermission) String() string {
	if s, ok := EClanPermission_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EClanPermission_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EChatPermission int32

const (
	EChatPermission_Close                    EChatPermission = 1
	EChatPermission_Invite                   EChatPermission = 2
	EChatPermission_Talk                     EChatPermission = 8
	EChatPermission_Kick                     EChatPermission = 16
	EChatPermission_Mute                     EChatPermission = 32
	EChatPermission_SetMetadata              EChatPermission = 64
	EChatPermission_ChangePermissions        EChatPermission = 128
	EChatPermission_Ban                      EChatPermission = 256
	EChatPermission_ChangeAccess             EChatPermission = 512
	EChatPermission_EveryoneNotInClanDefault EChatPermission = EChatPermission_Talk
	EChatPermission_EveryoneDefault          EChatPermission = EChatPermission_Talk | EChatPermission_Invite
	EChatPermission_MemberDefault            EChatPermission = EChatPermission_Ban | EChatPermission_Kick | EChatPermission_Talk | EChatPermission_Invite
	EChatPermission_OfficerDefault           EChatPermission = EChatPermission_Ban | EChatPermission_Kick | EChatPermission_Talk | EChatPermission_Invite
	EChatPermission_OwnerDefault             EChatPermission = EChatPermission_ChangeAccess | EChatPermission_Ban | EChatPermission_SetMetadata | EChatPermission_Mute | EChatPermission_Kick | EChatPermission_Talk | EChatPermission_Invite | EChatPermission_Close
	EChatPermission_Mask                     EChatPermission = 1019
)

var EChatPermission_name = map[EChatPermission]string{
	EChatPermission_Close:             "EChatPermission_Close",
	EChatPermission_Invite:            "EChatPermission_Invite",
	EChatPermission_Talk:              "EChatPermission_Talk",
	EChatPermission_Kick:              "EChatPermission_Kick",
	EChatPermission_Mute:              "EChatPermission_Mute",
	EChatPermission_SetMetadata:       "EChatPermission_SetMetadata",
	EChatPermission_ChangePermissions: "EChatPermission_ChangePermissions",
	EChatPermission_Ban:               "EChatPermission_Ban",
	EChatPermission_ChangeAccess:      "EChatPermission_ChangeAccess",
	EChatPermission_EveryoneDefault:   "EChatPermission_EveryoneDefault",
	EChatPermission_MemberDefault:     "EChatPermission_MemberDefault",
	EChatPermission_OwnerDefault:      "EChatPermission_OwnerDefault",
	EChatPermission_Mask:              "EChatPermission_Mask",
}

func (e EChatPermission) String() string {
	if s, ok := EChatPermission_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EChatPermission_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EFriendFlags int32

const (
	EFriendFlags_None                 EFriendFlags = 0
	EFriendFlags_Blocked              EFriendFlags = 1
	EFriendFlags_FriendshipRequested  EFriendFlags = 2
	EFriendFlags_Immediate            EFriendFlags = 4
	EFriendFlags_ClanMember           EFriendFlags = 8
	EFriendFlags_OnGameServer         EFriendFlags = 16
	EFriendFlags_RequestingFriendship EFriendFlags = 128
	EFriendFlags_RequestingInfo       EFriendFlags = 256
	EFriendFlags_Ignored              EFriendFlags = 512
	EFriendFlags_IgnoredFriend        EFriendFlags = 1024
	EFriendFlags_Suggested            EFriendFlags = 2048
	EFriendFlags_ChatMember           EFriendFlags = 4096
	EFriendFlags_FlagAll              EFriendFlags = 65535
)

var EFriendFlags_name = map[EFriendFlags]string{
	EFriendFlags_None:                 "EFriendFlags_None",
	EFriendFlags_Blocked:              "EFriendFlags_Blocked",
	EFriendFlags_FriendshipRequested:  "EFriendFlags_FriendshipRequested",
	EFriendFlags_Immediate:            "EFriendFlags_Immediate",
	EFriendFlags_ClanMember:           "EFriendFlags_ClanMember",
	EFriendFlags_OnGameServer:         "EFriendFlags_OnGameServer",
	EFriendFlags_RequestingFriendship: "EFriendFlags_RequestingFriendship",
	EFriendFlags_RequestingInfo:       "EFriendFlags_RequestingInfo",
	EFriendFlags_Ignored:              "EFriendFlags_Ignored",
	EFriendFlags_IgnoredFriend:        "EFriendFlags_IgnoredFriend",
	EFriendFlags_Suggested:            "EFriendFlags_Suggested",
	EFriendFlags_ChatMember:           "EFriendFlags_ChatMember",
	EFriendFlags_FlagAll:              "EFriendFlags_FlagAll",
}

func (e EFriendFlags) String() string {
	if s, ok := EFriendFlags_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EFriendFlags_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EPersonaStateFlag int32

const (
	EPersonaStateFlag_HasRichPresence       EPersonaStateFlag = 1
	EPersonaStateFlag_InJoinableGame        EPersonaStateFlag = 2
	EPersonaStateFlag_Golden                EPersonaStateFlag = 4   // removed: no longer has any effect
	EPersonaStateFlag_OnlineUsingWeb        EPersonaStateFlag = 256 // removed: renamed to ClientTypeWeb
	EPersonaStateFlag_ClientTypeWeb         EPersonaStateFlag = 256
	EPersonaStateFlag_OnlineUsingMobile     EPersonaStateFlag = 512 // removed: renamed to ClientTypeMobile
	EPersonaStateFlag_ClientTypeMobile      EPersonaStateFlag = 512
	EPersonaStateFlag_OnlineUsingBigPicture EPersonaStateFlag = 1024 // removed: renamed to ClientTypeTenfoot
	EPersonaStateFlag_ClientTypeTenfoot     EPersonaStateFlag = 1024
	EPersonaStateFlag_OnlineUsingVR         EPersonaStateFlag = 2048 // removed: renamed to ClientTypeVR
	EPersonaStateFlag_ClientTypeVR          EPersonaStateFlag = 2048
	EPersonaStateFlag_LaunchTypeGamepad     EPersonaStateFlag = 4096
)

var EPersonaStateFlag_name = map[EPersonaStateFlag]string{
	EPersonaStateFlag_HasRichPresence:   "EPersonaStateFlag_HasRichPresence",
	EPersonaStateFlag_InJoinableGame:    "EPersonaStateFlag_InJoinableGame",
	EPersonaStateFlag_Golden:            "EPersonaStateFlag_Golden",
	EPersonaStateFlag_ClientTypeWeb:     "EPersonaStateFlag_ClientTypeWeb",
	EPersonaStateFlag_ClientTypeMobile:  "EPersonaStateFlag_ClientTypeMobile",
	EPersonaStateFlag_ClientTypeTenfoot: "EPersonaStateFlag_ClientTypeTenfoot",
	EPersonaStateFlag_ClientTypeVR:      "EPersonaStateFlag_ClientTypeVR",
	EPersonaStateFlag_LaunchTypeGamepad: "EPersonaStateFlag_LaunchTypeGamepad",
}

func (e EPersonaStateFlag) String() string {
	if s, ok := EPersonaStateFlag_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EPersonaStateFlag_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EClientPersonaStateFlag int32

const (
	EClientPersonaStateFlag_Status        EClientPersonaStateFlag = 1
	EClientPersonaStateFlag_PlayerName    EClientPersonaStateFlag = 2
	EClientPersonaStateFlag_QueryPort     EClientPersonaStateFlag = 4
	EClientPersonaStateFlag_SourceID      EClientPersonaStateFlag = 8
	EClientPersonaStateFlag_Presence      EClientPersonaStateFlag = 16
	EClientPersonaStateFlag_Metadata      EClientPersonaStateFlag = 32 // removed
	EClientPersonaStateFlag_LastSeen      EClientPersonaStateFlag = 64
	EClientPersonaStateFlag_ClanInfo      EClientPersonaStateFlag = 128
	EClientPersonaStateFlag_GameExtraInfo EClientPersonaStateFlag = 256
	EClientPersonaStateFlag_GameDataBlob  EClientPersonaStateFlag = 512
	EClientPersonaStateFlag_ClanTag       EClientPersonaStateFlag = 1024
	EClientPersonaStateFlag_Facebook      EClientPersonaStateFlag = 2048
)

var EClientPersonaStateFlag_name = map[EClientPersonaStateFlag]string{
	EClientPersonaStateFlag_Status:        "EClientPersonaStateFlag_Status",
	EClientPersonaStateFlag_PlayerName:    "EClientPersonaStateFlag_PlayerName",
	EClientPersonaStateFlag_QueryPort:     "EClientPersonaStateFlag_QueryPort",
	EClientPersonaStateFlag_SourceID:      "EClientPersonaStateFlag_SourceID",
	EClientPersonaStateFlag_Presence:      "EClientPersonaStateFlag_Presence",
	EClientPersonaStateFlag_Metadata:      "EClientPersonaStateFlag_Metadata",
	EClientPersonaStateFlag_LastSeen:      "EClientPersonaStateFlag_LastSeen",
	EClientPersonaStateFlag_ClanInfo:      "EClientPersonaStateFlag_ClanInfo",
	EClientPersonaStateFlag_GameExtraInfo: "EClientPersonaStateFlag_GameExtraInfo",
	EClientPersonaStateFlag_GameDataBlob:  "EClientPersonaStateFlag_GameDataBlob",
	EClientPersonaStateFlag_ClanTag:       "EClientPersonaStateFlag_ClanTag",
	EClientPersonaStateFlag_Facebook:      "EClientPersonaStateFlag_Facebook",
}

func (e EClientPersonaStateFlag) String() string {
	if s, ok := EClientPersonaStateFlag_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EClientPersonaStateFlag_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EAppUsageEvent int32

const (
	EAppUsageEvent_GameLaunch            EAppUsageEvent = 1
	EAppUsageEvent_GameLaunchTrial       EAppUsageEvent = 2
	EAppUsageEvent_Media                 EAppUsageEvent = 3
	EAppUsageEvent_PreloadStart          EAppUsageEvent = 4
	EAppUsageEvent_PreloadFinish         EAppUsageEvent = 5
	EAppUsageEvent_MarketingMessageView  EAppUsageEvent = 6
	EAppUsageEvent_InGameAdViewed        EAppUsageEvent = 7
	EAppUsageEvent_GameLaunchFreeWeekend EAppUsageEvent = 8
)

var EAppUsageEvent_name = map[EAppUsageEvent]string{
	EAppUsageEvent_GameLaunch:            "EAppUsageEvent_GameLaunch",
	EAppUsageEvent_GameLaunchTrial:       "EAppUsageEvent_GameLaunchTrial",
	EAppUsageEvent_Media:                 "EAppUsageEvent_Media",
	EAppUsageEvent_PreloadStart:          "EAppUsageEvent_PreloadStart",
	EAppUsageEvent_PreloadFinish:         "EAppUsageEvent_PreloadFinish",
	EAppUsageEvent_MarketingMessageView:  "EAppUsageEvent_MarketingMessageView",
	EAppUsageEvent_InGameAdViewed:        "EAppUsageEvent_InGameAdViewed",
	EAppUsageEvent_GameLaunchFreeWeekend: "EAppUsageEvent_GameLaunchFreeWeekend",
}

func (e EAppUsageEvent) String() string {
	if s, ok := EAppUsageEvent_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EAppUsageEvent_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type ELicenseFlags int32

const (
	ELicenseFlags_None                         ELicenseFlags = 0
	ELicenseFlags_Renew                        ELicenseFlags = 0x01
	ELicenseFlags_RenewalFailed                ELicenseFlags = 0x02
	ELicenseFlags_Pending                      ELicenseFlags = 0x04
	ELicenseFlags_Expired                      ELicenseFlags = 0x08
	ELicenseFlags_CancelledByUser              ELicenseFlags = 0x10
	ELicenseFlags_CancelledByAdmin             ELicenseFlags = 0x20
	ELicenseFlags_LowViolenceContent           ELicenseFlags = 0x40
	ELicenseFlags_ImportedFromSteam2           ELicenseFlags = 0x80
	ELicenseFlags_ForceRunRestriction          ELicenseFlags = 0x100
	ELicenseFlags_RegionRestrictionExpired     ELicenseFlags = 0x200
	ELicenseFlags_CancelledByFriendlyFraudLock ELicenseFlags = 0x400
	ELicenseFlags_NotActivated                 ELicenseFlags = 0x800
)

var ELicenseFlags_name = map[ELicenseFlags]string{
	ELicenseFlags_None:                         "ELicenseFlags_None",
	ELicenseFlags_Renew:                        "ELicenseFlags_Renew",
	ELicenseFlags_RenewalFailed:                "ELicenseFlags_RenewalFailed",
	ELicenseFlags_Pending:                      "ELicenseFlags_Pending",
	ELicenseFlags_Expired:                      "ELicenseFlags_Expired",
	ELicenseFlags_CancelledByUser:              "ELicenseFlags_CancelledByUser",
	ELicenseFlags_CancelledByAdmin:             "ELicenseFlags_CancelledByAdmin",
	ELicenseFlags_LowViolenceContent:           "ELicenseFlags_LowViolenceContent",
	ELicenseFlags_ImportedFromSteam2:           "ELicenseFlags_ImportedFromSteam2",
	ELicenseFlags_ForceRunRestriction:          "ELicenseFlags_ForceRunRestriction",
	ELicenseFlags_RegionRestrictionExpired:     "ELicenseFlags_RegionRestrictionExpired",
	ELicenseFlags_CancelledByFriendlyFraudLock: "ELicenseFlags_CancelledByFriendlyFraudLock",
	ELicenseFlags_NotActivated:                 "ELicenseFlags_NotActivated",
}

func (e ELicenseFlags) String() string {
	if s, ok := ELicenseFlags_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range ELicenseFlags_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type ELicenseType int32

const (
	ELicenseType_NoLicense                             ELicenseType = 0
	ELicenseType_SinglePurchase                        ELicenseType = 1
	ELicenseType_SinglePurchaseLimitedUse              ELicenseType = 2
	ELicenseType_RecurringCharge                       ELicenseType = 3
	ELicenseType_RecurringChargeLimitedUse             ELicenseType = 4
	ELicenseType_RecurringChargeLimitedUseWithOverages ELicenseType = 5
	ELicenseType_RecurringOption                       ELicenseType = 6
	ELicenseType_LimitedUseDelayedActivation           ELicenseType = 7
)

var ELicenseType_name = map[ELicenseType]string{
	ELicenseType_NoLicense:                             "ELicenseType_NoLicense",
	ELicenseType_SinglePurchase:                        "ELicenseType_SinglePurchase",
	ELicenseType_SinglePurchaseLimitedUse:              "ELicenseType_SinglePurchaseLimitedUse",
	ELicenseType_RecurringCharge:                       "ELicenseType_RecurringCharge",
	ELicenseType_RecurringChargeLimitedUse:             "ELicenseType_RecurringChargeLimitedUse",
	ELicenseType_RecurringChargeLimitedUseWithOverages: "ELicenseType_RecurringChargeLimitedUseWithOverages",
	ELicenseType_RecurringOption:                       "ELicenseType_RecurringOption",
	ELicenseType_LimitedUseDelayedActivation:           "ELicenseType_LimitedUseDelayedActivation",
}

func (e ELicenseType) String() string {
	if s, ok := ELicenseType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range ELicenseType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EPaymentMethod int32

const (
	EPaymentMethod_None                   EPaymentMethod = 0
	EPaymentMethod_ActivationCode         EPaymentMethod = 1
	EPaymentMethod_CreditCard             EPaymentMethod = 2
	EPaymentMethod_Giropay                EPaymentMethod = 3
	EPaymentMethod_PayPal                 EPaymentMethod = 4
	EPaymentMethod_Ideal                  EPaymentMethod = 5
	EPaymentMethod_PaySafeCard            EPaymentMethod = 6
	EPaymentMethod_Sofort                 EPaymentMethod = 7
	EPaymentMethod_GuestPass              EPaymentMethod = 8
	EPaymentMethod_WebMoney               EPaymentMethod = 9
	EPaymentMethod_MoneyBookers           EPaymentMethod = 10
	EPaymentMethod_AliPay                 EPaymentMethod = 11
	EPaymentMethod_Yandex                 EPaymentMethod = 12
	EPaymentMethod_Kiosk                  EPaymentMethod = 13
	EPaymentMethod_Qiwi                   EPaymentMethod = 14
	EPaymentMethod_GameStop               EPaymentMethod = 15
	EPaymentMethod_HardwarePromo          EPaymentMethod = 16
	EPaymentMethod_MoPay                  EPaymentMethod = 17
	EPaymentMethod_BoletoBancario         EPaymentMethod = 18
	EPaymentMethod_BoaCompraGold          EPaymentMethod = 19
	EPaymentMethod_BancoDoBrasilOnline    EPaymentMethod = 20
	EPaymentMethod_ItauOnline             EPaymentMethod = 21
	EPaymentMethod_BradescoOnline         EPaymentMethod = 22
	EPaymentMethod_Pagseguro              EPaymentMethod = 23
	EPaymentMethod_VisaBrazil             EPaymentMethod = 24
	EPaymentMethod_AmexBrazil             EPaymentMethod = 25
	EPaymentMethod_Aura                   EPaymentMethod = 26
	EPaymentMethod_Hipercard              EPaymentMethod = 27
	EPaymentMethod_MastercardBrazil       EPaymentMethod = 28
	EPaymentMethod_DinersCardBrazil       EPaymentMethod = 29
	EPaymentMethod_AuthorizedDevice       EPaymentMethod = 30
	EPaymentMethod_MOLPoints              EPaymentMethod = 31
	EPaymentMethod_ClickAndBuy            EPaymentMethod = 32
	EPaymentMethod_Beeline                EPaymentMethod = 33
	EPaymentMethod_Konbini                EPaymentMethod = 34
	EPaymentMethod_EClubPoints            EPaymentMethod = 35
	EPaymentMethod_CreditCardJapan        EPaymentMethod = 36
	EPaymentMethod_BankTransferJapan      EPaymentMethod = 37
	EPaymentMethod_PayEasyJapan           EPaymentMethod = 38 // removed: renamed to PayEasy
	EPaymentMethod_PayEasy                EPaymentMethod = 38
	EPaymentMethod_Zong                   EPaymentMethod = 39
	EPaymentMethod_CultureVoucher         EPaymentMethod = 40
	EPaymentMethod_BookVoucher            EPaymentMethod = 41
	EPaymentMethod_HappymoneyVoucher      EPaymentMethod = 42
	EPaymentMethod_ConvenientStoreVoucher EPaymentMethod = 43
	EPaymentMethod_GameVoucher            EPaymentMethod = 44
	EPaymentMethod_Multibanco             EPaymentMethod = 45
	EPaymentMethod_Payshop                EPaymentMethod = 46
	EPaymentMethod_Maestro                EPaymentMethod = 47 // removed: renamed to MaestroBoaCompra
	EPaymentMethod_MaestroBoaCompra       EPaymentMethod = 47
	EPaymentMethod_OXXO                   EPaymentMethod = 48
	EPaymentMethod_ToditoCash             EPaymentMethod = 49
	EPaymentMethod_Carnet                 EPaymentMethod = 50
	EPaymentMethod_SPEI                   EPaymentMethod = 51
	EPaymentMethod_ThreePay               EPaymentMethod = 52
	EPaymentMethod_IsBank                 EPaymentMethod = 53
	EPaymentMethod_Garanti                EPaymentMethod = 54
	EPaymentMethod_Akbank                 EPaymentMethod = 55
	EPaymentMethod_YapiKredi              EPaymentMethod = 56
	EPaymentMethod_Halkbank               EPaymentMethod = 57
	EPaymentMethod_BankAsya               EPaymentMethod = 58
	EPaymentMethod_Finansbank             EPaymentMethod = 59
	EPaymentMethod_DenizBank              EPaymentMethod = 60
	EPaymentMethod_PTT                    EPaymentMethod = 61
	EPaymentMethod_CashU                  EPaymentMethod = 62
	EPaymentMethod_AutoGrant              EPaymentMethod = 64
	EPaymentMethod_WebMoneyJapan          EPaymentMethod = 65
	EPaymentMethod_OneCard                EPaymentMethod = 66
	EPaymentMethod_PSE                    EPaymentMethod = 67
	EPaymentMethod_Exito                  EPaymentMethod = 68
	EPaymentMethod_Efecty                 EPaymentMethod = 69
	EPaymentMethod_Paloto                 EPaymentMethod = 70
	EPaymentMethod_PinValidda             EPaymentMethod = 71
	EPaymentMethod_MangirKart             EPaymentMethod = 72
	EPaymentMethod_BancoCreditoDePeru     EPaymentMethod = 73
	EPaymentMethod_BBVAContinental        EPaymentMethod = 74
	EPaymentMethod_SafetyPay              EPaymentMethod = 75
	EPaymentMethod_PagoEfectivo           EPaymentMethod = 76
	EPaymentMethod_Trustly                EPaymentMethod = 77
	EPaymentMethod_UnionPay               EPaymentMethod = 78
	EPaymentMethod_BitCoin                EPaymentMethod = 79
	EPaymentMethod_Wallet                 EPaymentMethod = 128
	EPaymentMethod_Valve                  EPaymentMethod = 129
	EPaymentMethod_SteamPressMaster       EPaymentMethod = 130 // removed: renamed to MasterComp
	EPaymentMethod_MasterComp             EPaymentMethod = 130
	EPaymentMethod_StorePromotion         EPaymentMethod = 131 // removed: renamed to Promotional
	EPaymentMethod_Promotional            EPaymentMethod = 131
	EPaymentMethod_OEMTicket              EPaymentMethod = 256
	EPaymentMethod_Split                  EPaymentMethod = 512
	EPaymentMethod_Complimentary          EPaymentMethod = 1024
)

var EPaymentMethod_name = map[EPaymentMethod]string{
	EPaymentMethod_None:                   "EPaymentMethod_None",
	EPaymentMethod_ActivationCode:         "EPaymentMethod_ActivationCode",
	EPaymentMethod_CreditCard:             "EPaymentMethod_CreditCard",
	EPaymentMethod_Giropay:                "EPaymentMethod_Giropay",
	EPaymentMethod_PayPal:                 "EPaymentMethod_PayPal",
	EPaymentMethod_Ideal:                  "EPaymentMethod_Ideal",
	EPaymentMethod_PaySafeCard:            "EPaymentMethod_PaySafeCard",
	EPaymentMethod_Sofort:                 "EPaymentMethod_Sofort",
	EPaymentMethod_GuestPass:              "EPaymentMethod_GuestPass",
	EPaymentMethod_WebMoney:               "EPaymentMethod_WebMoney",
	EPaymentMethod_MoneyBookers:           "EPaymentMethod_MoneyBookers",
	EPaymentMethod_AliPay:                 "EPaymentMethod_AliPay",
	EPaymentMethod_Yandex:                 "EPaymentMethod_Yandex",
	EPaymentMethod_Kiosk:                  "EPaymentMethod_Kiosk",
	EPaymentMethod_Qiwi:                   "EPaymentMethod_Qiwi",
	EPaymentMethod_GameStop:               "EPaymentMethod_GameStop",
	EPaymentMethod_HardwarePromo:          "EPaymentMethod_HardwarePromo",
	EPaymentMethod_MoPay:                  "EPaymentMethod_MoPay",
	EPaymentMethod_BoletoBancario:         "EPaymentMethod_BoletoBancario",
	EPaymentMethod_BoaCompraGold:          "EPaymentMethod_BoaCompraGold",
	EPaymentMethod_BancoDoBrasilOnline:    "EPaymentMethod_BancoDoBrasilOnline",
	EPaymentMethod_ItauOnline:             "EPaymentMethod_ItauOnline",
	EPaymentMethod_BradescoOnline:         "EPaymentMethod_BradescoOnline",
	EPaymentMethod_Pagseguro:              "EPaymentMethod_Pagseguro",
	EPaymentMethod_VisaBrazil:             "EPaymentMethod_VisaBrazil",
	EPaymentMethod_AmexBrazil:             "EPaymentMethod_AmexBrazil",
	EPaymentMethod_Aura:                   "EPaymentMethod_Aura",
	EPaymentMethod_Hipercard:              "EPaymentMethod_Hipercard",
	EPaymentMethod_MastercardBrazil:       "EPaymentMethod_MastercardBrazil",
	EPaymentMethod_DinersCardBrazil:       "EPaymentMethod_DinersCardBrazil",
	EPaymentMethod_AuthorizedDevice:       "EPaymentMethod_AuthorizedDevice",
	EPaymentMethod_MOLPoints:              "EPaymentMethod_MOLPoints",
	EPaymentMethod_ClickAndBuy:            "EPaymentMethod_ClickAndBuy",
	EPaymentMethod_Beeline:                "EPaymentMethod_Beeline",
	EPaymentMethod_Konbini:                "EPaymentMethod_Konbini",
	EPaymentMethod_EClubPoints:            "EPaymentMethod_EClubPoints",
	EPaymentMethod_CreditCardJapan:        "EPaymentMethod_CreditCardJapan",
	EPaymentMethod_BankTransferJapan:      "EPaymentMethod_BankTransferJapan",
	EPaymentMethod_PayEasy:                "EPaymentMethod_PayEasy",
	EPaymentMethod_Zong:                   "EPaymentMethod_Zong",
	EPaymentMethod_CultureVoucher:         "EPaymentMethod_CultureVoucher",
	EPaymentMethod_BookVoucher:            "EPaymentMethod_BookVoucher",
	EPaymentMethod_HappymoneyVoucher:      "EPaymentMethod_HappymoneyVoucher",
	EPaymentMethod_ConvenientStoreVoucher: "EPaymentMethod_ConvenientStoreVoucher",
	EPaymentMethod_GameVoucher:            "EPaymentMethod_GameVoucher",
	EPaymentMethod_Multibanco:             "EPaymentMethod_Multibanco",
	EPaymentMethod_Payshop:                "EPaymentMethod_Payshop",
	EPaymentMethod_MaestroBoaCompra:       "EPaymentMethod_MaestroBoaCompra",
	EPaymentMethod_OXXO:                   "EPaymentMethod_OXXO",
	EPaymentMethod_ToditoCash:             "EPaymentMethod_ToditoCash",
	EPaymentMethod_Carnet:                 "EPaymentMethod_Carnet",
	EPaymentMethod_SPEI:                   "EPaymentMethod_SPEI",
	EPaymentMethod_ThreePay:               "EPaymentMethod_ThreePay",
	EPaymentMethod_IsBank:                 "EPaymentMethod_IsBank",
	EPaymentMethod_Garanti:                "EPaymentMethod_Garanti",
	EPaymentMethod_Akbank:                 "EPaymentMethod_Akbank",
	EPaymentMethod_YapiKredi:              "EPaymentMethod_YapiKredi",
	EPaymentMethod_Halkbank:               "EPaymentMethod_Halkbank",
	EPaymentMethod_BankAsya:               "EPaymentMethod_BankAsya",
	EPaymentMethod_Finansbank:             "EPaymentMethod_Finansbank",
	EPaymentMethod_DenizBank:              "EPaymentMethod_DenizBank",
	EPaymentMethod_PTT:                    "EPaymentMethod_PTT",
	EPaymentMethod_CashU:                  "EPaymentMethod_CashU",
	EPaymentMethod_AutoGrant:              "EPaymentMethod_AutoGrant",
	EPaymentMethod_WebMoneyJapan:          "EPaymentMethod_WebMoneyJapan",
	EPaymentMethod_OneCard:                "EPaymentMethod_OneCard",
	EPaymentMethod_PSE:                    "EPaymentMethod_PSE",
	EPaymentMethod_Exito:                  "EPaymentMethod_Exito",
	EPaymentMethod_Efecty:                 "EPaymentMethod_Efecty",
	EPaymentMethod_Paloto:                 "EPaymentMethod_Paloto",
	EPaymentMethod_PinValidda:             "EPaymentMethod_PinValidda",
	EPaymentMethod_MangirKart:             "EPaymentMethod_MangirKart",
	EPaymentMethod_BancoCreditoDePeru:     "EPaymentMethod_BancoCreditoDePeru",
	EPaymentMethod_BBVAContinental:        "EPaymentMethod_BBVAContinental",
	EPaymentMethod_SafetyPay:              "EPaymentMethod_SafetyPay",
	EPaymentMethod_PagoEfectivo:           "EPaymentMethod_PagoEfectivo",
	EPaymentMethod_Trustly:                "EPaymentMethod_Trustly",
	EPaymentMethod_UnionPay:               "EPaymentMethod_UnionPay",
	EPaymentMethod_BitCoin:                "EPaymentMethod_BitCoin",
	EPaymentMethod_Wallet:                 "EPaymentMethod_Wallet",
	EPaymentMethod_Valve:                  "EPaymentMethod_Valve",
	EPaymentMethod_MasterComp:             "EPaymentMethod_MasterComp",
	EPaymentMethod_Promotional:            "EPaymentMethod_Promotional",
	EPaymentMethod_OEMTicket:              "EPaymentMethod_OEMTicket",
	EPaymentMethod_Split:                  "EPaymentMethod_Split",
	EPaymentMethod_Complimentary:          "EPaymentMethod_Complimentary",
}

func (e EPaymentMethod) String() string {
	if s, ok := EPaymentMethod_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EPaymentMethod_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EPurchaseResultDetail int32

const (
	EPurchaseResultDetail_NoDetail                                EPurchaseResultDetail = 0
	EPurchaseResultDetail_AVSFailure                              EPurchaseResultDetail = 1
	EPurchaseResultDetail_InsufficientFunds                       EPurchaseResultDetail = 2
	EPurchaseResultDetail_ContactSupport                          EPurchaseResultDetail = 3
	EPurchaseResultDetail_Timeout                                 EPurchaseResultDetail = 4
	EPurchaseResultDetail_InvalidPackage                          EPurchaseResultDetail = 5
	EPurchaseResultDetail_InvalidPaymentMethod                    EPurchaseResultDetail = 6
	EPurchaseResultDetail_InvalidData                             EPurchaseResultDetail = 7
	EPurchaseResultDetail_OthersInProgress                        EPurchaseResultDetail = 8
	EPurchaseResultDetail_AlreadyPurchased                        EPurchaseResultDetail = 9
	EPurchaseResultDetail_WrongPrice                              EPurchaseResultDetail = 10
	EPurchaseResultDetail_FraudCheckFailed                        EPurchaseResultDetail = 11
	EPurchaseResultDetail_CancelledByUser                         EPurchaseResultDetail = 12
	EPurchaseResultDetail_RestrictedCountry                       EPurchaseResultDetail = 13
	EPurchaseResultDetail_BadActivationCode                       EPurchaseResultDetail = 14
	EPurchaseResultDetail_DuplicateActivationCode                 EPurchaseResultDetail = 15
	EPurchaseResultDetail_UseOtherPaymentMethod                   EPurchaseResultDetail = 16
	EPurchaseResultDetail_UseOtherFunctionSource                  EPurchaseResultDetail = 17
	EPurchaseResultDetail_InvalidShippingAddress                  EPurchaseResultDetail = 18
	EPurchaseResultDetail_RegionNotSupported                      EPurchaseResultDetail = 19
	EPurchaseResultDetail_AcctIsBlocked                           EPurchaseResultDetail = 20
	EPurchaseResultDetail_AcctNotVerified                         EPurchaseResultDetail = 21
	EPurchaseResultDetail_InvalidAccount                          EPurchaseResultDetail = 22
	EPurchaseResultDetail_StoreBillingCountryMismatch             EPurchaseResultDetail = 23
	EPurchaseResultDetail_DoesNotOwnRequiredApp                   EPurchaseResultDetail = 24
	EPurchaseResultDetail_CanceledByNewTransaction                EPurchaseResultDetail = 25
	EPurchaseResultDetail_ForceCanceledPending                    EPurchaseResultDetail = 26
	EPurchaseResultDetail_FailCurrencyTransProvider               EPurchaseResultDetail = 27
	EPurchaseResultDetail_FailedCyberCafe                         EPurchaseResultDetail = 28
	EPurchaseResultDetail_NeedsPreApproval                        EPurchaseResultDetail = 29
	EPurchaseResultDetail_PreApprovalDenied                       EPurchaseResultDetail = 30
	EPurchaseResultDetail_WalletCurrencyMismatch                  EPurchaseResultDetail = 31
	EPurchaseResultDetail_EmailNotValidated                       EPurchaseResultDetail = 32
	EPurchaseResultDetail_ExpiredCard                             EPurchaseResultDetail = 33
	EPurchaseResultDetail_TransactionExpired                      EPurchaseResultDetail = 34
	EPurchaseResultDetail_WouldExceedMaxWallet                    EPurchaseResultDetail = 35
	EPurchaseResultDetail_MustLoginPS3AppForPurchase              EPurchaseResultDetail = 36
	EPurchaseResultDetail_CannotShipToPOBox                       EPurchaseResultDetail = 37
	EPurchaseResultDetail_InsufficientInventory                   EPurchaseResultDetail = 38
	EPurchaseResultDetail_CannotGiftShippedGoods                  EPurchaseResultDetail = 39
	EPurchaseResultDetail_CannotShipInternationally               EPurchaseResultDetail = 40
	EPurchaseResultDetail_BillingAgreementCancelled               EPurchaseResultDetail = 41
	EPurchaseResultDetail_InvalidCoupon                           EPurchaseResultDetail = 42
	EPurchaseResultDetail_ExpiredCoupon                           EPurchaseResultDetail = 43
	EPurchaseResultDetail_AccountLocked                           EPurchaseResultDetail = 44
	EPurchaseResultDetail_OtherAbortableInProgress                EPurchaseResultDetail = 45
	EPurchaseResultDetail_ExceededSteamLimit                      EPurchaseResultDetail = 46
	EPurchaseResultDetail_OverlappingPackagesInCart               EPurchaseResultDetail = 47
	EPurchaseResultDetail_NoWallet                                EPurchaseResultDetail = 48
	EPurchaseResultDetail_NoCachedPaymentMethod                   EPurchaseResultDetail = 49
	EPurchaseResultDetail_CannotRedeemCodeFromClient              EPurchaseResultDetail = 50
	EPurchaseResultDetail_PurchaseAmountNoSupportedByProvider     EPurchaseResultDetail = 51
	EPurchaseResultDetail_OverlappingPackagesInPendingTransaction EPurchaseResultDetail = 52
	EPurchaseResultDetail_RateLimited                             EPurchaseResultDetail = 53
	EPurchaseResultDetail_OwnsExcludedApp                         EPurchaseResultDetail = 54
	EPurchaseResultDetail_CreditCardBinMismatchesType             EPurchaseResultDetail = 55
	EPurchaseResultDetail_CartValueTooHigh                        EPurchaseResultDetail = 56
	EPurchaseResultDetail_BillingAgreementAlreadyExists           EPurchaseResultDetail = 57
	EPurchaseResultDetail_POSACodeNotActivated                    EPurchaseResultDetail = 58
	EPurchaseResultDetail_CannotShipToCountry                     EPurchaseResultDetail = 59
	EPurchaseResultDetail_HungTransactionCancelled                EPurchaseResultDetail = 60
	EPurchaseResultDetail_PaypalInternalError                     EPurchaseResultDetail = 61
	EPurchaseResultDetail_UnknownGlobalCollectError               EPurchaseResultDetail = 62
	EPurchaseResultDetail_InvalidTaxAddress                       EPurchaseResultDetail = 63
	EPurchaseResultDetail_PhysicalProductLimitExceeded            EPurchaseResultDetail = 64
	EPurchaseResultDetail_PurchaseCannotBeReplayed                EPurchaseResultDetail = 65
	EPurchaseResultDetail_DelayedCompletion                       EPurchaseResultDetail = 66
	EPurchaseResultDetail_BundleTypeCannotBeGifted                EPurchaseResultDetail = 67
)

var EPurchaseResultDetail_name = map[EPurchaseResultDetail]string{
	EPurchaseResultDetail_NoDetail:                                "EPurchaseResultDetail_NoDetail",
	EPurchaseResultDetail_AVSFailure:                              "EPurchaseResultDetail_AVSFailure",
	EPurchaseResultDetail_InsufficientFunds:                       "EPurchaseResultDetail_InsufficientFunds",
	EPurchaseResultDetail_ContactSupport:                          "EPurchaseResultDetail_ContactSupport",
	EPurchaseResultDetail_Timeout:                                 "EPurchaseResultDetail_Timeout",
	EPurchaseResultDetail_InvalidPackage:                          "EPurchaseResultDetail_InvalidPackage",
	EPurchaseResultDetail_InvalidPaymentMethod:                    "EPurchaseResultDetail_InvalidPaymentMethod",
	EPurchaseResultDetail_InvalidData:                             "EPurchaseResultDetail_InvalidData",
	EPurchaseResultDetail_OthersInProgress:                        "EPurchaseResultDetail_OthersInProgress",
	EPurchaseResultDetail_AlreadyPurchased:                        "EPurchaseResultDetail_AlreadyPurchased",
	EPurchaseResultDetail_WrongPrice:                              "EPurchaseResultDetail_WrongPrice",
	EPurchaseResultDetail_FraudCheckFailed:                        "EPurchaseResultDetail_FraudCheckFailed",
	EPurchaseResultDetail_CancelledByUser:                         "EPurchaseResultDetail_CancelledByUser",
	EPurchaseResultDetail_RestrictedCountry:                       "EPurchaseResultDetail_RestrictedCountry",
	EPurchaseResultDetail_BadActivationCode:                       "EPurchaseResultDetail_BadActivationCode",
	EPurchaseResultDetail_DuplicateActivationCode:                 "EPurchaseResultDetail_DuplicateActivationCode",
	EPurchaseResultDetail_UseOtherPaymentMethod:                   "EPurchaseResultDetail_UseOtherPaymentMethod",
	EPurchaseResultDetail_UseOtherFunctionSource:                  "EPurchaseResultDetail_UseOtherFunctionSource",
	EPurchaseResultDetail_InvalidShippingAddress:                  "EPurchaseResultDetail_InvalidShippingAddress",
	EPurchaseResultDetail_RegionNotSupported:                      "EPurchaseResultDetail_RegionNotSupported",
	EPurchaseResultDetail_AcctIsBlocked:                           "EPurchaseResultDetail_AcctIsBlocked",
	EPurchaseResultDetail_AcctNotVerified:                         "EPurchaseResultDetail_AcctNotVerified",
	EPurchaseResultDetail_InvalidAccount:                          "EPurchaseResultDetail_InvalidAccount",
	EPurchaseResultDetail_StoreBillingCountryMismatch:             "EPurchaseResultDetail_StoreBillingCountryMismatch",
	EPurchaseResultDetail_DoesNotOwnRequiredApp:                   "EPurchaseResultDetail_DoesNotOwnRequiredApp",
	EPurchaseResultDetail_CanceledByNewTransaction:                "EPurchaseResultDetail_CanceledByNewTransaction",
	EPurchaseResultDetail_ForceCanceledPending:                    "EPurchaseResultDetail_ForceCanceledPending",
	EPurchaseResultDetail_FailCurrencyTransProvider:               "EPurchaseResultDetail_FailCurrencyTransProvider",
	EPurchaseResultDetail_FailedCyberCafe:                         "EPurchaseResultDetail_FailedCyberCafe",
	EPurchaseResultDetail_NeedsPreApproval:                        "EPurchaseResultDetail_NeedsPreApproval",
	EPurchaseResultDetail_PreApprovalDenied:                       "EPurchaseResultDetail_PreApprovalDenied",
	EPurchaseResultDetail_WalletCurrencyMismatch:                  "EPurchaseResultDetail_WalletCurrencyMismatch",
	EPurchaseResultDetail_EmailNotValidated:                       "EPurchaseResultDetail_EmailNotValidated",
	EPurchaseResultDetail_ExpiredCard:                             "EPurchaseResultDetail_ExpiredCard",
	EPurchaseResultDetail_TransactionExpired:                      "EPurchaseResultDetail_TransactionExpired",
	EPurchaseResultDetail_WouldExceedMaxWallet:                    "EPurchaseResultDetail_WouldExceedMaxWallet",
	EPurchaseResultDetail_MustLoginPS3AppForPurchase:              "EPurchaseResultDetail_MustLoginPS3AppForPurchase",
	EPurchaseResultDetail_CannotShipToPOBox:                       "EPurchaseResultDetail_CannotShipToPOBox",
	EPurchaseResultDetail_InsufficientInventory:                   "EPurchaseResultDetail_InsufficientInventory",
	EPurchaseResultDetail_CannotGiftShippedGoods:                  "EPurchaseResultDetail_CannotGiftShippedGoods",
	EPurchaseResultDetail_CannotShipInternationally:               "EPurchaseResultDetail_CannotShipInternationally",
	EPurchaseResultDetail_BillingAgreementCancelled:               "EPurchaseResultDetail_BillingAgreementCancelled",
	EPurchaseResultDetail_InvalidCoupon:                           "EPurchaseResultDetail_InvalidCoupon",
	EPurchaseResultDetail_ExpiredCoupon:                           "EPurchaseResultDetail_ExpiredCoupon",
	EPurchaseResultDetail_AccountLocked:                           "EPurchaseResultDetail_AccountLocked",
	EPurchaseResultDetail_OtherAbortableInProgress:                "EPurchaseResultDetail_OtherAbortableInProgress",
	EPurchaseResultDetail_ExceededSteamLimit:                      "EPurchaseResultDetail_ExceededSteamLimit",
	EPurchaseResultDetail_OverlappingPackagesInCart:               "EPurchaseResultDetail_OverlappingPackagesInCart",
	EPurchaseResultDetail_NoWallet:                                "EPurchaseResultDetail_NoWallet",
	EPurchaseResultDetail_NoCachedPaymentMethod:                   "EPurchaseResultDetail_NoCachedPaymentMethod",
	EPurchaseResultDetail_CannotRedeemCodeFromClient:              "EPurchaseResultDetail_CannotRedeemCodeFromClient",
	EPurchaseResultDetail_PurchaseAmountNoSupportedByProvider:     "EPurchaseResultDetail_PurchaseAmountNoSupportedByProvider",
	EPurchaseResultDetail_OverlappingPackagesInPendingTransaction: "EPurchaseResultDetail_OverlappingPackagesInPendingTransaction",
	EPurchaseResultDetail_RateLimited:                             "EPurchaseResultDetail_RateLimited",
	EPurchaseResultDetail_OwnsExcludedApp:                         "EPurchaseResultDetail_OwnsExcludedApp",
	EPurchaseResultDetail_CreditCardBinMismatchesType:             "EPurchaseResultDetail_CreditCardBinMismatchesType",
	EPurchaseResultDetail_CartValueTooHigh:                        "EPurchaseResultDetail_CartValueTooHigh",
	EPurchaseResultDetail_BillingAgreementAlreadyExists:           "EPurchaseResultDetail_BillingAgreementAlreadyExists",
	EPurchaseResultDetail_POSACodeNotActivated:                    "EPurchaseResultDetail_POSACodeNotActivated",
	EPurchaseResultDetail_CannotShipToCountry:                     "EPurchaseResultDetail_CannotShipToCountry",
	EPurchaseResultDetail_HungTransactionCancelled:                "EPurchaseResultDetail_HungTransactionCancelled",
	EPurchaseResultDetail_PaypalInternalError:                     "EPurchaseResultDetail_PaypalInternalError",
	EPurchaseResultDetail_UnknownGlobalCollectError:               "EPurchaseResultDetail_UnknownGlobalCollectError",
	EPurchaseResultDetail_InvalidTaxAddress:                       "EPurchaseResultDetail_InvalidTaxAddress",
	EPurchaseResultDetail_PhysicalProductLimitExceeded:            "EPurchaseResultDetail_PhysicalProductLimitExceeded",
	EPurchaseResultDetail_PurchaseCannotBeReplayed:                "EPurchaseResultDetail_PurchaseCannotBeReplayed",
	EPurchaseResultDetail_DelayedCompletion:                       "EPurchaseResultDetail_DelayedCompletion",
	EPurchaseResultDetail_BundleTypeCannotBeGifted:                "EPurchaseResultDetail_BundleTypeCannotBeGifted",
}

func (e EPurchaseResultDetail) String() string {
	if s, ok := EPurchaseResultDetail_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EPurchaseResultDetail_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EIntroducerRouting int32

const (
	EIntroducerRouting_FileShare     EIntroducerRouting = 0 // removed
	EIntroducerRouting_P2PVoiceChat  EIntroducerRouting = 1
	EIntroducerRouting_P2PNetworking EIntroducerRouting = 2
)

var EIntroducerRouting_name = map[EIntroducerRouting]string{
	EIntroducerRouting_FileShare:     "EIntroducerRouting_FileShare",
	EIntroducerRouting_P2PVoiceChat:  "EIntroducerRouting_P2PVoiceChat",
	EIntroducerRouting_P2PNetworking: "EIntroducerRouting_P2PNetworking",
}

func (e EIntroducerRouting) String() string {
	if s, ok := EIntroducerRouting_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EIntroducerRouting_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EServerFlags int32

const (
	EServerFlags_None       EServerFlags = 0
	EServerFlags_Active     EServerFlags = 1
	EServerFlags_Secure     EServerFlags = 2
	EServerFlags_Dedicated  EServerFlags = 4
	EServerFlags_Linux      EServerFlags = 8
	EServerFlags_Passworded EServerFlags = 16
	EServerFlags_Private    EServerFlags = 32
)

var EServerFlags_name = map[EServerFlags]string{
	EServerFlags_None:       "EServerFlags_None",
	EServerFlags_Active:     "EServerFlags_Active",
	EServerFlags_Secure:     "EServerFlags_Secure",
	EServerFlags_Dedicated:  "EServerFlags_Dedicated",
	EServerFlags_Linux:      "EServerFlags_Linux",
	EServerFlags_Passworded: "EServerFlags_Passworded",
	EServerFlags_Private:    "EServerFlags_Private",
}

func (e EServerFlags) String() string {
	if s, ok := EServerFlags_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EServerFlags_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EDenyReason int32

const (
	EDenyReason_InvalidVersion          EDenyReason = 1
	EDenyReason_Generic                 EDenyReason = 2
	EDenyReason_NotLoggedOn             EDenyReason = 3
	EDenyReason_NoLicense               EDenyReason = 4
	EDenyReason_Cheater                 EDenyReason = 5
	EDenyReason_LoggedInElseWhere       EDenyReason = 6
	EDenyReason_UnknownText             EDenyReason = 7
	EDenyReason_IncompatibleAnticheat   EDenyReason = 8
	EDenyReason_MemoryCorruption        EDenyReason = 9
	EDenyReason_IncompatibleSoftware    EDenyReason = 10
	EDenyReason_SteamConnectionLost     EDenyReason = 11
	EDenyReason_SteamConnectionError    EDenyReason = 12
	EDenyReason_SteamResponseTimedOut   EDenyReason = 13
	EDenyReason_SteamValidationStalled  EDenyReason = 14
	EDenyReason_SteamOwnerLeftGuestUser EDenyReason = 15
)

var EDenyReason_name = map[EDenyReason]string{
	EDenyReason_InvalidVersion:          "EDenyReason_InvalidVersion",
	EDenyReason_Generic:                 "EDenyReason_Generic",
	EDenyReason_NotLoggedOn:             "EDenyReason_NotLoggedOn",
	EDenyReason_NoLicense:               "EDenyReason_NoLicense",
	EDenyReason_Cheater:                 "EDenyReason_Cheater",
	EDenyReason_LoggedInElseWhere:       "EDenyReason_LoggedInElseWhere",
	EDenyReason_UnknownText:             "EDenyReason_UnknownText",
	EDenyReason_IncompatibleAnticheat:   "EDenyReason_IncompatibleAnticheat",
	EDenyReason_MemoryCorruption:        "EDenyReason_MemoryCorruption",
	EDenyReason_IncompatibleSoftware:    "EDenyReason_IncompatibleSoftware",
	EDenyReason_SteamConnectionLost:     "EDenyReason_SteamConnectionLost",
	EDenyReason_SteamConnectionError:    "EDenyReason_SteamConnectionError",
	EDenyReason_SteamResponseTimedOut:   "EDenyReason_SteamResponseTimedOut",
	EDenyReason_SteamValidationStalled:  "EDenyReason_SteamValidationStalled",
	EDenyReason_SteamOwnerLeftGuestUser: "EDenyReason_SteamOwnerLeftGuestUser",
}

func (e EDenyReason) String() string {
	if s, ok := EDenyReason_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EDenyReason_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EClanRank int32

const (
	EClanRank_None      EClanRank = 0
	EClanRank_Owner     EClanRank = 1
	EClanRank_Officer   EClanRank = 2
	EClanRank_Member    EClanRank = 3
	EClanRank_Moderator EClanRank = 4
)

var EClanRank_name = map[EClanRank]string{
	EClanRank_None:      "EClanRank_None",
	EClanRank_Owner:     "EClanRank_Owner",
	EClanRank_Officer:   "EClanRank_Officer",
	EClanRank_Member:    "EClanRank_Member",
	EClanRank_Moderator: "EClanRank_Moderator",
}

func (e EClanRank) String() string {
	if s, ok := EClanRank_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EClanRank_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EClanRelationship int32

const (
	EClanRelationship_None             EClanRelationship = 0
	EClanRelationship_Blocked          EClanRelationship = 1
	EClanRelationship_Invited          EClanRelationship = 2
	EClanRelationship_Member           EClanRelationship = 3
	EClanRelationship_Kicked           EClanRelationship = 4
	EClanRelationship_KickAcknowledged EClanRelationship = 5
)

var EClanRelationship_name = map[EClanRelationship]string{
	EClanRelationship_None:             "EClanRelationship_None",
	EClanRelationship_Blocked:          "EClanRelationship_Blocked",
	EClanRelationship_Invited:          "EClanRelationship_Invited",
	EClanRelationship_Member:           "EClanRelationship_Member",
	EClanRelationship_Kicked:           "EClanRelationship_Kicked",
	EClanRelationship_KickAcknowledged: "EClanRelationship_KickAcknowledged",
}

func (e EClanRelationship) String() string {
	if s, ok := EClanRelationship_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EClanRelationship_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EAuthSessionResponse int32

const (
	EAuthSessionResponse_OK                           EAuthSessionResponse = 0
	EAuthSessionResponse_UserNotConnectedToSteam      EAuthSessionResponse = 1
	EAuthSessionResponse_NoLicenseOrExpired           EAuthSessionResponse = 2
	EAuthSessionResponse_VACBanned                    EAuthSessionResponse = 3
	EAuthSessionResponse_LoggedInElseWhere            EAuthSessionResponse = 4
	EAuthSessionResponse_VACCheckTimedOut             EAuthSessionResponse = 5
	EAuthSessionResponse_AuthTicketCanceled           EAuthSessionResponse = 6
	EAuthSessionResponse_AuthTicketInvalidAlreadyUsed EAuthSessionResponse = 7
	EAuthSessionResponse_AuthTicketInvalid            EAuthSessionResponse = 8
	EAuthSessionResponse_PublisherIssuedBan           EAuthSessionResponse = 9
)

var EAuthSessionResponse_name = map[EAuthSessionResponse]string{
	EAuthSessionResponse_OK:                           "EAuthSessionResponse_OK",
	EAuthSessionResponse_UserNotConnectedToSteam:      "EAuthSessionResponse_UserNotConnectedToSteam",
	EAuthSessionResponse_NoLicenseOrExpired:           "EAuthSessionResponse_NoLicenseOrExpired",
	EAuthSessionResponse_VACBanned:                    "EAuthSessionResponse_VACBanned",
	EAuthSessionResponse_LoggedInElseWhere:            "EAuthSessionResponse_LoggedInElseWhere",
	EAuthSessionResponse_VACCheckTimedOut:             "EAuthSessionResponse_VACCheckTimedOut",
	EAuthSessionResponse_AuthTicketCanceled:           "EAuthSessionResponse_AuthTicketCanceled",
	EAuthSessionResponse_AuthTicketInvalidAlreadyUsed: "EAuthSessionResponse_AuthTicketInvalidAlreadyUsed",
	EAuthSessionResponse_AuthTicketInvalid:            "EAuthSessionResponse_AuthTicketInvalid",
	EAuthSessionResponse_PublisherIssuedBan:           "EAuthSessionResponse_PublisherIssuedBan",
}

func (e EAuthSessionResponse) String() string {
	if s, ok := EAuthSessionResponse_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EAuthSessionResponse_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EChatRoomEnterResponse int32

const (
	EChatRoomEnterResponse_Success            EChatRoomEnterResponse = 1
	EChatRoomEnterResponse_DoesntExist        EChatRoomEnterResponse = 2
	EChatRoomEnterResponse_NotAllowed         EChatRoomEnterResponse = 3
	EChatRoomEnterResponse_Full               EChatRoomEnterResponse = 4
	EChatRoomEnterResponse_Error              EChatRoomEnterResponse = 5
	EChatRoomEnterResponse_Banned             EChatRoomEnterResponse = 6
	EChatRoomEnterResponse_Limited            EChatRoomEnterResponse = 7
	EChatRoomEnterResponse_ClanDisabled       EChatRoomEnterResponse = 8
	EChatRoomEnterResponse_CommunityBan       EChatRoomEnterResponse = 9
	EChatRoomEnterResponse_MemberBlockedYou   EChatRoomEnterResponse = 10
	EChatRoomEnterResponse_YouBlockedMember   EChatRoomEnterResponse = 11
	EChatRoomEnterResponse_NoRankingDataLobby EChatRoomEnterResponse = 12 // removed
	EChatRoomEnterResponse_NoRankingDataUser  EChatRoomEnterResponse = 13 // removed
	EChatRoomEnterResponse_RankOutOfRange     EChatRoomEnterResponse = 14 // removed
)

var EChatRoomEnterResponse_name = map[EChatRoomEnterResponse]string{
	EChatRoomEnterResponse_Success:            "EChatRoomEnterResponse_Success",
	EChatRoomEnterResponse_DoesntExist:        "EChatRoomEnterResponse_DoesntExist",
	EChatRoomEnterResponse_NotAllowed:         "EChatRoomEnterResponse_NotAllowed",
	EChatRoomEnterResponse_Full:               "EChatRoomEnterResponse_Full",
	EChatRoomEnterResponse_Error:              "EChatRoomEnterResponse_Error",
	EChatRoomEnterResponse_Banned:             "EChatRoomEnterResponse_Banned",
	EChatRoomEnterResponse_Limited:            "EChatRoomEnterResponse_Limited",
	EChatRoomEnterResponse_ClanDisabled:       "EChatRoomEnterResponse_ClanDisabled",
	EChatRoomEnterResponse_CommunityBan:       "EChatRoomEnterResponse_CommunityBan",
	EChatRoomEnterResponse_MemberBlockedYou:   "EChatRoomEnterResponse_MemberBlockedYou",
	EChatRoomEnterResponse_YouBlockedMember:   "EChatRoomEnterResponse_YouBlockedMember",
	EChatRoomEnterResponse_NoRankingDataLobby: "EChatRoomEnterResponse_NoRankingDataLobby",
	EChatRoomEnterResponse_NoRankingDataUser:  "EChatRoomEnterResponse_NoRankingDataUser",
	EChatRoomEnterResponse_RankOutOfRange:     "EChatRoomEnterResponse_RankOutOfRange",
}

func (e EChatRoomEnterResponse) String() string {
	if s, ok := EChatRoomEnterResponse_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EChatRoomEnterResponse_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EChatRoomType int32

const (
	EChatRoomType_Friend EChatRoomType = 1
	EChatRoomType_MUC    EChatRoomType = 2
	EChatRoomType_Lobby  EChatRoomType = 3
)

var EChatRoomType_name = map[EChatRoomType]string{
	EChatRoomType_Friend: "EChatRoomType_Friend",
	EChatRoomType_MUC:    "EChatRoomType_MUC",
	EChatRoomType_Lobby:  "EChatRoomType_Lobby",
}

func (e EChatRoomType) String() string {
	if s, ok := EChatRoomType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EChatRoomType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EChatInfoType int32

const (
	EChatInfoType_StateChange       EChatInfoType = 1
	EChatInfoType_InfoUpdate        EChatInfoType = 2
	EChatInfoType_MemberLimitChange EChatInfoType = 3
)

var EChatInfoType_name = map[EChatInfoType]string{
	EChatInfoType_StateChange:       "EChatInfoType_StateChange",
	EChatInfoType_InfoUpdate:        "EChatInfoType_InfoUpdate",
	EChatInfoType_MemberLimitChange: "EChatInfoType_MemberLimitChange",
}

func (e EChatInfoType) String() string {
	if s, ok := EChatInfoType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EChatInfoType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EChatAction int32

const (
	EChatAction_InviteChat            EChatAction = 1
	EChatAction_Kick                  EChatAction = 2
	EChatAction_Ban                   EChatAction = 3
	EChatAction_UnBan                 EChatAction = 4
	EChatAction_StartVoiceSpeak       EChatAction = 5
	EChatAction_EndVoiceSpeak         EChatAction = 6
	EChatAction_LockChat              EChatAction = 7
	EChatAction_UnlockChat            EChatAction = 8
	EChatAction_CloseChat             EChatAction = 9
	EChatAction_SetJoinable           EChatAction = 10
	EChatAction_SetUnjoinable         EChatAction = 11
	EChatAction_SetOwner              EChatAction = 12
	EChatAction_SetInvisibleToFriends EChatAction = 13
	EChatAction_SetVisibleToFriends   EChatAction = 14
	EChatAction_SetModerated          EChatAction = 15
	EChatAction_SetUnmoderated        EChatAction = 16
)

var EChatAction_name = map[EChatAction]string{
	EChatAction_InviteChat:            "EChatAction_InviteChat",
	EChatAction_Kick:                  "EChatAction_Kick",
	EChatAction_Ban:                   "EChatAction_Ban",
	EChatAction_UnBan:                 "EChatAction_UnBan",
	EChatAction_StartVoiceSpeak:       "EChatAction_StartVoiceSpeak",
	EChatAction_EndVoiceSpeak:         "EChatAction_EndVoiceSpeak",
	EChatAction_LockChat:              "EChatAction_LockChat",
	EChatAction_UnlockChat:            "EChatAction_UnlockChat",
	EChatAction_CloseChat:             "EChatAction_CloseChat",
	EChatAction_SetJoinable:           "EChatAction_SetJoinable",
	EChatAction_SetUnjoinable:         "EChatAction_SetUnjoinable",
	EChatAction_SetOwner:              "EChatAction_SetOwner",
	EChatAction_SetInvisibleToFriends: "EChatAction_SetInvisibleToFriends",
	EChatAction_SetVisibleToFriends:   "EChatAction_SetVisibleToFriends",
	EChatAction_SetModerated:          "EChatAction_SetModerated",
	EChatAction_SetUnmoderated:        "EChatAction_SetUnmoderated",
}

func (e EChatAction) String() string {
	if s, ok := EChatAction_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EChatAction_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EChatActionResult int32

const (
	EChatActionResult_Success                EChatActionResult = 1
	EChatActionResult_Error                  EChatActionResult = 2
	EChatActionResult_NotPermitted           EChatActionResult = 3
	EChatActionResult_NotAllowedOnClanMember EChatActionResult = 4
	EChatActionResult_NotAllowedOnBannedUser EChatActionResult = 5
	EChatActionResult_NotAllowedOnChatOwner  EChatActionResult = 6
	EChatActionResult_NotAllowedOnSelf       EChatActionResult = 7
	EChatActionResult_ChatDoesntExist        EChatActionResult = 8
	EChatActionResult_ChatFull               EChatActionResult = 9
	EChatActionResult_VoiceSlotsFull         EChatActionResult = 10
)

var EChatActionResult_name = map[EChatActionResult]string{
	EChatActionResult_Success:                "EChatActionResult_Success",
	EChatActionResult_Error:                  "EChatActionResult_Error",
	EChatActionResult_NotPermitted:           "EChatActionResult_NotPermitted",
	EChatActionResult_NotAllowedOnClanMember: "EChatActionResult_NotAllowedOnClanMember",
	EChatActionResult_NotAllowedOnBannedUser: "EChatActionResult_NotAllowedOnBannedUser",
	EChatActionResult_NotAllowedOnChatOwner:  "EChatActionResult_NotAllowedOnChatOwner",
	EChatActionResult_NotAllowedOnSelf:       "EChatActionResult_NotAllowedOnSelf",
	EChatActionResult_ChatDoesntExist:        "EChatActionResult_ChatDoesntExist",
	EChatActionResult_ChatFull:               "EChatActionResult_ChatFull",
	EChatActionResult_VoiceSlotsFull:         "EChatActionResult_VoiceSlotsFull",
}

func (e EChatActionResult) String() string {
	if s, ok := EChatActionResult_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EChatActionResult_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EAppInfoSection int32

const (
	EAppInfoSection_Unknown      EAppInfoSection = 0
	EAppInfoSection_All          EAppInfoSection = 1
	EAppInfoSection_First        EAppInfoSection = 2
	EAppInfoSection_Common       EAppInfoSection = 2
	EAppInfoSection_Extended     EAppInfoSection = 3
	EAppInfoSection_Config       EAppInfoSection = 4
	EAppInfoSection_Stats        EAppInfoSection = 5
	EAppInfoSection_Install      EAppInfoSection = 6
	EAppInfoSection_Depots       EAppInfoSection = 7
	EAppInfoSection_VAC          EAppInfoSection = 8 // removed
	EAppInfoSection_VAC_UNUSED   EAppInfoSection = 8 // removed
	EAppInfoSection_DRM          EAppInfoSection = 9 // removed
	EAppInfoSection_DRM_UNUSED   EAppInfoSection = 9 // removed
	EAppInfoSection_UFS          EAppInfoSection = 10
	EAppInfoSection_OGG          EAppInfoSection = 11
	EAppInfoSection_Items        EAppInfoSection = 12 // removed
	EAppInfoSection_ItemsUNUSED  EAppInfoSection = 12 // removed
	EAppInfoSection_Policies     EAppInfoSection = 13
	EAppInfoSection_SysReqs      EAppInfoSection = 14
	EAppInfoSection_Community    EAppInfoSection = 15
	EAppInfoSection_Store        EAppInfoSection = 16
	EAppInfoSection_Localization EAppInfoSection = 17
	EAppInfoSection_Max          EAppInfoSection = 18
)

var EAppInfoSection_name = map[EAppInfoSection]string{
	EAppInfoSection_Unknown:      "EAppInfoSection_Unknown",
	EAppInfoSection_All:          "EAppInfoSection_All",
	EAppInfoSection_First:        "EAppInfoSection_First",
	EAppInfoSection_Extended:     "EAppInfoSection_Extended",
	EAppInfoSection_Config:       "EAppInfoSection_Config",
	EAppInfoSection_Stats:        "EAppInfoSection_Stats",
	EAppInfoSection_Install:      "EAppInfoSection_Install",
	EAppInfoSection_Depots:       "EAppInfoSection_Depots",
	EAppInfoSection_VAC:          "EAppInfoSection_VAC",
	EAppInfoSection_DRM:          "EAppInfoSection_DRM",
	EAppInfoSection_UFS:          "EAppInfoSection_UFS",
	EAppInfoSection_OGG:          "EAppInfoSection_OGG",
	EAppInfoSection_Items:        "EAppInfoSection_Items",
	EAppInfoSection_Policies:     "EAppInfoSection_Policies",
	EAppInfoSection_SysReqs:      "EAppInfoSection_SysReqs",
	EAppInfoSection_Community:    "EAppInfoSection_Community",
	EAppInfoSection_Store:        "EAppInfoSection_Store",
	EAppInfoSection_Localization: "EAppInfoSection_Localization",
	EAppInfoSection_Max:          "EAppInfoSection_Max",
}

func (e EAppInfoSection) String() string {
	if s, ok := EAppInfoSection_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EAppInfoSection_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EContentDownloadSourceType int32

const (
	EContentDownloadSourceType_Invalid    EContentDownloadSourceType = 0
	EContentDownloadSourceType_CS         EContentDownloadSourceType = 1
	EContentDownloadSourceType_CDN        EContentDownloadSourceType = 2
	EContentDownloadSourceType_LCS        EContentDownloadSourceType = 3
	EContentDownloadSourceType_ProxyCache EContentDownloadSourceType = 4
	EContentDownloadSourceType_LANPeer    EContentDownloadSourceType = 5
	EContentDownloadSourceType_SLS        EContentDownloadSourceType = 6
	EContentDownloadSourceType_SteamCache EContentDownloadSourceType = 7
	EContentDownloadSourceType_OpenCache  EContentDownloadSourceType = 8
	EContentDownloadSourceType_Max        EContentDownloadSourceType = 9
)

var EContentDownloadSourceType_name = map[EContentDownloadSourceType]string{
	EContentDownloadSourceType_Invalid:    "EContentDownloadSourceType_Invalid",
	EContentDownloadSourceType_CS:         "EContentDownloadSourceType_CS",
	EContentDownloadSourceType_CDN:        "EContentDownloadSourceType_CDN",
	EContentDownloadSourceType_LCS:        "EContentDownloadSourceType_LCS",
	EContentDownloadSourceType_ProxyCache: "EContentDownloadSourceType_ProxyCache",
	EContentDownloadSourceType_LANPeer:    "EContentDownloadSourceType_LANPeer",
	EContentDownloadSourceType_SLS:        "EContentDownloadSourceType_SLS",
	EContentDownloadSourceType_SteamCache: "EContentDownloadSourceType_SteamCache",
	EContentDownloadSourceType_OpenCache:  "EContentDownloadSourceType_OpenCache",
	EContentDownloadSourceType_Max:        "EContentDownloadSourceType_Max",
}

func (e EContentDownloadSourceType) String() string {
	if s, ok := EContentDownloadSourceType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EContentDownloadSourceType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EPlatformType int32

const (
	EPlatformType_Unknown EPlatformType = 0
	EPlatformType_Win32   EPlatformType = 1
	EPlatformType_Win64   EPlatformType = 2
	EPlatformType_Linux   EPlatformType = 3 // removed: split to Linux64 and Linux32
	EPlatformType_Linux64 EPlatformType = 3
	EPlatformType_OSX     EPlatformType = 4
	EPlatformType_PS3     EPlatformType = 5
	EPlatformType_Linux32 EPlatformType = 6
	EPlatformType_Max     EPlatformType = 7
)

var EPlatformType_name = map[EPlatformType]string{
	EPlatformType_Unknown: "EPlatformType_Unknown",
	EPlatformType_Win32:   "EPlatformType_Win32",
	EPlatformType_Win64:   "EPlatformType_Win64",
	EPlatformType_Linux64: "EPlatformType_Linux64",
	EPlatformType_OSX:     "EPlatformType_OSX",
	EPlatformType_PS3:     "EPlatformType_PS3",
	EPlatformType_Linux32: "EPlatformType_Linux32",
	EPlatformType_Max:     "EPlatformType_Max",
}

func (e EPlatformType) String() string {
	if s, ok := EPlatformType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EPlatformType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EOSType int32

const (
	EOSType_Unknown        EOSType = -1
	EOSType_Web            EOSType = -700
	EOSType_IOSUnknown     EOSType = -600
	EOSType_AndroidUnknown EOSType = -500
	EOSType_UMQ            EOSType = -400
	EOSType_PS3            EOSType = -300
	EOSType_MacOSUnknown   EOSType = -102
	EOSType_MacOS104       EOSType = -101
	EOSType_MacOS105       EOSType = -100
	EOSType_MacOS1058      EOSType = -99
	EOSType_MacOS106       EOSType = -95
	EOSType_MacOS1063      EOSType = -94
	EOSType_MacOS1064_slgu EOSType = -93
	EOSType_MacOS1067      EOSType = -92
	EOSType_MacOS107       EOSType = -90
	EOSType_MacOS108       EOSType = -89
	EOSType_MacOS109       EOSType = -88
	EOSType_MacOS1010      EOSType = -87
	EOSType_MacOS1011      EOSType = -86
	EOSType_MacOS1012      EOSType = -85
	EOSType_Macos1013      EOSType = -84
	EOSType_Macos1014      EOSType = -83
	EOSType_MacOSMax       EOSType = -1
	EOSType_LinuxUnknown   EOSType = -203
	EOSType_Linux22        EOSType = -202
	EOSType_Linux24        EOSType = -201
	EOSType_Linux26        EOSType = -200
	EOSType_Linux32        EOSType = -199
	EOSType_Linux35        EOSType = -198
	EOSType_Linux36        EOSType = -197
	EOSType_Linux310       EOSType = -196
	EOSType_Linux316       EOSType = -195
	EOSType_Linux318       EOSType = -194
	EOSType_Linux3x        EOSType = -193
	EOSType_Linux4x        EOSType = -192
	EOSType_Linux41        EOSType = -191
	EOSType_Linux44        EOSType = -190
	EOSType_Linux49        EOSType = -189
	EOSType_LinuxMax       EOSType = -101
	EOSType_WinUnknown     EOSType = 0
	EOSType_Win311         EOSType = 1
	EOSType_Win95          EOSType = 2
	EOSType_Win98          EOSType = 3
	EOSType_WinME          EOSType = 4
	EOSType_WinNT          EOSType = 5
	EOSType_Win200         EOSType = 6 // removed: renamed to Win2000
	EOSType_Win2000        EOSType = 6
	EOSType_WinXP          EOSType = 7
	EOSType_Win2003        EOSType = 8
	EOSType_WinVista       EOSType = 9
	EOSType_Win7           EOSType = 10 // removed: renamed to Windows7
	EOSType_Windows7       EOSType = 10
	EOSType_Win2008        EOSType = 11
	EOSType_Win2012        EOSType = 12
	EOSType_Win8           EOSType = 13 // removed: renamed to Windows8
	EOSType_Windows8       EOSType = 13
	EOSType_Win81          EOSType = 14 // removed: renamed to Windows81
	EOSType_Windows81      EOSType = 14
	EOSType_Win2012R2      EOSType = 15
	EOSType_Win10          EOSType = 16 // removed: renamed to Windows10
	EOSType_Windows10      EOSType = 16
	EOSType_Win2016        EOSType = 17
	EOSType_WinMAX         EOSType = 18
	EOSType_Max            EOSType = 26
)

var EOSType_name = map[EOSType]string{
	EOSType_Unknown:        "EOSType_Unknown",
	EOSType_Web:            "EOSType_Web",
	EOSType_IOSUnknown:     "EOSType_IOSUnknown",
	EOSType_AndroidUnknown: "EOSType_AndroidUnknown",
	EOSType_UMQ:            "EOSType_UMQ",
	EOSType_PS3:            "EOSType_PS3",
	EOSType_MacOSUnknown:   "EOSType_MacOSUnknown",
	EOSType_MacOS104:       "EOSType_MacOS104",
	EOSType_MacOS105:       "EOSType_MacOS105",
	EOSType_MacOS1058:      "EOSType_MacOS1058",
	EOSType_MacOS106:       "EOSType_MacOS106",
	EOSType_MacOS1063:      "EOSType_MacOS1063",
	EOSType_MacOS1064_slgu: "EOSType_MacOS1064_slgu",
	EOSType_MacOS1067:      "EOSType_MacOS1067",
	EOSType_MacOS107:       "EOSType_MacOS107",
	EOSType_MacOS108:       "EOSType_MacOS108",
	EOSType_MacOS109:       "EOSType_MacOS109",
	EOSType_MacOS1010:      "EOSType_MacOS1010",
	EOSType_MacOS1011:      "EOSType_MacOS1011",
	EOSType_MacOS1012:      "EOSType_MacOS1012",
	EOSType_Macos1013:      "EOSType_Macos1013",
	EOSType_Macos1014:      "EOSType_Macos1014",
	EOSType_LinuxUnknown:   "EOSType_LinuxUnknown",
	EOSType_Linux22:        "EOSType_Linux22",
	EOSType_Linux24:        "EOSType_Linux24",
	EOSType_Linux26:        "EOSType_Linux26",
	EOSType_Linux32:        "EOSType_Linux32",
	EOSType_Linux35:        "EOSType_Linux35",
	EOSType_Linux36:        "EOSType_Linux36",
	EOSType_Linux310:       "EOSType_Linux310",
	EOSType_Linux316:       "EOSType_Linux316",
	EOSType_Linux318:       "EOSType_Linux318",
	EOSType_Linux3x:        "EOSType_Linux3x",
	EOSType_Linux4x:        "EOSType_Linux4x",
	EOSType_Linux41:        "EOSType_Linux41",
	EOSType_Linux44:        "EOSType_Linux44",
	EOSType_Linux49:        "EOSType_Linux49",
	EOSType_WinUnknown:     "EOSType_WinUnknown",
	EOSType_Win311:         "EOSType_Win311",
	EOSType_Win95:          "EOSType_Win95",
	EOSType_Win98:          "EOSType_Win98",
	EOSType_WinME:          "EOSType_WinME",
	EOSType_WinNT:          "EOSType_WinNT",
	EOSType_Win2000:        "EOSType_Win2000",
	EOSType_WinXP:          "EOSType_WinXP",
	EOSType_Win2003:        "EOSType_Win2003",
	EOSType_WinVista:       "EOSType_WinVista",
	EOSType_Windows7:       "EOSType_Windows7",
	EOSType_Win2008:        "EOSType_Win2008",
	EOSType_Win2012:        "EOSType_Win2012",
	EOSType_Windows8:       "EOSType_Windows8",
	EOSType_Windows81:      "EOSType_Windows81",
	EOSType_Win2012R2:      "EOSType_Win2012R2",
	EOSType_Windows10:      "EOSType_Windows10",
	EOSType_Win2016:        "EOSType_Win2016",
	EOSType_WinMAX:         "EOSType_WinMAX",
	EOSType_Max:            "EOSType_Max",
}

func (e EOSType) String() string {
	if s, ok := EOSType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EOSType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EServerType int32

const (
	EServerType_Invalid             EServerType = -1
	EServerType_First               EServerType = 0
	EServerType_Shell               EServerType = 0
	EServerType_GM                  EServerType = 1
	EServerType_BUM                 EServerType = 2 // removed
	EServerType_BUMOBOSOLETE        EServerType = 2 // removed
	EServerType_AM                  EServerType = 3
	EServerType_BS                  EServerType = 4
	EServerType_VS                  EServerType = 5
	EServerType_ATS                 EServerType = 6
	EServerType_CM                  EServerType = 7
	EServerType_FBS                 EServerType = 8
	EServerType_FG                  EServerType = 9 // removed: renamed to BoxMonitor
	EServerType_BoxMonitor          EServerType = 9
	EServerType_SS                  EServerType = 10
	EServerType_DRMS                EServerType = 11
	EServerType_HubOBSOLETE         EServerType = 12 // removed
	EServerType_Console             EServerType = 13
	EServerType_ASBOBSOLETE         EServerType = 14 // removed
	EServerType_PICS                EServerType = 14
	EServerType_Client              EServerType = 15
	EServerType_BootstrapOBSOLETE   EServerType = 16 // removed
	EServerType_DP                  EServerType = 17
	EServerType_WG                  EServerType = 18
	EServerType_SM                  EServerType = 19
	EServerType_SLC                 EServerType = 20
	EServerType_UFS                 EServerType = 21
	EServerType_Util                EServerType = 23
	EServerType_DSS                 EServerType = 24 // removed: renamed to Community
	EServerType_Community           EServerType = 24
	EServerType_P2PRelayOBSOLETE    EServerType = 25 // removed
	EServerType_AppInformation      EServerType = 26
	EServerType_Spare               EServerType = 27
	EServerType_FTS                 EServerType = 28
	EServerType_EPM                 EServerType = 29 // removed
	EServerType_EPMOBSOLETE         EServerType = 29 // removed
	EServerType_PS                  EServerType = 30
	EServerType_IS                  EServerType = 31
	EServerType_CCS                 EServerType = 32
	EServerType_DFS                 EServerType = 33
	EServerType_LBS                 EServerType = 34
	EServerType_MDS                 EServerType = 35
	EServerType_CS                  EServerType = 36
	EServerType_GC                  EServerType = 37
	EServerType_NS                  EServerType = 38
	EServerType_OGS                 EServerType = 39
	EServerType_WebAPI              EServerType = 40
	EServerType_UDS                 EServerType = 41
	EServerType_MMS                 EServerType = 42
	EServerType_GMS                 EServerType = 43
	EServerType_KGS                 EServerType = 44
	EServerType_UCM                 EServerType = 45
	EServerType_RM                  EServerType = 46
	EServerType_FS                  EServerType = 47
	EServerType_Econ                EServerType = 48
	EServerType_Backpack            EServerType = 49
	EServerType_UGS                 EServerType = 50
	EServerType_Store               EServerType = 51 // removed: renamed to StoreFeature
	EServerType_StoreFeature        EServerType = 51
	EServerType_MoneyStats          EServerType = 52
	EServerType_CRE                 EServerType = 53
	EServerType_UMQ                 EServerType = 54
	EServerType_Workshop            EServerType = 55
	EServerType_BRP                 EServerType = 56
	EServerType_GCH                 EServerType = 57
	EServerType_MPAS                EServerType = 58
	EServerType_Trade               EServerType = 59
	EServerType_Secrets             EServerType = 60
	EServerType_Logsink             EServerType = 61
	EServerType_Market              EServerType = 62
	EServerType_Quest               EServerType = 63
	EServerType_WDS                 EServerType = 64
	EServerType_ACS                 EServerType = 65
	EServerType_PNP                 EServerType = 66
	EServerType_TaxForm             EServerType = 67
	EServerType_ExternalMonitor     EServerType = 68
	EServerType_Parental            EServerType = 69
	EServerType_PartnerUpload       EServerType = 70
	EServerType_Partner             EServerType = 71
	EServerType_ES                  EServerType = 72
	EServerType_DepotWebContent     EServerType = 73
	EServerType_ExternalConfig      EServerType = 74
	EServerType_GameNotifications   EServerType = 75
	EServerType_MarketRepl          EServerType = 76
	EServerType_MarketSearch        EServerType = 77
	EServerType_Localization        EServerType = 78
	EServerType_Steam2Emulator      EServerType = 79
	EServerType_PublicTest          EServerType = 80
	EServerType_SolrMgr             EServerType = 81
	EServerType_BroadcastRelay      EServerType = 82
	EServerType_BroadcastDirectory  EServerType = 83
	EServerType_VideoManager        EServerType = 84
	EServerType_TradeOffer          EServerType = 85
	EServerType_BroadcastChat       EServerType = 86
	EServerType_Phone               EServerType = 87
	EServerType_AccountScore        EServerType = 88
	EServerType_Support             EServerType = 89
	EServerType_LogRequest          EServerType = 90
	EServerType_LogWorker           EServerType = 91
	EServerType_EmailDelivery       EServerType = 92
	EServerType_InventoryManagement EServerType = 93
	EServerType_Auth                EServerType = 94
	EServerType_StoreCatalog        EServerType = 95
	EServerType_HLTVRelay           EServerType = 96
	EServerType_Max                 EServerType = 97
)

var EServerType_name = map[EServerType]string{
	EServerType_Invalid:             "EServerType_Invalid",
	EServerType_First:               "EServerType_First",
	EServerType_GM:                  "EServerType_GM",
	EServerType_BUM:                 "EServerType_BUM",
	EServerType_AM:                  "EServerType_AM",
	EServerType_BS:                  "EServerType_BS",
	EServerType_VS:                  "EServerType_VS",
	EServerType_ATS:                 "EServerType_ATS",
	EServerType_CM:                  "EServerType_CM",
	EServerType_FBS:                 "EServerType_FBS",
	EServerType_BoxMonitor:          "EServerType_BoxMonitor",
	EServerType_SS:                  "EServerType_SS",
	EServerType_DRMS:                "EServerType_DRMS",
	EServerType_HubOBSOLETE:         "EServerType_HubOBSOLETE",
	EServerType_Console:             "EServerType_Console",
	EServerType_PICS:                "EServerType_PICS",
	EServerType_Client:              "EServerType_Client",
	EServerType_BootstrapOBSOLETE:   "EServerType_BootstrapOBSOLETE",
	EServerType_DP:                  "EServerType_DP",
	EServerType_WG:                  "EServerType_WG",
	EServerType_SM:                  "EServerType_SM",
	EServerType_SLC:                 "EServerType_SLC",
	EServerType_UFS:                 "EServerType_UFS",
	EServerType_Util:                "EServerType_Util",
	EServerType_Community:           "EServerType_Community",
	EServerType_P2PRelayOBSOLETE:    "EServerType_P2PRelayOBSOLETE",
	EServerType_AppInformation:      "EServerType_AppInformation",
	EServerType_Spare:               "EServerType_Spare",
	EServerType_FTS:                 "EServerType_FTS",
	EServerType_EPM:                 "EServerType_EPM",
	EServerType_PS:                  "EServerType_PS",
	EServerType_IS:                  "EServerType_IS",
	EServerType_CCS:                 "EServerType_CCS",
	EServerType_DFS:                 "EServerType_DFS",
	EServerType_LBS:                 "EServerType_LBS",
	EServerType_MDS:                 "EServerType_MDS",
	EServerType_CS:                  "EServerType_CS",
	EServerType_GC:                  "EServerType_GC",
	EServerType_NS:                  "EServerType_NS",
	EServerType_OGS:                 "EServerType_OGS",
	EServerType_WebAPI:              "EServerType_WebAPI",
	EServerType_UDS:                 "EServerType_UDS",
	EServerType_MMS:                 "EServerType_MMS",
	EServerType_GMS:                 "EServerType_GMS",
	EServerType_KGS:                 "EServerType_KGS",
	EServerType_UCM:                 "EServerType_UCM",
	EServerType_RM:                  "EServerType_RM",
	EServerType_FS:                  "EServerType_FS",
	EServerType_Econ:                "EServerType_Econ",
	EServerType_Backpack:            "EServerType_Backpack",
	EServerType_UGS:                 "EServerType_UGS",
	EServerType_StoreFeature:        "EServerType_StoreFeature",
	EServerType_MoneyStats:          "EServerType_MoneyStats",
	EServerType_CRE:                 "EServerType_CRE",
	EServerType_UMQ:                 "EServerType_UMQ",
	EServerType_Workshop:            "EServerType_Workshop",
	EServerType_BRP:                 "EServerType_BRP",
	EServerType_GCH:                 "EServerType_GCH",
	EServerType_MPAS:                "EServerType_MPAS",
	EServerType_Trade:               "EServerType_Trade",
	EServerType_Secrets:             "EServerType_Secrets",
	EServerType_Logsink:             "EServerType_Logsink",
	EServerType_Market:              "EServerType_Market",
	EServerType_Quest:               "EServerType_Quest",
	EServerType_WDS:                 "EServerType_WDS",
	EServerType_ACS:                 "EServerType_ACS",
	EServerType_PNP:                 "EServerType_PNP",
	EServerType_TaxForm:             "EServerType_TaxForm",
	EServerType_ExternalMonitor:     "EServerType_ExternalMonitor",
	EServerType_Parental:            "EServerType_Parental",
	EServerType_PartnerUpload:       "EServerType_PartnerUpload",
	EServerType_Partner:             "EServerType_Partner",
	EServerType_ES:                  "EServerType_ES",
	EServerType_DepotWebContent:     "EServerType_DepotWebContent",
	EServerType_ExternalConfig:      "EServerType_ExternalConfig",
	EServerType_GameNotifications:   "EServerType_GameNotifications",
	EServerType_MarketRepl:          "EServerType_MarketRepl",
	EServerType_MarketSearch:        "EServerType_MarketSearch",
	EServerType_Localization:        "EServerType_Localization",
	EServerType_Steam2Emulator:      "EServerType_Steam2Emulator",
	EServerType_PublicTest:          "EServerType_PublicTest",
	EServerType_SolrMgr:             "EServerType_SolrMgr",
	EServerType_BroadcastRelay:      "EServerType_BroadcastRelay",
	EServerType_BroadcastDirectory:  "EServerType_BroadcastDirectory",
	EServerType_VideoManager:        "EServerType_VideoManager",
	EServerType_TradeOffer:          "EServerType_TradeOffer",
	EServerType_BroadcastChat:       "EServerType_BroadcastChat",
	EServerType_Phone:               "EServerType_Phone",
	EServerType_AccountScore:        "EServerType_AccountScore",
	EServerType_Support:             "EServerType_Support",
	EServerType_LogRequest:          "EServerType_LogRequest",
	EServerType_LogWorker:           "EServerType_LogWorker",
	EServerType_EmailDelivery:       "EServerType_EmailDelivery",
	EServerType_InventoryManagement: "EServerType_InventoryManagement",
	EServerType_Auth:                "EServerType_Auth",
	EServerType_StoreCatalog:        "EServerType_StoreCatalog",
	EServerType_HLTVRelay:           "EServerType_HLTVRelay",
	EServerType_Max:                 "EServerType_Max",
}

func (e EServerType) String() string {
	if s, ok := EServerType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EServerType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EBillingType int32

const (
	EBillingType_NoCost                 EBillingType = 0
	EBillingType_BillOnceOnly           EBillingType = 1
	EBillingType_BillMonthly            EBillingType = 2
	EBillingType_ProofOfPrepurchaseOnly EBillingType = 3
	EBillingType_GuestPass              EBillingType = 4
	EBillingType_HardwarePromo          EBillingType = 5
	EBillingType_Gift                   EBillingType = 6
	EBillingType_AutoGrant              EBillingType = 7
	EBillingType_OEMTicket              EBillingType = 8
	EBillingType_RecurringOption        EBillingType = 9
	EBillingType_BillOnceOrCDKey        EBillingType = 10
	EBillingType_Repurchaseable         EBillingType = 11
	EBillingType_FreeOnDemand           EBillingType = 12
	EBillingType_Rental                 EBillingType = 13
	EBillingType_CommercialLicense      EBillingType = 14
	EBillingType_FreeCommercialLicense  EBillingType = 15
	EBillingType_NumBillingTypes        EBillingType = 16
)

var EBillingType_name = map[EBillingType]string{
	EBillingType_NoCost:                 "EBillingType_NoCost",
	EBillingType_BillOnceOnly:           "EBillingType_BillOnceOnly",
	EBillingType_BillMonthly:            "EBillingType_BillMonthly",
	EBillingType_ProofOfPrepurchaseOnly: "EBillingType_ProofOfPrepurchaseOnly",
	EBillingType_GuestPass:              "EBillingType_GuestPass",
	EBillingType_HardwarePromo:          "EBillingType_HardwarePromo",
	EBillingType_Gift:                   "EBillingType_Gift",
	EBillingType_AutoGrant:              "EBillingType_AutoGrant",
	EBillingType_OEMTicket:              "EBillingType_OEMTicket",
	EBillingType_RecurringOption:        "EBillingType_RecurringOption",
	EBillingType_BillOnceOrCDKey:        "EBillingType_BillOnceOrCDKey",
	EBillingType_Repurchaseable:         "EBillingType_Repurchaseable",
	EBillingType_FreeOnDemand:           "EBillingType_FreeOnDemand",
	EBillingType_Rental:                 "EBillingType_Rental",
	EBillingType_CommercialLicense:      "EBillingType_CommercialLicense",
	EBillingType_FreeCommercialLicense:  "EBillingType_FreeCommercialLicense",
	EBillingType_NumBillingTypes:        "EBillingType_NumBillingTypes",
}

func (e EBillingType) String() string {
	if s, ok := EBillingType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EBillingType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EActivationCodeClass uint32

const (
	EActivationCodeClass_WonCDKey     EActivationCodeClass = 0
	EActivationCodeClass_ValveCDKey   EActivationCodeClass = 1
	EActivationCodeClass_Doom3CDKey   EActivationCodeClass = 2
	EActivationCodeClass_DBLookup     EActivationCodeClass = 3
	EActivationCodeClass_Steam2010Key EActivationCodeClass = 4
	EActivationCodeClass_Max          EActivationCodeClass = 5
	EActivationCodeClass_Test         EActivationCodeClass = 2147483647
	EActivationCodeClass_Invalid      EActivationCodeClass = 4294967295
)

var EActivationCodeClass_name = map[EActivationCodeClass]string{
	EActivationCodeClass_WonCDKey:     "EActivationCodeClass_WonCDKey",
	EActivationCodeClass_ValveCDKey:   "EActivationCodeClass_ValveCDKey",
	EActivationCodeClass_Doom3CDKey:   "EActivationCodeClass_Doom3CDKey",
	EActivationCodeClass_DBLookup:     "EActivationCodeClass_DBLookup",
	EActivationCodeClass_Steam2010Key: "EActivationCodeClass_Steam2010Key",
	EActivationCodeClass_Max:          "EActivationCodeClass_Max",
	EActivationCodeClass_Test:         "EActivationCodeClass_Test",
	EActivationCodeClass_Invalid:      "EActivationCodeClass_Invalid",
}

func (e EActivationCodeClass) String() string {
	if s, ok := EActivationCodeClass_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EActivationCodeClass_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EChatMemberStateChange int32

const (
	EChatMemberStateChange_Entered           EChatMemberStateChange = 0x01
	EChatMemberStateChange_Left              EChatMemberStateChange = 0x02
	EChatMemberStateChange_Disconnected      EChatMemberStateChange = 0x04
	EChatMemberStateChange_Kicked            EChatMemberStateChange = 0x08
	EChatMemberStateChange_Banned            EChatMemberStateChange = 0x10
	EChatMemberStateChange_VoiceSpeaking     EChatMemberStateChange = 0x1000
	EChatMemberStateChange_VoiceDoneSpeaking EChatMemberStateChange = 0x2000
)

var EChatMemberStateChange_name = map[EChatMemberStateChange]string{
	EChatMemberStateChange_Entered:           "EChatMemberStateChange_Entered",
	EChatMemberStateChange_Left:              "EChatMemberStateChange_Left",
	EChatMemberStateChange_Disconnected:      "EChatMemberStateChange_Disconnected",
	EChatMemberStateChange_Kicked:            "EChatMemberStateChange_Kicked",
	EChatMemberStateChange_Banned:            "EChatMemberStateChange_Banned",
	EChatMemberStateChange_VoiceSpeaking:     "EChatMemberStateChange_VoiceSpeaking",
	EChatMemberStateChange_VoiceDoneSpeaking: "EChatMemberStateChange_VoiceDoneSpeaking",
}

func (e EChatMemberStateChange) String() string {
	if s, ok := EChatMemberStateChange_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EChatMemberStateChange_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type ERegionCode uint8

const (
	ERegionCode_USEast       ERegionCode = 0x00
	ERegionCode_USWest       ERegionCode = 0x01
	ERegionCode_SouthAmerica ERegionCode = 0x02
	ERegionCode_Europe       ERegionCode = 0x03
	ERegionCode_Asia         ERegionCode = 0x04
	ERegionCode_Australia    ERegionCode = 0x05
	ERegionCode_MiddleEast   ERegionCode = 0x06
	ERegionCode_Africa       ERegionCode = 0x07
	ERegionCode_World        ERegionCode = 0xFF
)

var ERegionCode_name = map[ERegionCode]string{
	ERegionCode_USEast:       "ERegionCode_USEast",
	ERegionCode_USWest:       "ERegionCode_USWest",
	ERegionCode_SouthAmerica: "ERegionCode_SouthAmerica",
	ERegionCode_Europe:       "ERegionCode_Europe",
	ERegionCode_Asia:         "ERegionCode_Asia",
	ERegionCode_Australia:    "ERegionCode_Australia",
	ERegionCode_MiddleEast:   "ERegionCode_MiddleEast",
	ERegionCode_Africa:       "ERegionCode_Africa",
	ERegionCode_World:        "ERegionCode_World",
}

func (e ERegionCode) String() string {
	if s, ok := ERegionCode_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range ERegionCode_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type ECurrencyCode int32

const (
	ECurrencyCode_Invalid ECurrencyCode = 0
	ECurrencyCode_USD     ECurrencyCode = 1
	ECurrencyCode_GBP     ECurrencyCode = 2
	ECurrencyCode_EUR     ECurrencyCode = 3
	ECurrencyCode_CHF     ECurrencyCode = 4
	ECurrencyCode_RUB     ECurrencyCode = 5
	ECurrencyCode_PLN     ECurrencyCode = 6
	ECurrencyCode_BRL     ECurrencyCode = 7
	ECurrencyCode_JPY     ECurrencyCode = 8
	ECurrencyCode_NOK     ECurrencyCode = 9
	ECurrencyCode_IDR     ECurrencyCode = 10
	ECurrencyCode_MYR     ECurrencyCode = 11
	ECurrencyCode_PHP     ECurrencyCode = 12
	ECurrencyCode_SGD     ECurrencyCode = 13
	ECurrencyCode_THB     ECurrencyCode = 14
	ECurrencyCode_VND     ECurrencyCode = 15
	ECurrencyCode_KRW     ECurrencyCode = 16
	ECurrencyCode_TRY     ECurrencyCode = 17
	ECurrencyCode_UAH     ECurrencyCode = 18
	ECurrencyCode_MXN     ECurrencyCode = 19
	ECurrencyCode_CAD     ECurrencyCode = 20
	ECurrencyCode_AUD     ECurrencyCode = 21
	ECurrencyCode_NZD     ECurrencyCode = 22
	ECurrencyCode_CNY     ECurrencyCode = 23
	ECurrencyCode_INR     ECurrencyCode = 24
	ECurrencyCode_CLP     ECurrencyCode = 25
	ECurrencyCode_PEN     ECurrencyCode = 26
	ECurrencyCode_COP     ECurrencyCode = 27
	ECurrencyCode_ZAR     ECurrencyCode = 28
	ECurrencyCode_HKD     ECurrencyCode = 29
	ECurrencyCode_TWD     ECurrencyCode = 30
	ECurrencyCode_SAR     ECurrencyCode = 31
	ECurrencyCode_AED     ECurrencyCode = 32
	ECurrencyCode_ARS     ECurrencyCode = 34
	ECurrencyCode_ILS     ECurrencyCode = 35
	ECurrencyCode_BYN     ECurrencyCode = 36
	ECurrencyCode_KZT     ECurrencyCode = 37
	ECurrencyCode_KWD     ECurrencyCode = 38
	ECurrencyCode_QAR     ECurrencyCode = 39
	ECurrencyCode_CRC     ECurrencyCode = 40
	ECurrencyCode_UYU     ECurrencyCode = 41
	ECurrencyCode_Max     ECurrencyCode = 42
)

var ECurrencyCode_name = map[ECurrencyCode]string{
	ECurrencyCode_Invalid: "ECurrencyCode_Invalid",
	ECurrencyCode_USD:     "ECurrencyCode_USD",
	ECurrencyCode_GBP:     "ECurrencyCode_GBP",
	ECurrencyCode_EUR:     "ECurrencyCode_EUR",
	ECurrencyCode_CHF:     "ECurrencyCode_CHF",
	ECurrencyCode_RUB:     "ECurrencyCode_RUB",
	ECurrencyCode_PLN:     "ECurrencyCode_PLN",
	ECurrencyCode_BRL:     "ECurrencyCode_BRL",
	ECurrencyCode_JPY:     "ECurrencyCode_JPY",
	ECurrencyCode_NOK:     "ECurrencyCode_NOK",
	ECurrencyCode_IDR:     "ECurrencyCode_IDR",
	ECurrencyCode_MYR:     "ECurrencyCode_MYR",
	ECurrencyCode_PHP:     "ECurrencyCode_PHP",
	ECurrencyCode_SGD:     "ECurrencyCode_SGD",
	ECurrencyCode_THB:     "ECurrencyCode_THB",
	ECurrencyCode_VND:     "ECurrencyCode_VND",
	ECurrencyCode_KRW:     "ECurrencyCode_KRW",
	ECurrencyCode_TRY:     "ECurrencyCode_TRY",
	ECurrencyCode_UAH:     "ECurrencyCode_UAH",
	ECurrencyCode_MXN:     "ECurrencyCode_MXN",
	ECurrencyCode_CAD:     "ECurrencyCode_CAD",
	ECurrencyCode_AUD:     "ECurrencyCode_AUD",
	ECurrencyCode_NZD:     "ECurrencyCode_NZD",
	ECurrencyCode_CNY:     "ECurrencyCode_CNY",
	ECurrencyCode_INR:     "ECurrencyCode_INR",
	ECurrencyCode_CLP:     "ECurrencyCode_CLP",
	ECurrencyCode_PEN:     "ECurrencyCode_PEN",
	ECurrencyCode_COP:     "ECurrencyCode_COP",
	ECurrencyCode_ZAR:     "ECurrencyCode_ZAR",
	ECurrencyCode_HKD:     "ECurrencyCode_HKD",
	ECurrencyCode_TWD:     "ECurrencyCode_TWD",
	ECurrencyCode_SAR:     "ECurrencyCode_SAR",
	ECurrencyCode_AED:     "ECurrencyCode_AED",
	ECurrencyCode_ARS:     "ECurrencyCode_ARS",
	ECurrencyCode_ILS:     "ECurrencyCode_ILS",
	ECurrencyCode_BYN:     "ECurrencyCode_BYN",
	ECurrencyCode_KZT:     "ECurrencyCode_KZT",
	ECurrencyCode_KWD:     "ECurrencyCode_KWD",
	ECurrencyCode_QAR:     "ECurrencyCode_QAR",
	ECurrencyCode_CRC:     "ECurrencyCode_CRC",
	ECurrencyCode_UYU:     "ECurrencyCode_UYU",
	ECurrencyCode_Max:     "ECurrencyCode_Max",
}

func (e ECurrencyCode) String() string {
	if s, ok := ECurrencyCode_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range ECurrencyCode_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EDepotFileFlag int32

const (
	EDepotFileFlag_UserConfig          EDepotFileFlag = 1
	EDepotFileFlag_VersionedUserConfig EDepotFileFlag = 2
	EDepotFileFlag_Encrypted           EDepotFileFlag = 4
	EDepotFileFlag_ReadOnly            EDepotFileFlag = 8
	EDepotFileFlag_Hidden              EDepotFileFlag = 16
	EDepotFileFlag_Executable          EDepotFileFlag = 32
	EDepotFileFlag_Directory           EDepotFileFlag = 64
	EDepotFileFlag_CustomExecutable    EDepotFileFlag = 128
	EDepotFileFlag_InstallScript       EDepotFileFlag = 256
	EDepotFileFlag_Symlink             EDepotFileFlag = 512
)

var EDepotFileFlag_name = map[EDepotFileFlag]string{
	EDepotFileFlag_UserConfig:          "EDepotFileFlag_UserConfig",
	EDepotFileFlag_VersionedUserConfig: "EDepotFileFlag_VersionedUserConfig",
	EDepotFileFlag_Encrypted:           "EDepotFileFlag_Encrypted",
	EDepotFileFlag_ReadOnly:            "EDepotFileFlag_ReadOnly",
	EDepotFileFlag_Hidden:              "EDepotFileFlag_Hidden",
	EDepotFileFlag_Executable:          "EDepotFileFlag_Executable",
	EDepotFileFlag_Directory:           "EDepotFileFlag_Directory",
	EDepotFileFlag_CustomExecutable:    "EDepotFileFlag_CustomExecutable",
	EDepotFileFlag_InstallScript:       "EDepotFileFlag_InstallScript",
	EDepotFileFlag_Symlink:             "EDepotFileFlag_Symlink",
}

func (e EDepotFileFlag) String() string {
	if s, ok := EDepotFileFlag_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EDepotFileFlag_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EWorkshopEnumerationType int32

const (
	EWorkshopEnumerationType_RankedByVote            EWorkshopEnumerationType = 0
	EWorkshopEnumerationType_Recent                  EWorkshopEnumerationType = 1
	EWorkshopEnumerationType_Trending                EWorkshopEnumerationType = 2
	EWorkshopEnumerationType_FavoriteOfFriends       EWorkshopEnumerationType = 3
	EWorkshopEnumerationType_VotedByFriends          EWorkshopEnumerationType = 4
	EWorkshopEnumerationType_ContentByFriends        EWorkshopEnumerationType = 5
	EWorkshopEnumerationType_RecentFromFollowedUsers EWorkshopEnumerationType = 6
)

var EWorkshopEnumerationType_name = map[EWorkshopEnumerationType]string{
	EWorkshopEnumerationType_RankedByVote:            "EWorkshopEnumerationType_RankedByVote",
	EWorkshopEnumerationType_Recent:                  "EWorkshopEnumerationType_Recent",
	EWorkshopEnumerationType_Trending:                "EWorkshopEnumerationType_Trending",
	EWorkshopEnumerationType_FavoriteOfFriends:       "EWorkshopEnumerationType_FavoriteOfFriends",
	EWorkshopEnumerationType_VotedByFriends:          "EWorkshopEnumerationType_VotedByFriends",
	EWorkshopEnumerationType_ContentByFriends:        "EWorkshopEnumerationType_ContentByFriends",
	EWorkshopEnumerationType_RecentFromFollowedUsers: "EWorkshopEnumerationType_RecentFromFollowedUsers",
}

func (e EWorkshopEnumerationType) String() string {
	if s, ok := EWorkshopEnumerationType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EWorkshopEnumerationType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EPublishedFileVisibility int32

const (
	EPublishedFileVisibility_Public      EPublishedFileVisibility = 0
	EPublishedFileVisibility_FriendsOnly EPublishedFileVisibility = 1
	EPublishedFileVisibility_Private     EPublishedFileVisibility = 2
)

var EPublishedFileVisibility_name = map[EPublishedFileVisibility]string{
	EPublishedFileVisibility_Public:      "EPublishedFileVisibility_Public",
	EPublishedFileVisibility_FriendsOnly: "EPublishedFileVisibility_FriendsOnly",
	EPublishedFileVisibility_Private:     "EPublishedFileVisibility_Private",
}

func (e EPublishedFileVisibility) String() string {
	if s, ok := EPublishedFileVisibility_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EPublishedFileVisibility_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EWorkshopFileType int32

const (
	EWorkshopFileType_First                  EWorkshopFileType = 0
	EWorkshopFileType_Community              EWorkshopFileType = 0
	EWorkshopFileType_Microtransaction       EWorkshopFileType = 1
	EWorkshopFileType_Collection             EWorkshopFileType = 2
	EWorkshopFileType_Art                    EWorkshopFileType = 3
	EWorkshopFileType_Video                  EWorkshopFileType = 4
	EWorkshopFileType_Screenshot             EWorkshopFileType = 5
	EWorkshopFileType_Game                   EWorkshopFileType = 6
	EWorkshopFileType_Software               EWorkshopFileType = 7
	EWorkshopFileType_Concept                EWorkshopFileType = 8
	EWorkshopFileType_WebGuide               EWorkshopFileType = 9
	EWorkshopFileType_IntegratedGuide        EWorkshopFileType = 10
	EWorkshopFileType_Merch                  EWorkshopFileType = 11
	EWorkshopFileType_ControllerBinding      EWorkshopFileType = 12
	EWorkshopFileType_SteamworksAccessInvite EWorkshopFileType = 13
	EWorkshopFileType_SteamVideo             EWorkshopFileType = 14
	EWorkshopFileType_GameManagedItem        EWorkshopFileType = 15
	EWorkshopFileType_Max                    EWorkshopFileType = 16
)

var EWorkshopFileType_name = map[EWorkshopFileType]string{
	EWorkshopFileType_First:                  "EWorkshopFileType_First",
	EWorkshopFileType_Microtransaction:       "EWorkshopFileType_Microtransaction",
	EWorkshopFileType_Collection:             "EWorkshopFileType_Collection",
	EWorkshopFileType_Art:                    "EWorkshopFileType_Art",
	EWorkshopFileType_Video:                  "EWorkshopFileType_Video",
	EWorkshopFileType_Screenshot:             "EWorkshopFileType_Screenshot",
	EWorkshopFileType_Game:                   "EWorkshopFileType_Game",
	EWorkshopFileType_Software:               "EWorkshopFileType_Software",
	EWorkshopFileType_Concept:                "EWorkshopFileType_Concept",
	EWorkshopFileType_WebGuide:               "EWorkshopFileType_WebGuide",
	EWorkshopFileType_IntegratedGuide:        "EWorkshopFileType_IntegratedGuide",
	EWorkshopFileType_Merch:                  "EWorkshopFileType_Merch",
	EWorkshopFileType_ControllerBinding:      "EWorkshopFileType_ControllerBinding",
	EWorkshopFileType_SteamworksAccessInvite: "EWorkshopFileType_SteamworksAccessInvite",
	EWorkshopFileType_SteamVideo:             "EWorkshopFileType_SteamVideo",
	EWorkshopFileType_GameManagedItem:        "EWorkshopFileType_GameManagedItem",
	EWorkshopFileType_Max:                    "EWorkshopFileType_Max",
}

func (e EWorkshopFileType) String() string {
	if s, ok := EWorkshopFileType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EWorkshopFileType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EWorkshopFileAction int32

const (
	EWorkshopFileAction_Played    EWorkshopFileAction = 0
	EWorkshopFileAction_Completed EWorkshopFileAction = 1
)

var EWorkshopFileAction_name = map[EWorkshopFileAction]string{
	EWorkshopFileAction_Played:    "EWorkshopFileAction_Played",
	EWorkshopFileAction_Completed: "EWorkshopFileAction_Completed",
}

func (e EWorkshopFileAction) String() string {
	if s, ok := EWorkshopFileAction_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EWorkshopFileAction_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EEconTradeResponse int32

const (
	EEconTradeResponse_Accepted                                  EEconTradeResponse = 0
	EEconTradeResponse_Declined                                  EEconTradeResponse = 1
	EEconTradeResponse_TradeBannedInitiator                      EEconTradeResponse = 2
	EEconTradeResponse_TradeBannedTarget                         EEconTradeResponse = 3
	EEconTradeResponse_TargetAlreadyTrading                      EEconTradeResponse = 4
	EEconTradeResponse_Disabled                                  EEconTradeResponse = 5
	EEconTradeResponse_NotLoggedIn                               EEconTradeResponse = 6
	EEconTradeResponse_Cancel                                    EEconTradeResponse = 7
	EEconTradeResponse_TooSoon                                   EEconTradeResponse = 8
	EEconTradeResponse_TooSoonPenalty                            EEconTradeResponse = 9
	EEconTradeResponse_ConnectionFailed                          EEconTradeResponse = 10
	EEconTradeResponse_AlreadyTrading                            EEconTradeResponse = 11
	EEconTradeResponse_AlreadyHasTradeRequest                    EEconTradeResponse = 12
	EEconTradeResponse_NoResponse                                EEconTradeResponse = 13
	EEconTradeResponse_CyberCafeInitiator                        EEconTradeResponse = 14
	EEconTradeResponse_CyberCafeTarget                           EEconTradeResponse = 15
	EEconTradeResponse_SchoolLabInitiator                        EEconTradeResponse = 16
	EEconTradeResponse_SchoolLabTarget                           EEconTradeResponse = 16
	EEconTradeResponse_InitiatorBlockedTarget                    EEconTradeResponse = 18
	EEconTradeResponse_InitiatorNeedsVerifiedEmail               EEconTradeResponse = 20
	EEconTradeResponse_InitiatorNeedsSteamGuard                  EEconTradeResponse = 21
	EEconTradeResponse_TargetAccountCannotTrade                  EEconTradeResponse = 22
	EEconTradeResponse_InitiatorSteamGuardDuration               EEconTradeResponse = 23
	EEconTradeResponse_InitiatorPasswordResetProbation           EEconTradeResponse = 24
	EEconTradeResponse_InitiatorNewDeviceCooldown                EEconTradeResponse = 25
	EEconTradeResponse_InitiatorSentInvalidCookie                EEconTradeResponse = 26
	EEconTradeResponse_NeedsEmailConfirmation                    EEconTradeResponse = 27
	EEconTradeResponse_InitiatorRecentEmailChange                EEconTradeResponse = 28
	EEconTradeResponse_NeedsMobileConfirmation                   EEconTradeResponse = 29
	EEconTradeResponse_TradingHoldForClearedTradeOffersInitiator EEconTradeResponse = 30
	EEconTradeResponse_WouldExceedMaxAssetCount                  EEconTradeResponse = 31
	EEconTradeResponse_OKToDeliver                               EEconTradeResponse = 50
)

var EEconTradeResponse_name = map[EEconTradeResponse]string{
	EEconTradeResponse_Accepted:                                  "EEconTradeResponse_Accepted",
	EEconTradeResponse_Declined:                                  "EEconTradeResponse_Declined",
	EEconTradeResponse_TradeBannedInitiator:                      "EEconTradeResponse_TradeBannedInitiator",
	EEconTradeResponse_TradeBannedTarget:                         "EEconTradeResponse_TradeBannedTarget",
	EEconTradeResponse_TargetAlreadyTrading:                      "EEconTradeResponse_TargetAlreadyTrading",
	EEconTradeResponse_Disabled:                                  "EEconTradeResponse_Disabled",
	EEconTradeResponse_NotLoggedIn:                               "EEconTradeResponse_NotLoggedIn",
	EEconTradeResponse_Cancel:                                    "EEconTradeResponse_Cancel",
	EEconTradeResponse_TooSoon:                                   "EEconTradeResponse_TooSoon",
	EEconTradeResponse_TooSoonPenalty:                            "EEconTradeResponse_TooSoonPenalty",
	EEconTradeResponse_ConnectionFailed:                          "EEconTradeResponse_ConnectionFailed",
	EEconTradeResponse_AlreadyTrading:                            "EEconTradeResponse_AlreadyTrading",
	EEconTradeResponse_AlreadyHasTradeRequest:                    "EEconTradeResponse_AlreadyHasTradeRequest",
	EEconTradeResponse_NoResponse:                                "EEconTradeResponse_NoResponse",
	EEconTradeResponse_CyberCafeInitiator:                        "EEconTradeResponse_CyberCafeInitiator",
	EEconTradeResponse_CyberCafeTarget:                           "EEconTradeResponse_CyberCafeTarget",
	EEconTradeResponse_SchoolLabInitiator:                        "EEconTradeResponse_SchoolLabInitiator",
	EEconTradeResponse_InitiatorBlockedTarget:                    "EEconTradeResponse_InitiatorBlockedTarget",
	EEconTradeResponse_InitiatorNeedsVerifiedEmail:               "EEconTradeResponse_InitiatorNeedsVerifiedEmail",
	EEconTradeResponse_InitiatorNeedsSteamGuard:                  "EEconTradeResponse_InitiatorNeedsSteamGuard",
	EEconTradeResponse_TargetAccountCannotTrade:                  "EEconTradeResponse_TargetAccountCannotTrade",
	EEconTradeResponse_InitiatorSteamGuardDuration:               "EEconTradeResponse_InitiatorSteamGuardDuration",
	EEconTradeResponse_InitiatorPasswordResetProbation:           "EEconTradeResponse_InitiatorPasswordResetProbation",
	EEconTradeResponse_InitiatorNewDeviceCooldown:                "EEconTradeResponse_InitiatorNewDeviceCooldown",
	EEconTradeResponse_InitiatorSentInvalidCookie:                "EEconTradeResponse_InitiatorSentInvalidCookie",
	EEconTradeResponse_NeedsEmailConfirmation:                    "EEconTradeResponse_NeedsEmailConfirmation",
	EEconTradeResponse_InitiatorRecentEmailChange:                "EEconTradeResponse_InitiatorRecentEmailChange",
	EEconTradeResponse_NeedsMobileConfirmation:                   "EEconTradeResponse_NeedsMobileConfirmation",
	EEconTradeResponse_TradingHoldForClearedTradeOffersInitiator: "EEconTradeResponse_TradingHoldForClearedTradeOffersInitiator",
	EEconTradeResponse_WouldExceedMaxAssetCount:                  "EEconTradeResponse_WouldExceedMaxAssetCount",
	EEconTradeResponse_OKToDeliver:                               "EEconTradeResponse_OKToDeliver",
}

func (e EEconTradeResponse) String() string {
	if s, ok := EEconTradeResponse_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EEconTradeResponse_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EMarketingMessageFlags int32

const (
	EMarketingMessageFlags_None                 EMarketingMessageFlags = 0
	EMarketingMessageFlags_HighPriority         EMarketingMessageFlags = 1
	EMarketingMessageFlags_PlatformWindows      EMarketingMessageFlags = 2
	EMarketingMessageFlags_PlatformMac          EMarketingMessageFlags = 4
	EMarketingMessageFlags_PlatformLinux        EMarketingMessageFlags = 8
	EMarketingMessageFlags_PlatformRestrictions EMarketingMessageFlags = EMarketingMessageFlags_PlatformWindows | EMarketingMessageFlags_PlatformMac | EMarketingMessageFlags_PlatformLinux
)

var EMarketingMessageFlags_name = map[EMarketingMessageFlags]string{
	EMarketingMessageFlags_None:                 "EMarketingMessageFlags_None",
	EMarketingMessageFlags_HighPriority:         "EMarketingMessageFlags_HighPriority",
	EMarketingMessageFlags_PlatformWindows:      "EMarketingMessageFlags_PlatformWindows",
	EMarketingMessageFlags_PlatformMac:          "EMarketingMessageFlags_PlatformMac",
	EMarketingMessageFlags_PlatformLinux:        "EMarketingMessageFlags_PlatformLinux",
	EMarketingMessageFlags_PlatformRestrictions: "EMarketingMessageFlags_PlatformRestrictions",
}

func (e EMarketingMessageFlags) String() string {
	if s, ok := EMarketingMessageFlags_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EMarketingMessageFlags_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type ENewsUpdateType int32

const (
	ENewsUpdateType_AppNews      ENewsUpdateType = 0
	ENewsUpdateType_SteamAds     ENewsUpdateType = 1
	ENewsUpdateType_SteamNews    ENewsUpdateType = 2
	ENewsUpdateType_CDDBUpdate   ENewsUpdateType = 3
	ENewsUpdateType_ClientUpdate ENewsUpdateType = 4
)

var ENewsUpdateType_name = map[ENewsUpdateType]string{
	ENewsUpdateType_AppNews:      "ENewsUpdateType_AppNews",
	ENewsUpdateType_SteamAds:     "ENewsUpdateType_SteamAds",
	ENewsUpdateType_SteamNews:    "ENewsUpdateType_SteamNews",
	ENewsUpdateType_CDDBUpdate:   "ENewsUpdateType_CDDBUpdate",
	ENewsUpdateType_ClientUpdate: "ENewsUpdateType_ClientUpdate",
}

func (e ENewsUpdateType) String() string {
	if s, ok := ENewsUpdateType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range ENewsUpdateType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type ESystemIMType int32

const (
	ESystemIMType_RawText                  ESystemIMType = 0
	ESystemIMType_InvalidCard              ESystemIMType = 1
	ESystemIMType_RecurringPurchaseFailed  ESystemIMType = 2
	ESystemIMType_CardWillExpire           ESystemIMType = 3
	ESystemIMType_SubscriptionExpired      ESystemIMType = 4
	ESystemIMType_GuestPassReceived        ESystemIMType = 5
	ESystemIMType_GuestPassGranted         ESystemIMType = 6
	ESystemIMType_GiftRevoked              ESystemIMType = 7
	ESystemIMType_SupportMessage           ESystemIMType = 8
	ESystemIMType_SupportMessageClearAlert ESystemIMType = 9
	ESystemIMType_Max                      ESystemIMType = 10
)

var ESystemIMType_name = map[ESystemIMType]string{
	ESystemIMType_RawText:                  "ESystemIMType_RawText",
	ESystemIMType_InvalidCard:              "ESystemIMType_InvalidCard",
	ESystemIMType_RecurringPurchaseFailed:  "ESystemIMType_RecurringPurchaseFailed",
	ESystemIMType_CardWillExpire:           "ESystemIMType_CardWillExpire",
	ESystemIMType_SubscriptionExpired:      "ESystemIMType_SubscriptionExpired",
	ESystemIMType_GuestPassReceived:        "ESystemIMType_GuestPassReceived",
	ESystemIMType_GuestPassGranted:         "ESystemIMType_GuestPassGranted",
	ESystemIMType_GiftRevoked:              "ESystemIMType_GiftRevoked",
	ESystemIMType_SupportMessage:           "ESystemIMType_SupportMessage",
	ESystemIMType_SupportMessageClearAlert: "ESystemIMType_SupportMessageClearAlert",
	ESystemIMType_Max:                      "ESystemIMType_Max",
}

func (e ESystemIMType) String() string {
	if s, ok := ESystemIMType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range ESystemIMType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EChatFlags int32

const (
	EChatFlags_Locked             EChatFlags = 1
	EChatFlags_InvisibleToFriends EChatFlags = 2
	EChatFlags_Moderated          EChatFlags = 4
	EChatFlags_Unjoinable         EChatFlags = 8
)

var EChatFlags_name = map[EChatFlags]string{
	EChatFlags_Locked:             "EChatFlags_Locked",
	EChatFlags_InvisibleToFriends: "EChatFlags_InvisibleToFriends",
	EChatFlags_Moderated:          "EChatFlags_Moderated",
	EChatFlags_Unjoinable:         "EChatFlags_Unjoinable",
}

func (e EChatFlags) String() string {
	if s, ok := EChatFlags_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EChatFlags_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type ERemoteStoragePlatform int32

const (
	ERemoteStoragePlatform_None      ERemoteStoragePlatform = 0
	ERemoteStoragePlatform_Windows   ERemoteStoragePlatform = 1
	ERemoteStoragePlatform_OSX       ERemoteStoragePlatform = 2
	ERemoteStoragePlatform_PS3       ERemoteStoragePlatform = 4
	ERemoteStoragePlatform_Linux     ERemoteStoragePlatform = 8
	ERemoteStoragePlatform_Reserved1 ERemoteStoragePlatform = 8 // removed
	ERemoteStoragePlatform_Reserved2 ERemoteStoragePlatform = 16
	ERemoteStoragePlatform_All       ERemoteStoragePlatform = -1
)

var ERemoteStoragePlatform_name = map[ERemoteStoragePlatform]string{
	ERemoteStoragePlatform_None:      "ERemoteStoragePlatform_None",
	ERemoteStoragePlatform_Windows:   "ERemoteStoragePlatform_Windows",
	ERemoteStoragePlatform_OSX:       "ERemoteStoragePlatform_OSX",
	ERemoteStoragePlatform_PS3:       "ERemoteStoragePlatform_PS3",
	ERemoteStoragePlatform_Linux:     "ERemoteStoragePlatform_Linux",
	ERemoteStoragePlatform_Reserved2: "ERemoteStoragePlatform_Reserved2",
	ERemoteStoragePlatform_All:       "ERemoteStoragePlatform_All",
}

func (e ERemoteStoragePlatform) String() string {
	if s, ok := ERemoteStoragePlatform_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range ERemoteStoragePlatform_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EDRMBlobDownloadType int32

const (
	EDRMBlobDownloadType_Error        EDRMBlobDownloadType = 0
	EDRMBlobDownloadType_File         EDRMBlobDownloadType = 1
	EDRMBlobDownloadType_Parts        EDRMBlobDownloadType = 2
	EDRMBlobDownloadType_Compressed   EDRMBlobDownloadType = 4
	EDRMBlobDownloadType_AllMask      EDRMBlobDownloadType = 7
	EDRMBlobDownloadType_IsJob        EDRMBlobDownloadType = 8
	EDRMBlobDownloadType_HighPriority EDRMBlobDownloadType = 16
	EDRMBlobDownloadType_AddTimestamp EDRMBlobDownloadType = 32
	EDRMBlobDownloadType_LowPriority  EDRMBlobDownloadType = 64
)

var EDRMBlobDownloadType_name = map[EDRMBlobDownloadType]string{
	EDRMBlobDownloadType_Error:        "EDRMBlobDownloadType_Error",
	EDRMBlobDownloadType_File:         "EDRMBlobDownloadType_File",
	EDRMBlobDownloadType_Parts:        "EDRMBlobDownloadType_Parts",
	EDRMBlobDownloadType_Compressed:   "EDRMBlobDownloadType_Compressed",
	EDRMBlobDownloadType_AllMask:      "EDRMBlobDownloadType_AllMask",
	EDRMBlobDownloadType_IsJob:        "EDRMBlobDownloadType_IsJob",
	EDRMBlobDownloadType_HighPriority: "EDRMBlobDownloadType_HighPriority",
	EDRMBlobDownloadType_AddTimestamp: "EDRMBlobDownloadType_AddTimestamp",
	EDRMBlobDownloadType_LowPriority:  "EDRMBlobDownloadType_LowPriority",
}

func (e EDRMBlobDownloadType) String() string {
	if s, ok := EDRMBlobDownloadType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EDRMBlobDownloadType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EDRMBlobDownloadErrorDetail int32

const (
	EDRMBlobDownloadErrorDetail_None                      EDRMBlobDownloadErrorDetail = 0
	EDRMBlobDownloadErrorDetail_DownloadFailed            EDRMBlobDownloadErrorDetail = 1
	EDRMBlobDownloadErrorDetail_TargetLocked              EDRMBlobDownloadErrorDetail = 2
	EDRMBlobDownloadErrorDetail_OpenZip                   EDRMBlobDownloadErrorDetail = 3
	EDRMBlobDownloadErrorDetail_ReadZipDirectory          EDRMBlobDownloadErrorDetail = 4
	EDRMBlobDownloadErrorDetail_UnexpectedZipEntry        EDRMBlobDownloadErrorDetail = 5
	EDRMBlobDownloadErrorDetail_UnzipFullFile             EDRMBlobDownloadErrorDetail = 6
	EDRMBlobDownloadErrorDetail_UnknownBlobType           EDRMBlobDownloadErrorDetail = 7
	EDRMBlobDownloadErrorDetail_UnzipStrips               EDRMBlobDownloadErrorDetail = 8
	EDRMBlobDownloadErrorDetail_UnzipMergeGuid            EDRMBlobDownloadErrorDetail = 9
	EDRMBlobDownloadErrorDetail_UnzipSignature            EDRMBlobDownloadErrorDetail = 10
	EDRMBlobDownloadErrorDetail_ApplyStrips               EDRMBlobDownloadErrorDetail = 11
	EDRMBlobDownloadErrorDetail_ApplyMergeGuid            EDRMBlobDownloadErrorDetail = 12
	EDRMBlobDownloadErrorDetail_ApplySignature            EDRMBlobDownloadErrorDetail = 13
	EDRMBlobDownloadErrorDetail_AppIdMismatch             EDRMBlobDownloadErrorDetail = 14
	EDRMBlobDownloadErrorDetail_AppIdUnexpected           EDRMBlobDownloadErrorDetail = 15
	EDRMBlobDownloadErrorDetail_AppliedSignatureCorrupt   EDRMBlobDownloadErrorDetail = 16
	EDRMBlobDownloadErrorDetail_ApplyValveSignatureHeader EDRMBlobDownloadErrorDetail = 17
	EDRMBlobDownloadErrorDetail_UnzipValveSignatureHeader EDRMBlobDownloadErrorDetail = 18
	EDRMBlobDownloadErrorDetail_PathManipulationError     EDRMBlobDownloadErrorDetail = 19
	EDRMBlobDownloadErrorDetail_TargetLocked_Base         EDRMBlobDownloadErrorDetail = 65536
	EDRMBlobDownloadErrorDetail_TargetLocked_Max          EDRMBlobDownloadErrorDetail = 131071
	EDRMBlobDownloadErrorDetail_NextBase                  EDRMBlobDownloadErrorDetail = 131072
)

var EDRMBlobDownloadErrorDetail_name = map[EDRMBlobDownloadErrorDetail]string{
	EDRMBlobDownloadErrorDetail_None:                      "EDRMBlobDownloadErrorDetail_None",
	EDRMBlobDownloadErrorDetail_DownloadFailed:            "EDRMBlobDownloadErrorDetail_DownloadFailed",
	EDRMBlobDownloadErrorDetail_TargetLocked:              "EDRMBlobDownloadErrorDetail_TargetLocked",
	EDRMBlobDownloadErrorDetail_OpenZip:                   "EDRMBlobDownloadErrorDetail_OpenZip",
	EDRMBlobDownloadErrorDetail_ReadZipDirectory:          "EDRMBlobDownloadErrorDetail_ReadZipDirectory",
	EDRMBlobDownloadErrorDetail_UnexpectedZipEntry:        "EDRMBlobDownloadErrorDetail_UnexpectedZipEntry",
	EDRMBlobDownloadErrorDetail_UnzipFullFile:             "EDRMBlobDownloadErrorDetail_UnzipFullFile",
	EDRMBlobDownloadErrorDetail_UnknownBlobType:           "EDRMBlobDownloadErrorDetail_UnknownBlobType",
	EDRMBlobDownloadErrorDetail_UnzipStrips:               "EDRMBlobDownloadErrorDetail_UnzipStrips",
	EDRMBlobDownloadErrorDetail_UnzipMergeGuid:            "EDRMBlobDownloadErrorDetail_UnzipMergeGuid",
	EDRMBlobDownloadErrorDetail_UnzipSignature:            "EDRMBlobDownloadErrorDetail_UnzipSignature",
	EDRMBlobDownloadErrorDetail_ApplyStrips:               "EDRMBlobDownloadErrorDetail_ApplyStrips",
	EDRMBlobDownloadErrorDetail_ApplyMergeGuid:            "EDRMBlobDownloadErrorDetail_ApplyMergeGuid",
	EDRMBlobDownloadErrorDetail_ApplySignature:            "EDRMBlobDownloadErrorDetail_ApplySignature",
	EDRMBlobDownloadErrorDetail_AppIdMismatch:             "EDRMBlobDownloadErrorDetail_AppIdMismatch",
	EDRMBlobDownloadErrorDetail_AppIdUnexpected:           "EDRMBlobDownloadErrorDetail_AppIdUnexpected",
	EDRMBlobDownloadErrorDetail_AppliedSignatureCorrupt:   "EDRMBlobDownloadErrorDetail_AppliedSignatureCorrupt",
	EDRMBlobDownloadErrorDetail_ApplyValveSignatureHeader: "EDRMBlobDownloadErrorDetail_ApplyValveSignatureHeader",
	EDRMBlobDownloadErrorDetail_UnzipValveSignatureHeader: "EDRMBlobDownloadErrorDetail_UnzipValveSignatureHeader",
	EDRMBlobDownloadErrorDetail_PathManipulationError:     "EDRMBlobDownloadErrorDetail_PathManipulationError",
	EDRMBlobDownloadErrorDetail_TargetLocked_Base:         "EDRMBlobDownloadErrorDetail_TargetLocked_Base",
	EDRMBlobDownloadErrorDetail_TargetLocked_Max:          "EDRMBlobDownloadErrorDetail_TargetLocked_Max",
	EDRMBlobDownloadErrorDetail_NextBase:                  "EDRMBlobDownloadErrorDetail_NextBase",
}

func (e EDRMBlobDownloadErrorDetail) String() string {
	if s, ok := EDRMBlobDownloadErrorDetail_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EDRMBlobDownloadErrorDetail_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EClientStat int32

const (
	EClientStat_P2PConnectionsUDP   EClientStat = 0
	EClientStat_P2PConnectionsRelay EClientStat = 1
	EClientStat_P2PGameConnections  EClientStat = 2
	EClientStat_P2PVoiceConnections EClientStat = 3
	EClientStat_BytesDownloaded     EClientStat = 4
	EClientStat_Max                 EClientStat = 5
)

var EClientStat_name = map[EClientStat]string{
	EClientStat_P2PConnectionsUDP:   "EClientStat_P2PConnectionsUDP",
	EClientStat_P2PConnectionsRelay: "EClientStat_P2PConnectionsRelay",
	EClientStat_P2PGameConnections:  "EClientStat_P2PGameConnections",
	EClientStat_P2PVoiceConnections: "EClientStat_P2PVoiceConnections",
	EClientStat_BytesDownloaded:     "EClientStat_BytesDownloaded",
	EClientStat_Max:                 "EClientStat_Max",
}

func (e EClientStat) String() string {
	if s, ok := EClientStat_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EClientStat_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EClientStatAggregateMethod int32

const (
	EClientStatAggregateMethod_LatestOnly EClientStatAggregateMethod = 0
	EClientStatAggregateMethod_Sum        EClientStatAggregateMethod = 1
	EClientStatAggregateMethod_Event      EClientStatAggregateMethod = 2
	EClientStatAggregateMethod_Scalar     EClientStatAggregateMethod = 3
)

var EClientStatAggregateMethod_name = map[EClientStatAggregateMethod]string{
	EClientStatAggregateMethod_LatestOnly: "EClientStatAggregateMethod_LatestOnly",
	EClientStatAggregateMethod_Sum:        "EClientStatAggregateMethod_Sum",
	EClientStatAggregateMethod_Event:      "EClientStatAggregateMethod_Event",
	EClientStatAggregateMethod_Scalar:     "EClientStatAggregateMethod_Scalar",
}

func (e EClientStatAggregateMethod) String() string {
	if s, ok := EClientStatAggregateMethod_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EClientStatAggregateMethod_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type ELeaderboardDataRequest int32

const (
	ELeaderboardDataRequest_Global           ELeaderboardDataRequest = 0
	ELeaderboardDataRequest_GlobalAroundUser ELeaderboardDataRequest = 1
	ELeaderboardDataRequest_Friends          ELeaderboardDataRequest = 2
	ELeaderboardDataRequest_Users            ELeaderboardDataRequest = 3
)

var ELeaderboardDataRequest_name = map[ELeaderboardDataRequest]string{
	ELeaderboardDataRequest_Global:           "ELeaderboardDataRequest_Global",
	ELeaderboardDataRequest_GlobalAroundUser: "ELeaderboardDataRequest_GlobalAroundUser",
	ELeaderboardDataRequest_Friends:          "ELeaderboardDataRequest_Friends",
	ELeaderboardDataRequest_Users:            "ELeaderboardDataRequest_Users",
}

func (e ELeaderboardDataRequest) String() string {
	if s, ok := ELeaderboardDataRequest_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range ELeaderboardDataRequest_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type ELeaderboardSortMethod int32

const (
	ELeaderboardSortMethod_None       ELeaderboardSortMethod = 0
	ELeaderboardSortMethod_Ascending  ELeaderboardSortMethod = 1
	ELeaderboardSortMethod_Descending ELeaderboardSortMethod = 2
)

var ELeaderboardSortMethod_name = map[ELeaderboardSortMethod]string{
	ELeaderboardSortMethod_None:       "ELeaderboardSortMethod_None",
	ELeaderboardSortMethod_Ascending:  "ELeaderboardSortMethod_Ascending",
	ELeaderboardSortMethod_Descending: "ELeaderboardSortMethod_Descending",
}

func (e ELeaderboardSortMethod) String() string {
	if s, ok := ELeaderboardSortMethod_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range ELeaderboardSortMethod_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type ELeaderboardDisplayType int32

const (
	ELeaderboardDisplayType_None             ELeaderboardDisplayType = 0
	ELeaderboardDisplayType_Numeric          ELeaderboardDisplayType = 1
	ELeaderboardDisplayType_TimeSeconds      ELeaderboardDisplayType = 2
	ELeaderboardDisplayType_TimeMilliSeconds ELeaderboardDisplayType = 3
)

var ELeaderboardDisplayType_name = map[ELeaderboardDisplayType]string{
	ELeaderboardDisplayType_None:             "ELeaderboardDisplayType_None",
	ELeaderboardDisplayType_Numeric:          "ELeaderboardDisplayType_Numeric",
	ELeaderboardDisplayType_TimeSeconds:      "ELeaderboardDisplayType_TimeSeconds",
	ELeaderboardDisplayType_TimeMilliSeconds: "ELeaderboardDisplayType_TimeMilliSeconds",
}

func (e ELeaderboardDisplayType) String() string {
	if s, ok := ELeaderboardDisplayType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range ELeaderboardDisplayType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type ELeaderboardUploadScoreMethod int32

const (
	ELeaderboardUploadScoreMethod_None        ELeaderboardUploadScoreMethod = 0
	ELeaderboardUploadScoreMethod_KeepBest    ELeaderboardUploadScoreMethod = 1
	ELeaderboardUploadScoreMethod_ForceUpdate ELeaderboardUploadScoreMethod = 2
)

var ELeaderboardUploadScoreMethod_name = map[ELeaderboardUploadScoreMethod]string{
	ELeaderboardUploadScoreMethod_None:        "ELeaderboardUploadScoreMethod_None",
	ELeaderboardUploadScoreMethod_KeepBest:    "ELeaderboardUploadScoreMethod_KeepBest",
	ELeaderboardUploadScoreMethod_ForceUpdate: "ELeaderboardUploadScoreMethod_ForceUpdate",
}

func (e ELeaderboardUploadScoreMethod) String() string {
	if s, ok := ELeaderboardUploadScoreMethod_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range ELeaderboardUploadScoreMethod_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EUCMFilePrivacyState int32

const (
	EUCMFilePrivacyState_Invalid     EUCMFilePrivacyState = -1
	EUCMFilePrivacyState_Private     EUCMFilePrivacyState = 2
	EUCMFilePrivacyState_FriendsOnly EUCMFilePrivacyState = 4
	EUCMFilePrivacyState_Public      EUCMFilePrivacyState = 8
	EUCMFilePrivacyState_All         EUCMFilePrivacyState = EUCMFilePrivacyState_Public | EUCMFilePrivacyState_FriendsOnly | EUCMFilePrivacyState_Private
)

var EUCMFilePrivacyState_name = map[EUCMFilePrivacyState]string{
	EUCMFilePrivacyState_Invalid:     "EUCMFilePrivacyState_Invalid",
	EUCMFilePrivacyState_Private:     "EUCMFilePrivacyState_Private",
	EUCMFilePrivacyState_FriendsOnly: "EUCMFilePrivacyState_FriendsOnly",
	EUCMFilePrivacyState_Public:      "EUCMFilePrivacyState_Public",
	EUCMFilePrivacyState_All:         "EUCMFilePrivacyState_All",
}

func (e EUCMFilePrivacyState) String() string {
	if s, ok := EUCMFilePrivacyState_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EUCMFilePrivacyState_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

type EUdpPacketType uint8

const (
	EUdpPacketType_Invalid      EUdpPacketType = 0
	EUdpPacketType_ChallengeReq EUdpPacketType = 1
	EUdpPacketType_Challenge    EUdpPacketType = 2
	EUdpPacketType_Connect      EUdpPacketType = 3
	EUdpPacketType_Accept       EUdpPacketType = 4
	EUdpPacketType_Disconnect   EUdpPacketType = 5
	EUdpPacketType_Data         EUdpPacketType = 6
	EUdpPacketType_Datagram     EUdpPacketType = 7
	EUdpPacketType_Max          EUdpPacketType = 8
)

var EUdpPacketType_name = map[EUdpPacketType]string{
	EUdpPacketType_Invalid:      "EUdpPacketType_Invalid",
	EUdpPacketType_ChallengeReq: "EUdpPacketType_ChallengeReq",
	EUdpPacketType_Challenge:    "EUdpPacketType_Challenge",
	EUdpPacketType_Connect:      "EUdpPacketType_Connect",
	EUdpPacketType_Accept:       "EUdpPacketType_Accept",
	EUdpPacketType_Disconnect:   "EUdpPacketType_Disconnect",
	EUdpPacketType_Data:         "EUdpPacketType_Data",
	EUdpPacketType_Datagram:     "EUdpPacketType_Datagram",
	EUdpPacketType_Max:          "EUdpPacketType_Max",
}

func (e EUdpPacketType) String() string {
	if s, ok := EUdpPacketType_name[e]; ok {
		return s
	}
	var flags []string
	for k, v := range EUdpPacketType_name {
		if e&k != 0 {
			flags = append(flags, v)
		}
	}
	if flags == nil {
		return fmt.Sprintf("%d", e)
	}
	sort.Strings(flags)
	return strings.Join(flags, " | ")
}

const (
	UdpHeader_MAGIC uint32 = 0x31305356
)

type UdpHeader struct {
	Magic        uint32
	PayloadSize  uint16
	PacketType   EUdpPacketType
	Flags        uint8
	SourceConnID uint32
	DestConnID   uint32
	SeqThis      uint32
	SeqAck       uint32
	PacketsInMsg uint32
	MsgStartSeq  uint32
	MsgSize      uint32
}

func NewUdpHeader() *UdpHeader {
	return &UdpHeader{
		Magic:        UdpHeader_MAGIC,
		PacketType:   EUdpPacketType_Invalid,
		SourceConnID: 512,
	}
}

func (m *UdpHeader) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.Magic); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.PayloadSize); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.PacketType); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.Flags); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SourceConnID); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.DestConnID); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SeqThis); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SeqAck); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.PacketsInMsg); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.MsgStartSeq); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.MsgSize); err != nil {
		return
	}
	return
}

func (m *UdpHeader) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.Magic); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.PayloadSize); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.PacketType); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.Flags); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SourceConnID); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.DestConnID); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SeqThis); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SeqAck); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.PacketsInMsg); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.MsgStartSeq); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.MsgSize); err != nil {
		return
	}
	return
}

const (
	ChallengeData_CHALLENGE_MASK uint32 = 0xA426DF2B
)

type ChallengeData struct {
	ChallengeValue uint32
	ServerLoad     uint32
}

func NewChallengeData() *ChallengeData {
	return new(ChallengeData)
}

func (m *ChallengeData) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.ChallengeValue); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.ServerLoad); err != nil {
		return
	}
	return
}

func (m *ChallengeData) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.ChallengeValue); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.ServerLoad); err != nil {
		return
	}
	return
}

const (
	ConnectData_CHALLENGE_MASK uint32 = ChallengeData_CHALLENGE_MASK
)

type ConnectData struct {
	ChallengeValue uint32
}

func NewConnectData() *ConnectData {
	return new(ConnectData)
}

func (m *ConnectData) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.ChallengeValue); err != nil {
		return
	}
	return
}

func (m *ConnectData) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.ChallengeValue); err != nil {
		return
	}
	return
}

type Accept struct {
}

func NewAccept() *Accept {
	return new(Accept)
}

func (m *Accept) Serialize(w io.Writer) (err error) {
	return
}

func (m *Accept) Deserialize(r io.Reader) (err error) {
	return
}

type Datagram struct {
}

func NewDatagram() *Datagram {
	return new(Datagram)
}

func (m *Datagram) Serialize(w io.Writer) (err error) {
	return
}

func (m *Datagram) Deserialize(r io.Reader) (err error) {
	return
}

type Disconnect struct {
}

func NewDisconnect() *Disconnect {
	return new(Disconnect)
}

func (m *Disconnect) Serialize(w io.Writer) (err error) {
	return
}

func (m *Disconnect) Deserialize(r io.Reader) (err error) {
	return
}

type MsgHdr struct {
	Msg         EMsg
	TargetJobID uint64
	SourceJobID uint64
}

func NewMsgHdr() *MsgHdr {
	return &MsgHdr{
		Msg:         EMsg_Invalid,
		TargetJobID: ^uint64(0),
		SourceJobID: ^uint64(0),
	}
}

func (m *MsgHdr) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.Msg); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.TargetJobID); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SourceJobID); err != nil {
		return
	}
	return
}

func (m *MsgHdr) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.Msg); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.TargetJobID); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SourceJobID); err != nil {
		return
	}
	return
}

type ExtendedClientMsgHdr struct {
	Msg           EMsg
	HeaderSize    uint8
	HeaderVersion uint16
	TargetJobID   uint64
	SourceJobID   uint64
	HeaderCanary  uint8
	SteamID       uint64
	SessionID     int32
}

func NewExtendedClientMsgHdr() *ExtendedClientMsgHdr {
	return &ExtendedClientMsgHdr{
		Msg:           EMsg_Invalid,
		HeaderSize:    36,
		HeaderVersion: 2,
		TargetJobID:   ^uint64(0),
		SourceJobID:   ^uint64(0),
		HeaderCanary:  239,
	}
}

func (m *ExtendedClientMsgHdr) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.Msg); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.HeaderSize); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.HeaderVersion); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.TargetJobID); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SourceJobID); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.HeaderCanary); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SteamID); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SessionID); err != nil {
		return
	}
	return
}

func (m *ExtendedClientMsgHdr) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.Msg); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.HeaderSize); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.HeaderVersion); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.TargetJobID); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SourceJobID); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.HeaderCanary); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SteamID); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SessionID); err != nil {
		return
	}
	return
}

type MsgHdrProtoBuf struct {
	Msg          EMsg
	HeaderLength int32
	Proto        *steam.CMsgProtoBufHeader
}

func NewMsgHdrProtoBuf() *MsgHdrProtoBuf {
	return &MsgHdrProtoBuf{
		Msg: EMsg_Invalid,
	}
}

func (m *MsgHdrProtoBuf) Serialize(w io.Writer) (err error) {
	data0 := MaskProto(uint32(m.Msg))
	if err = binary.Write(w, binary.LittleEndian, data0); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.HeaderLength); err != nil {
		return
	}
	var data2 []byte
	if data2, err = proto.Marshal(m.Proto); err != nil {
		return
	}
	m.HeaderLength = int32(len(data2))
	if _, err = w.Write(data2); err != nil {
		return
	}
	return
}

func (m *MsgHdrProtoBuf) Deserialize(r io.Reader) (err error) {
	var data0 uint32
	if err = binary.Read(r, binary.LittleEndian, &data0); err != nil {
		return
	}
	m.Msg = MakeEMsg(data0)
	if err = binary.Read(r, binary.LittleEndian, &m.HeaderLength); err != nil {
		return
	}
	data2 := make([]byte, m.HeaderLength, m.HeaderLength)
	if _, err = io.ReadFull(r, data2); err != nil {
		return
	}
	if err = proto.Unmarshal(data2, m.Proto); err != nil {
		return
	}
	return
}

type MsgGCHdrProtoBuf struct {
	Msg          uint32
	HeaderLength int32
	Proto        *steam.CMsgProtoBufHeader
}

func NewMsgGCHdrProtoBuf() *MsgGCHdrProtoBuf {
	return &MsgGCHdrProtoBuf{
		Msg: 0,
	}
}

func (m *MsgGCHdrProtoBuf) Serialize(w io.Writer) (err error) {
	data0 := MaskProto(uint32(m.Msg))
	if err = binary.Write(w, binary.LittleEndian, data0); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.HeaderLength); err != nil {
		return
	}
	var data2 []byte
	if data2, err = proto.Marshal(m.Proto); err != nil {
		return
	}
	m.HeaderLength = int32(len(data2))
	if _, err = w.Write(data2); err != nil {
		return
	}
	return
}

func (m *MsgGCHdrProtoBuf) Deserialize(r io.Reader) (err error) {
	var data0 uint32
	if err = binary.Read(r, binary.LittleEndian, &data0); err != nil {
		return
	}
	m.Msg = uint32(MakeEMsg(data0))
	if err = binary.Read(r, binary.LittleEndian, &m.HeaderLength); err != nil {
		return
	}
	data2 := make([]byte, m.HeaderLength, m.HeaderLength)
	if _, err = io.ReadFull(r, data2); err != nil {
		return
	}
	if err = proto.Unmarshal(data2, m.Proto); err != nil {
		return
	}
	return
}

type MsgGCHdr struct {
	HeaderVersion uint16
	TargetJobID   uint64
	SourceJobID   uint64
}

func NewMsgGCHdr() *MsgGCHdr {
	return &MsgGCHdr{
		HeaderVersion: 1,
		TargetJobID:   ^uint64(0),
		SourceJobID:   ^uint64(0),
	}
}

func (m *MsgGCHdr) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.HeaderVersion); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.TargetJobID); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SourceJobID); err != nil {
		return
	}
	return
}

func (m *MsgGCHdr) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.HeaderVersion); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.TargetJobID); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SourceJobID); err != nil {
		return
	}
	return
}

type MsgClientJustStrings struct {
}

func NewMsgClientJustStrings() *MsgClientJustStrings {
	return new(MsgClientJustStrings)
}

func (m *MsgClientJustStrings) GetEMsg() EMsg {
	return EMsg_Invalid
}

func (m *MsgClientJustStrings) Serialize(w io.Writer) (err error) {
	return
}

func (m *MsgClientJustStrings) Deserialize(r io.Reader) (err error) {
	return
}

type MsgClientGenericResponse struct {
	Result EResult
}

func NewMsgClientGenericResponse() *MsgClientGenericResponse {
	return new(MsgClientGenericResponse)
}

func (m *MsgClientGenericResponse) GetEMsg() EMsg {
	return EMsg_Invalid
}

func (m *MsgClientGenericResponse) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.Result); err != nil {
		return
	}
	return
}

func (m *MsgClientGenericResponse) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.Result); err != nil {
		return
	}
	return
}

const (
	MsgChannelEncryptRequest_PROTOCOL_VERSION uint32 = 1
)

type MsgChannelEncryptRequest struct {
	ProtocolVersion uint32
	Universe        EUniverse
}

func NewMsgChannelEncryptRequest() *MsgChannelEncryptRequest {
	return &MsgChannelEncryptRequest{
		ProtocolVersion: MsgChannelEncryptRequest_PROTOCOL_VERSION,
		Universe:        EUniverse_Invalid,
	}
}

func (m *MsgChannelEncryptRequest) GetEMsg() EMsg {
	return EMsg_ChannelEncryptRequest
}

func (m *MsgChannelEncryptRequest) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.ProtocolVersion); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.Universe); err != nil {
		return
	}
	return
}

func (m *MsgChannelEncryptRequest) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.ProtocolVersion); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.Universe); err != nil {
		return
	}
	return
}

type MsgChannelEncryptResponse struct {
	ProtocolVersion uint32
	KeySize         uint32
}

func NewMsgChannelEncryptResponse() *MsgChannelEncryptResponse {
	return &MsgChannelEncryptResponse{
		ProtocolVersion: MsgChannelEncryptRequest_PROTOCOL_VERSION,
		KeySize:         128,
	}
}

func (m *MsgChannelEncryptResponse) GetEMsg() EMsg {
	return EMsg_ChannelEncryptResponse
}

func (m *MsgChannelEncryptResponse) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.ProtocolVersion); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.KeySize); err != nil {
		return
	}
	return
}

func (m *MsgChannelEncryptResponse) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.ProtocolVersion); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.KeySize); err != nil {
		return
	}
	return
}

type MsgChannelEncryptResult struct {
	Result EResult
}

func NewMsgChannelEncryptResult() *MsgChannelEncryptResult {
	return &MsgChannelEncryptResult{
		Result: EResult_Invalid,
	}
}

func (m *MsgChannelEncryptResult) GetEMsg() EMsg {
	return EMsg_ChannelEncryptResult
}

func (m *MsgChannelEncryptResult) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.Result); err != nil {
		return
	}
	return
}

func (m *MsgChannelEncryptResult) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.Result); err != nil {
		return
	}
	return
}

type MsgClientNewLoginKey struct {
	UniqueID uint32
	LoginKey [20]byte
}

func NewMsgClientNewLoginKey() *MsgClientNewLoginKey {
	return new(MsgClientNewLoginKey)
}

func (m *MsgClientNewLoginKey) GetEMsg() EMsg {
	return EMsg_ClientNewLoginKey
}

func (m *MsgClientNewLoginKey) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.UniqueID); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.LoginKey); err != nil {
		return
	}
	return
}

func (m *MsgClientNewLoginKey) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.UniqueID); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.LoginKey); err != nil {
		return
	}
	return
}

type MsgClientNewLoginKeyAccepted struct {
	UniqueID uint32
}

func NewMsgClientNewLoginKeyAccepted() *MsgClientNewLoginKeyAccepted {
	return new(MsgClientNewLoginKeyAccepted)
}

func (m *MsgClientNewLoginKeyAccepted) GetEMsg() EMsg {
	return EMsg_ClientNewLoginKeyAccepted
}

func (m *MsgClientNewLoginKeyAccepted) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.UniqueID); err != nil {
		return
	}
	return
}

func (m *MsgClientNewLoginKeyAccepted) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.UniqueID); err != nil {
		return
	}
	return
}

const (
	MsgClientLogon_ObfuscationMask                                      uint32 = 0xBAADF00D
	MsgClientLogon_CurrentProtocol                                      uint32 = 65580
	MsgClientLogon_ProtocolVerMajorMask                                 uint32 = 0xFFFF0000
	MsgClientLogon_ProtocolVerMinorMask                                 uint32 = 0xFFFF
	MsgClientLogon_ProtocolVerMinorMinGameServers                       uint16 = 4
	MsgClientLogon_ProtocolVerMinorMinForSupportingEMsgMulti            uint16 = 12
	MsgClientLogon_ProtocolVerMinorMinForSupportingEMsgClientEncryptPct uint16 = 14
	MsgClientLogon_ProtocolVerMinorMinForExtendedMsgHdr                 uint16 = 17
	MsgClientLogon_ProtocolVerMinorMinForCellId                         uint16 = 18
	MsgClientLogon_ProtocolVerMinorMinForSessionIDLast                  uint16 = 19
	MsgClientLogon_ProtocolVerMinorMinForServerAvailablityMsgs          uint16 = 24
	MsgClientLogon_ProtocolVerMinorMinClients                           uint16 = 25
	MsgClientLogon_ProtocolVerMinorMinForOSType                         uint16 = 26
	MsgClientLogon_ProtocolVerMinorMinForCegApplyPESig                  uint16 = 27
	MsgClientLogon_ProtocolVerMinorMinForMarketingMessages2             uint16 = 27
	MsgClientLogon_ProtocolVerMinorMinForAnyProtoBufMessages            uint16 = 28
	MsgClientLogon_ProtocolVerMinorMinForProtoBufLoggedOffMessage       uint16 = 28
	MsgClientLogon_ProtocolVerMinorMinForProtoBufMultiMessages          uint16 = 28
	MsgClientLogon_ProtocolVerMinorMinForSendingProtocolToUFS           uint16 = 30
	MsgClientLogon_ProtocolVerMinorMinForMachineAuth                    uint16 = 33
	MsgClientLogon_ProtocolVerMinorMinForSessionIDLastAnon              uint16 = 36
	MsgClientLogon_ProtocolVerMinorMinForEnhancedAppList                uint16 = 40
	MsgClientLogon_ProtocolVerMinorMinForSteamGuardNotificationUI       uint16 = 41
	MsgClientLogon_ProtocolVerMinorMinForProtoBufServiceModuleCalls     uint16 = 42
	MsgClientLogon_ProtocolVerMinorMinForGzipMultiMessages              uint16 = 43
	MsgClientLogon_ProtocolVerMinorMinForNewVoiceCallAuthorize          uint16 = 44
	MsgClientLogon_ProtocolVerMinorMinForClientInstanceIDs              uint16 = 44
)

type MsgClientLogon struct {
}

func NewMsgClientLogon() *MsgClientLogon {
	return new(MsgClientLogon)
}

func (m *MsgClientLogon) GetEMsg() EMsg {
	return EMsg_ClientLogon
}

func (m *MsgClientLogon) Serialize(w io.Writer) (err error) {
	return
}

func (m *MsgClientLogon) Deserialize(r io.Reader) (err error) {
	return
}

type MsgClientVACBanStatus struct {
	NumBans uint32
}

func NewMsgClientVACBanStatus() *MsgClientVACBanStatus {
	return new(MsgClientVACBanStatus)
}

func (m *MsgClientVACBanStatus) GetEMsg() EMsg {
	return EMsg_ClientVACBanStatus
}

func (m *MsgClientVACBanStatus) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.NumBans); err != nil {
		return
	}
	return
}

func (m *MsgClientVACBanStatus) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.NumBans); err != nil {
		return
	}
	return
}

type MsgClientAppUsageEvent struct {
	AppUsageEvent EAppUsageEvent
	GameID        uint64
	Offline       uint16
}

func NewMsgClientAppUsageEvent() *MsgClientAppUsageEvent {
	return new(MsgClientAppUsageEvent)
}

func (m *MsgClientAppUsageEvent) GetEMsg() EMsg {
	return EMsg_ClientAppUsageEvent
}

func (m *MsgClientAppUsageEvent) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.AppUsageEvent); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.GameID); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.Offline); err != nil {
		return
	}
	return
}

func (m *MsgClientAppUsageEvent) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.AppUsageEvent); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.GameID); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.Offline); err != nil {
		return
	}
	return
}

type MsgClientEmailAddrInfo struct {
	PasswordStrength           uint32
	FlagsAccountSecurityPolicy uint32
	Validated                  bool
}

func NewMsgClientEmailAddrInfo() *MsgClientEmailAddrInfo {
	return new(MsgClientEmailAddrInfo)
}

func (m *MsgClientEmailAddrInfo) GetEMsg() EMsg {
	return EMsg_ClientEmailAddrInfo
}

func (m *MsgClientEmailAddrInfo) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.PasswordStrength); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.FlagsAccountSecurityPolicy); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.Validated); err != nil {
		return
	}
	return
}

func (m *MsgClientEmailAddrInfo) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.PasswordStrength); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.FlagsAccountSecurityPolicy); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.Validated); err != nil {
		return
	}
	return
}

type MsgClientUpdateGuestPassesList struct {
	Result                   EResult
	CountGuestPassesToGive   int32
	CountGuestPassesToRedeem int32
}

func NewMsgClientUpdateGuestPassesList() *MsgClientUpdateGuestPassesList {
	return new(MsgClientUpdateGuestPassesList)
}

func (m *MsgClientUpdateGuestPassesList) GetEMsg() EMsg {
	return EMsg_ClientUpdateGuestPassesList
}

func (m *MsgClientUpdateGuestPassesList) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.Result); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.CountGuestPassesToGive); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.CountGuestPassesToRedeem); err != nil {
		return
	}
	return
}

func (m *MsgClientUpdateGuestPassesList) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.Result); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.CountGuestPassesToGive); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.CountGuestPassesToRedeem); err != nil {
		return
	}
	return
}

type MsgClientRequestedClientStats struct {
	CountStats int32
}

func NewMsgClientRequestedClientStats() *MsgClientRequestedClientStats {
	return new(MsgClientRequestedClientStats)
}

func (m *MsgClientRequestedClientStats) GetEMsg() EMsg {
	return EMsg_ClientRequestedClientStats
}

func (m *MsgClientRequestedClientStats) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.CountStats); err != nil {
		return
	}
	return
}

func (m *MsgClientRequestedClientStats) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.CountStats); err != nil {
		return
	}
	return
}

type MsgClientP2PIntroducerMessage struct {
	SteamID     uint64
	RoutingType EIntroducerRouting
	Data        [1450]byte
	DataLen     uint32
}

func NewMsgClientP2PIntroducerMessage() *MsgClientP2PIntroducerMessage {
	return new(MsgClientP2PIntroducerMessage)
}

func (m *MsgClientP2PIntroducerMessage) GetEMsg() EMsg {
	return EMsg_ClientP2PIntroducerMessage
}

func (m *MsgClientP2PIntroducerMessage) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.SteamID); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.RoutingType); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.Data); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.DataLen); err != nil {
		return
	}
	return
}

func (m *MsgClientP2PIntroducerMessage) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.SteamID); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.RoutingType); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.Data); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.DataLen); err != nil {
		return
	}
	return
}

type MsgClientOGSBeginSession struct {
	AccountType uint8
	AccountId   uint64
	AppId       uint32
	TimeStarted uint32
}

func NewMsgClientOGSBeginSession() *MsgClientOGSBeginSession {
	return new(MsgClientOGSBeginSession)
}

func (m *MsgClientOGSBeginSession) GetEMsg() EMsg {
	return EMsg_ClientOGSBeginSession
}

func (m *MsgClientOGSBeginSession) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.AccountType); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.AccountId); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.AppId); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.TimeStarted); err != nil {
		return
	}
	return
}

func (m *MsgClientOGSBeginSession) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.AccountType); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.AccountId); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.AppId); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.TimeStarted); err != nil {
		return
	}
	return
}

type MsgClientOGSBeginSessionResponse struct {
	Result            EResult
	CollectingAny     bool
	CollectingDetails bool
	SessionId         uint64
}

func NewMsgClientOGSBeginSessionResponse() *MsgClientOGSBeginSessionResponse {
	return new(MsgClientOGSBeginSessionResponse)
}

func (m *MsgClientOGSBeginSessionResponse) GetEMsg() EMsg {
	return EMsg_ClientOGSBeginSessionResponse
}

func (m *MsgClientOGSBeginSessionResponse) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.Result); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.CollectingAny); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.CollectingDetails); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SessionId); err != nil {
		return
	}
	return
}

func (m *MsgClientOGSBeginSessionResponse) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.Result); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.CollectingAny); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.CollectingDetails); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SessionId); err != nil {
		return
	}
	return
}

type MsgClientOGSEndSession struct {
	SessionId       uint64
	TimeEnded       uint32
	ReasonCode      int32
	CountAttributes int32
}

func NewMsgClientOGSEndSession() *MsgClientOGSEndSession {
	return new(MsgClientOGSEndSession)
}

func (m *MsgClientOGSEndSession) GetEMsg() EMsg {
	return EMsg_ClientOGSEndSession
}

func (m *MsgClientOGSEndSession) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.SessionId); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.TimeEnded); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.ReasonCode); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.CountAttributes); err != nil {
		return
	}
	return
}

func (m *MsgClientOGSEndSession) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.SessionId); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.TimeEnded); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.ReasonCode); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.CountAttributes); err != nil {
		return
	}
	return
}

type MsgClientOGSEndSessionResponse struct {
	Result EResult
}

func NewMsgClientOGSEndSessionResponse() *MsgClientOGSEndSessionResponse {
	return new(MsgClientOGSEndSessionResponse)
}

func (m *MsgClientOGSEndSessionResponse) GetEMsg() EMsg {
	return EMsg_ClientOGSEndSessionResponse
}

func (m *MsgClientOGSEndSessionResponse) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.Result); err != nil {
		return
	}
	return
}

func (m *MsgClientOGSEndSessionResponse) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.Result); err != nil {
		return
	}
	return
}

type MsgClientOGSWriteRow struct {
	SessionId       uint64
	CountAttributes int32
}

func NewMsgClientOGSWriteRow() *MsgClientOGSWriteRow {
	return new(MsgClientOGSWriteRow)
}

func (m *MsgClientOGSWriteRow) GetEMsg() EMsg {
	return EMsg_ClientOGSWriteRow
}

func (m *MsgClientOGSWriteRow) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.SessionId); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.CountAttributes); err != nil {
		return
	}
	return
}

func (m *MsgClientOGSWriteRow) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.SessionId); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.CountAttributes); err != nil {
		return
	}
	return
}

type MsgClientGetFriendsWhoPlayGame struct {
	GameId uint64
}

func NewMsgClientGetFriendsWhoPlayGame() *MsgClientGetFriendsWhoPlayGame {
	return new(MsgClientGetFriendsWhoPlayGame)
}

func (m *MsgClientGetFriendsWhoPlayGame) GetEMsg() EMsg {
	return EMsg_ClientGetFriendsWhoPlayGame
}

func (m *MsgClientGetFriendsWhoPlayGame) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.GameId); err != nil {
		return
	}
	return
}

func (m *MsgClientGetFriendsWhoPlayGame) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.GameId); err != nil {
		return
	}
	return
}

type MsgClientGetFriendsWhoPlayGameResponse struct {
	Result       EResult
	GameId       uint64
	CountFriends uint32
}

func NewMsgClientGetFriendsWhoPlayGameResponse() *MsgClientGetFriendsWhoPlayGameResponse {
	return new(MsgClientGetFriendsWhoPlayGameResponse)
}

func (m *MsgClientGetFriendsWhoPlayGameResponse) GetEMsg() EMsg {
	return EMsg_ClientGetFriendsWhoPlayGameResponse
}

func (m *MsgClientGetFriendsWhoPlayGameResponse) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.Result); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.GameId); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.CountFriends); err != nil {
		return
	}
	return
}

func (m *MsgClientGetFriendsWhoPlayGameResponse) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.Result); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.GameId); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.CountFriends); err != nil {
		return
	}
	return
}

type MsgGSPerformHardwareSurvey struct {
	Flags uint32
}

func NewMsgGSPerformHardwareSurvey() *MsgGSPerformHardwareSurvey {
	return new(MsgGSPerformHardwareSurvey)
}

func (m *MsgGSPerformHardwareSurvey) GetEMsg() EMsg {
	return EMsg_GSPerformHardwareSurvey
}

func (m *MsgGSPerformHardwareSurvey) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.Flags); err != nil {
		return
	}
	return
}

func (m *MsgGSPerformHardwareSurvey) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.Flags); err != nil {
		return
	}
	return
}

type MsgGSGetPlayStatsResponse struct {
	Result                EResult
	Rank                  int32
	LifetimeConnects      uint32
	LifetimeMinutesPlayed uint32
}

func NewMsgGSGetPlayStatsResponse() *MsgGSGetPlayStatsResponse {
	return new(MsgGSGetPlayStatsResponse)
}

func (m *MsgGSGetPlayStatsResponse) GetEMsg() EMsg {
	return EMsg_GSGetPlayStatsResponse
}

func (m *MsgGSGetPlayStatsResponse) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.Result); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.Rank); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.LifetimeConnects); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.LifetimeMinutesPlayed); err != nil {
		return
	}
	return
}

func (m *MsgGSGetPlayStatsResponse) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.Result); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.Rank); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.LifetimeConnects); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.LifetimeMinutesPlayed); err != nil {
		return
	}
	return
}

type MsgGSGetReputationResponse struct {
	Result          EResult
	ReputationScore uint32
	Banned          bool
	BannedIp        uint32
	BannedPort      uint16
	BannedGameId    uint64
	TimeBanExpires  uint32
}

func NewMsgGSGetReputationResponse() *MsgGSGetReputationResponse {
	return new(MsgGSGetReputationResponse)
}

func (m *MsgGSGetReputationResponse) GetEMsg() EMsg {
	return EMsg_GSGetReputationResponse
}

func (m *MsgGSGetReputationResponse) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.Result); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.ReputationScore); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.Banned); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.BannedIp); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.BannedPort); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.BannedGameId); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.TimeBanExpires); err != nil {
		return
	}
	return
}

func (m *MsgGSGetReputationResponse) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.Result); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.ReputationScore); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.Banned); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.BannedIp); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.BannedPort); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.BannedGameId); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.TimeBanExpires); err != nil {
		return
	}
	return
}

type MsgGSDeny struct {
	SteamId    uint64
	DenyReason EDenyReason
}

func NewMsgGSDeny() *MsgGSDeny {
	return new(MsgGSDeny)
}

func (m *MsgGSDeny) GetEMsg() EMsg {
	return EMsg_GSDeny
}

func (m *MsgGSDeny) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.SteamId); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.DenyReason); err != nil {
		return
	}
	return
}

func (m *MsgGSDeny) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.SteamId); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.DenyReason); err != nil {
		return
	}
	return
}

type MsgGSApprove struct {
	SteamId uint64
}

func NewMsgGSApprove() *MsgGSApprove {
	return new(MsgGSApprove)
}

func (m *MsgGSApprove) GetEMsg() EMsg {
	return EMsg_GSApprove
}

func (m *MsgGSApprove) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.SteamId); err != nil {
		return
	}
	return
}

func (m *MsgGSApprove) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.SteamId); err != nil {
		return
	}
	return
}

type MsgGSKick struct {
	SteamId          uint64
	DenyReason       EDenyReason
	WaitTilMapChange int32
}

func NewMsgGSKick() *MsgGSKick {
	return new(MsgGSKick)
}

func (m *MsgGSKick) GetEMsg() EMsg {
	return EMsg_GSKick
}

func (m *MsgGSKick) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.SteamId); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.DenyReason); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.WaitTilMapChange); err != nil {
		return
	}
	return
}

func (m *MsgGSKick) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.SteamId); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.DenyReason); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.WaitTilMapChange); err != nil {
		return
	}
	return
}

type MsgGSGetUserGroupStatus struct {
	SteamIdUser  uint64
	SteamIdGroup uint64
}

func NewMsgGSGetUserGroupStatus() *MsgGSGetUserGroupStatus {
	return new(MsgGSGetUserGroupStatus)
}

func (m *MsgGSGetUserGroupStatus) GetEMsg() EMsg {
	return EMsg_GSGetUserGroupStatus
}

func (m *MsgGSGetUserGroupStatus) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdUser); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdGroup); err != nil {
		return
	}
	return
}

func (m *MsgGSGetUserGroupStatus) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdUser); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdGroup); err != nil {
		return
	}
	return
}

type MsgGSGetUserGroupStatusResponse struct {
	SteamIdUser      uint64
	SteamIdGroup     uint64
	ClanRelationship EClanRelationship
	ClanRank         EClanRank
}

func NewMsgGSGetUserGroupStatusResponse() *MsgGSGetUserGroupStatusResponse {
	return new(MsgGSGetUserGroupStatusResponse)
}

func (m *MsgGSGetUserGroupStatusResponse) GetEMsg() EMsg {
	return EMsg_GSGetUserGroupStatusResponse
}

func (m *MsgGSGetUserGroupStatusResponse) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdUser); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdGroup); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.ClanRelationship); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.ClanRank); err != nil {
		return
	}
	return
}

func (m *MsgGSGetUserGroupStatusResponse) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdUser); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdGroup); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.ClanRelationship); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.ClanRank); err != nil {
		return
	}
	return
}

type MsgClientJoinChat struct {
	SteamIdChat    uint64
	IsVoiceSpeaker bool
}

func NewMsgClientJoinChat() *MsgClientJoinChat {
	return new(MsgClientJoinChat)
}

func (m *MsgClientJoinChat) GetEMsg() EMsg {
	return EMsg_ClientJoinChat
}

func (m *MsgClientJoinChat) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdChat); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.IsVoiceSpeaker); err != nil {
		return
	}
	return
}

func (m *MsgClientJoinChat) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdChat); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.IsVoiceSpeaker); err != nil {
		return
	}
	return
}

type MsgClientChatEnter struct {
	SteamIdChat   uint64
	SteamIdFriend uint64
	ChatRoomType  EChatRoomType
	SteamIdOwner  uint64
	SteamIdClan   uint64
	ChatFlags     uint8
	EnterResponse EChatRoomEnterResponse
	NumMembers    int32
}

func NewMsgClientChatEnter() *MsgClientChatEnter {
	return new(MsgClientChatEnter)
}

func (m *MsgClientChatEnter) GetEMsg() EMsg {
	return EMsg_ClientChatEnter
}

func (m *MsgClientChatEnter) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdChat); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdFriend); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.ChatRoomType); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdOwner); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdClan); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.ChatFlags); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.EnterResponse); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.NumMembers); err != nil {
		return
	}
	return
}

func (m *MsgClientChatEnter) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdChat); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdFriend); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.ChatRoomType); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdOwner); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdClan); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.ChatFlags); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.EnterResponse); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.NumMembers); err != nil {
		return
	}
	return
}

type MsgClientChatMsg struct {
	SteamIdChatter  uint64
	SteamIdChatRoom uint64
	ChatMsgType     EChatEntryType
}

func NewMsgClientChatMsg() *MsgClientChatMsg {
	return new(MsgClientChatMsg)
}

func (m *MsgClientChatMsg) GetEMsg() EMsg {
	return EMsg_ClientChatMsg
}

func (m *MsgClientChatMsg) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdChatter); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdChatRoom); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.ChatMsgType); err != nil {
		return
	}
	return
}

func (m *MsgClientChatMsg) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdChatter); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdChatRoom); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.ChatMsgType); err != nil {
		return
	}
	return
}

type MsgClientChatMemberInfo struct {
	SteamIdChat uint64
	Type        EChatInfoType
}

func NewMsgClientChatMemberInfo() *MsgClientChatMemberInfo {
	return new(MsgClientChatMemberInfo)
}

func (m *MsgClientChatMemberInfo) GetEMsg() EMsg {
	return EMsg_ClientChatMemberInfo
}

func (m *MsgClientChatMemberInfo) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdChat); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.Type); err != nil {
		return
	}
	return
}

func (m *MsgClientChatMemberInfo) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdChat); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.Type); err != nil {
		return
	}
	return
}

type MsgClientChatAction struct {
	SteamIdChat        uint64
	SteamIdUserToActOn uint64
	ChatAction         EChatAction
}

func NewMsgClientChatAction() *MsgClientChatAction {
	return new(MsgClientChatAction)
}

func (m *MsgClientChatAction) GetEMsg() EMsg {
	return EMsg_ClientChatAction
}

func (m *MsgClientChatAction) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdChat); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdUserToActOn); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.ChatAction); err != nil {
		return
	}
	return
}

func (m *MsgClientChatAction) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdChat); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdUserToActOn); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.ChatAction); err != nil {
		return
	}
	return
}

type MsgClientChatActionResult struct {
	SteamIdChat        uint64
	SteamIdUserActedOn uint64
	ChatAction         EChatAction
	ActionResult       EChatActionResult
}

func NewMsgClientChatActionResult() *MsgClientChatActionResult {
	return new(MsgClientChatActionResult)
}

func (m *MsgClientChatActionResult) GetEMsg() EMsg {
	return EMsg_ClientChatActionResult
}

func (m *MsgClientChatActionResult) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdChat); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdUserActedOn); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.ChatAction); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.ActionResult); err != nil {
		return
	}
	return
}

func (m *MsgClientChatActionResult) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdChat); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdUserActedOn); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.ChatAction); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.ActionResult); err != nil {
		return
	}
	return
}

type MsgClientChatRoomInfo struct {
	SteamIdChat uint64
	Type        EChatInfoType
}

func NewMsgClientChatRoomInfo() *MsgClientChatRoomInfo {
	return new(MsgClientChatRoomInfo)
}

func (m *MsgClientChatRoomInfo) GetEMsg() EMsg {
	return EMsg_ClientChatRoomInfo
}

func (m *MsgClientChatRoomInfo) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdChat); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.Type); err != nil {
		return
	}
	return
}

func (m *MsgClientChatRoomInfo) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdChat); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.Type); err != nil {
		return
	}
	return
}

type MsgClientSetIgnoreFriend struct {
	MySteamId     uint64
	SteamIdFriend uint64
	Ignore        uint8
}

func NewMsgClientSetIgnoreFriend() *MsgClientSetIgnoreFriend {
	return new(MsgClientSetIgnoreFriend)
}

func (m *MsgClientSetIgnoreFriend) GetEMsg() EMsg {
	return EMsg_ClientSetIgnoreFriend
}

func (m *MsgClientSetIgnoreFriend) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.MySteamId); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdFriend); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.Ignore); err != nil {
		return
	}
	return
}

func (m *MsgClientSetIgnoreFriend) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.MySteamId); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdFriend); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.Ignore); err != nil {
		return
	}
	return
}

type MsgClientSetIgnoreFriendResponse struct {
	FriendId uint64
	Result   EResult
}

func NewMsgClientSetIgnoreFriendResponse() *MsgClientSetIgnoreFriendResponse {
	return new(MsgClientSetIgnoreFriendResponse)
}

func (m *MsgClientSetIgnoreFriendResponse) GetEMsg() EMsg {
	return EMsg_ClientSetIgnoreFriendResponse
}

func (m *MsgClientSetIgnoreFriendResponse) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.FriendId); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.Result); err != nil {
		return
	}
	return
}

func (m *MsgClientSetIgnoreFriendResponse) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.FriendId); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.Result); err != nil {
		return
	}
	return
}

type MsgClientLoggedOff struct {
	Result              EResult
	SecMinReconnectHint int32
	SecMaxReconnectHint int32
}

func NewMsgClientLoggedOff() *MsgClientLoggedOff {
	return new(MsgClientLoggedOff)
}

func (m *MsgClientLoggedOff) GetEMsg() EMsg {
	return EMsg_ClientLoggedOff
}

func (m *MsgClientLoggedOff) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.Result); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SecMinReconnectHint); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SecMaxReconnectHint); err != nil {
		return
	}
	return
}

func (m *MsgClientLoggedOff) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.Result); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SecMinReconnectHint); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SecMaxReconnectHint); err != nil {
		return
	}
	return
}

type MsgClientLogOnResponse struct {
	Result                    EResult
	OutOfGameHeartbeatRateSec int32
	InGameHeartbeatRateSec    int32
	ClientSuppliedSteamId     uint64
	IpPublic                  uint32
	ServerRealTime            uint32
}

func NewMsgClientLogOnResponse() *MsgClientLogOnResponse {
	return new(MsgClientLogOnResponse)
}

func (m *MsgClientLogOnResponse) GetEMsg() EMsg {
	return EMsg_ClientLogOnResponse
}

func (m *MsgClientLogOnResponse) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.Result); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.OutOfGameHeartbeatRateSec); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.InGameHeartbeatRateSec); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.ClientSuppliedSteamId); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.IpPublic); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.ServerRealTime); err != nil {
		return
	}
	return
}

func (m *MsgClientLogOnResponse) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.Result); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.OutOfGameHeartbeatRateSec); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.InGameHeartbeatRateSec); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.ClientSuppliedSteamId); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.IpPublic); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.ServerRealTime); err != nil {
		return
	}
	return
}

type MsgClientSendGuestPass struct {
	GiftId    uint64
	GiftType  uint8
	AccountId uint32
}

func NewMsgClientSendGuestPass() *MsgClientSendGuestPass {
	return new(MsgClientSendGuestPass)
}

func (m *MsgClientSendGuestPass) GetEMsg() EMsg {
	return EMsg_ClientSendGuestPass
}

func (m *MsgClientSendGuestPass) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.GiftId); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.GiftType); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.AccountId); err != nil {
		return
	}
	return
}

func (m *MsgClientSendGuestPass) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.GiftId); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.GiftType); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.AccountId); err != nil {
		return
	}
	return
}

type MsgClientSendGuestPassResponse struct {
	Result EResult
}

func NewMsgClientSendGuestPassResponse() *MsgClientSendGuestPassResponse {
	return new(MsgClientSendGuestPassResponse)
}

func (m *MsgClientSendGuestPassResponse) GetEMsg() EMsg {
	return EMsg_ClientSendGuestPassResponse
}

func (m *MsgClientSendGuestPassResponse) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.Result); err != nil {
		return
	}
	return
}

func (m *MsgClientSendGuestPassResponse) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.Result); err != nil {
		return
	}
	return
}

type MsgClientServerUnavailable struct {
	JobidSent              uint64
	EMsgSent               uint32
	EServerTypeUnavailable EServerType
}

func NewMsgClientServerUnavailable() *MsgClientServerUnavailable {
	return new(MsgClientServerUnavailable)
}

func (m *MsgClientServerUnavailable) GetEMsg() EMsg {
	return EMsg_ClientServerUnavailable
}

func (m *MsgClientServerUnavailable) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.JobidSent); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.EMsgSent); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.EServerTypeUnavailable); err != nil {
		return
	}
	return
}

func (m *MsgClientServerUnavailable) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.JobidSent); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.EMsgSent); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.EServerTypeUnavailable); err != nil {
		return
	}
	return
}

type MsgClientCreateChat struct {
	ChatRoomType      EChatRoomType
	GameId            uint64
	SteamIdClan       uint64
	PermissionOfficer EChatPermission
	PermissionMember  EChatPermission
	PermissionAll     EChatPermission
	MembersMax        uint32
	ChatFlags         uint8
	SteamIdFriendChat uint64
	SteamIdInvited    uint64
}

func NewMsgClientCreateChat() *MsgClientCreateChat {
	return new(MsgClientCreateChat)
}

func (m *MsgClientCreateChat) GetEMsg() EMsg {
	return EMsg_ClientCreateChat
}

func (m *MsgClientCreateChat) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.ChatRoomType); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.GameId); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdClan); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.PermissionOfficer); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.PermissionMember); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.PermissionAll); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.MembersMax); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.ChatFlags); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdFriendChat); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdInvited); err != nil {
		return
	}
	return
}

func (m *MsgClientCreateChat) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.ChatRoomType); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.GameId); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdClan); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.PermissionOfficer); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.PermissionMember); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.PermissionAll); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.MembersMax); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.ChatFlags); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdFriendChat); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdInvited); err != nil {
		return
	}
	return
}

type MsgClientCreateChatResponse struct {
	Result            EResult
	SteamIdChat       uint64
	ChatRoomType      EChatRoomType
	SteamIdFriendChat uint64
}

func NewMsgClientCreateChatResponse() *MsgClientCreateChatResponse {
	return new(MsgClientCreateChatResponse)
}

func (m *MsgClientCreateChatResponse) GetEMsg() EMsg {
	return EMsg_ClientCreateChatResponse
}

func (m *MsgClientCreateChatResponse) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.Result); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdChat); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.ChatRoomType); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.SteamIdFriendChat); err != nil {
		return
	}
	return
}

func (m *MsgClientCreateChatResponse) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.Result); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdChat); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.ChatRoomType); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.SteamIdFriendChat); err != nil {
		return
	}
	return
}

type MsgClientMarketingMessageUpdate2 struct {
	MarketingMessageUpdateTime uint32
	Count                      uint32
}

func NewMsgClientMarketingMessageUpdate2() *MsgClientMarketingMessageUpdate2 {
	return new(MsgClientMarketingMessageUpdate2)
}

func (m *MsgClientMarketingMessageUpdate2) GetEMsg() EMsg {
	return EMsg_ClientMarketingMessageUpdate2
}

func (m *MsgClientMarketingMessageUpdate2) Serialize(w io.Writer) (err error) {
	if err = binary.Write(w, binary.LittleEndian, m.MarketingMessageUpdateTime); err != nil {
		return
	}
	if err = binary.Write(w, binary.LittleEndian, m.Count); err != nil {
		return
	}
	return
}

func (m *MsgClientMarketingMessageUpdate2) Deserialize(r io.Reader) (err error) {
	if err = binary.Read(r, binary.LittleEndian, &m.MarketingMessageUpdateTime); err != nil {
		return
	}
	if err = binary.Read(r, binary.LittleEndian, &m.Count); err != nil {
		return
	}
	return
}
