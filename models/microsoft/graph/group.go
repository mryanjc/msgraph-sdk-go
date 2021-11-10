package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Group struct {
    DirectoryObject
    // The list of users or groups that are allowed to create post's or calendar events in this group. If this list is non-empty then only users or groups listed here are allowed to post.
    acceptedSenders []DirectoryObject;
    // Indicates if people external to the organization can send messages to the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    allowExternalSenders *bool;
    // Represents the app roles a group has been granted for an application. Supports $expand.
    appRoleAssignments []AppRoleAssignment;
    // The list of sensitivity label pairs (label ID, label name) associated with a Microsoft 365 group. Returned only on $select. Read-only.
    assignedLabels []AssignedLabel;
    // The licenses that are assigned to the group. Returned only on $select. Supports $filter (eq).Read-only.
    assignedLicenses []AssignedLicense;
    // Indicates if new members added to the group will be auto-subscribed to receive email notifications. You can set this property in a PATCH request for the group; do not set it in the initial POST request that creates the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    autoSubscribeNewMembers *bool;
    // The group's calendar. Read-only.
    calendar *Calendar;
    // The calendar view for the calendar. Read-only.
    calendarView []Event;
    // Describes a classification for the group (such as low, medium or high business impact). Valid values for this property are defined by creating a ClassificationList setting value, based on the template definition.Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith).
    classification *string;
    // The group's conversations.
    conversations []Conversation;
    // Timestamp of when the group was created. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The user (or application) that created the group. NOTE: This is not set if the user is an administrator. Read-only.
    createdOnBehalfOf *DirectoryObject;
    // An optional description for the group. Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith) and $search.
    description *string;
    // The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith), $search, and $orderBy.
    displayName *string;
    // The group's default drive. Read-only.
    drive *Drive;
    // The group's drives. Read-only.
    drives []Drive;
    // The group's calendar events.
    events []Event;
    // Timestamp of when the group is set to expire. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The collection of open extensions defined for the group. Read-only. Nullable.
    extensions []Extension;
    // The collection of lifecycle policies for this group. Read-only. Nullable.
    groupLifecyclePolicies []GroupLifecyclePolicy;
    // Specifies the group type and its membership.  If the collection contains Unified, the group is a Microsoft 365 group; otherwise, it's either a security group or distribution group. For details, see groups overview.If the collection includes DynamicMembership, the group has dynamic membership; otherwise, membership is static.  Returned by default. Supports $filter (eq, NOT).
    groupTypes []string;
    // Indicates whether there are members in this group that have license errors from its group-based license assignment. This property is never returned on a GET operation. You can use it as a $filter argument to get groups that have members with license errors (that is, filter for this property being true). See an example. Supports $filter (eq).
    hasMembersWithLicenseErrors *bool;
    // True if the group is not displayed in certain parts of the Outlook UI: the Address Book, address lists for selecting message recipients, and the Browse Groups dialog for searching groups; otherwise, false. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    hideFromAddressLists *bool;
    // True if the group is not displayed in Outlook clients, such as Outlook for Windows and Outlook on the web; otherwise, false. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    hideFromOutlookClients *bool;
    // 
    isArchived *bool;
    // Indicates whether this group can be assigned to an Azure Active Directory role or not. Optional. This property can only be set while creating the group and is immutable. If set to true, the securityEnabled property must also be set to true and the group cannot be a dynamic group (that is, groupTypes cannot contain DynamicMembership). Only callers in Global administrator and Privileged role administrator roles can set this property. The caller must be assigned the RoleManagement.ReadWrite.Directory permission to set this property or update the membership of such groups. For more, see Using a group to manage Azure AD role assignmentsReturned by default. Supports $filter (eq, ne, NOT).
    isAssignableToRole *bool;
    // Indicates whether the signed-in user is subscribed to receive email conversations. Default value is true. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    isSubscribedByMail *bool;
    // Indicates status of the group license assignment to all members of the group. Default value is false. Read-only. Possible values: QueuedForProcessing, ProcessingInProgress, and ProcessingComplete.Returned only on $select. Read-only.
    licenseProcessingState *LicenseProcessingState;
    // The SMTP address for the group, for example, 'serviceadmins@contoso.onmicrosoft.com'. Returned by default. Read-only. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    mail *string;
    // Specifies whether the group is mail-enabled. Required. Returned by default. Supports $filter (eq, ne, NOT).
    mailEnabled *bool;
    // The mail alias for the group, unique in the organization. Maximum length is 64 characters. This property can contain only characters in the ASCII character set 0 - 127 except the following: @ () / [] ' ; : . <> , SPACE. Required. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    mailNickname *string;
    // Groups that this group is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObject;
    // Users and groups that are members of this group. HTTP Methods: GET (supported for all groups), POST (supported for Microsoft 365 groups, security groups and mail-enabled security groups), DELETE (supported for Microsoft 365 groups and security groups). Nullable. Supports $expand.
    members []DirectoryObject;
    // The rule that determines members for this group if the group is a dynamic group (groupTypes contains DynamicMembership). For more information about the syntax of the membership rule, see Membership Rules syntax. Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith).
    membershipRule *string;
    // Indicates whether the dynamic membership processing is on or paused. Possible values are On or Paused. Returned by default. Supports $filter (eq, ne, NOT, in).
    membershipRuleProcessingState *string;
    // A list of group members with license errors from this group-based license assignment. Read-only.
    membersWithLicenseErrors []DirectoryObject;
    // Read-only.
    onenote *Onenote;
    // Contains the on-premises domain FQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only.
    onPremisesDomainName *string;
    // Indicates the last time at which the group was synced with the on-premises directory.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only. Supports $filter (eq, ne, NOT, ge, le, in).
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Contains the on-premises netBios name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only.
    onPremisesNetBiosName *string;
    // Errors when using Microsoft synchronization product during provisioning. Returned by default. Supports $filter (eq, NOT).
    onPremisesProvisioningErrors []OnPremisesProvisioningError;
    // Contains the on-premises SAM account name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith). Read-only.
    onPremisesSamAccountName *string;
    // Contains the on-premises security identifier (SID) for the group that was synchronized from on-premises to the cloud. Returned by default. Supports $filter on null values. Read-only.
    onPremisesSecurityIdentifier *string;
    // true if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Returned by default. Read-only. Supports $filter (eq, ne, NOT, in).
    onPremisesSyncEnabled *bool;
    // The owners of the group. The owners are a set of non-admin users who are allowed to modify this object. Limited to 100 owners. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand.
    owners []DirectoryObject;
    // The permission that has been granted for a group to a specific application. Supports $expand.
    permissionGrants []ResourceSpecificPermissionGrant;
    // The group's profile photo
    photo *ProfilePhoto;
    // The profile photos owned by the group. Read-only. Nullable.
    photos []ProfilePhoto;
    // Entry-point to Planner resource that might exist for a Unified Group.
    planner *PlannerGroup;
    // The preferred data location for the Microsoft 365 group. By default, the group inherits the group creator's preferred data location. To set this property, the calling user must be assigned one of the following Azure AD roles:  Global Administrator  User Account Administrator Directory Writer  Exchange Administrator  SharePoint Administrator  For more information about this property, see  OneDrive Online Multi-Geo. Nullable. Returned by default.
    preferredDataLocation *string;
    // The preferred language for a Microsoft 365 group. Should follow ISO 639-1 Code; for example en-US. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
    preferredLanguage *string;
    // Email addresses for the group that direct to the same group mailbox. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. The any operator is required to filter expressions on multi-valued properties. Returned by default. Read-only. Not nullable. Supports $filter (eq, NOT, ge, le, startsWith).
    proxyAddresses []string;
    // The list of users or groups that are not allowed to create posts or calendar events in this group. Nullable
    rejectedSenders []DirectoryObject;
    // Timestamp of when the group was last renewed. This cannot be modified directly and is only updated via the renew service action. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
    renewedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Specifies whether the group is a security group. Required. Returned by default. Supports $filter (eq, ne, NOT, in).
    securityEnabled *bool;
    // Security identifier of the group, used in Windows scenarios. Returned by default.
    securityIdentifier *string;
    // Read-only. Nullable.
    settings []GroupSetting;
    // The list of SharePoint sites in this group. Access the default site with /sites/root.
    sites []Site;
    // 
    team *Team;
    // Specifies a Microsoft 365 group's color theme. Possible values are Teal, Purple, Green, Blue, Pink, Orange or Red. Returned by default.
    theme *string;
    // The group's conversation threads. Nullable.
    threads []ConversationThread;
    // 
    transitiveMemberOf []DirectoryObject;
    // 
    transitiveMembers []DirectoryObject;
    // Count of conversations that have received new posts since the signed-in user last visited the group. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    unseenCount *int32;
    // Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or Hiddenmembership. Hiddenmembership can be set only for Microsoft 365 groups, when the groups are created. It can't be updated later. Other values of visibility can be updated after group creation. If visibility value is not specified during group creation on Microsoft Graph, a security group is created as Private by default and Microsoft 365 group is Public. Groups assignable to roles are always Private. See group visibility options to learn more. Returned by default. Nullable.
    visibility *string;
}
// Instantiates a new group and sets the default values.
func NewGroup()(*Group) {
    m := &Group{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// Gets the acceptedSenders property value. The list of users or groups that are allowed to create post's or calendar events in this group. If this list is non-empty then only users or groups listed here are allowed to post.
func (m *Group) GetAcceptedSenders()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.acceptedSenders
    }
}
// Gets the allowExternalSenders property value. Indicates if people external to the organization can send messages to the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) GetAllowExternalSenders()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowExternalSenders
    }
}
// Gets the appRoleAssignments property value. Represents the app roles a group has been granted for an application. Supports $expand.
func (m *Group) GetAppRoleAssignments()([]AppRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.appRoleAssignments
    }
}
// Gets the assignedLabels property value. The list of sensitivity label pairs (label ID, label name) associated with a Microsoft 365 group. Returned only on $select. Read-only.
func (m *Group) GetAssignedLabels()([]AssignedLabel) {
    if m == nil {
        return nil
    } else {
        return m.assignedLabels
    }
}
// Gets the assignedLicenses property value. The licenses that are assigned to the group. Returned only on $select. Supports $filter (eq).Read-only.
func (m *Group) GetAssignedLicenses()([]AssignedLicense) {
    if m == nil {
        return nil
    } else {
        return m.assignedLicenses
    }
}
// Gets the autoSubscribeNewMembers property value. Indicates if new members added to the group will be auto-subscribed to receive email notifications. You can set this property in a PATCH request for the group; do not set it in the initial POST request that creates the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) GetAutoSubscribeNewMembers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoSubscribeNewMembers
    }
}
// Gets the calendar property value. The group's calendar. Read-only.
func (m *Group) GetCalendar()(*Calendar) {
    if m == nil {
        return nil
    } else {
        return m.calendar
    }
}
// Gets the calendarView property value. The calendar view for the calendar. Read-only.
func (m *Group) GetCalendarView()([]Event) {
    if m == nil {
        return nil
    } else {
        return m.calendarView
    }
}
// Gets the classification property value. Describes a classification for the group (such as low, medium or high business impact). Valid values for this property are defined by creating a ClassificationList setting value, based on the template definition.Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith).
func (m *Group) GetClassification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
// Gets the conversations property value. The group's conversations.
func (m *Group) GetConversations()([]Conversation) {
    if m == nil {
        return nil
    } else {
        return m.conversations
    }
}
// Gets the createdDateTime property value. Timestamp of when the group was created. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
func (m *Group) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the createdOnBehalfOf property value. The user (or application) that created the group. NOTE: This is not set if the user is an administrator. Read-only.
func (m *Group) GetCreatedOnBehalfOf()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.createdOnBehalfOf
    }
}
// Gets the description property value. An optional description for the group. Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith) and $search.
func (m *Group) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith), $search, and $orderBy.
func (m *Group) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the drive property value. The group's default drive. Read-only.
func (m *Group) GetDrive()(*Drive) {
    if m == nil {
        return nil
    } else {
        return m.drive
    }
}
// Gets the drives property value. The group's drives. Read-only.
func (m *Group) GetDrives()([]Drive) {
    if m == nil {
        return nil
    } else {
        return m.drives
    }
}
// Gets the events property value. The group's calendar events.
func (m *Group) GetEvents()([]Event) {
    if m == nil {
        return nil
    } else {
        return m.events
    }
}
// Gets the expirationDateTime property value. Timestamp of when the group is set to expire. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
func (m *Group) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the extensions property value. The collection of open extensions defined for the group. Read-only. Nullable.
func (m *Group) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// Gets the groupLifecyclePolicies property value. The collection of lifecycle policies for this group. Read-only. Nullable.
func (m *Group) GetGroupLifecyclePolicies()([]GroupLifecyclePolicy) {
    if m == nil {
        return nil
    } else {
        return m.groupLifecyclePolicies
    }
}
// Gets the groupTypes property value. Specifies the group type and its membership.  If the collection contains Unified, the group is a Microsoft 365 group; otherwise, it's either a security group or distribution group. For details, see groups overview.If the collection includes DynamicMembership, the group has dynamic membership; otherwise, membership is static.  Returned by default. Supports $filter (eq, NOT).
func (m *Group) GetGroupTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.groupTypes
    }
}
// Gets the hasMembersWithLicenseErrors property value. Indicates whether there are members in this group that have license errors from its group-based license assignment. This property is never returned on a GET operation. You can use it as a $filter argument to get groups that have members with license errors (that is, filter for this property being true). See an example. Supports $filter (eq).
func (m *Group) GetHasMembersWithLicenseErrors()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasMembersWithLicenseErrors
    }
}
// Gets the hideFromAddressLists property value. True if the group is not displayed in certain parts of the Outlook UI: the Address Book, address lists for selecting message recipients, and the Browse Groups dialog for searching groups; otherwise, false. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) GetHideFromAddressLists()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideFromAddressLists
    }
}
// Gets the hideFromOutlookClients property value. True if the group is not displayed in Outlook clients, such as Outlook for Windows and Outlook on the web; otherwise, false. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) GetHideFromOutlookClients()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideFromOutlookClients
    }
}
// Gets the isArchived property value. 
func (m *Group) GetIsArchived()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isArchived
    }
}
// Gets the isAssignableToRole property value. Indicates whether this group can be assigned to an Azure Active Directory role or not. Optional. This property can only be set while creating the group and is immutable. If set to true, the securityEnabled property must also be set to true and the group cannot be a dynamic group (that is, groupTypes cannot contain DynamicMembership). Only callers in Global administrator and Privileged role administrator roles can set this property. The caller must be assigned the RoleManagement.ReadWrite.Directory permission to set this property or update the membership of such groups. For more, see Using a group to manage Azure AD role assignmentsReturned by default. Supports $filter (eq, ne, NOT).
func (m *Group) GetIsAssignableToRole()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAssignableToRole
    }
}
// Gets the isSubscribedByMail property value. Indicates whether the signed-in user is subscribed to receive email conversations. Default value is true. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) GetIsSubscribedByMail()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSubscribedByMail
    }
}
// Gets the licenseProcessingState property value. Indicates status of the group license assignment to all members of the group. Default value is false. Read-only. Possible values: QueuedForProcessing, ProcessingInProgress, and ProcessingComplete.Returned only on $select. Read-only.
func (m *Group) GetLicenseProcessingState()(*LicenseProcessingState) {
    if m == nil {
        return nil
    } else {
        return m.licenseProcessingState
    }
}
// Gets the mail property value. The SMTP address for the group, for example, 'serviceadmins@contoso.onmicrosoft.com'. Returned by default. Read-only. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *Group) GetMail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mail
    }
}
// Gets the mailEnabled property value. Specifies whether the group is mail-enabled. Required. Returned by default. Supports $filter (eq, ne, NOT).
func (m *Group) GetMailEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.mailEnabled
    }
}
// Gets the mailNickname property value. The mail alias for the group, unique in the organization. Maximum length is 64 characters. This property can contain only characters in the ASCII character set 0 - 127 except the following: @ () / [] ' ; : . <> , SPACE. Required. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *Group) GetMailNickname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mailNickname
    }
}
// Gets the memberOf property value. Groups that this group is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports $expand.
func (m *Group) GetMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.memberOf
    }
}
// Gets the members property value. Users and groups that are members of this group. HTTP Methods: GET (supported for all groups), POST (supported for Microsoft 365 groups, security groups and mail-enabled security groups), DELETE (supported for Microsoft 365 groups and security groups). Nullable. Supports $expand.
func (m *Group) GetMembers()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// Gets the membershipRule property value. The rule that determines members for this group if the group is a dynamic group (groupTypes contains DynamicMembership). For more information about the syntax of the membership rule, see Membership Rules syntax. Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith).
func (m *Group) GetMembershipRule()(*string) {
    if m == nil {
        return nil
    } else {
        return m.membershipRule
    }
}
// Gets the membershipRuleProcessingState property value. Indicates whether the dynamic membership processing is on or paused. Possible values are On or Paused. Returned by default. Supports $filter (eq, ne, NOT, in).
func (m *Group) GetMembershipRuleProcessingState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.membershipRuleProcessingState
    }
}
// Gets the membersWithLicenseErrors property value. A list of group members with license errors from this group-based license assignment. Read-only.
func (m *Group) GetMembersWithLicenseErrors()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.membersWithLicenseErrors
    }
}
// Gets the onenote property value. Read-only.
func (m *Group) GetOnenote()(*Onenote) {
    if m == nil {
        return nil
    } else {
        return m.onenote
    }
}
// Gets the onPremisesDomainName property value. Contains the on-premises domain FQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only.
func (m *Group) GetOnPremisesDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesDomainName
    }
}
// Gets the onPremisesLastSyncDateTime property value. Indicates the last time at which the group was synced with the on-premises directory.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only. Supports $filter (eq, ne, NOT, ge, le, in).
func (m *Group) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesLastSyncDateTime
    }
}
// Gets the onPremisesNetBiosName property value. Contains the on-premises netBios name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only.
func (m *Group) GetOnPremisesNetBiosName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesNetBiosName
    }
}
// Gets the onPremisesProvisioningErrors property value. Errors when using Microsoft synchronization product during provisioning. Returned by default. Supports $filter (eq, NOT).
func (m *Group) GetOnPremisesProvisioningErrors()([]OnPremisesProvisioningError) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesProvisioningErrors
    }
}
// Gets the onPremisesSamAccountName property value. Contains the on-premises SAM account name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith). Read-only.
func (m *Group) GetOnPremisesSamAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSamAccountName
    }
}
// Gets the onPremisesSecurityIdentifier property value. Contains the on-premises security identifier (SID) for the group that was synchronized from on-premises to the cloud. Returned by default. Supports $filter on null values. Read-only.
func (m *Group) GetOnPremisesSecurityIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSecurityIdentifier
    }
}
// Gets the onPremisesSyncEnabled property value. true if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Returned by default. Read-only. Supports $filter (eq, ne, NOT, in).
func (m *Group) GetOnPremisesSyncEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSyncEnabled
    }
}
// Gets the owners property value. The owners of the group. The owners are a set of non-admin users who are allowed to modify this object. Limited to 100 owners. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand.
func (m *Group) GetOwners()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.owners
    }
}
// Gets the permissionGrants property value. The permission that has been granted for a group to a specific application. Supports $expand.
func (m *Group) GetPermissionGrants()([]ResourceSpecificPermissionGrant) {
    if m == nil {
        return nil
    } else {
        return m.permissionGrants
    }
}
// Gets the photo property value. The group's profile photo
func (m *Group) GetPhoto()(*ProfilePhoto) {
    if m == nil {
        return nil
    } else {
        return m.photo
    }
}
// Gets the photos property value. The profile photos owned by the group. Read-only. Nullable.
func (m *Group) GetPhotos()([]ProfilePhoto) {
    if m == nil {
        return nil
    } else {
        return m.photos
    }
}
// Gets the planner property value. Entry-point to Planner resource that might exist for a Unified Group.
func (m *Group) GetPlanner()(*PlannerGroup) {
    if m == nil {
        return nil
    } else {
        return m.planner
    }
}
// Gets the preferredDataLocation property value. The preferred data location for the Microsoft 365 group. By default, the group inherits the group creator's preferred data location. To set this property, the calling user must be assigned one of the following Azure AD roles:  Global Administrator  User Account Administrator Directory Writer  Exchange Administrator  SharePoint Administrator  For more information about this property, see  OneDrive Online Multi-Geo. Nullable. Returned by default.
func (m *Group) GetPreferredDataLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredDataLocation
    }
}
// Gets the preferredLanguage property value. The preferred language for a Microsoft 365 group. Should follow ISO 639-1 Code; for example en-US. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
func (m *Group) GetPreferredLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredLanguage
    }
}
// Gets the proxyAddresses property value. Email addresses for the group that direct to the same group mailbox. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. The any operator is required to filter expressions on multi-valued properties. Returned by default. Read-only. Not nullable. Supports $filter (eq, NOT, ge, le, startsWith).
func (m *Group) GetProxyAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.proxyAddresses
    }
}
// Gets the rejectedSenders property value. The list of users or groups that are not allowed to create posts or calendar events in this group. Nullable
func (m *Group) GetRejectedSenders()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.rejectedSenders
    }
}
// Gets the renewedDateTime property value. Timestamp of when the group was last renewed. This cannot be modified directly and is only updated via the renew service action. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
func (m *Group) GetRenewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.renewedDateTime
    }
}
// Gets the securityEnabled property value. Specifies whether the group is a security group. Required. Returned by default. Supports $filter (eq, ne, NOT, in).
func (m *Group) GetSecurityEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityEnabled
    }
}
// Gets the securityIdentifier property value. Security identifier of the group, used in Windows scenarios. Returned by default.
func (m *Group) GetSecurityIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.securityIdentifier
    }
}
// Gets the settings property value. Read-only. Nullable.
func (m *Group) GetSettings()([]GroupSetting) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// Gets the sites property value. The list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *Group) GetSites()([]Site) {
    if m == nil {
        return nil
    } else {
        return m.sites
    }
}
// Gets the team property value. 
func (m *Group) GetTeam()(*Team) {
    if m == nil {
        return nil
    } else {
        return m.team
    }
}
// Gets the theme property value. Specifies a Microsoft 365 group's color theme. Possible values are Teal, Purple, Green, Blue, Pink, Orange or Red. Returned by default.
func (m *Group) GetTheme()(*string) {
    if m == nil {
        return nil
    } else {
        return m.theme
    }
}
// Gets the threads property value. The group's conversation threads. Nullable.
func (m *Group) GetThreads()([]ConversationThread) {
    if m == nil {
        return nil
    } else {
        return m.threads
    }
}
// Gets the transitiveMemberOf property value. 
func (m *Group) GetTransitiveMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.transitiveMemberOf
    }
}
// Gets the transitiveMembers property value. 
func (m *Group) GetTransitiveMembers()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.transitiveMembers
    }
}
// Gets the unseenCount property value. Count of conversations that have received new posts since the signed-in user last visited the group. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) GetUnseenCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unseenCount
    }
}
// Gets the visibility property value. Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or Hiddenmembership. Hiddenmembership can be set only for Microsoft 365 groups, when the groups are created. It can't be updated later. Other values of visibility can be updated after group creation. If visibility value is not specified during group creation on Microsoft Graph, a security group is created as Private by default and Microsoft 365 group is Public. Groups assignable to roles are always Private. See group visibility options to learn more. Returned by default. Nullable.
func (m *Group) GetVisibility()(*string) {
    if m == nil {
        return nil
    } else {
        return m.visibility
    }
}
// The deserialization information for the current model
func (m *Group) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["acceptedSenders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetAcceptedSenders(res)
        }
        return nil
    }
    res["allowExternalSenders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowExternalSenders(val)
        }
        return nil
    }
    res["appRoleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppRoleAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppRoleAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*AppRoleAssignment))
            }
            m.SetAppRoleAssignments(res)
        }
        return nil
    }
    res["assignedLabels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignedLabel() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedLabel, len(val))
            for i, v := range val {
                res[i] = *(v.(*AssignedLabel))
            }
            m.SetAssignedLabels(res)
        }
        return nil
    }
    res["assignedLicenses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignedLicense() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedLicense, len(val))
            for i, v := range val {
                res[i] = *(v.(*AssignedLicense))
            }
            m.SetAssignedLicenses(res)
        }
        return nil
    }
    res["autoSubscribeNewMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoSubscribeNewMembers(val)
        }
        return nil
    }
    res["calendar"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCalendar() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalendar(val.(*Calendar))
        }
        return nil
    }
    res["calendarView"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Event, len(val))
            for i, v := range val {
                res[i] = *(v.(*Event))
            }
            m.SetCalendarView(res)
        }
        return nil
    }
    res["classification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassification(val)
        }
        return nil
    }
    res["conversations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConversation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Conversation, len(val))
            for i, v := range val {
                res[i] = *(v.(*Conversation))
            }
            m.SetConversations(res)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["createdOnBehalfOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedOnBehalfOf(val.(*DirectoryObject))
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["drive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDrive() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDrive(val.(*Drive))
        }
        return nil
    }
    res["drives"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDrive() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Drive, len(val))
            for i, v := range val {
                res[i] = *(v.(*Drive))
            }
            m.SetDrives(res)
        }
        return nil
    }
    res["events"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Event, len(val))
            for i, v := range val {
                res[i] = *(v.(*Event))
            }
            m.SetEvents(res)
        }
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtension() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extension, len(val))
            for i, v := range val {
                res[i] = *(v.(*Extension))
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["groupLifecyclePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupLifecyclePolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupLifecyclePolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*GroupLifecyclePolicy))
            }
            m.SetGroupLifecyclePolicies(res)
        }
        return nil
    }
    res["groupTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGroupTypes(res)
        }
        return nil
    }
    res["hasMembersWithLicenseErrors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasMembersWithLicenseErrors(val)
        }
        return nil
    }
    res["hideFromAddressLists"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideFromAddressLists(val)
        }
        return nil
    }
    res["hideFromOutlookClients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideFromOutlookClients(val)
        }
        return nil
    }
    res["isArchived"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsArchived(val)
        }
        return nil
    }
    res["isAssignableToRole"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAssignableToRole(val)
        }
        return nil
    }
    res["isSubscribedByMail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSubscribedByMail(val)
        }
        return nil
    }
    res["licenseProcessingState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLicenseProcessingState() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseProcessingState(val.(*LicenseProcessingState))
        }
        return nil
    }
    res["mail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMail(val)
        }
        return nil
    }
    res["mailEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailEnabled(val)
        }
        return nil
    }
    res["mailNickname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailNickname(val)
        }
        return nil
    }
    res["memberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetMemberOf(res)
        }
        return nil
    }
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["membershipRule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipRule(val)
        }
        return nil
    }
    res["membershipRuleProcessingState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipRuleProcessingState(val)
        }
        return nil
    }
    res["membersWithLicenseErrors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetMembersWithLicenseErrors(res)
        }
        return nil
    }
    res["onenote"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnenote() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnenote(val.(*Onenote))
        }
        return nil
    }
    res["onPremisesDomainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesDomainName(val)
        }
        return nil
    }
    res["onPremisesLastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesLastSyncDateTime(val)
        }
        return nil
    }
    res["onPremisesNetBiosName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesNetBiosName(val)
        }
        return nil
    }
    res["onPremisesProvisioningErrors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesProvisioningError() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesProvisioningError, len(val))
            for i, v := range val {
                res[i] = *(v.(*OnPremisesProvisioningError))
            }
            m.SetOnPremisesProvisioningErrors(res)
        }
        return nil
    }
    res["onPremisesSamAccountName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSamAccountName(val)
        }
        return nil
    }
    res["onPremisesSecurityIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSecurityIdentifier(val)
        }
        return nil
    }
    res["onPremisesSyncEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSyncEnabled(val)
        }
        return nil
    }
    res["owners"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetOwners(res)
        }
        return nil
    }
    res["permissionGrants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResourceSpecificPermissionGrant() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResourceSpecificPermissionGrant, len(val))
            for i, v := range val {
                res[i] = *(v.(*ResourceSpecificPermissionGrant))
            }
            m.SetPermissionGrants(res)
        }
        return nil
    }
    res["photo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProfilePhoto() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoto(val.(*ProfilePhoto))
        }
        return nil
    }
    res["photos"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProfilePhoto() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProfilePhoto, len(val))
            for i, v := range val {
                res[i] = *(v.(*ProfilePhoto))
            }
            m.SetPhotos(res)
        }
        return nil
    }
    res["planner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlanner(val.(*PlannerGroup))
        }
        return nil
    }
    res["preferredDataLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredDataLocation(val)
        }
        return nil
    }
    res["preferredLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredLanguage(val)
        }
        return nil
    }
    res["proxyAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetProxyAddresses(res)
        }
        return nil
    }
    res["rejectedSenders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetRejectedSenders(res)
        }
        return nil
    }
    res["renewedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenewedDateTime(val)
        }
        return nil
    }
    res["securityEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityEnabled(val)
        }
        return nil
    }
    res["securityIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityIdentifier(val)
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupSetting, len(val))
            for i, v := range val {
                res[i] = *(v.(*GroupSetting))
            }
            m.SetSettings(res)
        }
        return nil
    }
    res["sites"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSite() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Site, len(val))
            for i, v := range val {
                res[i] = *(v.(*Site))
            }
            m.SetSites(res)
        }
        return nil
    }
    res["team"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeam() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeam(val.(*Team))
        }
        return nil
    }
    res["theme"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTheme(val)
        }
        return nil
    }
    res["threads"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConversationThread() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConversationThread, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConversationThread))
            }
            m.SetThreads(res)
        }
        return nil
    }
    res["transitiveMemberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetTransitiveMemberOf(res)
        }
        return nil
    }
    res["transitiveMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetTransitiveMembers(res)
        }
        return nil
    }
    res["unseenCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnseenCount(val)
        }
        return nil
    }
    res["visibility"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibility(val)
        }
        return nil
    }
    return res
}
func (m *Group) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Group) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAcceptedSenders()))
        for i, v := range m.GetAcceptedSenders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("acceptedSenders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowExternalSenders", m.GetAllowExternalSenders())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppRoleAssignments()))
        for i, v := range m.GetAppRoleAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appRoleAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignedLabels()))
        for i, v := range m.GetAssignedLabels() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignedLabels", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignedLicenses()))
        for i, v := range m.GetAssignedLicenses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignedLicenses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("autoSubscribeNewMembers", m.GetAutoSubscribeNewMembers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("calendar", m.GetCalendar())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCalendarView()))
        for i, v := range m.GetCalendarView() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("calendarView", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("classification", m.GetClassification())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConversations()))
        for i, v := range m.GetConversations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("conversations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdOnBehalfOf", m.GetCreatedOnBehalfOf())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("drive", m.GetDrive())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDrives()))
        for i, v := range m.GetDrives() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("drives", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEvents()))
        for i, v := range m.GetEvents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("events", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupLifecyclePolicies()))
        for i, v := range m.GetGroupLifecyclePolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groupLifecyclePolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("groupTypes", m.GetGroupTypes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasMembersWithLicenseErrors", m.GetHasMembersWithLicenseErrors())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hideFromAddressLists", m.GetHideFromAddressLists())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hideFromOutlookClients", m.GetHideFromOutlookClients())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isArchived", m.GetIsArchived())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAssignableToRole", m.GetIsAssignableToRole())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSubscribedByMail", m.GetIsSubscribedByMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("licenseProcessingState", m.GetLicenseProcessingState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mail", m.GetMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("mailEnabled", m.GetMailEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mailNickname", m.GetMailNickname())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMemberOf()))
        for i, v := range m.GetMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("memberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("membershipRule", m.GetMembershipRule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("membershipRuleProcessingState", m.GetMembershipRuleProcessingState())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembersWithLicenseErrors()))
        for i, v := range m.GetMembersWithLicenseErrors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("membersWithLicenseErrors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onenote", m.GetOnenote())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesDomainName", m.GetOnPremisesDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("onPremisesLastSyncDateTime", m.GetOnPremisesLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesNetBiosName", m.GetOnPremisesNetBiosName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOnPremisesProvisioningErrors()))
        for i, v := range m.GetOnPremisesProvisioningErrors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("onPremisesProvisioningErrors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesSamAccountName", m.GetOnPremisesSamAccountName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesSecurityIdentifier", m.GetOnPremisesSecurityIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onPremisesSyncEnabled", m.GetOnPremisesSyncEnabled())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOwners()))
        for i, v := range m.GetOwners() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("owners", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPermissionGrants()))
        for i, v := range m.GetPermissionGrants() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("permissionGrants", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("photo", m.GetPhoto())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPhotos()))
        for i, v := range m.GetPhotos() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("photos", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("planner", m.GetPlanner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredDataLocation", m.GetPreferredDataLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredLanguage", m.GetPreferredLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("proxyAddresses", m.GetProxyAddresses())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRejectedSenders()))
        for i, v := range m.GetRejectedSenders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("rejectedSenders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("renewedDateTime", m.GetRenewedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityEnabled", m.GetSecurityEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("securityIdentifier", m.GetSecurityIdentifier())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("settings", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSites()))
        for i, v := range m.GetSites() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sites", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("team", m.GetTeam())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("theme", m.GetTheme())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetThreads()))
        for i, v := range m.GetThreads() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("threads", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTransitiveMemberOf()))
        for i, v := range m.GetTransitiveMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("transitiveMemberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTransitiveMembers()))
        for i, v := range m.GetTransitiveMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("transitiveMembers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unseenCount", m.GetUnseenCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("visibility", m.GetVisibility())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the acceptedSenders property value. The list of users or groups that are allowed to create post's or calendar events in this group. If this list is non-empty then only users or groups listed here are allowed to post.
// Parameters:
//  - value : Value to set for the acceptedSenders property.
func (m *Group) SetAcceptedSenders(value []DirectoryObject)() {
    m.acceptedSenders = value
}
// Sets the allowExternalSenders property value. Indicates if people external to the organization can send messages to the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
// Parameters:
//  - value : Value to set for the allowExternalSenders property.
func (m *Group) SetAllowExternalSenders(value *bool)() {
    m.allowExternalSenders = value
}
// Sets the appRoleAssignments property value. Represents the app roles a group has been granted for an application. Supports $expand.
// Parameters:
//  - value : Value to set for the appRoleAssignments property.
func (m *Group) SetAppRoleAssignments(value []AppRoleAssignment)() {
    m.appRoleAssignments = value
}
// Sets the assignedLabels property value. The list of sensitivity label pairs (label ID, label name) associated with a Microsoft 365 group. Returned only on $select. Read-only.
// Parameters:
//  - value : Value to set for the assignedLabels property.
func (m *Group) SetAssignedLabels(value []AssignedLabel)() {
    m.assignedLabels = value
}
// Sets the assignedLicenses property value. The licenses that are assigned to the group. Returned only on $select. Supports $filter (eq).Read-only.
// Parameters:
//  - value : Value to set for the assignedLicenses property.
func (m *Group) SetAssignedLicenses(value []AssignedLicense)() {
    m.assignedLicenses = value
}
// Sets the autoSubscribeNewMembers property value. Indicates if new members added to the group will be auto-subscribed to receive email notifications. You can set this property in a PATCH request for the group; do not set it in the initial POST request that creates the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
// Parameters:
//  - value : Value to set for the autoSubscribeNewMembers property.
func (m *Group) SetAutoSubscribeNewMembers(value *bool)() {
    m.autoSubscribeNewMembers = value
}
// Sets the calendar property value. The group's calendar. Read-only.
// Parameters:
//  - value : Value to set for the calendar property.
func (m *Group) SetCalendar(value *Calendar)() {
    m.calendar = value
}
// Sets the calendarView property value. The calendar view for the calendar. Read-only.
// Parameters:
//  - value : Value to set for the calendarView property.
func (m *Group) SetCalendarView(value []Event)() {
    m.calendarView = value
}
// Sets the classification property value. Describes a classification for the group (such as low, medium or high business impact). Valid values for this property are defined by creating a ClassificationList setting value, based on the template definition.Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith).
// Parameters:
//  - value : Value to set for the classification property.
func (m *Group) SetClassification(value *string)() {
    m.classification = value
}
// Sets the conversations property value. The group's conversations.
// Parameters:
//  - value : Value to set for the conversations property.
func (m *Group) SetConversations(value []Conversation)() {
    m.conversations = value
}
// Sets the createdDateTime property value. Timestamp of when the group was created. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *Group) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the createdOnBehalfOf property value. The user (or application) that created the group. NOTE: This is not set if the user is an administrator. Read-only.
// Parameters:
//  - value : Value to set for the createdOnBehalfOf property.
func (m *Group) SetCreatedOnBehalfOf(value *DirectoryObject)() {
    m.createdOnBehalfOf = value
}
// Sets the description property value. An optional description for the group. Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith) and $search.
// Parameters:
//  - value : Value to set for the description property.
func (m *Group) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith), $search, and $orderBy.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Group) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the drive property value. The group's default drive. Read-only.
// Parameters:
//  - value : Value to set for the drive property.
func (m *Group) SetDrive(value *Drive)() {
    m.drive = value
}
// Sets the drives property value. The group's drives. Read-only.
// Parameters:
//  - value : Value to set for the drives property.
func (m *Group) SetDrives(value []Drive)() {
    m.drives = value
}
// Sets the events property value. The group's calendar events.
// Parameters:
//  - value : Value to set for the events property.
func (m *Group) SetEvents(value []Event)() {
    m.events = value
}
// Sets the expirationDateTime property value. Timestamp of when the group is set to expire. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *Group) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the extensions property value. The collection of open extensions defined for the group. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the extensions property.
func (m *Group) SetExtensions(value []Extension)() {
    m.extensions = value
}
// Sets the groupLifecyclePolicies property value. The collection of lifecycle policies for this group. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the groupLifecyclePolicies property.
func (m *Group) SetGroupLifecyclePolicies(value []GroupLifecyclePolicy)() {
    m.groupLifecyclePolicies = value
}
// Sets the groupTypes property value. Specifies the group type and its membership.  If the collection contains Unified, the group is a Microsoft 365 group; otherwise, it's either a security group or distribution group. For details, see groups overview.If the collection includes DynamicMembership, the group has dynamic membership; otherwise, membership is static.  Returned by default. Supports $filter (eq, NOT).
// Parameters:
//  - value : Value to set for the groupTypes property.
func (m *Group) SetGroupTypes(value []string)() {
    m.groupTypes = value
}
// Sets the hasMembersWithLicenseErrors property value. Indicates whether there are members in this group that have license errors from its group-based license assignment. This property is never returned on a GET operation. You can use it as a $filter argument to get groups that have members with license errors (that is, filter for this property being true). See an example. Supports $filter (eq).
// Parameters:
//  - value : Value to set for the hasMembersWithLicenseErrors property.
func (m *Group) SetHasMembersWithLicenseErrors(value *bool)() {
    m.hasMembersWithLicenseErrors = value
}
// Sets the hideFromAddressLists property value. True if the group is not displayed in certain parts of the Outlook UI: the Address Book, address lists for selecting message recipients, and the Browse Groups dialog for searching groups; otherwise, false. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
// Parameters:
//  - value : Value to set for the hideFromAddressLists property.
func (m *Group) SetHideFromAddressLists(value *bool)() {
    m.hideFromAddressLists = value
}
// Sets the hideFromOutlookClients property value. True if the group is not displayed in Outlook clients, such as Outlook for Windows and Outlook on the web; otherwise, false. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
// Parameters:
//  - value : Value to set for the hideFromOutlookClients property.
func (m *Group) SetHideFromOutlookClients(value *bool)() {
    m.hideFromOutlookClients = value
}
// Sets the isArchived property value. 
// Parameters:
//  - value : Value to set for the isArchived property.
func (m *Group) SetIsArchived(value *bool)() {
    m.isArchived = value
}
// Sets the isAssignableToRole property value. Indicates whether this group can be assigned to an Azure Active Directory role or not. Optional. This property can only be set while creating the group and is immutable. If set to true, the securityEnabled property must also be set to true and the group cannot be a dynamic group (that is, groupTypes cannot contain DynamicMembership). Only callers in Global administrator and Privileged role administrator roles can set this property. The caller must be assigned the RoleManagement.ReadWrite.Directory permission to set this property or update the membership of such groups. For more, see Using a group to manage Azure AD role assignmentsReturned by default. Supports $filter (eq, ne, NOT).
// Parameters:
//  - value : Value to set for the isAssignableToRole property.
func (m *Group) SetIsAssignableToRole(value *bool)() {
    m.isAssignableToRole = value
}
// Sets the isSubscribedByMail property value. Indicates whether the signed-in user is subscribed to receive email conversations. Default value is true. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
// Parameters:
//  - value : Value to set for the isSubscribedByMail property.
func (m *Group) SetIsSubscribedByMail(value *bool)() {
    m.isSubscribedByMail = value
}
// Sets the licenseProcessingState property value. Indicates status of the group license assignment to all members of the group. Default value is false. Read-only. Possible values: QueuedForProcessing, ProcessingInProgress, and ProcessingComplete.Returned only on $select. Read-only.
// Parameters:
//  - value : Value to set for the licenseProcessingState property.
func (m *Group) SetLicenseProcessingState(value *LicenseProcessingState)() {
    m.licenseProcessingState = value
}
// Sets the mail property value. The SMTP address for the group, for example, 'serviceadmins@contoso.onmicrosoft.com'. Returned by default. Read-only. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the mail property.
func (m *Group) SetMail(value *string)() {
    m.mail = value
}
// Sets the mailEnabled property value. Specifies whether the group is mail-enabled. Required. Returned by default. Supports $filter (eq, ne, NOT).
// Parameters:
//  - value : Value to set for the mailEnabled property.
func (m *Group) SetMailEnabled(value *bool)() {
    m.mailEnabled = value
}
// Sets the mailNickname property value. The mail alias for the group, unique in the organization. Maximum length is 64 characters. This property can contain only characters in the ASCII character set 0 - 127 except the following: @ () / [] ' ; : . <> , SPACE. Required. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the mailNickname property.
func (m *Group) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// Sets the memberOf property value. Groups that this group is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the memberOf property.
func (m *Group) SetMemberOf(value []DirectoryObject)() {
    m.memberOf = value
}
// Sets the members property value. Users and groups that are members of this group. HTTP Methods: GET (supported for all groups), POST (supported for Microsoft 365 groups, security groups and mail-enabled security groups), DELETE (supported for Microsoft 365 groups and security groups). Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the members property.
func (m *Group) SetMembers(value []DirectoryObject)() {
    m.members = value
}
// Sets the membershipRule property value. The rule that determines members for this group if the group is a dynamic group (groupTypes contains DynamicMembership). For more information about the syntax of the membership rule, see Membership Rules syntax. Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith).
// Parameters:
//  - value : Value to set for the membershipRule property.
func (m *Group) SetMembershipRule(value *string)() {
    m.membershipRule = value
}
// Sets the membershipRuleProcessingState property value. Indicates whether the dynamic membership processing is on or paused. Possible values are On or Paused. Returned by default. Supports $filter (eq, ne, NOT, in).
// Parameters:
//  - value : Value to set for the membershipRuleProcessingState property.
func (m *Group) SetMembershipRuleProcessingState(value *string)() {
    m.membershipRuleProcessingState = value
}
// Sets the membersWithLicenseErrors property value. A list of group members with license errors from this group-based license assignment. Read-only.
// Parameters:
//  - value : Value to set for the membersWithLicenseErrors property.
func (m *Group) SetMembersWithLicenseErrors(value []DirectoryObject)() {
    m.membersWithLicenseErrors = value
}
// Sets the onenote property value. Read-only.
// Parameters:
//  - value : Value to set for the onenote property.
func (m *Group) SetOnenote(value *Onenote)() {
    m.onenote = value
}
// Sets the onPremisesDomainName property value. Contains the on-premises domain FQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only.
// Parameters:
//  - value : Value to set for the onPremisesDomainName property.
func (m *Group) SetOnPremisesDomainName(value *string)() {
    m.onPremisesDomainName = value
}
// Sets the onPremisesLastSyncDateTime property value. Indicates the last time at which the group was synced with the on-premises directory.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only. Supports $filter (eq, ne, NOT, ge, le, in).
// Parameters:
//  - value : Value to set for the onPremisesLastSyncDateTime property.
func (m *Group) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onPremisesLastSyncDateTime = value
}
// Sets the onPremisesNetBiosName property value. Contains the on-premises netBios name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only.
// Parameters:
//  - value : Value to set for the onPremisesNetBiosName property.
func (m *Group) SetOnPremisesNetBiosName(value *string)() {
    m.onPremisesNetBiosName = value
}
// Sets the onPremisesProvisioningErrors property value. Errors when using Microsoft synchronization product during provisioning. Returned by default. Supports $filter (eq, NOT).
// Parameters:
//  - value : Value to set for the onPremisesProvisioningErrors property.
func (m *Group) SetOnPremisesProvisioningErrors(value []OnPremisesProvisioningError)() {
    m.onPremisesProvisioningErrors = value
}
// Sets the onPremisesSamAccountName property value. Contains the on-premises SAM account name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith). Read-only.
// Parameters:
//  - value : Value to set for the onPremisesSamAccountName property.
func (m *Group) SetOnPremisesSamAccountName(value *string)() {
    m.onPremisesSamAccountName = value
}
// Sets the onPremisesSecurityIdentifier property value. Contains the on-premises security identifier (SID) for the group that was synchronized from on-premises to the cloud. Returned by default. Supports $filter on null values. Read-only.
// Parameters:
//  - value : Value to set for the onPremisesSecurityIdentifier property.
func (m *Group) SetOnPremisesSecurityIdentifier(value *string)() {
    m.onPremisesSecurityIdentifier = value
}
// Sets the onPremisesSyncEnabled property value. true if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Returned by default. Read-only. Supports $filter (eq, ne, NOT, in).
// Parameters:
//  - value : Value to set for the onPremisesSyncEnabled property.
func (m *Group) SetOnPremisesSyncEnabled(value *bool)() {
    m.onPremisesSyncEnabled = value
}
// Sets the owners property value. The owners of the group. The owners are a set of non-admin users who are allowed to modify this object. Limited to 100 owners. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand.
// Parameters:
//  - value : Value to set for the owners property.
func (m *Group) SetOwners(value []DirectoryObject)() {
    m.owners = value
}
// Sets the permissionGrants property value. The permission that has been granted for a group to a specific application. Supports $expand.
// Parameters:
//  - value : Value to set for the permissionGrants property.
func (m *Group) SetPermissionGrants(value []ResourceSpecificPermissionGrant)() {
    m.permissionGrants = value
}
// Sets the photo property value. The group's profile photo
// Parameters:
//  - value : Value to set for the photo property.
func (m *Group) SetPhoto(value *ProfilePhoto)() {
    m.photo = value
}
// Sets the photos property value. The profile photos owned by the group. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the photos property.
func (m *Group) SetPhotos(value []ProfilePhoto)() {
    m.photos = value
}
// Sets the planner property value. Entry-point to Planner resource that might exist for a Unified Group.
// Parameters:
//  - value : Value to set for the planner property.
func (m *Group) SetPlanner(value *PlannerGroup)() {
    m.planner = value
}
// Sets the preferredDataLocation property value. The preferred data location for the Microsoft 365 group. By default, the group inherits the group creator's preferred data location. To set this property, the calling user must be assigned one of the following Azure AD roles:  Global Administrator  User Account Administrator Directory Writer  Exchange Administrator  SharePoint Administrator  For more information about this property, see  OneDrive Online Multi-Geo. Nullable. Returned by default.
// Parameters:
//  - value : Value to set for the preferredDataLocation property.
func (m *Group) SetPreferredDataLocation(value *string)() {
    m.preferredDataLocation = value
}
// Sets the preferredLanguage property value. The preferred language for a Microsoft 365 group. Should follow ISO 639-1 Code; for example en-US. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith).
// Parameters:
//  - value : Value to set for the preferredLanguage property.
func (m *Group) SetPreferredLanguage(value *string)() {
    m.preferredLanguage = value
}
// Sets the proxyAddresses property value. Email addresses for the group that direct to the same group mailbox. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. The any operator is required to filter expressions on multi-valued properties. Returned by default. Read-only. Not nullable. Supports $filter (eq, NOT, ge, le, startsWith).
// Parameters:
//  - value : Value to set for the proxyAddresses property.
func (m *Group) SetProxyAddresses(value []string)() {
    m.proxyAddresses = value
}
// Sets the rejectedSenders property value. The list of users or groups that are not allowed to create posts or calendar events in this group. Nullable
// Parameters:
//  - value : Value to set for the rejectedSenders property.
func (m *Group) SetRejectedSenders(value []DirectoryObject)() {
    m.rejectedSenders = value
}
// Sets the renewedDateTime property value. Timestamp of when the group was last renewed. This cannot be modified directly and is only updated via the renew service action. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
// Parameters:
//  - value : Value to set for the renewedDateTime property.
func (m *Group) SetRenewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.renewedDateTime = value
}
// Sets the securityEnabled property value. Specifies whether the group is a security group. Required. Returned by default. Supports $filter (eq, ne, NOT, in).
// Parameters:
//  - value : Value to set for the securityEnabled property.
func (m *Group) SetSecurityEnabled(value *bool)() {
    m.securityEnabled = value
}
// Sets the securityIdentifier property value. Security identifier of the group, used in Windows scenarios. Returned by default.
// Parameters:
//  - value : Value to set for the securityIdentifier property.
func (m *Group) SetSecurityIdentifier(value *string)() {
    m.securityIdentifier = value
}
// Sets the settings property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the settings property.
func (m *Group) SetSettings(value []GroupSetting)() {
    m.settings = value
}
// Sets the sites property value. The list of SharePoint sites in this group. Access the default site with /sites/root.
// Parameters:
//  - value : Value to set for the sites property.
func (m *Group) SetSites(value []Site)() {
    m.sites = value
}
// Sets the team property value. 
// Parameters:
//  - value : Value to set for the team property.
func (m *Group) SetTeam(value *Team)() {
    m.team = value
}
// Sets the theme property value. Specifies a Microsoft 365 group's color theme. Possible values are Teal, Purple, Green, Blue, Pink, Orange or Red. Returned by default.
// Parameters:
//  - value : Value to set for the theme property.
func (m *Group) SetTheme(value *string)() {
    m.theme = value
}
// Sets the threads property value. The group's conversation threads. Nullable.
// Parameters:
//  - value : Value to set for the threads property.
func (m *Group) SetThreads(value []ConversationThread)() {
    m.threads = value
}
// Sets the transitiveMemberOf property value. 
// Parameters:
//  - value : Value to set for the transitiveMemberOf property.
func (m *Group) SetTransitiveMemberOf(value []DirectoryObject)() {
    m.transitiveMemberOf = value
}
// Sets the transitiveMembers property value. 
// Parameters:
//  - value : Value to set for the transitiveMembers property.
func (m *Group) SetTransitiveMembers(value []DirectoryObject)() {
    m.transitiveMembers = value
}
// Sets the unseenCount property value. Count of conversations that have received new posts since the signed-in user last visited the group. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
// Parameters:
//  - value : Value to set for the unseenCount property.
func (m *Group) SetUnseenCount(value *int32)() {
    m.unseenCount = value
}
// Sets the visibility property value. Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or Hiddenmembership. Hiddenmembership can be set only for Microsoft 365 groups, when the groups are created. It can't be updated later. Other values of visibility can be updated after group creation. If visibility value is not specified during group creation on Microsoft Graph, a security group is created as Private by default and Microsoft 365 group is Public. Groups assignable to roles are always Private. See group visibility options to learn more. Returned by default. Nullable.
// Parameters:
//  - value : Value to set for the visibility property.
func (m *Group) SetVisibility(value *string)() {
    m.visibility = value
}
