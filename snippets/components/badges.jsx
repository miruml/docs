export const MutableBadge = ({ size = "sm" }) => {
    return <Badge icon="pencil" color="orange" size={size}>mutable</Badge>;
};

export const ImmutableBadge = ({ size = "sm" }) => {
    return <Badge icon="lock" color="gray" size={size}>immutable</Badge>;
};
